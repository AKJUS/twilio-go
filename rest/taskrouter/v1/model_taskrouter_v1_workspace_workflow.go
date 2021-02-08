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
import (
	"time"
)
// TaskrouterV1WorkspaceWorkflow struct for TaskrouterV1WorkspaceWorkflow
type TaskrouterV1WorkspaceWorkflow struct {
	AccountSid string `json:"AccountSid,omitempty"`
	AssignmentCallbackUrl string `json:"AssignmentCallbackUrl,omitempty"`
	Configuration string `json:"Configuration,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	DocumentContentType string `json:"DocumentContentType,omitempty"`
	FallbackAssignmentCallbackUrl string `json:"FallbackAssignmentCallbackUrl,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	Links map[string]interface{} `json:"Links,omitempty"`
	Sid string `json:"Sid,omitempty"`
	TaskReservationTimeout int32 `json:"TaskReservationTimeout,omitempty"`
	Url string `json:"Url,omitempty"`
	WorkspaceSid string `json:"WorkspaceSid,omitempty"`
}
