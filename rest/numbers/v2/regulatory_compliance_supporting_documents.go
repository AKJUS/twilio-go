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

// Optional parameters for the method 'CreateSupportingDocument'
type CreateSupportingDocumentParams struct {
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The type of the Supporting Document.
	Type *string `json:"Type,omitempty"`
	// The set of parameters that are the attributes of the Supporting Documents resource which are derived Supporting Document Types.
	Attributes *map[string]interface{} `json:"Attributes,omitempty"`
}

func (params *CreateSupportingDocumentParams) SetFriendlyName(FriendlyName string) *CreateSupportingDocumentParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateSupportingDocumentParams) SetType(Type string) *CreateSupportingDocumentParams {
	params.Type = &Type
	return params
}
func (params *CreateSupportingDocumentParams) SetAttributes(Attributes map[string]interface{}) *CreateSupportingDocumentParams {
	params.Attributes = &Attributes
	return params
}

// Create a new Supporting Document.
func (c *ApiService) CreateSupportingDocument(params *CreateSupportingDocumentParams) (*NumbersV2SupportingDocument, error) {
	path := "/v2/RegulatoryCompliance/SupportingDocuments"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Type != nil {
		data.Set("Type", *params.Type)
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

	ps := &NumbersV2SupportingDocument{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific Supporting Document.
func (c *ApiService) DeleteSupportingDocument(Sid string) error {
	path := "/v2/RegulatoryCompliance/SupportingDocuments/{Sid}"
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

// Fetch specific Supporting Document Instance.
func (c *ApiService) FetchSupportingDocument(Sid string) (*NumbersV2SupportingDocument, error) {
	path := "/v2/RegulatoryCompliance/SupportingDocuments/{Sid}"
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

	ps := &NumbersV2SupportingDocument{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSupportingDocument'
type ListSupportingDocumentParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSupportingDocumentParams) SetPageSize(PageSize int) *ListSupportingDocumentParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSupportingDocumentParams) SetLimit(Limit int) *ListSupportingDocumentParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of SupportingDocument records from the API. Request is executed immediately.
func (c *ApiService) PageSupportingDocument(params *ListSupportingDocumentParams, pageToken, pageNumber string) (*ListSupportingDocumentResponse, error) {
	path := "/v2/RegulatoryCompliance/SupportingDocuments"

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

	ps := &ListSupportingDocumentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists SupportingDocument records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSupportingDocument(params *ListSupportingDocumentParams) ([]NumbersV2SupportingDocument, error) {
	response, errors := c.StreamSupportingDocument(params)

	records := make([]NumbersV2SupportingDocument, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams SupportingDocument records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSupportingDocument(params *ListSupportingDocumentParams) (chan NumbersV2SupportingDocument, chan error) {
	if params == nil {
		params = &ListSupportingDocumentParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan NumbersV2SupportingDocument, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageSupportingDocument(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamSupportingDocument(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamSupportingDocument(response *ListSupportingDocumentResponse, params *ListSupportingDocumentParams, recordChannel chan NumbersV2SupportingDocument, errorChannel chan error) {
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

		record, err := client.GetNext(c.baseURL, response, c.getNextListSupportingDocumentResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListSupportingDocumentResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListSupportingDocumentResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSupportingDocumentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateSupportingDocument'
type UpdateSupportingDocumentParams struct {
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The set of parameters that are the attributes of the Supporting Document resource which are derived Supporting Document Types.
	Attributes *map[string]interface{} `json:"Attributes,omitempty"`
}

func (params *UpdateSupportingDocumentParams) SetFriendlyName(FriendlyName string) *UpdateSupportingDocumentParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateSupportingDocumentParams) SetAttributes(Attributes map[string]interface{}) *UpdateSupportingDocumentParams {
	params.Attributes = &Attributes
	return params
}

// Update an existing Supporting Document.
func (c *ApiService) UpdateSupportingDocument(Sid string, params *UpdateSupportingDocumentParams) (*NumbersV2SupportingDocument, error) {
	path := "/v2/RegulatoryCompliance/SupportingDocuments/{Sid}"
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

	ps := &NumbersV2SupportingDocument{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
