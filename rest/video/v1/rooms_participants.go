/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Video
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
	"time"

	"github.com/twilio/twilio-go/client"
)

//
func (c *ApiService) FetchRoomParticipant(RoomSid string, Sid string) (*VideoV1RoomParticipant, error) {
	path := "/v1/Rooms/{RoomSid}/Participants/{Sid}"
	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)
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

	ps := &VideoV1RoomParticipant{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListRoomParticipant'
type ListRoomParticipantParams struct {
	// Read only the participants with this status. Can be: `connected` or `disconnected`. For `in-progress` Rooms the default Status is `connected`, for `completed` Rooms only `disconnected` Participants are returned.
	Status *string `json:"Status,omitempty"`
	// Read only the Participants with this [User](https://www.twilio.com/docs/chat/rest/user-resource) `identity` value.
	Identity *string `json:"Identity,omitempty"`
	// Read only Participants that started after this date in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#UTC) format.
	DateCreatedAfter *time.Time `json:"DateCreatedAfter,omitempty"`
	// Read only Participants that started before this date in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#UTC) format.
	DateCreatedBefore *time.Time `json:"DateCreatedBefore,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListRoomParticipantParams) SetStatus(Status string) *ListRoomParticipantParams {
	params.Status = &Status
	return params
}
func (params *ListRoomParticipantParams) SetIdentity(Identity string) *ListRoomParticipantParams {
	params.Identity = &Identity
	return params
}
func (params *ListRoomParticipantParams) SetDateCreatedAfter(DateCreatedAfter time.Time) *ListRoomParticipantParams {
	params.DateCreatedAfter = &DateCreatedAfter
	return params
}
func (params *ListRoomParticipantParams) SetDateCreatedBefore(DateCreatedBefore time.Time) *ListRoomParticipantParams {
	params.DateCreatedBefore = &DateCreatedBefore
	return params
}
func (params *ListRoomParticipantParams) SetPageSize(PageSize int) *ListRoomParticipantParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListRoomParticipantParams) SetLimit(Limit int) *ListRoomParticipantParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of RoomParticipant records from the API. Request is executed immediately.
func (c *ApiService) PageRoomParticipant(RoomSid string, params *ListRoomParticipantParams, pageToken, pageNumber string) (*ListRoomParticipantResponse, error) {
	path := "/v1/Rooms/{RoomSid}/Participants"

	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Status != nil {
		data.Set("Status", fmt.Sprint(*params.Status))
	}
	if params != nil && params.Identity != nil {
		data.Set("Identity", *params.Identity)
	}
	if params != nil && params.DateCreatedAfter != nil {
		data.Set("DateCreatedAfter", fmt.Sprint((*params.DateCreatedAfter).Format(time.RFC3339)))
	}
	if params != nil && params.DateCreatedBefore != nil {
		data.Set("DateCreatedBefore", fmt.Sprint((*params.DateCreatedBefore).Format(time.RFC3339)))
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

	ps := &ListRoomParticipantResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists RoomParticipant records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListRoomParticipant(RoomSid string, params *ListRoomParticipantParams) ([]VideoV1RoomParticipant, error) {
	response, errors := c.StreamRoomParticipant(RoomSid, params)

	records := make([]VideoV1RoomParticipant, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams RoomParticipant records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamRoomParticipant(RoomSid string, params *ListRoomParticipantParams) (chan VideoV1RoomParticipant, chan error) {
	if params == nil {
		params = &ListRoomParticipantParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan VideoV1RoomParticipant, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageRoomParticipant(RoomSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamRoomParticipant(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamRoomParticipant(response *ListRoomParticipantResponse, params *ListRoomParticipantParams, recordChannel chan VideoV1RoomParticipant, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Participants
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListRoomParticipantResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListRoomParticipantResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListRoomParticipantResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListRoomParticipantResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateRoomParticipant'
type UpdateRoomParticipantParams struct {
	//
	Status *string `json:"Status,omitempty"`
}

func (params *UpdateRoomParticipantParams) SetStatus(Status string) *UpdateRoomParticipantParams {
	params.Status = &Status
	return params
}

//
func (c *ApiService) UpdateRoomParticipant(RoomSid string, Sid string, params *UpdateRoomParticipantParams) (*VideoV1RoomParticipant, error) {
	path := "/v1/Rooms/{RoomSid}/Participants/{Sid}"
	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Status != nil {
		data.Set("Status", fmt.Sprint(*params.Status))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VideoV1RoomParticipant{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
