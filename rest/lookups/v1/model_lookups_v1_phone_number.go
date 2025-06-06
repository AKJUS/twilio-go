/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Lookups
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// LookupsV1PhoneNumber struct for LookupsV1PhoneNumber
type LookupsV1PhoneNumber struct {
	// The name of the phone number's owner. If `null`, that information was not available.
	CallerName *map[string]interface{} `json:"caller_name,omitempty"`
	// The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) for the phone number.
	CountryCode *string `json:"country_code,omitempty"`
	// The phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, which consists of a + followed by the country code and subscriber number.
	PhoneNumber *string `json:"phone_number,omitempty"`
	// The phone number, in national format.
	NationalFormat *string `json:"national_format,omitempty"`
	// The telecom company that provides the phone number.
	Carrier *map[string]interface{} `json:"carrier,omitempty"`
	// A JSON string with the results of the Add-ons you specified in the `add_ons` parameters. For the format of the object, see [Using Add-ons](https://www.twilio.com/docs/add-ons).
	AddOns *map[string]interface{} `json:"add_ons,omitempty"`
	// The absolute URL of the resource.
	Url *string `json:"url,omitempty"`
}
