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

// FlexV1FlexTeamMembers struct for FlexV1FlexTeamMembers
type FlexV1FlexTeamMembers struct {
	FlexUserSid  *string `json:"flex_user_sid,omitempty"`
	FriendlyName *string `json:"friendly_name,omitempty"`
	Email        *string `json:"email,omitempty"`
	WorkerSid    *string `json:"worker_sid,omitempty"`
	TeamSid      *string `json:"team_sid,omitempty"`
	InstanceSid  *string `json:"instance_sid,omitempty"`
	AccountSid   *string `json:"account_sid,omitempty"`
}
