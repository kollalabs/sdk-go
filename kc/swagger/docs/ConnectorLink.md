# ConnectorLink

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | full name of the created linked account link | [optional] [default to null]
**Uri** | **string** | uri to share with the consumer to install the connector | [optional] [default to null]
**Connector** | **string** | connector that the link belongs to, added as a separate field for convenience in not having to extract it from the ConnectorLink name | [optional] [default to null]
**ConsumerId** | **string** | consumer id | [default to null]
**ConsumerMetadata** | [***ConsumerMetadata**](ConsumerMetadata.md) |  | [optional] [default to null]
**CreateTime** | [**time.Time**](time.Time.md) | Timestamps. See: https://aip.kolla.dev/kolla/9001 create time | [optional] [default to null]
**ExpireTime** | [**time.Time**](time.Time.md) | expiration time of the uri | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

