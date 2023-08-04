/*
 * Connect API
 *
 * Service for interacting with Connection Manager resources
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Oauth2 config message
type ConnectorOauth2Config struct {
	// oauth2 client id
	ClientId string `json:"client_id,omitempty"`
	// oauth2 client secret
	ClientSecret string `json:"client_secret,omitempty"`
	// oauth2 redirect uri, this should be left blank to have Kolla Connect manage the auth code and resulting credentials
	RedirectUri string `json:"redirect_uri,omitempty"`
	// oauth2 authorization uri
	AuthorizationUri string `json:"authorization_uri,omitempty"`
	// oauth2 token uri
	TokenUri string `json:"token_uri,omitempty"`
	// oauth2 token revocation uri
	RevocationUri string `json:"revocation_uri,omitempty"`
	// oauth2 scopes
	Scopes []string `json:"scopes,omitempty"`
	// Extra authorization_uri parameters
	ExtraAuthorizationUriParams map[string]string `json:"extra_authorization_uri_params,omitempty"`
	// Extra token_uri parameters
	ExtraTokenUriParams map[string]string `json:"extra_token_uri_params,omitempty"`
	// PKCE configuration, if set to S256 PKCE will be used during token exchange
	PkceMethod string `json:"pkce_method,omitempty"`
	// token exchange/refresh auth style, default is \"params\", options are \"params\", \"header\", or \"detect\"
	AuthStyle string `json:"auth_style,omitempty"`
}
