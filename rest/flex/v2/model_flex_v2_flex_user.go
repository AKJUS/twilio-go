/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Flex
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

// FlexV2FlexUser struct for FlexV2FlexUser
type FlexV2FlexUser struct {
	// The unique SID of the account that created the resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The unique ID created by Twilio to identify a Flex instance.
	InstanceSid *string `json:"instance_sid,omitempty"`
	// The unique SID identifier of the Twilio Unified User.
	UserSid *string `json:"user_sid,omitempty"`
	// The unique SID identifier of the Flex User.
	FlexUserSid *string `json:"flex_user_sid,omitempty"`
	// The unique SID identifier of the worker.
	WorkerSid *string `json:"worker_sid,omitempty"`
	// The unique SID identifier of the workspace.
	WorkspaceSid *string `json:"workspace_sid,omitempty"`
	// The unique SID identifier of the Flex Team.
	FlexTeamSid *string `json:"flex_team_sid,omitempty"`
	// Username of the User.
	Username *string `json:"username,omitempty"`
	// Email of the User.
	Email *string `json:"email,omitempty"`
	// The locale preference of the user.
	Locale *string `json:"locale,omitempty"`
	// The roles of the user.
	Roles *[]string `json:"roles,omitempty"`
	// The date that this user was created, given in ISO 8601 format.
	CreatedDate *time.Time `json:"created_date,omitempty"`
	// The date that this user was updated, given in ISO 8601 format.
	UpdatedDate *time.Time `json:"updated_date,omitempty"`
	// The current version of the user.
	Version int     `json:"version,omitempty"`
	Url     *string `json:"url,omitempty"`
}
