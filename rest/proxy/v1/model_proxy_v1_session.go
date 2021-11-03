/*
 * Twilio - Proxy
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

// ProxyV1Session struct for ProxyV1Session
type ProxyV1Session struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The reason the Session ended
	ClosedReason *string `json:"closed_reason,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date when the Session ended
	DateEnded *time.Time `json:"date_ended,omitempty"`
	// The ISO 8601 date when the Session should expire
	DateExpiry *time.Time `json:"date_expiry,omitempty"`
	// The ISO 8601 date when the Session last had an interaction
	DateLastInteraction *time.Time `json:"date_last_interaction,omitempty"`
	// The ISO 8601 date when the Session started
	DateStarted *time.Time `json:"date_started,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The URLs of resources related to the Session
	Links *map[string]interface{} `json:"links,omitempty"`
	// The Mode of the Session
	Mode *string `json:"mode,omitempty"`
	// The SID of the resource's parent Service
	ServiceSid *string `json:"service_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The status of the Session
	Status *string `json:"status,omitempty"`
	// When the session will expire
	Ttl *int `json:"ttl,omitempty"`
	// An application-defined string that uniquely identifies the resource
	UniqueName *string `json:"unique_name,omitempty"`
	// The absolute URL of the Session resource
	Url *string `json:"url,omitempty"`
}
