/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ListFieldTypeResponse struct for ListFieldTypeResponse
type ListFieldTypeResponse struct {
	FieldTypes []AutopilotV1AssistantFieldType `json:"FieldTypes,omitempty"`
	Meta ListAssistantResponseMeta `json:"Meta,omitempty"`
}
