# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectConsumerToken**](ConnectApi.md#ConnectConsumerToken) | **Post** /v1/consumers:consumerToken | Consumer Token
[**ConnectCreateConnector**](ConnectApi.md#ConnectCreateConnector) | **Post** /v1/connectors | Create connector
[**ConnectCreateLinkedAccount**](ConnectApi.md#ConnectCreateLinkedAccount) | **Post** /v1/connectors/{connector}/linkedaccounts | Create LinkedAccount
[**ConnectCredentials**](ConnectApi.md#ConnectCredentials) | **Post** /v1/connectors/{connector}/linkedaccounts/{linkedaccount}:credentials | LinkedAccount Credentials
[**ConnectDisableLinkedAccount**](ConnectApi.md#ConnectDisableLinkedAccount) | **Post** /v1/connectors/{connector}/linkedaccounts/{linkedaccount}:disable | DisableLinkedAccount
[**ConnectGetConnector**](ConnectApi.md#ConnectGetConnector) | **Get** /v1/connectors/{connector} | Get Connector
[**ConnectGetLinkedAccount**](ConnectApi.md#ConnectGetLinkedAccount) | **Get** /v1/connectors/{connector}/linkedaccounts/{linkedaccount} | Get LinkedAccount
[**ConnectListConnectors**](ConnectApi.md#ConnectListConnectors) | **Get** /v1/connectors | List Connectors
[**ConnectListLinkedAccountLogs**](ConnectApi.md#ConnectListLinkedAccountLogs) | **Get** /v1/connectors/{connector}/linkedaccounts/{linkedaccount}/logs | ListLinkedAccountLogs
[**ConnectListLinkedAccounts**](ConnectApi.md#ConnectListLinkedAccounts) | **Get** /v1/connectors/{connector}/linkedaccounts | List LinkedAccounts
[**ConnectUpdateConnector**](ConnectApi.md#ConnectUpdateConnector) | **Patch** /v1/connectors/{connector} | Update Connector

# **ConnectConsumerToken**
> ConsumerTokenResponse ConnectConsumerToken(ctx, body)
Consumer Token

Returns a short lived access token that a consumer can use to access the  API

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConsumerTokenRequest**](ConsumerTokenRequest.md)|  | 

### Return type

[**ConsumerTokenResponse**](ConsumerTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConnectCreateConnector**
> Connector ConnectCreateConnector(ctx, body)
Create connector

Creates a new connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Connector**](Connector.md)|  | 

### Return type

[**Connector**](Connector.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConnectCreateLinkedAccount**
> LinkedAccount ConnectCreateLinkedAccount(ctx, body, connector)
Create LinkedAccount

Creates a new LinkedAccount

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LinkedAccount**](LinkedAccount.md)|  | 
  **connector** | **string**| The connector id. | 

### Return type

[**LinkedAccount**](LinkedAccount.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConnectCredentials**
> CredentialsResponse ConnectCredentials(ctx, body, connector, linkedaccount)
LinkedAccount Credentials

Returns the current auth credentials and expiry time  for a given LinkedAccount, to use consumer_id in place of the linked_account_id,  the linked_account_id in the url path should be a `-`  and the consumer_id specified in the request body  {connector} can be either the connector id or slug  /v1/{linked_account=connectors/{connector}/linkedaccounts/-}:credentials

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CredentialsRequest**](CredentialsRequest.md)|  | 
  **connector** | **string**| The connector id. | 
  **linkedaccount** | **string**| The linkedaccount id. | 

### Return type

[**CredentialsResponse**](CredentialsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConnectDisableLinkedAccount**
> LinkedAccount ConnectDisableLinkedAccount(ctx, body, connector, linkedaccount)
DisableLinkedAccount

Disable a linked account (used when disconnecting by a consumer)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DisableLinkedAccountRequest**](DisableLinkedAccountRequest.md)|  | 
  **connector** | **string**| The connector id. | 
  **linkedaccount** | **string**| The linkedaccount id. | 

### Return type

[**LinkedAccount**](LinkedAccount.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConnectGetConnector**
> Connector ConnectGetConnector(ctx, connector)
Get Connector

This gets a connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connector** | **string**| The connector id. | 

### Return type

[**Connector**](Connector.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConnectGetLinkedAccount**
> LinkedAccount ConnectGetLinkedAccount(ctx, connector, linkedaccount)
Get LinkedAccount

This gets a LinkedAccount.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connector** | **string**| The connector id. | 
  **linkedaccount** | **string**| The linkedaccount id. | 

### Return type

[**LinkedAccount**](LinkedAccount.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConnectListConnectors**
> ListConnectorsResponse ConnectListConnectors(ctx, optional)
List Connectors

Lists connectors

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConnectApiConnectListConnectorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectApiConnectListConnectorsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The maximum number of apis to return. The service may return fewer than this value. If unspecified, at most 50 will be returned. The maximum value is 1000; values above 1000 will be coerced to 1000. | 
 **pageToken** | **optional.String**| A page token, received from a previous list call. Provide this to retrieve the subsequent page. When paginating, all other parameters provided to &#x60;ListConnectors&#x60; must match the call that provided the page token. | 
 **orderBy** | **optional.String**| Order By Parameter | 
 **filter** | **optional.String**| Filter parameter- See https://aip.kolla.dev/160 for more info | 

### Return type

[**ListConnectorsResponse**](ListConnectorsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConnectListLinkedAccountLogs**
> ListLinkedAccountLogsResponse ConnectListLinkedAccountLogs(ctx, connector, linkedaccount, optional)
ListLinkedAccountLogs

Retrieve the logs for a given LinkedAccount

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connector** | **string**| The connector id. | 
  **linkedaccount** | **string**| The linkedaccount id. | 
 **optional** | ***ConnectApiConnectListLinkedAccountLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectApiConnectListLinkedAccountLogsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| The maximum number of apis to return. The service may return fewer than this value. If unspecified, at most 50 will be returned. The maximum value is 1000; values above 1000 will be coerced to 1000. | 
 **pageToken** | **optional.String**| A page token, received from a previous list call. Provide this to retrieve the subsequent page. When paginating, all other parameters provided to &#x60;ListConnectors&#x60; must match the call that provided the page token. | 
 **filter** | **optional.String**| Filter parameter, no filters currently supported | 

### Return type

[**ListLinkedAccountLogsResponse**](ListLinkedAccountLogsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConnectListLinkedAccounts**
> ListLinkedAccountsResponse ConnectListLinkedAccounts(ctx, connector, optional)
List LinkedAccounts

This gets a list of linked accounts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connector** | **string**| The connector id. | 
 **optional** | ***ConnectApiConnectListLinkedAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectApiConnectListLinkedAccountsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| The maximum number of items to return. The service may return fewer than this value. If unspecified, at most 50 will be returned. The maximum value is 1000; values above 1000 will be coerced to 1000. | 
 **pageToken** | **optional.String**| A page token, received from a previous list call. Provide this to retrieve the subsequent page. When paginating, all other parameters provided to &#x60;ListConnectors&#x60; must match the call that provided the page token. | 
 **orderBy** | **optional.String**| Order By Parameter | 
 **filter** | **optional.String**| Filter parameter- See https://aip.kolla.dev/160 for more info | 

### Return type

[**ListLinkedAccountsResponse**](ListLinkedAccountsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConnectUpdateConnector**
> Connector ConnectUpdateConnector(ctx, body, connector, optional)
Update Connector

Update a connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Connector**](Connector.md)|  | 
  **connector** | **string**| The connector id. | 
 **optional** | ***ConnectApiConnectUpdateConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectApiConnectUpdateConnectorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateMask** | **optional.**| The list of fields to update. | 

### Return type

[**Connector**](Connector.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

