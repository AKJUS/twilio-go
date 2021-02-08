/*
 * Twilio - Voice
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// UpdateSourceIpMappingRequest struct for UpdateSourceIpMappingRequest
type UpdateSourceIpMappingRequest struct {
	// The SID of the SIP Domain that the IP Record should be mapped to.
	SipDomainSid string `json:"SipDomainSid"`
}
