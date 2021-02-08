/*
 * Twilio - Chat
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
// ChatV2ServiceChannelChannelWebhook struct for ChatV2ServiceChannelChannelWebhook
type ChatV2ServiceChannelChannelWebhook struct {
	AccountSid string `json:"AccountSid,omitempty"`
	ChannelSid string `json:"ChannelSid,omitempty"`
	Configuration map[string]interface{} `json:"Configuration,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	ServiceSid string `json:"ServiceSid,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Type string `json:"Type,omitempty"`
	Url string `json:"Url,omitempty"`
}
