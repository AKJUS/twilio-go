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
// ListQueryResponse struct for ListQueryResponse
type ListQueryResponse struct {
	Meta ListDayResponseMeta `json:"Meta,omitempty"`
	Queries []PreviewUnderstandAssistantQuery `json:"Queries,omitempty"`
}
