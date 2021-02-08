/*
 * Twilio - Preview
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
// PreviewHostedNumbersAuthorizationDocument struct for PreviewHostedNumbersAuthorizationDocument
type PreviewHostedNumbersAuthorizationDocument struct {
	AddressSid string `json:"AddressSid,omitempty"`
	CcEmails []string `json:"CcEmails,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	Email string `json:"Email,omitempty"`
	Links map[string]interface{} `json:"Links,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Status Status `json:"Status,omitempty"`
	Url string `json:"Url,omitempty"`
}
