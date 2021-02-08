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
// ServerlessV1ServiceFunctionFunctionVersion struct for ServerlessV1ServiceFunctionFunctionVersion
type ServerlessV1ServiceFunctionFunctionVersion struct {
	AccountSid string `json:"AccountSid,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	FunctionSid string `json:"FunctionSid,omitempty"`
	Links map[string]interface{} `json:"Links,omitempty"`
	Path string `json:"Path,omitempty"`
	ServiceSid string `json:"ServiceSid,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Url string `json:"Url,omitempty"`
	Visibility Visibility `json:"Visibility,omitempty"`
}
