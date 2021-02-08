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
// ApiV2010AccountUsageUsageTrigger struct for ApiV2010AccountUsageUsageTrigger
type ApiV2010AccountUsageUsageTrigger struct {
	AccountSid string `json:"AccountSid,omitempty"`
	ApiVersion string `json:"ApiVersion,omitempty"`
	CallbackMethod HttpMethod `json:"CallbackMethod,omitempty"`
	CallbackUrl string `json:"CallbackUrl,omitempty"`
	CurrentValue string `json:"CurrentValue,omitempty"`
	DateCreated string `json:"DateCreated,omitempty"`
	DateFired string `json:"DateFired,omitempty"`
	DateUpdated string `json:"DateUpdated,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	Recurring Recurring `json:"Recurring,omitempty"`
	Sid string `json:"Sid,omitempty"`
	TriggerBy TriggerField `json:"TriggerBy,omitempty"`
	TriggerValue string `json:"TriggerValue,omitempty"`
	Uri string `json:"Uri,omitempty"`
	UsageCategory UsageCategory `json:"UsageCategory,omitempty"`
	UsageRecordUri string `json:"UsageRecordUri,omitempty"`
}
