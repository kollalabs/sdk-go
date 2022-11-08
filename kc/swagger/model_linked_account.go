/*
 * Connect API
 *
 * Service for interacting with Connection Manager resources
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"time"
)

// LinkedAccount to a connector
type LinkedAccount struct {
	// Resource name of the connector
	Name string `json:"name,omitempty"`
	// The system ID of the resource
	Uid string `json:"uid,omitempty"`
	// The consumer that the LinkedAccount belongs to
	ConsumerId string `json:"consumer_id,omitempty"`
	// Install URL that the end user can use to install the connector The install_uri is only set if the user has not yet completed the install.
	InstallUri string `json:"install_uri,omitempty"`
	// current state of the LinkedAccount
	State string `json:"state,omitempty"`
	// A user displayable message about the linked_account state
	StateMessage string `json:"state_message,omitempty"`
	// Current credentials for the linked account, only needs to be supplied if migrating existing credentials into Kolla. Credentials can only be accessed through the Credentials endpoint
	Credentials map[string]string `json:"credentials,omitempty"`
	// Additional auth data received from the connected provider during consumer authentication
	AuthData *interface{} `json:"auth_data,omitempty"`
	// current state of the embedded credentials, can be used to determine if the user needs to re-auth before the credentials expire or need to be manually refreshed, typically a sub-state of the state field
	AuthState string `json:"auth_state,omitempty"`
	// A user displayable message about the auth state
	AuthStateDescription string `json:"auth_state_description,omitempty"`
	// Scopes associated with the linked account credentials, only needs to be provided if migrating credentials into Kolla
	LinkedAccountScopes []string `json:"linked_account_scopes,omitempty"`
	// Labels associated with the linked account
	Labels map[string]string `json:"labels,omitempty"`
	// Configuration values provided by the consumer
	ConsumerConfigValues map[string]string `json:"consumer_config_values,omitempty"`
	// Timestamps. See: https://aip.kolla.dev/kolla/9001 create time
	CreateTime time.Time `json:"create_time,omitempty"`
	// time of last update
	UpdateTime time.Time `json:"update_time,omitempty"`
}
