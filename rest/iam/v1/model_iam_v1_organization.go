/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Iam
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

// IamV1Organization struct for IamV1Organization
type IamV1Organization struct {
	// Unique Twilio organization sid
	OrganizationSid string `json:"organizationSid,omitempty"`
	// Organization friendly name
	FriendlyName string `json:"friendlyName,omitempty"`
	// User sign-up type for Organization domains
	DomainSignupType string `json:"domainSignupType,omitempty"`
	// Creation date of the organization
	DateCreated time.Time `json:"dateCreated,omitempty"`
}
