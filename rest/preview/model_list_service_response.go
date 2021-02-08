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
// ListServiceResponse struct for ListServiceResponse
type ListServiceResponse struct {
	Meta ListDayResponseMeta `json:"Meta,omitempty"`
	Services []PreviewSyncService `json:"Services,omitempty"`
}
