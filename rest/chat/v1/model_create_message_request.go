/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CreateMessageRequest struct for CreateMessageRequest
type CreateMessageRequest struct {
	// A valid JSON string that contains application-specific data.
	Attributes string `json:"Attributes,omitempty"`
	// The message to send to the channel. Can also be an empty string or `null`, which sets the value as an empty string. You can send structured data in the body by serializing it as a string.
	Body string `json:"Body"`
	// The [identity](https://www.twilio.com/docs/api/chat/guides/identity) of the new message's author. The default value is `system`.
	From string `json:"From,omitempty"`
}
