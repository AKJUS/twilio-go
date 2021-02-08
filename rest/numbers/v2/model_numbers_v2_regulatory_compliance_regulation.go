/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// NumbersV2RegulatoryComplianceRegulation struct for NumbersV2RegulatoryComplianceRegulation
type NumbersV2RegulatoryComplianceRegulation struct {
	EndUserType EndUserType `json:"EndUserType,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	IsoCountry string `json:"IsoCountry,omitempty"`
	NumberType string `json:"NumberType,omitempty"`
	Requirements map[string]interface{} `json:"Requirements,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Url string `json:"Url,omitempty"`
}
