/*
 * Twilio - Supersim
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ListNetworkAccessProfileNetworkResponse struct for ListNetworkAccessProfileNetworkResponse
type ListNetworkAccessProfileNetworkResponse struct {
	Meta ListCommandResponseMeta `json:"Meta,omitempty"`
	Networks []SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork `json:"Networks,omitempty"`
}
