/*
 * Twilio - Serverless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CreateBuildRequest struct for CreateBuildRequest
type CreateBuildRequest struct {
	// The list of Asset Version resource SIDs to include in the Build.
	AssetVersions []string `json:"AssetVersions,omitempty"`
	// A list of objects that describe the Dependencies included in the Build. Each object contains the `name` and `version` of the dependency.
	Dependencies string `json:"Dependencies,omitempty"`
	// The list of the Function Version resource SIDs to include in the Build.
	FunctionVersions []string `json:"FunctionVersions,omitempty"`
}
