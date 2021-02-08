/*
 * Twilio - Fax
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ListFaxResponse struct for ListFaxResponse
type ListFaxResponse struct {
	Faxes []FaxV1Fax `json:"Faxes,omitempty"`
	Meta ListFaxResponseMeta `json:"Meta,omitempty"`
}
