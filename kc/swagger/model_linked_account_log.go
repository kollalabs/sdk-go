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

// LinkedAccountLogs
type LinkedAccountLog struct {
	// The name of the linked account log Format: connectors/{connector}/linkedaccounts/{linked_account}/logs/{linked_account_log}
	Name string `json:"name,omitempty"`
	// The system ID of the resource
	Uid string `json:"uid,omitempty"`
	// request id of the request that generated this log
	RequestId string `json:"request_id,omitempty"`
	// action that triggered the workflow that generated the log message
	Action string `json:"action,omitempty"`
	// code associated with the log message
	Code string `json:"code,omitempty"`
	// The log level
	Level string `json:"level,omitempty"`
	// The log message
	Message string `json:"message,omitempty"`
	// capture the current linkedaccount state in the log messages? linked account state
	State string `json:"state,omitempty"`
	// linked account auth state
	AuthState string `json:"auth_state,omitempty"`
	// The log timestamp
	CreateTime time.Time `json:"create_time,omitempty"`
	// Update timestamp, not used
	UpdateTime time.Time `json:"update_time,omitempty"`
}
