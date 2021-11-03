/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// TaskrouterV1WorkerStatistics struct for TaskrouterV1WorkerStatistics
type TaskrouterV1WorkerStatistics struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// An object that contains the cumulative statistics for the Worker
	Cumulative *map[string]interface{} `json:"cumulative,omitempty"`
	// An object that contains the real-time statistics for the Worker
	Realtime *map[string]interface{} `json:"realtime,omitempty"`
	// The absolute URL of the Worker statistics resource
	Url *string `json:"url,omitempty"`
	// The SID of the Workspace that contains the Worker
	WorkspaceSid *string `json:"workspace_sid,omitempty"`
}
