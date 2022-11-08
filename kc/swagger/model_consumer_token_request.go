/*
 * Connect API
 *
 * Service for interacting with Connection Manager resources
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Request for a consumer token
type ConsumerTokenRequest struct {
	// id that links back to the consumer, can be a user_id, tenant_id, machine_id, etc
	ConsumerId string `json:"consumer_id,omitempty"`
	Metadata *ConsumerTokenRequestConsumerMetadata `json:"metadata,omitempty"`
}
