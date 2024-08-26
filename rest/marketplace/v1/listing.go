/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Marketplace
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"net/url"
	"strings"
)

// This endpoint returns the data of a given Listing. To find a Listing's SID, use the [Available Add-ons resource](/docs/marketplace/api/available-add-ons) or view its Listing details page in the Console by visiting the [Catalog](https://console.twilio.com/us1/develop/add-ons/catalog) or the [My Listings tab](https://console.twilio.com/us1/develop/add-ons/publish/my-listings) and selecting the Listing.
func (c *ApiService) FetchModuleDataManagement(Sid string) (*MarketplaceV1ModuleDataManagement, error) {
	path := "/v1/Listing/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MarketplaceV1ModuleDataManagement{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateModuleDataManagement'
type UpdateModuleDataManagementParams struct {
	//
	ModuleInfo *string `json:"ModuleInfo,omitempty"`
	//
	Description *string `json:"Description,omitempty"`
	//
	Documentation *string `json:"Documentation,omitempty"`
	//
	Policies *string `json:"Policies,omitempty"`
	//
	Support *string `json:"Support,omitempty"`
	//
	Configuration *string `json:"Configuration,omitempty"`
}

func (params *UpdateModuleDataManagementParams) SetModuleInfo(ModuleInfo string) *UpdateModuleDataManagementParams {
	params.ModuleInfo = &ModuleInfo
	return params
}
func (params *UpdateModuleDataManagementParams) SetDescription(Description string) *UpdateModuleDataManagementParams {
	params.Description = &Description
	return params
}
func (params *UpdateModuleDataManagementParams) SetDocumentation(Documentation string) *UpdateModuleDataManagementParams {
	params.Documentation = &Documentation
	return params
}
func (params *UpdateModuleDataManagementParams) SetPolicies(Policies string) *UpdateModuleDataManagementParams {
	params.Policies = &Policies
	return params
}
func (params *UpdateModuleDataManagementParams) SetSupport(Support string) *UpdateModuleDataManagementParams {
	params.Support = &Support
	return params
}
func (params *UpdateModuleDataManagementParams) SetConfiguration(Configuration string) *UpdateModuleDataManagementParams {
	params.Configuration = &Configuration
	return params
}

//
func (c *ApiService) UpdateModuleDataManagement(Sid string, params *UpdateModuleDataManagementParams) (*MarketplaceV1ModuleDataManagement, error) {
	path := "/v1/Listing/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.ModuleInfo != nil {
		data.Set("ModuleInfo", *params.ModuleInfo)
	}
	if params != nil && params.Description != nil {
		data.Set("Description", *params.Description)
	}
	if params != nil && params.Documentation != nil {
		data.Set("Documentation", *params.Documentation)
	}
	if params != nil && params.Policies != nil {
		data.Set("Policies", *params.Policies)
	}
	if params != nil && params.Support != nil {
		data.Set("Support", *params.Support)
	}
	if params != nil && params.Configuration != nil {
		data.Set("Configuration", *params.Configuration)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MarketplaceV1ModuleDataManagement{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
