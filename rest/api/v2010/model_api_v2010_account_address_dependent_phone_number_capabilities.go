/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// ApiV2010AccountAddressDependentPhoneNumberCapabilities The set of Boolean properties that indicates whether a phone number can receive calls or messages.  Capabilities are  `Voice`, `SMS`, and `MMS` and each capability can be: `true` or `false`.
type ApiV2010AccountAddressDependentPhoneNumberCapabilities struct {
	Mms   bool `json:"mms,omitempty"`
	Sms   bool `json:"sms,omitempty"`
	Voice bool `json:"voice,omitempty"`
	Fax   bool `json:"fax,omitempty"`
}
