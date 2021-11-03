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

// Optional parameters for the method 'ListUsageRecordMonthly'
type ListUsageRecordMonthlyParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved.
	Category *string `json:"Category,omitempty"`
	// Only include usage that has occurred on or after this date. Specify the date in GMT and format as `YYYY-MM-DD`. You can also specify offsets from the current date, such as: `-30days`, which will set the start date to be 30 days before the current date.
	StartDate *string `json:"StartDate,omitempty"`
	// Only include usage that occurred on or before this date. Specify the date in GMT and format as `YYYY-MM-DD`.  You can also specify offsets from the current date, such as: `+30days`, which will set the end date to 30 days from the current date.
	EndDate *string `json:"EndDate,omitempty"`
	// Whether to include usage from the master account and all its subaccounts. Can be: `true` (the default) to include usage from the master account and all subaccounts or `false` to retrieve usage from only the specified account.
	IncludeSubaccounts *bool `json:"IncludeSubaccounts,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListUsageRecordMonthlyParams) SetPathAccountSid(PathAccountSid string) *ListUsageRecordMonthlyParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListUsageRecordMonthlyParams) SetCategory(Category string) *ListUsageRecordMonthlyParams {
	params.Category = &Category
	return params
}
func (params *ListUsageRecordMonthlyParams) SetStartDate(StartDate string) *ListUsageRecordMonthlyParams {
	params.StartDate = &StartDate
	return params
}
func (params *ListUsageRecordMonthlyParams) SetEndDate(EndDate string) *ListUsageRecordMonthlyParams {
	params.EndDate = &EndDate
	return params
}
func (params *ListUsageRecordMonthlyParams) SetIncludeSubaccounts(IncludeSubaccounts bool) *ListUsageRecordMonthlyParams {
	params.IncludeSubaccounts = &IncludeSubaccounts
	return params
}
func (params *ListUsageRecordMonthlyParams) SetPageSize(PageSize int) *ListUsageRecordMonthlyParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListUsageRecordMonthlyParams) SetLimit(Limit int) *ListUsageRecordMonthlyParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of UsageRecordMonthly records from the API. Request is executed immediately.
func (c *ApiService) PageUsageRecordMonthly(params *ListUsageRecordMonthlyParams, pageToken, pageNumber string) (*ListUsageRecordMonthlyResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Usage/Records/Monthly.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Category != nil {
		data.Set("Category", *params.Category)
	}
	if params != nil && params.StartDate != nil {
		data.Set("StartDate", fmt.Sprint(*params.StartDate))
	}
	if params != nil && params.EndDate != nil {
		data.Set("EndDate", fmt.Sprint(*params.EndDate))
	}
	if params != nil && params.IncludeSubaccounts != nil {
		data.Set("IncludeSubaccounts", fmt.Sprint(*params.IncludeSubaccounts))
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

	ps := &ListUsageRecordMonthlyResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists UsageRecordMonthly records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListUsageRecordMonthly(params *ListUsageRecordMonthlyParams) ([]ApiV2010UsageRecordMonthly, error) {
	if params == nil {
		params = &ListUsageRecordMonthlyParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageUsageRecordMonthly(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ApiV2010UsageRecordMonthly

	for response != nil {
		records = append(records, response.UsageRecords...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListUsageRecordMonthlyResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListUsageRecordMonthlyResponse)
	}

	return records, err
}

// Streams UsageRecordMonthly records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamUsageRecordMonthly(params *ListUsageRecordMonthlyParams) (chan ApiV2010UsageRecordMonthly, error) {
	if params == nil {
		params = &ListUsageRecordMonthlyParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageUsageRecordMonthly(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010UsageRecordMonthly, 1)

	go func() {
		for response != nil {
			for item := range response.UsageRecords {
				channel <- response.UsageRecords[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListUsageRecordMonthlyResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListUsageRecordMonthlyResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListUsageRecordMonthlyResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListUsageRecordMonthlyResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
