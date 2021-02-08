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
// CreateServiceRequest struct for CreateServiceRequest
type CreateServiceRequest struct {
	AclEnabled bool `json:"AclEnabled,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	ReachabilityWebhooksEnabled bool `json:"ReachabilityWebhooksEnabled,omitempty"`
	WebhookUrl string `json:"WebhookUrl,omitempty"`
}
