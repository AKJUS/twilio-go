/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CreateFactorRequest struct for CreateFactorRequest
type CreateFactorRequest struct {
	// The algorithm used when `factor_type` is `push`. Algorithm supported: `ES256`
	BindingAlg string `json:"BindingAlg,omitempty"`
	// The Ecdsa public key in PKIX, ASN.1 DER format encoded in Base64
	BindingPublicKey string `json:"BindingPublicKey,omitempty"`
	// The ID that uniquely identifies your app in the Google or Apple store, such as `com.example.myapp`. Required when `factor_type` is `push`
	ConfigAppId string `json:"ConfigAppId,omitempty"`
	// The transport technology used to generate the Notification Token. Can be `apn` or `fcm`. Required when `factor_type` is `push`
	ConfigNotificationPlatform string `json:"ConfigNotificationPlatform,omitempty"`
	// For APN, the device token. For FCM the registration token. It used to send the push notifications. Required when `factor_type` is `push`
	ConfigNotificationToken string `json:"ConfigNotificationToken,omitempty"`
	// The Verify Push SDK version used to configure the factor
	ConfigSdkVersion string `json:"ConfigSdkVersion,omitempty"`
	// The Type of this Factor. Currently only `push` is supported
	FactorType string `json:"FactorType"`
	// The friendly name of this Factor
	FriendlyName string `json:"FriendlyName"`
}
