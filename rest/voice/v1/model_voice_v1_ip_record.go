/*
 * Twilio - Voice
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
// VoiceV1IpRecord struct for VoiceV1IpRecord
type VoiceV1IpRecord struct {
	AccountSid string `json:"AccountSid,omitempty"`
	CidrPrefixLength int32 `json:"CidrPrefixLength,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	IpAddress string `json:"IpAddress,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Url string `json:"Url,omitempty"`
}
