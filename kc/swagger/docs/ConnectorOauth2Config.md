# ConnectorOauth2Config

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | oauth2 client id | [optional] [default to null]
**ClientSecret** | **string** | oauth2 client secret | [optional] [default to null]
**RedirectUri** | **string** | oauth2 redirect uri, this should be left blank to have Kolla Connect manage the auth code and resulting credentials | [optional] [default to null]
**AuthorizationUri** | **string** | oauth2 authorization uri | [optional] [default to null]
**TokenUri** | **string** | oauth2 token uri | [optional] [default to null]
**RevocationUri** | **string** | oauth2 token revocation uri | [optional] [default to null]
**Scopes** | **[]string** | oauth2 scopes | [optional] [default to null]
**ExtraAuthorizationUriParams** | **map[string]string** | Extra authorization_uri parameters | [optional] [default to null]
**ExtraTokenUriParams** | **map[string]string** | Extra token_uri parameters | [optional] [default to null]
**PkceMethod** | **string** | PKCE configuration, if set to S256 PKCE will be used during token exchange | [optional] [default to null]
**AuthStyle** | **string** | token exchange/refresh auth style, default is \&quot;params\&quot;, options are \&quot;params\&quot;, \&quot;header\&quot;, or \&quot;detect\&quot; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

