# LinkedAccountLog

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the linked account log Format: connectors/{connector}/linkedaccounts/{linked_account}/logs/{linked_account_log} | [optional] [default to null]
**Uid** | **string** | The system ID of the resource | [optional] [default to null]
**RequestId** | **string** | request id of the request that generated this log | [optional] [default to null]
**Action** | **string** | action that triggered the workflow that generated the log message this should not be free-form text, it should be one of a small set of values that can be used to filter logs | [default to null]
**Code** | **string** | code associated with the log message | [optional] [default to null]
**Level** | **string** | The log level | [default to null]
**ErrorClass** | **string** | Error class associated with the log message | [optional] [default to null]
**Message** | **string** | The log message | [optional] [default to null]
**State** | **string** | capture the current linkedaccount state in the log messages? linked account state | [optional] [default to null]
**AuthState** | **string** | linked account auth state | [optional] [default to null]
**ConsumerId** | **string** | consumer id | [optional] [default to null]
**ConsumerTitle** | **string** | consumer title | [optional] [default to null]
**CreateTime** | [**time.Time**](time.Time.md) | The log timestamp | [optional] [default to null]
**UpdateTime** | [**time.Time**](time.Time.md) | Update timestamp, not used | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

