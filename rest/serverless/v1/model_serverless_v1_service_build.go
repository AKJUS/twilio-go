/*
 * Twilio - Serverless
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
// ServerlessV1ServiceBuild struct for ServerlessV1ServiceBuild
type ServerlessV1ServiceBuild struct {
	AccountSid string `json:"AccountSid,omitempty"`
	AssetVersions []map[string]interface{} `json:"AssetVersions,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	Dependencies []map[string]interface{} `json:"Dependencies,omitempty"`
	FunctionVersions []map[string]interface{} `json:"FunctionVersions,omitempty"`
	Links map[string]interface{} `json:"Links,omitempty"`
	ServiceSid string `json:"ServiceSid,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Status Status `json:"Status,omitempty"`
	Url string `json:"Url,omitempty"`
}
