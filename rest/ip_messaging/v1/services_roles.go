/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Ip_messaging
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateRole'
type CreateRoleParams struct {
	//
	FriendlyName *string `json:"FriendlyName,omitempty"`
	//
	Type *string `json:"Type,omitempty"`
	//
	Permission *[]string `json:"Permission,omitempty"`
}

func (params *CreateRoleParams) SetFriendlyName(FriendlyName string) *CreateRoleParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateRoleParams) SetType(Type string) *CreateRoleParams {
	params.Type = &Type
	return params
}
func (params *CreateRoleParams) SetPermission(Permission []string) *CreateRoleParams {
	params.Permission = &Permission
	return params
}

//
func (c *ApiService) CreateRole(ServiceSid string, params *CreateRoleParams) (*IpMessagingV1Role, error) {
	path := "/v1/Services/{ServiceSid}/Roles"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Type != nil {
		data.Set("Type", fmt.Sprint(*params.Type))
	}
	if params != nil && params.Permission != nil {
		for _, item := range *params.Permission {
			data.Add("Permission", item)
		}
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV1Role{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//
func (c *ApiService) DeleteRole(ServiceSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/Roles/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

//
func (c *ApiService) FetchRole(ServiceSid string, Sid string) (*IpMessagingV1Role, error) {
	path := "/v1/Services/{ServiceSid}/Roles/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
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

	ps := &IpMessagingV1Role{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListRole'
type ListRoleParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListRoleParams) SetPageSize(PageSize int) *ListRoleParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListRoleParams) SetLimit(Limit int) *ListRoleParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Role records from the API. Request is executed immediately.
func (c *ApiService) PageRole(ServiceSid string, params *ListRoleParams, pageToken, pageNumber string) (*ListRoleResponse, error) {
	path := "/v1/Services/{ServiceSid}/Roles"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListRoleResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Role records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListRole(ServiceSid string, params *ListRoleParams) ([]IpMessagingV1Role, error) {
	response, errors := c.StreamRole(ServiceSid, params)

	records := make([]IpMessagingV1Role, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Role records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamRole(ServiceSid string, params *ListRoleParams) (chan IpMessagingV1Role, chan error) {
	if params == nil {
		params = &ListRoleParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan IpMessagingV1Role, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageRole(ServiceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamRole(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamRole(response *ListRoleResponse, params *ListRoleParams, recordChannel chan IpMessagingV1Role, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Roles
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListRoleResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListRoleResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListRoleResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListRoleResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateRole'
type UpdateRoleParams struct {
	//
	Permission *[]string `json:"Permission,omitempty"`
}

func (params *UpdateRoleParams) SetPermission(Permission []string) *UpdateRoleParams {
	params.Permission = &Permission
	return params
}

//
func (c *ApiService) UpdateRole(ServiceSid string, Sid string, params *UpdateRoleParams) (*IpMessagingV1Role, error) {
	path := "/v1/Services/{ServiceSid}/Roles/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Permission != nil {
		for _, item := range *params.Permission {
			data.Add("Permission", item)
		}
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV1Role{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
