/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// TaskrouterV1WorkspaceWorkflowWorkflowStatistics struct for TaskrouterV1WorkspaceWorkflowWorkflowStatistics
type TaskrouterV1WorkspaceWorkflowWorkflowStatistics struct {
	AccountSid string `json:"AccountSid,omitempty"`
	Cumulative map[string]interface{} `json:"Cumulative,omitempty"`
	Realtime map[string]interface{} `json:"Realtime,omitempty"`
	Url string `json:"Url,omitempty"`
	WorkflowSid string `json:"WorkflowSid,omitempty"`
	WorkspaceSid string `json:"WorkspaceSid,omitempty"`
}
