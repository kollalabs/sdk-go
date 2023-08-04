# CredentialsResponseCredentials

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | LinkedAccount access token, if available | [optional] [default to null]
**ExpiryTime** | [**time.Time**](time.Time.md) | timestamp for when a new token should be requested | [optional] [default to null]
**Secrets** | **map[string]string** | secret key value pairs, available keys depend on the connector type. Basic auth connectors will have a username and password. OAuth2 and most APIKey connectors will not have any secrets | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

