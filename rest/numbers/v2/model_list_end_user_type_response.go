/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListEndUserTypeResponse struct for ListEndUserTypeResponse
type ListEndUserTypeResponse struct {
	EndUserTypes []NumbersV2EndUserType `json:"end_user_types,omitempty"`
	Meta         ListBundleResponseMeta `json:"meta,omitempty"`
}
