package dynamic

import (
	"github.com/traefik/paerser/types"
)

// Middleware holds the Middleware configuration.
type Middleware struct {
	AddPrefix         *AddPrefix         `json:"addPrefix,omitempty" toml:"addPrefix,omitempty" yaml:"addPrefix,omitempty" export:"true"`
	StripPrefix       *StripPrefix       `json:"stripPrefix,omitempty" toml:"stripPrefix,omitempty" yaml:"stripPrefix,omitempty" export:"true"`
	StripPrefixRegex  *StripPrefixRegex  `json:"stripPrefixRegex,omitempty" toml:"stripPrefixRegex,omitempty" yaml:"stripPrefixRegex,omitempty" export:"true"`
	ReplacePath       *ReplacePath       `json:"replacePath,omitempty" toml:"replacePath,omitempty" yaml:"replacePath,omitempty" export:"true"`
	ReplacePathRegex  *ReplacePathRegex  `json:"replacePathRegex,omitempty" toml:"replacePathRegex,omitempty" yaml:"replacePathRegex,omitempty" export:"true"`
	Chain             *Chain             `json:"chain,omitempty" toml:"chain,omitempty" yaml:"chain,omitempty" export:"true"`
	IPWhiteList       *IPWhiteList       `json:"ipWhiteList,omitempty" toml:"ipWhiteList,omitempty" yaml:"ipWhiteList,omitempty" export:"true"`
	Headers           *Headers           `json:"headers,omitempty" toml:"headers,omitempty" yaml:"headers,omitempty" export:"true"`
	Errors            *ErrorPage         `json:"errors,omitempty" toml:"errors,omitempty" yaml:"errors,omitempty" export:"true"`
	RateLimit         *RateLimit         `json:"rateLimit,omitempty" toml:"rateLimit,omitempty" yaml:"rateLimit,omitempty" export:"true"`
	RedirectRegex     *RedirectRegex     `json:"redirectRegex,omitempty" toml:"redirectRegex,omitempty" yaml:"redirectRegex,omitempty" export:"true"`
	RedirectScheme    *RedirectScheme    `json:"redirectScheme,omitempty" toml:"redirectScheme,omitempty" yaml:"redirectScheme,omitempty" export:"true"`
	BasicAuth         *BasicAuth         `json:"basicAuth,omitempty" toml:"basicAuth,omitempty" yaml:"basicAuth,omitempty" export:"true"`
	DigestAuth        *DigestAuth        `json:"digestAuth,omitempty" toml:"digestAuth,omitempty" yaml:"digestAuth,omitempty" export:"true"`
	ForwardAuth       *ForwardAuth       `json:"forwardAuth,omitempty" toml:"forwardAuth,omitempty" yaml:"forwardAuth,omitempty" export:"true"`
	MaxRequestBody    *MaxRequestBody    `json:"maxRequestBody,omitempty" toml:"maxRequestBody,omitempty" yaml:"maxRequestBody,omitempty" export:"true"`
	Buffering         *Buffering         `json:"buffering,omitempty" toml:"buffering,omitempty" yaml:"buffering,omitempty" export:"true"`
	CircuitBreaker    *CircuitBreaker    `json:"circuitBreaker,omitempty" toml:"circuitBreaker,omitempty" yaml:"circuitBreaker,omitempty" export:"true"`
	Compress          *Compress          `json:"compress,omitempty" toml:"compress,omitempty" yaml:"compress,omitempty" export:"true"`
	PassTLSClientCert *PassTLSClientCert `json:"passTLSClientCert,omitempty" toml:"passTLSClientCert,omitempty" yaml:"passTLSClientCert,omitempty" export:"true"`
	Retry             *Retry             `json:"retry,omitempty" toml:"retry,omitempty" yaml:"retry,omitempty" export:"true"`
	InFlightReq       *InFlightReq       `json:"inFlightReq,omitempty" toml:"inFlightReq,omitempty" yaml:"inFlightReq,omitempty" export:"true"`
}

// AddPrefix holds the select-route-by-prefix middleware configuration.
type AddPrefix struct {
	Prefix string `json:"prefix,omitempty" toml:"prefix,omitempty" yaml:"prefix,omitempty" export:"true"`
}

// StripPrefix holds the strip prefix middleware configuration.
type StripPrefix struct {
	Prefixes   []string `json:"prefixes,omitempty" toml:"prefixes,omitempty" yaml:"prefixes,omitempty" export:"true"`
	ForceSlash bool     `json:"forceSlash,omitempty" toml:"forceSlash,omitempty" yaml:"forceSlash,omitempty" export:"true"`
}

// StripPrefixRegex holds the strip prefix regex middleware configuration.
type StripPrefixRegex struct {
	Regex []string `json:"regex,omitempty" toml:"regex,omitempty" yaml:"regex,omitempty" export:"true"`
}

// ReplacePath holds the replace path middleware configuration.
type ReplacePath struct {
	Path string `json:"path,omitempty" toml:"path,omitempty" yaml:"path,omitempty" export:"true"`
}

// ReplacePathRegex holds the replace path regex middleware configuration.
type ReplacePathRegex struct {
	Regex       string `json:"regex,omitempty" toml:"regex,omitempty" yaml:"regex,omitempty" export:"true"`
	Replacement string `json:"replacement,omitempty" toml:"replacement,omitempty" yaml:"replacement,omitempty" export:"true"`
}

// Chain holds the chain middleware configuration.
type Chain struct {
	Middlewares []string `json:"middlewares,omitempty" toml:"middlewares,omitempty" yaml:"middlewares,omitempty" export:"true"`
}

// IPWhiteList holds the ip white list middleware configuration.
type IPWhiteList struct {
	SourceRange []string     `json:"sourceRange,omitempty" toml:"sourceRange,omitempty" yaml:"sourceRange,omitempty"`
	IPStrategy  *IPStrategy  `json:"ipStrategy,omitempty" toml:"ipStrategy,omitempty" yaml:"ipStrategy,omitempty" export:"true"`
}

// IPStrategy holds the IP strategy configuration.
type IPStrategy struct {
	Depth       int      `json:"depth,omitempty" toml:"depth,omitempty" yaml:"depth,omitempty" export:"true"`
	ExcludedIPs []string `json:"excludedIPs,omitempty" toml:"excludedIPs,omitempty" yaml:"excludedIPs,omitempty" export:"true"`
}

// Headers holds the headers middleware configuration.
type Headers struct {
	CustomRequestHeaders          map[string]string `json:"customRequestHeaders,omitempty" toml:"customRequestHeaders,omitempty" yaml:"customRequestHeaders,omitempty" export:"true"`
	CustomResponseHeaders         map[string]string `json:"customResponseHeaders,omitempty" toml:"customResponseHeaders,omitempty" yaml:"customResponseHeaders,omitempty" export:"true"`
	AccessControlAllowCredentials bool              `json:"accessControlAllowCredentials,omitempty" toml:"accessControlAllowCredentials,omitempty" yaml:"accessControlAllowCredentials,omitempty" export:"true"`
	AccessControlAllowHeaders     []string          `json:"accessControlAllowHeaders,omitempty" toml:"accessControlAllowHeaders,omitempty" yaml:"accessControlAllowHeaders,omitempty" export:"true"`
	AccessControlAllowMethods     []string          `json:"accessControlAllowMethods,omitempty" toml:"accessControlAllowMethods,omitempty" yaml:"accessControlAllowMethods,omitempty" export:"true"`
	AccessControlAllowOriginList  []string          `json:"accessControlAllowOriginList,omitempty" toml:"accessControlAllowOriginList,omitempty" yaml:"accessControlAllowOriginList,omitempty" export:"true"`
	AccessControlAllowOrigin      string            `json:"accessControlAllowOrigin,omitempty" toml:"accessControlAllowOrigin,omitempty" yaml:"accessControlAllowOrigin,omitempty" export:"true"`
	AccessControlExposeHeaders    []string          `json:"accessControlExposeHeaders,omitempty" toml:"accessControlExposeHeaders,omitempty" yaml:"accessControlExposeHeaders,omitempty" export:"true"`
	AccessControlMaxAge           int64             `json:"accessControlMaxAge,omitempty" toml:"accessControlMaxAge,omitempty" yaml:"accessControlMaxAge,omitempty" export:"true"`
	AddVaryHeader                 bool              `json:"addVaryHeader,omitempty" toml:"addVaryHeader,omitempty" yaml:"addVaryHeader,omitempty" export:"true"`
	AllowedHosts                  []string          `json:"allowedHosts,omitempty" toml:"allowedHosts,omitempty" yaml:"allowedHosts,omitempty" export:"true"`
	HostsProxyHeaders             []string          `json:"hostsProxyHeaders,omitempty" toml:"hostsProxyHeaders,omitempty" yaml:"hostsProxyHeaders,omitempty" export:"true"`
	SSLRedirect                   bool              `json:"sslRedirect,omitempty" toml:"sslRedirect,omitempty" yaml:"sslRedirect,omitempty" export:"true"`
	SSLTemporaryRedirect          bool              `json:"sslTemporaryRedirect,omitempty" toml:"sslTemporaryRedirect,omitempty" yaml:"sslTemporaryRedirect,omitempty" export:"true"`
	SSLHost                       string            `json:"sslHost,omitempty" toml:"sslHost,omitempty" yaml:"sslHost,omitempty" export:"true"`
	SSLForceHost                  bool              `json:"sslForceHost,omitempty" toml:"sslForceHost,omitempty" yaml:"sslForceHost,omitempty" export:"true"`
	SSLProxyHeaders               map[string]string `json:"sslProxyHeaders,omitempty" toml:"sslProxyHeaders,omitempty" yaml:"sslProxyHeaders,omitempty" export:"true"`
	STSSeconds                    int64             `json:"stsSeconds,omitempty" toml:"stsSeconds,omitempty" yaml:"stsSeconds,omitempty" export:"true"`
	STSIncludeSubdomains          bool              `json:"stsIncludeSubdomains,omitempty" toml:"stsIncludeSubdomains,omitempty" yaml:"stsIncludeSubdomains,omitempty" export:"true"`
	STSPreload                    bool              `json:"stsPreload,omitempty" toml:"stsPreload,omitempty" yaml:"stsPreload,omitempty" export:"true"`
	ForceSTSHeader                bool              `json:"forceSTSHeader,omitempty" toml:"forceSTSHeader,omitempty" yaml:"forceSTSHeader,omitempty" export:"true"`
	FrameDeny                     bool              `json:"frameDeny,omitempty" toml:"frameDeny,omitempty" yaml:"frameDeny,omitempty" export:"true"`
	CustomFrameOptionsValue       string            `json:"customFrameOptionsValue,omitempty" toml:"customFrameOptionsValue,omitempty" yaml:"customFrameOptionsValue,omitempty" export:"true"`
	ContentTypeNosniff            bool              `json:"contentTypeNosniff,omitempty" toml:"contentTypeNosniff,omitempty" yaml:"contentTypeNosniff,omitempty" export:"true"`
	BrowserXSSFilter              bool              `json:"browserXssFilter,omitempty" toml:"browserXssFilter,omitempty" yaml:"browserXssFilter,omitempty" export:"true"`
	CustomBrowserXSSValue         string            `json:"customBrowserXssValue,omitempty" toml:"customBrowserXssValue,omitempty" yaml:"customBrowserXssValue,omitempty" export:"true"`
	ContentSecurityPolicy         string            `json:"contentSecurityPolicy,omitempty" toml:"contentSecurityPolicy,omitempty" yaml:"contentSecurityPolicy,omitempty" export:"true"`
	PublicKey                     string            `json:"publicKey,omitempty" toml:"publicKey,omitempty" yaml:"publicKey,omitempty" export:"true"`
	ReferrerPolicy                string            `json:"referrerPolicy,omitempty" toml:"referrerPolicy,omitempty" yaml:"referrerPolicy,omitempty" export:"true"`
	FeaturePolicy                 string            `json:"featurePolicy,omitempty" toml:"featurePolicy,omitempty" yaml:"featurePolicy,omitempty" export:"true"`
	PermissionsPolicy             string            `json:"permissionsPolicy,omitempty" toml:"permissionsPolicy,omitempty" yaml:"permissionsPolicy,omitempty" export:"true"`
	IsDevelopment                 bool              `json:"isDevelopment,omitempty" toml:"isDevelopment,omitempty" yaml:"isDevelopment,omitempty" export:"true"`
}

// ErrorPage holds the custom error page middleware configuration.
type ErrorPage struct {
	Status  []string `json:"status,omitempty" toml:"status,omitempty" yaml:"status,omitempty" export:"true"`
	Service string   `json:"service,omitempty" toml:"service,omitempty" yaml:"service,omitempty" export:"true"`
	Query   string   `json:"query,omitempty" toml:"query,omitempty" yaml:"query,omitempty" export:"true"`
}

// RateLimit holds the rate limit middleware configuration.
type RateLimit struct {
	Average         int64            `json:"average,omitempty" toml:"average,omitempty" yaml:"average,omitempty" export:"true"`
	Period          types.Duration   `json:"period,omitempty" toml:"period,omitempty" yaml:"period,omitempty" export:"true"`
	Burst           int64            `json:"burst,omitempty" toml:"burst,omitempty" yaml:"burst,omitempty" export:"true"`
	SourceCriterion *SourceCriterion `json:"sourceCriterion,omitempty" toml:"sourceCriterion,omitempty" yaml:"sourceCriterion,omitempty" export:"true"`
}

// SourceCriterion holds the source criterion configuration.
type SourceCriterion struct {
	IPStrategy        *IPStrategy `json:"ipStrategy,omitempty" toml:"ipStrategy,omitempty" yaml:"ipStrategy,omitempty" export:"true"`
	RequestHeaderName string      `json:"requestHeaderName,omitempty" toml:"requestHeaderName,omitempty" yaml:"requestHeaderName,omitempty" export:"true"`
	RequestHost       bool        `json:"requestHost,omitempty" toml:"requestHost,omitempty" yaml:"requestHost,omitempty" export:"true"`
}

// RedirectRegex holds the redirect regex middleware configuration.
type RedirectRegex struct {
	Regex       string `json:"regex,omitempty" toml:"regex,omitempty" yaml:"regex,omitempty" export:"true"`
	Replacement string `json:"replacement,omitempty" toml:"replacement,omitempty" yaml:"replacement,omitempty" export:"true"`
	Permanent   bool   `json:"permanent,omitempty" toml:"permanent,omitempty" yaml:"permanent,omitempty" export:"true"`
}

// RedirectScheme holds the redirect scheme middleware configuration.
type RedirectScheme struct {
	Scheme    string `json:"scheme,omitempty" toml:"scheme,omitempty" yaml:"scheme,omitempty" export:"true"`
	Port      string `json:"port,omitempty" toml:"port,omitempty" yaml:"port,omitempty" export:"true"`
	Permanent bool   `json:"permanent,omitempty" toml:"permanent,omitempty" yaml:"permanent,omitempty" export:"true"`
}

// BasicAuth holds the basic auth middleware configuration.
type BasicAuth struct {
	Users        []string `json:"users,omitempty" toml:"users,omitempty" yaml:"users,omitempty"`
	UsersFile    string   `json:"usersFile,omitempty" toml:"usersFile,omitempty" yaml:"usersFile,omitempty"`
	Realm        string   `json:"realm,omitempty" toml:"realm,omitempty" yaml:"realm,omitempty" export:"true"`
	RemoveHeader bool     `json:"removeHeader,omitempty" toml:"removeHeader,omitempty" yaml:"removeHeader,omitempty" export:"true"`
}

// DigestAuth holds the digest auth middleware configuration.
type DigestAuth struct {
	Users        []string `json:"users,omitempty" toml:"users,omitempty" yaml:"users,omitempty"`
	UsersFile    string   `json:"usersFile,omitempty" toml:"usersFile,omitempty" yaml:"usersFile,omitempty"`
	Realm        string   `json:"realm,omitempty" toml:"realm,omitempty" yaml:"realm,omitempty" export:"true"`
	RemoveHeader bool     `json:"removeHeader,omitempty" toml:"removeHeader,omitempty" yaml:"removeHeader,omitempty" export:"true"`
}

// ForwardAuth holds the forward auth middleware configuration.
type ForwardAuth struct {
	Address                  string           `json:"address,omitempty" toml:"address,omitempty" yaml:"address,omitempty"`
	TLS                      *ClientTLS       `json:"tls,omitempty" toml:"tls,omitempty" yaml:"tls,omitempty" export:"true"`
	TrustForwardHeader       bool             `json:"trustForwardHeader,omitempty" toml:"trustForwardHeader,omitempty" yaml:"trustForwardHeader,omitempty" export:"true"`
	AuthResponseHeaders      []string         `json:"authResponseHeaders,omitempty" toml:"authResponseHeaders,omitempty" yaml:"authResponseHeaders,omitempty" export:"true"`
	AuthResponseHeadersRegex string           `json:"authResponseHeadersRegex,omitempty" toml:"authResponseHeadersRegex,omitempty" yaml:"authResponseHeadersRegex,omitempty" export:"true"`
	AuthRequestHeaders       []string         `json:"authRequestHeaders,omitempty" toml:"authRequestHeaders,omitempty" yaml:"authRequestHeaders,omitempty" export:"true"`
	Timeout                  types.Duration   `json:"timeout,omitempty" toml:"timeout,omitempty" yaml:"timeout,omitempty" export:"true"`
}

// MaxRequestBody holds the max request body middleware configuration.
type MaxRequestBody struct {
	Limit int64 `json:"limit,omitempty" toml:"limit,omitempty" yaml:"limit,omitempty" export:"true"`
}

// Buffering holds the buffering middleware configuration.
type Buffering struct {
	MaxRequestBodyBytes  int64  `json:"maxRequestBodyBytes,omitempty" toml:"maxRequestBodyBytes,omitempty" yaml:"maxRequestBodyBytes,omitempty" export:"true"`
	MemRequestBodyBytes  int64  `json:"memRequestBodyBytes,omitempty" toml:"memRequestBodyBytes,omitempty" yaml:"memRequestBodyBytes,omitempty" export:"true"`
	MaxResponseBodyBytes int64  `json:"maxResponseBodyBytes,omitempty" toml:"maxResponseBodyBytes,omitempty" yaml:"maxResponseBodyBytes,omitempty" export:"true"`
	MemResponseBodyBytes int64  `json:"memResponseBodyBytes,omitempty" toml:"memResponseBodyBytes,omitempty" yaml:"memResponseBodyBytes,omitempty" export:"true"`
	RetryExpression      string `json:"retryExpression,omitempty" toml:"retryExpression,omitempty" yaml:"retryExpression,omitempty" export:"true"`
}

// CircuitBreaker holds the circuit breaker middleware configuration.
type CircuitBreaker struct {
	Expression string `json:"expression,omitempty" toml:"expression,omitempty" yaml:"expression,omitempty" export:"true"`
}

// Compress holds the compress middleware configuration.
type Compress struct {
	ExcludedContentTypes []string `json:"excludedContentTypes,omitempty" toml:"excludedContentTypes,omitempty" yaml:"excludedContentTypes,omitempty" export:"true"`
}

// PassTLSClientCert holds the pass TLS client cert middleware configuration.
type PassTLSClientCert struct {
	PEM  bool              `json:"pem,omitempty" toml:"pem,omitempty" yaml:"pem,omitempty" export:"true"`
	Info *TLSClientHeaders `json:"info,omitempty" toml:"info,omitempty" yaml:"info,omitempty" export:"true"`
}

// TLSClientHeaders holds the TLS client cert headers configuration.
type TLSClientHeaders struct {
	NotAfter  bool                     `json:"notAfter,omitempty" toml:"notAfter,omitempty" yaml:"notAfter,omitempty" export:"true"`
	NotBefore bool                     `json:"notBefore,omitempty" toml:"notBefore,omitempty" yaml:"notBefore,omitempty" export:"true"`
	Sans      bool                     `json:"sans,omitempty" toml:"sans,omitempty" yaml:"sans,omitempty" export:"true"`
	Subject   *TLSClientSubjectHeaders `json:"subject,omitempty" toml:"subject,omitempty" yaml:"subject,omitempty" export:"true"`
	Issuer    *TLSClientIssuerHeaders  `json:"issuer,omitempty" toml:"issuer,omitempty" yaml:"issuer,omitempty" export:"true"`
}

// TLSClientSubjectHeaders holds the TLS client cert subject headers configuration.
type TLSClientSubjectHeaders struct {
	Country            bool `json:"country,omitempty" toml:"country,omitempty" yaml:"country,omitempty" export:"true"`
	Organization       bool `json:"organization,omitempty" toml:"organization,omitempty" yaml:"organization,omitempty" export:"true"`
	OrganizationalUnit bool `json:"organizationalUnit,omitempty" toml:"organizationalUnit,omitempty" yaml:"organizationalUnit,omitempty" export:"true"`
	Locality           bool `json:"locality,omitempty" toml:"locality,omitempty" yaml:"locality,omitempty" export:"true"`
	Province           bool `json:"province,omitempty" toml:"province,omitempty" yaml:"province,omitempty" export:"true"`
	SerialNumber       bool `json:"serialNumber,omitempty" toml:"serialNumber,omitempty" yaml:"serialNumber,omitempty" export:"true"`
	CommonName         bool `json:"commonName,omitempty" toml:"commonName,omitempty" yaml:"commonName,omitempty" export:"true"`
}

// TLSClientIssuerHeaders holds the TLS client cert issuer headers configuration.
type TLSClientIssuerHeaders struct {
	Country      bool `json:"country,omitempty" toml:"country,omitempty" yaml:"country,omitempty" export:"true"`
	Organization bool `json:"organization,omitempty" toml:"organization,omitempty" yaml:"organization,omitempty" export:"true"`
	CommonName   bool `json:"commonName,omitempty" toml:"commonName,omitempty" yaml:"commonName,omitempty" export:"true"`
}

// Retry holds the retry middleware configuration.
type Retry struct {
	Attempts        int            `json:"attempts,omitempty" toml:"attempts,omitempty" yaml:"attempts,omitempty" export:"true"`
	InitialInterval types.Duration `json:"initialInterval,omitempty" toml:"initialInterval,omitempty" yaml:"initialInterval,omitempty" export:"true"`
}

// InFlightReq holds the in-flight request middleware configuration.
type InFlightReq struct {
	Amount          int64            `json:"amount,omitempty" toml:"amount,omitempty" yaml:"amount,omitempty" export:"true"`
	SourceCriterion *SourceCriterion `json:"sourceCriterion,omitempty" toml:"sourceCriterion,omitempty" yaml:"sourceCriterion,omitempty" export:"true"`
}

// ClientTLS holds the client TLS configuration.
type ClientTLS struct {
	CA                 string `json:"ca,omitempty" toml:"ca,omitempty" yaml:"ca,omitempty"`
	Cert               string `json:"cert,omitempty" toml:"cert,omitempty" yaml:"cert,omitempty"`
	Key                string `json:"key,omitempty" toml:"key,omitempty" yaml:"key,omitempty"`
	InsecureSkipVerify bool   `json:"insecureSkipVerify,omitempty" toml:"insecureSkipVerify,omitempty" yaml:"insecureSkipVerify,omitempty" export:"true"`
}
