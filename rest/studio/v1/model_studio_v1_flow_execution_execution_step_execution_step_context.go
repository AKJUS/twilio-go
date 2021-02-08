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
// StudioV1FlowExecutionExecutionStepExecutionStepContext struct for StudioV1FlowExecutionExecutionStepExecutionStepContext
type StudioV1FlowExecutionExecutionStepExecutionStepContext struct {
	AccountSid string `json:"AccountSid,omitempty"`
	Context map[string]interface{} `json:"Context,omitempty"`
	ExecutionSid string `json:"ExecutionSid,omitempty"`
	FlowSid string `json:"FlowSid,omitempty"`
	StepSid string `json:"StepSid,omitempty"`
	Url string `json:"Url,omitempty"`
}
