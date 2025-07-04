/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Numbers
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

// Optional parameters for the method 'CreateEndUser'
type CreateEndUserParams struct {
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	//
	Type *string `json:"Type,omitempty"`
	// The set of parameters that are the attributes of the End User resource which are derived End User Types.
	Attributes *interface{} `json:"Attributes,omitempty"`
}

func (params *CreateEndUserParams) SetFriendlyName(FriendlyName string) *CreateEndUserParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateEndUserParams) SetType(Type string) *CreateEndUserParams {
	params.Type = &Type
	return params
}
func (params *CreateEndUserParams) SetAttributes(Attributes interface{}) *CreateEndUserParams {
	params.Attributes = &Attributes
	return params
}

// Create a new End User.
func (c *ApiService) CreateEndUser(params *CreateEndUserParams) (*NumbersV2EndUser, error) {
	path := "/v2/RegulatoryCompliance/EndUsers"

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
	if params != nil && params.Attributes != nil {
		v, err := json.Marshal(params.Attributes)

		if err != nil {
			return nil, err
		}

		data.Set("Attributes", string(v))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2EndUser{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific End User.
func (c *ApiService) DeleteEndUser(Sid string) error {
	path := "/v2/RegulatoryCompliance/EndUsers/{Sid}"
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

// Fetch specific End User Instance.
func (c *ApiService) FetchEndUser(Sid string) (*NumbersV2EndUser, error) {
	path := "/v2/RegulatoryCompliance/EndUsers/{Sid}"
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

	ps := &NumbersV2EndUser{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListEndUser'
type ListEndUserParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListEndUserParams) SetPageSize(PageSize int) *ListEndUserParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListEndUserParams) SetLimit(Limit int) *ListEndUserParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of EndUser records from the API. Request is executed immediately.
func (c *ApiService) PageEndUser(params *ListEndUserParams, pageToken, pageNumber string) (*ListEndUserResponse, error) {
	path := "/v2/RegulatoryCompliance/EndUsers"

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

	ps := &ListEndUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists EndUser records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListEndUser(params *ListEndUserParams) ([]NumbersV2EndUser, error) {
	response, errors := c.StreamEndUser(params)

	records := make([]NumbersV2EndUser, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams EndUser records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamEndUser(params *ListEndUserParams) (chan NumbersV2EndUser, chan error) {
	if params == nil {
		params = &ListEndUserParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan NumbersV2EndUser, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageEndUser(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamEndUser(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamEndUser(response *ListEndUserResponse, params *ListEndUserParams, recordChannel chan NumbersV2EndUser, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Results
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListEndUserResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListEndUserResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListEndUserResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListEndUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateEndUser'
type UpdateEndUserParams struct {
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The set of parameters that are the attributes of the End User resource which are derived End User Types.
	Attributes *interface{} `json:"Attributes,omitempty"`
}

func (params *UpdateEndUserParams) SetFriendlyName(FriendlyName string) *UpdateEndUserParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateEndUserParams) SetAttributes(Attributes interface{}) *UpdateEndUserParams {
	params.Attributes = &Attributes
	return params
}

// Update an existing End User.
func (c *ApiService) UpdateEndUser(Sid string, params *UpdateEndUserParams) (*NumbersV2EndUser, error) {
	path := "/v2/RegulatoryCompliance/EndUsers/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Attributes != nil {
		v, err := json.Marshal(params.Attributes)

		if err != nil {
			return nil, err
		}

		data.Set("Attributes", string(v))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2EndUser{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
