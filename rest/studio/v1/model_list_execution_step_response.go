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
// ListExecutionStepResponse struct for ListExecutionStepResponse
type ListExecutionStepResponse struct {
	Meta ListFlowResponseMeta `json:"Meta,omitempty"`
	Steps []StudioV1FlowExecutionExecutionStep `json:"Steps,omitempty"`
}
