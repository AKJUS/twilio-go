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
// TaskrouterV1WorkspaceTaskQueue struct for TaskrouterV1WorkspaceTaskQueue
type TaskrouterV1WorkspaceTaskQueue struct {
	AccountSid string `json:"AccountSid,omitempty"`
	AssignmentActivityName string `json:"AssignmentActivityName,omitempty"`
	AssignmentActivitySid string `json:"AssignmentActivitySid,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	Links map[string]interface{} `json:"Links,omitempty"`
	MaxReservedWorkers int32 `json:"MaxReservedWorkers,omitempty"`
	ReservationActivityName string `json:"ReservationActivityName,omitempty"`
	ReservationActivitySid string `json:"ReservationActivitySid,omitempty"`
	Sid string `json:"Sid,omitempty"`
	TargetWorkers string `json:"TargetWorkers,omitempty"`
	TaskOrder TaskOrder `json:"TaskOrder,omitempty"`
	Url string `json:"Url,omitempty"`
	WorkspaceSid string `json:"WorkspaceSid,omitempty"`
}
