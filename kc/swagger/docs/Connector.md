# Connector

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Resource name of the connector | [optional] [default to null]
**Uid** | **string** | The system ID of the resource | [optional] [default to null]
**Slug** | **string** | Short name that can be used to identify the connector in place of using the ID | [optional] [default to null]
**MarketplaceApp** | **string** | reference to the marketplace app that this connector is associated with | [optional] [default to null]
**AuthType** | **string** | auth type for the connector | [optional] [default to null]
**Oauth2Config** | [***ConnectorOauth2Config**](Connector_Oauth2Config.md) |  | [optional] [default to null]
**ApiKeyConfig** | [***ConnectorApiKeyConfig**](Connector_ApiKeyConfig.md) |  | [optional] [default to null]
**BasicAuthConfig** | [***ConnectorBasicAuthConfig**](Connector_BasicAuthConfig.md) |  | [optional] [default to null]
**AgentConfig** | [***ConnectorAgentConfig**](Connector_AgentConfig.md) |  | [optional] [default to null]
**CreateTime** | [**time.Time**](time.Time.md) | Timestamps. See: https://aip.kolla.dev/kolla/9001 create time | [optional] [default to null]
**UpdateTime** | [**time.Time**](time.Time.md) | time of last update | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

