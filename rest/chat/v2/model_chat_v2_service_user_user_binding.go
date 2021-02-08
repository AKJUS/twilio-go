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
// ChatV2ServiceUserUserBinding struct for ChatV2ServiceUserUserBinding
type ChatV2ServiceUserUserBinding struct {
	AccountSid string `json:"AccountSid,omitempty"`
	BindingType BindingType `json:"BindingType,omitempty"`
	CredentialSid string `json:"CredentialSid,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	Endpoint string `json:"Endpoint,omitempty"`
	Identity string `json:"Identity,omitempty"`
	MessageTypes []string `json:"MessageTypes,omitempty"`
	ServiceSid string `json:"ServiceSid,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Url string `json:"Url,omitempty"`
	UserSid string `json:"UserSid,omitempty"`
}
