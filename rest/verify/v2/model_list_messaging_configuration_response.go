/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ListMessagingConfigurationResponse struct for ListMessagingConfigurationResponse
type ListMessagingConfigurationResponse struct {
	MessagingConfigurations []VerifyV2ServiceMessagingConfiguration `json:"MessagingConfigurations,omitempty"`
	Meta ListServiceResponseMeta `json:"Meta,omitempty"`
}
