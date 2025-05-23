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

import (
	"encoding/json"
	"net/url"
)

// Optional parameters for the method 'CreateToken'
type CreateTokenParams struct {
	// Grant type is a credential representing resource owner's authorization which can be used by client to obtain access token.
	GrantType *string `json:"grant_type,omitempty"`
	// A 34 character string that uniquely identifies this OAuth App.
	ClientId *string `json:"client_id,omitempty"`
	// The credential for confidential OAuth App.
	ClientSecret *string `json:"client_secret,omitempty"`
	// JWT token related to the authorization code grant type.
	Code *string `json:"code,omitempty"`
	// The redirect uri
	RedirectUri *string `json:"redirect_uri,omitempty"`
	// The targeted audience uri
	Audience *string `json:"audience,omitempty"`
	// JWT token related to refresh access token.
	RefreshToken *string `json:"refresh_token,omitempty"`
	// The scope of token
	Scope *string `json:"scope,omitempty"`
}

func (params *CreateTokenParams) SetGrantType(GrantType string) *CreateTokenParams {
	params.GrantType = &GrantType
	return params
}
func (params *CreateTokenParams) SetClientId(ClientId string) *CreateTokenParams {
	params.ClientId = &ClientId
	return params
}
func (params *CreateTokenParams) SetClientSecret(ClientSecret string) *CreateTokenParams {
	params.ClientSecret = &ClientSecret
	return params
}
func (params *CreateTokenParams) SetCode(Code string) *CreateTokenParams {
	params.Code = &Code
	return params
}
func (params *CreateTokenParams) SetRedirectUri(RedirectUri string) *CreateTokenParams {
	params.RedirectUri = &RedirectUri
	return params
}
func (params *CreateTokenParams) SetAudience(Audience string) *CreateTokenParams {
	params.Audience = &Audience
	return params
}
func (params *CreateTokenParams) SetRefreshToken(RefreshToken string) *CreateTokenParams {
	params.RefreshToken = &RefreshToken
	return params
}
func (params *CreateTokenParams) SetScope(Scope string) *CreateTokenParams {
	params.Scope = &Scope
	return params
}

func (c *ApiService) CreateToken(params *CreateTokenParams) (*OauthV1Token, error) {
	path := "/v1/token"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.GrantType != nil {
		data.Set("grant_type", *params.GrantType)
	}
	if params != nil && params.ClientId != nil {
		data.Set("client_id", *params.ClientId)
	}
	if params != nil && params.ClientSecret != nil {
		data.Set("client_secret", *params.ClientSecret)
	}
	if params != nil && params.Code != nil {
		data.Set("code", *params.Code)
	}
	if params != nil && params.RedirectUri != nil {
		data.Set("redirect_uri", *params.RedirectUri)
	}
	if params != nil && params.Audience != nil {
		data.Set("audience", *params.Audience)
	}
	if params != nil && params.RefreshToken != nil {
		data.Set("refresh_token", *params.RefreshToken)
	}
	if params != nil && params.Scope != nil {
		data.Set("scope", *params.Scope)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &OauthV1Token{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
