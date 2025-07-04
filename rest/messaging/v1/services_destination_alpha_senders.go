/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Messaging
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

// Optional parameters for the method 'CreateDestinationAlphaSender'
type CreateDestinationAlphaSenderParams struct {
	// The Alphanumeric Sender ID string. Can be up to 11 characters long. Valid characters are A-Z, a-z, 0-9, space, hyphen `-`, plus `+`, underscore `_` and ampersand `&`. This value cannot contain only numbers.
	AlphaSender *string `json:"AlphaSender,omitempty"`
	// The Optional Two Character ISO Country Code the Alphanumeric Sender ID will be used for. If the IsoCountryCode is not provided, a default Alpha Sender will be created that can be used across all countries.
	IsoCountryCode *string `json:"IsoCountryCode,omitempty"`
}

func (params *CreateDestinationAlphaSenderParams) SetAlphaSender(AlphaSender string) *CreateDestinationAlphaSenderParams {
	params.AlphaSender = &AlphaSender
	return params
}
func (params *CreateDestinationAlphaSenderParams) SetIsoCountryCode(IsoCountryCode string) *CreateDestinationAlphaSenderParams {
	params.IsoCountryCode = &IsoCountryCode
	return params
}

//
func (c *ApiService) CreateDestinationAlphaSender(ServiceSid string, params *CreateDestinationAlphaSenderParams) (*MessagingV1DestinationAlphaSender, error) {
	path := "/v1/Services/{ServiceSid}/DestinationAlphaSenders"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.AlphaSender != nil {
		data.Set("AlphaSender", *params.AlphaSender)
	}
	if params != nil && params.IsoCountryCode != nil {
		data.Set("IsoCountryCode", *params.IsoCountryCode)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1DestinationAlphaSender{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//
func (c *ApiService) DeleteDestinationAlphaSender(ServiceSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/DestinationAlphaSenders/{Sid}"
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
func (c *ApiService) FetchDestinationAlphaSender(ServiceSid string, Sid string) (*MessagingV1DestinationAlphaSender, error) {
	path := "/v1/Services/{ServiceSid}/DestinationAlphaSenders/{Sid}"
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

	ps := &MessagingV1DestinationAlphaSender{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListDestinationAlphaSender'
type ListDestinationAlphaSenderParams struct {
	// Optional filter to return only alphanumeric sender IDs associated with the specified two-character ISO country code.
	IsoCountryCode *string `json:"IsoCountryCode,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListDestinationAlphaSenderParams) SetIsoCountryCode(IsoCountryCode string) *ListDestinationAlphaSenderParams {
	params.IsoCountryCode = &IsoCountryCode
	return params
}
func (params *ListDestinationAlphaSenderParams) SetPageSize(PageSize int) *ListDestinationAlphaSenderParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListDestinationAlphaSenderParams) SetLimit(Limit int) *ListDestinationAlphaSenderParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of DestinationAlphaSender records from the API. Request is executed immediately.
func (c *ApiService) PageDestinationAlphaSender(ServiceSid string, params *ListDestinationAlphaSenderParams, pageToken, pageNumber string) (*ListDestinationAlphaSenderResponse, error) {
	path := "/v1/Services/{ServiceSid}/DestinationAlphaSenders"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.IsoCountryCode != nil {
		data.Set("IsoCountryCode", *params.IsoCountryCode)
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

	ps := &ListDestinationAlphaSenderResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists DestinationAlphaSender records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListDestinationAlphaSender(ServiceSid string, params *ListDestinationAlphaSenderParams) ([]MessagingV1DestinationAlphaSender, error) {
	response, errors := c.StreamDestinationAlphaSender(ServiceSid, params)

	records := make([]MessagingV1DestinationAlphaSender, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams DestinationAlphaSender records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamDestinationAlphaSender(ServiceSid string, params *ListDestinationAlphaSenderParams) (chan MessagingV1DestinationAlphaSender, chan error) {
	if params == nil {
		params = &ListDestinationAlphaSenderParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan MessagingV1DestinationAlphaSender, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageDestinationAlphaSender(ServiceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamDestinationAlphaSender(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamDestinationAlphaSender(response *ListDestinationAlphaSenderResponse, params *ListDestinationAlphaSenderParams, recordChannel chan MessagingV1DestinationAlphaSender, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.AlphaSenders
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListDestinationAlphaSenderResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListDestinationAlphaSenderResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListDestinationAlphaSenderResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListDestinationAlphaSenderResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
