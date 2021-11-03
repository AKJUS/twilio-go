/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010ValidationRequest struct for ApiV2010ValidationRequest
type ApiV2010ValidationRequest struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Call the resource is associated with
	CallSid *string `json:"call_sid,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The phone number to verify in E.164 format
	PhoneNumber *string `json:"phone_number,omitempty"`
	// The 6 digit validation code that someone must enter to validate the Caller ID  when `phone_number` is called
	ValidationCode *string `json:"validation_code,omitempty"`
}
