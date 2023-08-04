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

// A representation of a Connector
type Connector struct {
	// Resource name of the connector
	Name string `json:"name,omitempty"`
	// The system ID of the resource
	Uid string `json:"uid,omitempty"`
	// Short name that can be used to identify the connector in place of using the ID
	Slug string `json:"slug,omitempty"`
	// reference to the marketplace app that this connector is associated with
	MarketplaceApp string `json:"marketplace_app,omitempty"`
	// auth type for the connector
	AuthType        string                    `json:"auth_type,omitempty"`
	Oauth2Config    *ConnectorOauth2Config    `json:"oauth2_config,omitempty"`
	ApiKeyConfig    *ConnectorApiKeyConfig    `json:"api_key_config,omitempty"`
	BasicAuthConfig *ConnectorBasicAuthConfig `json:"basic_auth_config,omitempty"`
	AgentConfig     *ConnectorAgentConfig     `json:"agent_config,omitempty"`
	// Timestamps. See: https://aip.kolla.dev/kolla/9001 create time
	CreateTime time.Time `json:"create_time,omitempty"`
	// time of last update
	UpdateTime time.Time `json:"update_time,omitempty"`
}
