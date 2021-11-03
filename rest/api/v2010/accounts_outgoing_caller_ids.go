/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateValidationRequest'
type CreateValidationRequestParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for the new caller ID resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The number of seconds to delay before initiating the verification call. Can be an integer between `0` and `60`, inclusive. The default is `0`.
	CallDelay *int `json:"CallDelay,omitempty"`
	// The digits to dial after connecting the verification call.
	Extension *string `json:"Extension,omitempty"`
	// A descriptive string that you create to describe the new caller ID resource. It can be up to 64 characters long. The default value is a formatted version of the phone number.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The phone number to verify in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, which consists of a + followed by the country code and subscriber number.
	PhoneNumber *string `json:"PhoneNumber,omitempty"`
	// The URL we should call using the `status_callback_method` to send status information about the verification process to your application.
	StatusCallback *string `json:"StatusCallback,omitempty"`
	// The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST`, and the default is `POST`.
	StatusCallbackMethod *string `json:"StatusCallbackMethod,omitempty"`
}

func (params *CreateValidationRequestParams) SetPathAccountSid(PathAccountSid string) *CreateValidationRequestParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateValidationRequestParams) SetCallDelay(CallDelay int) *CreateValidationRequestParams {
	params.CallDelay = &CallDelay
	return params
}
func (params *CreateValidationRequestParams) SetExtension(Extension string) *CreateValidationRequestParams {
	params.Extension = &Extension
	return params
}
func (params *CreateValidationRequestParams) SetFriendlyName(FriendlyName string) *CreateValidationRequestParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateValidationRequestParams) SetPhoneNumber(PhoneNumber string) *CreateValidationRequestParams {
	params.PhoneNumber = &PhoneNumber
	return params
}
func (params *CreateValidationRequestParams) SetStatusCallback(StatusCallback string) *CreateValidationRequestParams {
	params.StatusCallback = &StatusCallback
	return params
}
func (params *CreateValidationRequestParams) SetStatusCallbackMethod(StatusCallbackMethod string) *CreateValidationRequestParams {
	params.StatusCallbackMethod = &StatusCallbackMethod
	return params
}

func (c *ApiService) CreateValidationRequest(params *CreateValidationRequestParams) (*ApiV2010ValidationRequest, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/OutgoingCallerIds.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CallDelay != nil {
		data.Set("CallDelay", fmt.Sprint(*params.CallDelay))
	}
	if params != nil && params.Extension != nil {
		data.Set("Extension", *params.Extension)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.PhoneNumber != nil {
		data.Set("PhoneNumber", *params.PhoneNumber)
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	if params != nil && params.StatusCallbackMethod != nil {
		data.Set("StatusCallbackMethod", *params.StatusCallbackMethod)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010ValidationRequest{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteOutgoingCallerId'
type DeleteOutgoingCallerIdParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OutgoingCallerId resources to delete.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteOutgoingCallerIdParams) SetPathAccountSid(PathAccountSid string) *DeleteOutgoingCallerIdParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Delete the caller-id specified from the account
func (c *ApiService) DeleteOutgoingCallerId(Sid string, params *DeleteOutgoingCallerIdParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/OutgoingCallerIds/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Optional parameters for the method 'FetchOutgoingCallerId'
type FetchOutgoingCallerIdParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OutgoingCallerId resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchOutgoingCallerIdParams) SetPathAccountSid(PathAccountSid string) *FetchOutgoingCallerIdParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch an outgoing-caller-id belonging to the account used to make the request
func (c *ApiService) FetchOutgoingCallerId(Sid string, params *FetchOutgoingCallerIdParams) (*ApiV2010OutgoingCallerId, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/OutgoingCallerIds/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010OutgoingCallerId{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListOutgoingCallerId'
type ListOutgoingCallerIdParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OutgoingCallerId resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The phone number of the OutgoingCallerId resources to read.
	PhoneNumber *string `json:"PhoneNumber,omitempty"`
	// The string that identifies the OutgoingCallerId resources to read.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListOutgoingCallerIdParams) SetPathAccountSid(PathAccountSid string) *ListOutgoingCallerIdParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListOutgoingCallerIdParams) SetPhoneNumber(PhoneNumber string) *ListOutgoingCallerIdParams {
	params.PhoneNumber = &PhoneNumber
	return params
}
func (params *ListOutgoingCallerIdParams) SetFriendlyName(FriendlyName string) *ListOutgoingCallerIdParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *ListOutgoingCallerIdParams) SetPageSize(PageSize int) *ListOutgoingCallerIdParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListOutgoingCallerIdParams) SetLimit(Limit int) *ListOutgoingCallerIdParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of OutgoingCallerId records from the API. Request is executed immediately.
func (c *ApiService) PageOutgoingCallerId(params *ListOutgoingCallerIdParams, pageToken, pageNumber string) (*ListOutgoingCallerIdResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/OutgoingCallerIds.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PhoneNumber != nil {
		data.Set("PhoneNumber", *params.PhoneNumber)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
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

	ps := &ListOutgoingCallerIdResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists OutgoingCallerId records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListOutgoingCallerId(params *ListOutgoingCallerIdParams) ([]ApiV2010OutgoingCallerId, error) {
	if params == nil {
		params = &ListOutgoingCallerIdParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageOutgoingCallerId(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ApiV2010OutgoingCallerId

	for response != nil {
		records = append(records, response.OutgoingCallerIds...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListOutgoingCallerIdResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListOutgoingCallerIdResponse)
	}

	return records, err
}

// Streams OutgoingCallerId records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamOutgoingCallerId(params *ListOutgoingCallerIdParams) (chan ApiV2010OutgoingCallerId, error) {
	if params == nil {
		params = &ListOutgoingCallerIdParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageOutgoingCallerId(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010OutgoingCallerId, 1)

	go func() {
		for response != nil {
			for item := range response.OutgoingCallerIds {
				channel <- response.OutgoingCallerIds[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListOutgoingCallerIdResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListOutgoingCallerIdResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListOutgoingCallerIdResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListOutgoingCallerIdResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateOutgoingCallerId'
type UpdateOutgoingCallerIdParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OutgoingCallerId resources to update.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdateOutgoingCallerIdParams) SetPathAccountSid(PathAccountSid string) *UpdateOutgoingCallerIdParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateOutgoingCallerIdParams) SetFriendlyName(FriendlyName string) *UpdateOutgoingCallerIdParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Updates the caller-id
func (c *ApiService) UpdateOutgoingCallerId(Sid string, params *UpdateOutgoingCallerIdParams) (*ApiV2010OutgoingCallerId, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/OutgoingCallerIds/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010OutgoingCallerId{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
