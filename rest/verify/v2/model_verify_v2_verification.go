/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Verify
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// VerifyV2Verification struct for VerifyV2Verification
type VerifyV2Verification struct {
	// The unique string that we created to identify the Verification resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.
	ServiceSid *string `json:"service_sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Verification resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The phone number or [email](https://www.twilio.com/docs/verify/email) being verified. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164).
	To      *string `json:"to,omitempty"`
	Channel *string `json:"channel,omitempty"`
	// The status of the verification. Can be: `pending`, `approved`, `canceled`, `max_attempts_reached`, `deleted`, `failed` or `expired`.
	Status *string `json:"status,omitempty"`
	// Use \"status\" instead. Legacy property indicating whether the verification was successful.
	Valid *bool `json:"valid,omitempty"`
	// Information about the phone number being verified.
	Lookup *map[string]interface{} `json:"lookup,omitempty"`
	// The amount of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled.
	Amount *string `json:"amount,omitempty"`
	// The payee of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled.
	Payee *string `json:"payee,omitempty"`
	// An array of verification attempt objects containing the channel attempted and the channel-specific transaction SID.
	SendCodeAttempts *[]map[string]interface{} `json:"send_code_attempts,omitempty"`
	// The date and time in GMT when the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The set of fields used for a silent network auth (`sna`) verification. Contains a single field with the URL to be invoked to verify the phone number.
	Sna *map[string]interface{} `json:"sna,omitempty"`
	// The absolute URL of the Verification resource.
	Url *string `json:"url,omitempty"`
}
