/*
 * Twilio - Studio
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
// StudioV1FlowEngagement struct for StudioV1FlowEngagement
type StudioV1FlowEngagement struct {
	AccountSid string `json:"AccountSid,omitempty"`
	ContactChannelAddress string `json:"ContactChannelAddress,omitempty"`
	ContactSid string `json:"ContactSid,omitempty"`
	Context map[string]interface{} `json:"Context,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	FlowSid string `json:"FlowSid,omitempty"`
	Links map[string]interface{} `json:"Links,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Status Status `json:"Status,omitempty"`
	Url string `json:"Url,omitempty"`
}
