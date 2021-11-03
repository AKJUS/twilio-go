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

import (
	"time"
)

// TaskrouterV1TaskChannel struct for TaskrouterV1TaskChannel
type TaskrouterV1TaskChannel struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// Whether the Task Channel will prioritize Workers that have been idle
	ChannelOptimizedRouting *bool `json:"channel_optimized_routing,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The URLs of related resources
	Links *map[string]interface{} `json:"links,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// An application-defined string that uniquely identifies the Task Channel
	UniqueName *string `json:"unique_name,omitempty"`
	// The absolute URL of the Task Channel resource
	Url *string `json:"url,omitempty"`
	// The SID of the Workspace that contains the Task Channel
	WorkspaceSid *string `json:"workspace_sid,omitempty"`
}
