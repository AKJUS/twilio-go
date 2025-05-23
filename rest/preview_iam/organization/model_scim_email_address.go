/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Organization Public API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// ScimEmailAddress Email address list of the user. Primary email must be defined if there are more than 1 email. Primary email must match the username.
type ScimEmailAddress struct {
	// Indicates if this email address is the primary one
	Primary bool `json:"primary,omitempty"`
	// The actual email address value
	Value string `json:"value,omitempty"`
	// The type of email address (e.g., work, home, etc.)
	Type string `json:"type,omitempty"`
}
