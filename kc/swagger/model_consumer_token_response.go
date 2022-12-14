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

// Response for a consumer token
type ConsumerTokenResponse struct {
	// name of the token
	Name string `json:"name,omitempty"`
	// access token that the consumer can user
	Token string `json:"token,omitempty"`
	// expiration time of the consumer token
	ExpiryTime time.Time `json:"expiry_time,omitempty"`
}
