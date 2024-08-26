/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
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
)

// Optional parameters for the method 'CreateRealtimeTranscription'
type CreateRealtimeTranscriptionParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created this Transcription resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The user-specified name of this Transcription, if one was given when the Transcription was created. This may be used to stop the Transcription.
	Name *string `json:"Name,omitempty"`
	//
	Track *string `json:"Track,omitempty"`
	// Absolute URL of the status callback.
	StatusCallbackUrl *string `json:"StatusCallbackUrl,omitempty"`
	// The http method for the status_callback (one of GET, POST).
	StatusCallbackMethod *string `json:"StatusCallbackMethod,omitempty"`
	// Friendly name given to the Inbound Track
	InboundTrackLabel *string `json:"InboundTrackLabel,omitempty"`
	// Friendly name given to the Outbound Track
	OutboundTrackLabel *string `json:"OutboundTrackLabel,omitempty"`
	// Indicates if partial results are going to be sent to the customer
	PartialResults *bool `json:"PartialResults,omitempty"`
	// Language code used by the transcription engine, specified in [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt) format
	LanguageCode *string `json:"LanguageCode,omitempty"`
	// Definition of the transcription engine to be used, among those supported by Twilio
	TranscriptionEngine *string `json:"TranscriptionEngine,omitempty"`
	// indicates if the server will attempt to filter out profanities, replacing all but the initial character in each filtered word with asterisks
	ProfanityFilter *bool `json:"ProfanityFilter,omitempty"`
	// Recognition model used by the transcription engine, among those supported by the provider
	SpeechModel *string `json:"SpeechModel,omitempty"`
	// A Phrase contains words and phrase \\\"hints\\\" so that the speech recognition engine is more likely to recognize them.
	Hints *string `json:"Hints,omitempty"`
	// The provider will add punctuation to recognition result
	EnableAutomaticPunctuation *bool `json:"EnableAutomaticPunctuation,omitempty"`
}

func (params *CreateRealtimeTranscriptionParams) SetPathAccountSid(PathAccountSid string) *CreateRealtimeTranscriptionParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateRealtimeTranscriptionParams) SetName(Name string) *CreateRealtimeTranscriptionParams {
	params.Name = &Name
	return params
}
func (params *CreateRealtimeTranscriptionParams) SetTrack(Track string) *CreateRealtimeTranscriptionParams {
	params.Track = &Track
	return params
}
func (params *CreateRealtimeTranscriptionParams) SetStatusCallbackUrl(StatusCallbackUrl string) *CreateRealtimeTranscriptionParams {
	params.StatusCallbackUrl = &StatusCallbackUrl
	return params
}
func (params *CreateRealtimeTranscriptionParams) SetStatusCallbackMethod(StatusCallbackMethod string) *CreateRealtimeTranscriptionParams {
	params.StatusCallbackMethod = &StatusCallbackMethod
	return params
}
func (params *CreateRealtimeTranscriptionParams) SetInboundTrackLabel(InboundTrackLabel string) *CreateRealtimeTranscriptionParams {
	params.InboundTrackLabel = &InboundTrackLabel
	return params
}
func (params *CreateRealtimeTranscriptionParams) SetOutboundTrackLabel(OutboundTrackLabel string) *CreateRealtimeTranscriptionParams {
	params.OutboundTrackLabel = &OutboundTrackLabel
	return params
}
func (params *CreateRealtimeTranscriptionParams) SetPartialResults(PartialResults bool) *CreateRealtimeTranscriptionParams {
	params.PartialResults = &PartialResults
	return params
}
func (params *CreateRealtimeTranscriptionParams) SetLanguageCode(LanguageCode string) *CreateRealtimeTranscriptionParams {
	params.LanguageCode = &LanguageCode
	return params
}
func (params *CreateRealtimeTranscriptionParams) SetTranscriptionEngine(TranscriptionEngine string) *CreateRealtimeTranscriptionParams {
	params.TranscriptionEngine = &TranscriptionEngine
	return params
}
func (params *CreateRealtimeTranscriptionParams) SetProfanityFilter(ProfanityFilter bool) *CreateRealtimeTranscriptionParams {
	params.ProfanityFilter = &ProfanityFilter
	return params
}
func (params *CreateRealtimeTranscriptionParams) SetSpeechModel(SpeechModel string) *CreateRealtimeTranscriptionParams {
	params.SpeechModel = &SpeechModel
	return params
}
func (params *CreateRealtimeTranscriptionParams) SetHints(Hints string) *CreateRealtimeTranscriptionParams {
	params.Hints = &Hints
	return params
}
func (params *CreateRealtimeTranscriptionParams) SetEnableAutomaticPunctuation(EnableAutomaticPunctuation bool) *CreateRealtimeTranscriptionParams {
	params.EnableAutomaticPunctuation = &EnableAutomaticPunctuation
	return params
}

// Create a Transcription
func (c *ApiService) CreateRealtimeTranscription(CallSid string, params *CreateRealtimeTranscriptionParams) (*ApiV2010RealtimeTranscription, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Transcriptions.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Name != nil {
		data.Set("Name", *params.Name)
	}
	if params != nil && params.Track != nil {
		data.Set("Track", *params.Track)
	}
	if params != nil && params.StatusCallbackUrl != nil {
		data.Set("StatusCallbackUrl", *params.StatusCallbackUrl)
	}
	if params != nil && params.StatusCallbackMethod != nil {
		data.Set("StatusCallbackMethod", *params.StatusCallbackMethod)
	}
	if params != nil && params.InboundTrackLabel != nil {
		data.Set("InboundTrackLabel", *params.InboundTrackLabel)
	}
	if params != nil && params.OutboundTrackLabel != nil {
		data.Set("OutboundTrackLabel", *params.OutboundTrackLabel)
	}
	if params != nil && params.PartialResults != nil {
		data.Set("PartialResults", fmt.Sprint(*params.PartialResults))
	}
	if params != nil && params.LanguageCode != nil {
		data.Set("LanguageCode", *params.LanguageCode)
	}
	if params != nil && params.TranscriptionEngine != nil {
		data.Set("TranscriptionEngine", *params.TranscriptionEngine)
	}
	if params != nil && params.ProfanityFilter != nil {
		data.Set("ProfanityFilter", fmt.Sprint(*params.ProfanityFilter))
	}
	if params != nil && params.SpeechModel != nil {
		data.Set("SpeechModel", *params.SpeechModel)
	}
	if params != nil && params.Hints != nil {
		data.Set("Hints", *params.Hints)
	}
	if params != nil && params.EnableAutomaticPunctuation != nil {
		data.Set("EnableAutomaticPunctuation", fmt.Sprint(*params.EnableAutomaticPunctuation))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010RealtimeTranscription{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateRealtimeTranscription'
type UpdateRealtimeTranscriptionParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created this Transcription resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	//
	Status *string `json:"Status,omitempty"`
}

func (params *UpdateRealtimeTranscriptionParams) SetPathAccountSid(PathAccountSid string) *UpdateRealtimeTranscriptionParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateRealtimeTranscriptionParams) SetStatus(Status string) *UpdateRealtimeTranscriptionParams {
	params.Status = &Status
	return params
}

// Stop a Transcription using either the SID of the Transcription resource or the `name` used when creating the resource
func (c *ApiService) UpdateRealtimeTranscription(CallSid string, Sid string, params *UpdateRealtimeTranscriptionParams) (*ApiV2010RealtimeTranscription, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Transcriptions/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010RealtimeTranscription{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
