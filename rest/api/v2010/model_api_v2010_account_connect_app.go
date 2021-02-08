/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ApiV2010AccountConnectApp struct for ApiV2010AccountConnectApp
type ApiV2010AccountConnectApp struct {
	AccountSid string `json:"AccountSid,omitempty"`
	AuthorizeRedirectUrl string `json:"AuthorizeRedirectUrl,omitempty"`
	CompanyName string `json:"CompanyName,omitempty"`
	DeauthorizeCallbackMethod HttpMethod `json:"DeauthorizeCallbackMethod,omitempty"`
	DeauthorizeCallbackUrl string `json:"DeauthorizeCallbackUrl,omitempty"`
	Description string `json:"Description,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	HomepageUrl string `json:"HomepageUrl,omitempty"`
	Permissions []string `json:"Permissions,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Uri string `json:"Uri,omitempty"`
}
