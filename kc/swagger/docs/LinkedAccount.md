# LinkedAccount

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Resource name of the connector | [optional] [default to null]
**Uid** | **string** | The system ID of the resource | [optional] [default to null]
**ConsumerId** | **string** | The consumer that the LinkedAccount belongs to | [optional] [default to null]
**InstallUri** | **string** | Install URL that the end user can use to install the connector The install_uri is only set if the user has not yet completed the install. | [optional] [default to null]
**State** | **string** | current state of the LinkedAccount | [optional] [default to null]
**StateMessage** | **string** | A user displayable message about the linked_account state | [optional] [default to null]
**Credentials** | **map[string]string** | Current credentials for the linked account, only needs to be supplied if migrating existing credentials into Kolla. Credentials can only be accessed through the Credentials endpoint | [optional] [default to null]
**AuthData** | [***interface{}**](interface{}.md) | Additional auth data received from the connected provider during consumer authentication | [optional] [default to null]
**AuthState** | **string** | current state of the embedded credentials, can be used to determine if the user needs to re-auth before the credentials expire or need to be manually refreshed, typically a sub-state of the state field | [optional] [default to null]
**AuthStateDescription** | **string** | A user displayable message about the auth state | [optional] [default to null]
**LinkedAccountScopes** | **[]string** | Scopes associated with the linked account credentials, only needs to be provided if migrating credentials into Kolla | [optional] [default to null]
**Labels** | **map[string]string** | Labels associated with the linked account | [optional] [default to null]
**ConsumerConfigValues** | **map[string]string** | Configuration values provided by the consumer | [optional] [default to null]
**CreateTime** | [**time.Time**](time.Time.md) | Timestamps. See: https://aip.kolla.dev/kolla/9001 create time | [optional] [default to null]
**UpdateTime** | [**time.Time**](time.Time.md) | time of last update | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

