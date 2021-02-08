/*
 * Twilio - Pricing
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ListMessagingCountryResponse struct for ListMessagingCountryResponse
type ListMessagingCountryResponse struct {
	Countries []PricingV1MessagingMessagingCountry `json:"Countries,omitempty"`
	Meta ListMessagingCountryResponseMeta `json:"Meta,omitempty"`
}
