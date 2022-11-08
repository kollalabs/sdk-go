/*
 * Connect API
 *
 * Service for interacting with Connection Manager resources
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// List linked accounts logs response
type ListLinkedAccountLogsResponse struct {
	// The available connectors.
	LinkedAccountLogs []LinkedAccountLog `json:"linked_account_logs,omitempty"`
	// A token, which can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`
}
