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

// response message
type ConnectorLink struct {
	// full name of the created linked account link
	Name string `json:"name,omitempty"`
	// uri to share with the consumer to install the connector
	Uri string `json:"uri,omitempty"`
	// connector that the link belongs to, added as a separate field for convenience in not having to extract it from the ConnectorLink name
	Connector string `json:"connector,omitempty"`
	// consumer id
	ConsumerId       string            `json:"consumer_id"`
	ConsumerMetadata *ConsumerMetadata `json:"consumer_metadata,omitempty"`
	// Timestamps. See: https://aip.kolla.dev/kolla/9001 create time
	CreateTime time.Time `json:"create_time,omitempty"`
	// expiration time of the uri
	ExpireTime time.Time `json:"expire_time,omitempty"`
}
