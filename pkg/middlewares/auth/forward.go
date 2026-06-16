package auth

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/traefik/traefik/v2/pkg/config/dynamic"
	"github.com/traefik/traefik/v2/pkg/log"
	"github.com/traefik/traefik/v2/pkg/middlewares"
	"github.com/traefik/traefik/v2/pkg/tracing"
	"github.com/vulcand/oxy/utils"
)

const (
	xForwardedURI     = "X-Forwarded-Uri"
	xForwardedMethod  = "X-Forwarded-Method"
	xForwardedTLSUser = "X-Forwarded-Tls-Client-Cert-Info"
)

type forwardAuth struct {
	address                  string
	authResponseHeaders      []string
	authResponseHeadersRegex *regexp.Regexp
	authRequestHeaders       []string
	next                     http.Handler
	name                     string
	client                   *http.Client
	trustForwardHeader       bool
}

// New creates a forward auth middleware.
func New(ctx context.Context, next http.Handler, config dynamic.ForwardAuth, name string) (http.Handler, error) {
	logger := log.FromContext(middlewares.GetLoggerCtx(ctx, name))
	logger.Debug("Creating middleware")

	var authResponseHeadersRegex *regexp.Regexp
	if len(config.AuthResponseHeadersRegex) > 0 {
		var err error
		authResponseHeadersRegex, err = regexp.Compile(config.AuthResponseHeadersRegex)
		if err != nil {
			return nil, fmt.Errorf("error compiling response headers regex %q: %w", config.AuthResponseHeadersRegex, err)
		}
	}

	fa := &forwardAuth{
		address:                  config.Address,
		authResponseHeaders:      config.AuthResponseHeaders,
		authResponseHeadersRegex: authResponseHeadersRegex,
		authRequestHeaders:       config.AuthRequestHeaders,
		next:                     next,
		name:                     name,
		trustForwardHeader:       config.TrustForwardHeader,
		client: &http.Client{
			CheckRedirect: func(r *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Timeout: time.Duration(config.Timeout),
		},
	}

	if config.TLS != nil {
		tlsConfig, err := config.TLS.CreateTLSConfig(ctx)
		if err != nil {
			return nil, fmt.Errorf("unable to create client TLS: %w", err)
		}

		tr := http.DefaultTransport.(*http.Transport).Clone()
		tr.TLSClientConfig = tlsConfig
		fa.client.Transport = tr
	}

	return fa, nil
}

func (f *forwardAuth) GetName() string {
	return f.name
}

func (f *forwardAuth) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	logger := log.FromContext(middlewares.GetLoggerCtx(req.Context(), f.name))

	authReq, err := http.NewRequestWithContext(req.Context(), http.MethodGet, f.address, nil)
	if err != nil {
		logger.Errorf("Error calling %s: %s", f.address, err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	writeHeader(req, authReq, f.trustForwardHeader)

	// Copy request headers
	if len(f.authRequestHeaders) > 0 {
		for _, headerName := range f.authRequestHeaders {
			headerKey := http.CanonicalHeaderKey(headerName)
			if reqHeaderValue := req.Header.Get(headerKey); reqHeaderValue != "" {
				authReq.Header.Set(headerKey, reqHeaderValue)
			}
		}
	}

	resp, err := f.client.Do(authReq)
	if err != nil {
		logger.Errorf("Error calling %s: %s", f.address, err)
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			rw.WriteHeader(http.StatusGatewayTimeout)
		} else {
			rw.WriteHeader(http.StatusBadGateway)
		}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		// Copy response headers
		if f.authResponseHeadersRegex != nil {
			for headerName, headerValues := range resp.Header {
				if f.authResponseHeadersRegex.MatchString(headerName) {
					req.Header.Del(headerName)
					for _, headerValue := range headerValues {
						req.Header.Add(headerName, headerValue)
					}
				}
			}
		} else if len(f.authResponseHeaders) > 0 {
			for _, headerName := range f.authResponseHeaders {
				headerKey := http.CanonicalHeaderKey(headerName)
				req.Header.Del(headerKey)
				if respHeaderValue := resp.Header.Get(headerKey); respHeaderValue != "" {
					req.Header.Set(headerKey, respHeaderValue)
				}
			}
		}

		f.next.ServeHTTP(rw, req)
		return
	}

	// Copy response headers to client
	for headerName, headerValues := range resp.Header {
		rw.Header().Del(headerName)
		for _, headerValue := range headerValues {
			rw.Header().Add(headerName, headerValue)
		}
	}

	rw.WriteHeader(resp.StatusCode)
	_, err = io.Copy(rw, resp.Body)
	if err != nil {
		logger.Errorf("Error copying response body: %s", err)
	}
}

func writeHeader(req, authReq *http.Request, trustForwardHeader bool) {
	utils.CopyHeaders(authReq.Header, req.Header)
	authReq.Header.Set(xForwardedURI, req.URL.RequestURI())
	authReq.Header.Set(xForwardedMethod, req.Method)

	if clientURL, err := url.Parse(req.Header.Get("X-Forwarded-Proto")); err == nil && clientURL.Scheme != "" {
		authReq.Header.Set("X-Forwarded-Proto", clientURL.Scheme)
	}

	if trustForwardHeader {
		if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
			if prior, ok := req.Header["X-Forwarded-For"]; ok {
				clientIP = strings.Join(prior, ", ") + ", " + clientIP
			}
			authReq.Header.Set("X-Forwarded-For", clientIP)
		}
	} else {
		authReq.Header.Del("X-Forwarded-For")
		if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
			authReq.Header.Set("X-Forwarded-For", clientIP)
		}
	}

	if req.TLS != nil {
		authReq.Header.Set("X-Forwarded-Proto", "https")
		if len(req.TLS.PeerCertificates) > 0 {
			authReq.Header.Set(xForwardedTLSUser, url.QueryEscape(req.TLS.PeerCertificates[0].Subject.String()))
		}
	} else {
		authReq.Header.Set("X-Forwarded-Proto", "http")
	}
}
