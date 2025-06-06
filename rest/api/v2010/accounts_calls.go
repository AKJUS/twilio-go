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
	"time"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateCall'
type CreateCallParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The phone number, SIP address, or client identifier to call.
	To *string `json:"To,omitempty"`
	// The phone number or client identifier to use as the caller id. If using a phone number, it must be a Twilio number or a Verified [outgoing caller id](https://www.twilio.com/docs/voice/api/outgoing-caller-ids) for your account. If the `to` parameter is a phone number, `From` must also be a phone number.
	From *string `json:"From,omitempty"`
	// The HTTP method we should use when calling the `url` parameter's value. Can be: `GET` or `POST` and the default is `POST`. If an `application_sid` parameter is present, this parameter is ignored.
	Method *string `json:"Method,omitempty"`
	// The URL that we call using the `fallback_method` if an error occurs when requesting or executing the TwiML at `url`. If an `application_sid` parameter is present, this parameter is ignored.
	FallbackUrl *string `json:"FallbackUrl,omitempty"`
	// The HTTP method that we should use to request the `fallback_url`. Can be: `GET` or `POST` and the default is `POST`. If an `application_sid` parameter is present, this parameter is ignored.
	FallbackMethod *string `json:"FallbackMethod,omitempty"`
	// The URL we should call using the `status_callback_method` to send status information to your application. If no `status_callback_event` is specified, we will send the `completed` status. If an `application_sid` parameter is present, this parameter is ignored. URLs must contain a valid hostname (underscores are not permitted).
	StatusCallback *string `json:"StatusCallback,omitempty"`
	// The call progress events that we will send to the `status_callback` URL. Can be: `initiated`, `ringing`, `answered`, and `completed`. If no event is specified, we send the `completed` status. If you want to receive multiple events, specify each one in a separate `status_callback_event` parameter. See the code sample for [monitoring call progress](https://www.twilio.com/docs/voice/api/call-resource?code-sample=code-create-a-call-resource-and-specify-a-statuscallbackevent&code-sdk-version=json). If an `application_sid` is present, this parameter is ignored.
	StatusCallbackEvent *[]string `json:"StatusCallbackEvent,omitempty"`
	// The HTTP method we should use when calling the `status_callback` URL. Can be: `GET` or `POST` and the default is `POST`. If an `application_sid` parameter is present, this parameter is ignored.
	StatusCallbackMethod *string `json:"StatusCallbackMethod,omitempty"`
	// The string of keys to dial after connecting to the number, with a maximum length of 32 digits. Valid digits in the string include any digit (`0`-`9`), '`A`', '`B`', '`C`', '`D`', '`#`', and '`*`'. You can also use '`w`' to insert a half-second pause and '`W`' to insert a one-second pause. For example, to pause for one second after connecting and then dial extension 1234 followed by the # key, set this parameter to `W1234#`. Be sure to URL-encode this string because the '`#`' character has special meaning in a URL. If both `SendDigits` and `MachineDetection` parameters are provided, then `MachineDetection` will be ignored.
	SendDigits *string `json:"SendDigits,omitempty"`
	// The integer number of seconds that we should allow the phone to ring before assuming there is no answer. The default is `60` seconds and the maximum is `600` seconds. For some call flows, we will add a 5-second buffer to the timeout value you provide. For this reason, a timeout value of 10 seconds could result in an actual timeout closer to 15 seconds. You can set this to a short time, such as `15` seconds, to hang up before reaching an answering machine or voicemail.
	Timeout *int `json:"Timeout,omitempty"`
	// Whether to record the call. Can be `true` to record the phone call, or `false` to not. The default is `false`. The `recording_url` is sent to the `status_callback` URL.
	Record *bool `json:"Record,omitempty"`
	// The number of channels in the final recording. Can be: `mono` or `dual`. The default is `mono`. `mono` records both legs of the call in a single channel of the recording file. `dual` records each leg to a separate channel of the recording file. The first channel of a dual-channel recording contains the parent call and the second channel contains the child call.
	RecordingChannels *string `json:"RecordingChannels,omitempty"`
	// The URL that we call when the recording is available to be accessed.
	RecordingStatusCallback *string `json:"RecordingStatusCallback,omitempty"`
	// The HTTP method we should use when calling the `recording_status_callback` URL. Can be: `GET` or `POST` and the default is `POST`.
	RecordingStatusCallbackMethod *string `json:"RecordingStatusCallbackMethod,omitempty"`
	// The username used to authenticate the caller making a SIP call.
	SipAuthUsername *string `json:"SipAuthUsername,omitempty"`
	// The password required to authenticate the user account specified in `sip_auth_username`.
	SipAuthPassword *string `json:"SipAuthPassword,omitempty"`
	// Whether to detect if a human, answering machine, or fax has picked up the call. Can be: `Enable` or `DetectMessageEnd`. Use `Enable` if you would like us to return `AnsweredBy` as soon as the called party is identified. Use `DetectMessageEnd`, if you would like to leave a message on an answering machine. If `send_digits` is provided, this parameter is ignored. For more information, see [Answering Machine Detection](https://www.twilio.com/docs/voice/answering-machine-detection).
	MachineDetection *string `json:"MachineDetection,omitempty"`
	// The number of seconds that we should attempt to detect an answering machine before timing out and sending a voice request with `AnsweredBy` of `unknown`. The default timeout is 30 seconds.
	MachineDetectionTimeout *int `json:"MachineDetectionTimeout,omitempty"`
	// The recording status events that will trigger calls to the URL specified in `recording_status_callback`. Can be: `in-progress`, `completed` and `absent`. Defaults to `completed`. Separate  multiple values with a space.
	RecordingStatusCallbackEvent *[]string `json:"RecordingStatusCallbackEvent,omitempty"`
	// Whether to trim any leading and trailing silence from the recording. Can be: `trim-silence` or `do-not-trim` and the default is `trim-silence`.
	Trim *string `json:"Trim,omitempty"`
	// The phone number, SIP address, or Client identifier that made this call. Phone numbers are in [E.164 format](https://wwnw.twilio.com/docs/glossary/what-e164) (e.g., +16175551212). SIP addresses are formatted as `name@company.com`.
	CallerId *string `json:"CallerId,omitempty"`
	// The number of milliseconds that is used as the measuring stick for the length of the speech activity, where durations lower than this value will be interpreted as a human and longer than this value as a machine. Possible Values: 1000-6000. Default: 2400.
	MachineDetectionSpeechThreshold *int `json:"MachineDetectionSpeechThreshold,omitempty"`
	// The number of milliseconds of silence after speech activity at which point the speech activity is considered complete. Possible Values: 500-5000. Default: 1200.
	MachineDetectionSpeechEndThreshold *int `json:"MachineDetectionSpeechEndThreshold,omitempty"`
	// The number of milliseconds of initial silence after which an `unknown` AnsweredBy result will be returned. Possible Values: 2000-10000. Default: 5000.
	MachineDetectionSilenceTimeout *int `json:"MachineDetectionSilenceTimeout,omitempty"`
	// Select whether to perform answering machine detection in the background. Default, blocks the execution of the call until Answering Machine Detection is completed. Can be: `true` or `false`.
	AsyncAmd *string `json:"AsyncAmd,omitempty"`
	// The URL that we should call using the `async_amd_status_callback_method` to notify customer application whether the call was answered by human, machine or fax.
	AsyncAmdStatusCallback *string `json:"AsyncAmdStatusCallback,omitempty"`
	// The HTTP method we should use when calling the `async_amd_status_callback` URL. Can be: `GET` or `POST` and the default is `POST`.
	AsyncAmdStatusCallbackMethod *string `json:"AsyncAmdStatusCallbackMethod,omitempty"`
	// The SID of a BYOC (Bring Your Own Carrier) trunk to route this call with. Note that `byoc` is only meaningful when `to` is a phone number; it will otherwise be ignored. (Beta)
	Byoc *string `json:"Byoc,omitempty"`
	// The Reason for the outgoing call. Use it to specify the purpose of the call that is presented on the called party's phone. (Branded Calls Beta)
	CallReason *string `json:"CallReason,omitempty"`
	// A token string needed to invoke a forwarded call. A call_token is generated when an incoming call is received on a Twilio number. Pass an incoming call's call_token value to a forwarded call via the call_token parameter when creating a new call. A forwarded call should bear the same CallerID of the original incoming call.
	CallToken *string `json:"CallToken,omitempty"`
	// The audio track to record for the call. Can be: `inbound`, `outbound` or `both`. The default is `both`. `inbound` records the audio that is received by Twilio. `outbound` records the audio that is generated from Twilio. `both` records the audio that is received and generated by Twilio.
	RecordingTrack *string `json:"RecordingTrack,omitempty"`
	// The maximum duration of the call in seconds. Constraints depend on account and configuration.
	TimeLimit *int `json:"TimeLimit,omitempty"`
	// The absolute URL that returns the TwiML instructions for the call. We will call this URL using the `method` when the call connects. For more information, see the [Url Parameter](https://www.twilio.com/docs/voice/make-calls#specify-a-url-parameter) section in [Making Calls](https://www.twilio.com/docs/voice/make-calls).
	Url *string `json:"Url,omitempty"`
	// TwiML instructions for the call Twilio will use without fetching Twiml from url parameter. If both `twiml` and `url` are provided then `twiml` parameter will be ignored. Max 4000 characters.
	Twiml *string `json:"Twiml,omitempty"`
	// The SID of the Application resource that will handle the call, if the call will be handled by an application.
	ApplicationSid *string `json:"ApplicationSid,omitempty"`
}

func (params *CreateCallParams) SetPathAccountSid(PathAccountSid string) *CreateCallParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateCallParams) SetTo(To string) *CreateCallParams {
	params.To = &To
	return params
}
func (params *CreateCallParams) SetFrom(From string) *CreateCallParams {
	params.From = &From
	return params
}
func (params *CreateCallParams) SetMethod(Method string) *CreateCallParams {
	params.Method = &Method
	return params
}
func (params *CreateCallParams) SetFallbackUrl(FallbackUrl string) *CreateCallParams {
	params.FallbackUrl = &FallbackUrl
	return params
}
func (params *CreateCallParams) SetFallbackMethod(FallbackMethod string) *CreateCallParams {
	params.FallbackMethod = &FallbackMethod
	return params
}
func (params *CreateCallParams) SetStatusCallback(StatusCallback string) *CreateCallParams {
	params.StatusCallback = &StatusCallback
	return params
}
func (params *CreateCallParams) SetStatusCallbackEvent(StatusCallbackEvent []string) *CreateCallParams {
	params.StatusCallbackEvent = &StatusCallbackEvent
	return params
}
func (params *CreateCallParams) SetStatusCallbackMethod(StatusCallbackMethod string) *CreateCallParams {
	params.StatusCallbackMethod = &StatusCallbackMethod
	return params
}
func (params *CreateCallParams) SetSendDigits(SendDigits string) *CreateCallParams {
	params.SendDigits = &SendDigits
	return params
}
func (params *CreateCallParams) SetTimeout(Timeout int) *CreateCallParams {
	params.Timeout = &Timeout
	return params
}
func (params *CreateCallParams) SetRecord(Record bool) *CreateCallParams {
	params.Record = &Record
	return params
}
func (params *CreateCallParams) SetRecordingChannels(RecordingChannels string) *CreateCallParams {
	params.RecordingChannels = &RecordingChannels
	return params
}
func (params *CreateCallParams) SetRecordingStatusCallback(RecordingStatusCallback string) *CreateCallParams {
	params.RecordingStatusCallback = &RecordingStatusCallback
	return params
}
func (params *CreateCallParams) SetRecordingStatusCallbackMethod(RecordingStatusCallbackMethod string) *CreateCallParams {
	params.RecordingStatusCallbackMethod = &RecordingStatusCallbackMethod
	return params
}
func (params *CreateCallParams) SetSipAuthUsername(SipAuthUsername string) *CreateCallParams {
	params.SipAuthUsername = &SipAuthUsername
	return params
}
func (params *CreateCallParams) SetSipAuthPassword(SipAuthPassword string) *CreateCallParams {
	params.SipAuthPassword = &SipAuthPassword
	return params
}
func (params *CreateCallParams) SetMachineDetection(MachineDetection string) *CreateCallParams {
	params.MachineDetection = &MachineDetection
	return params
}
func (params *CreateCallParams) SetMachineDetectionTimeout(MachineDetectionTimeout int) *CreateCallParams {
	params.MachineDetectionTimeout = &MachineDetectionTimeout
	return params
}
func (params *CreateCallParams) SetRecordingStatusCallbackEvent(RecordingStatusCallbackEvent []string) *CreateCallParams {
	params.RecordingStatusCallbackEvent = &RecordingStatusCallbackEvent
	return params
}
func (params *CreateCallParams) SetTrim(Trim string) *CreateCallParams {
	params.Trim = &Trim
	return params
}
func (params *CreateCallParams) SetCallerId(CallerId string) *CreateCallParams {
	params.CallerId = &CallerId
	return params
}
func (params *CreateCallParams) SetMachineDetectionSpeechThreshold(MachineDetectionSpeechThreshold int) *CreateCallParams {
	params.MachineDetectionSpeechThreshold = &MachineDetectionSpeechThreshold
	return params
}
func (params *CreateCallParams) SetMachineDetectionSpeechEndThreshold(MachineDetectionSpeechEndThreshold int) *CreateCallParams {
	params.MachineDetectionSpeechEndThreshold = &MachineDetectionSpeechEndThreshold
	return params
}
func (params *CreateCallParams) SetMachineDetectionSilenceTimeout(MachineDetectionSilenceTimeout int) *CreateCallParams {
	params.MachineDetectionSilenceTimeout = &MachineDetectionSilenceTimeout
	return params
}
func (params *CreateCallParams) SetAsyncAmd(AsyncAmd string) *CreateCallParams {
	params.AsyncAmd = &AsyncAmd
	return params
}
func (params *CreateCallParams) SetAsyncAmdStatusCallback(AsyncAmdStatusCallback string) *CreateCallParams {
	params.AsyncAmdStatusCallback = &AsyncAmdStatusCallback
	return params
}
func (params *CreateCallParams) SetAsyncAmdStatusCallbackMethod(AsyncAmdStatusCallbackMethod string) *CreateCallParams {
	params.AsyncAmdStatusCallbackMethod = &AsyncAmdStatusCallbackMethod
	return params
}
func (params *CreateCallParams) SetByoc(Byoc string) *CreateCallParams {
	params.Byoc = &Byoc
	return params
}
func (params *CreateCallParams) SetCallReason(CallReason string) *CreateCallParams {
	params.CallReason = &CallReason
	return params
}
func (params *CreateCallParams) SetCallToken(CallToken string) *CreateCallParams {
	params.CallToken = &CallToken
	return params
}
func (params *CreateCallParams) SetRecordingTrack(RecordingTrack string) *CreateCallParams {
	params.RecordingTrack = &RecordingTrack
	return params
}
func (params *CreateCallParams) SetTimeLimit(TimeLimit int) *CreateCallParams {
	params.TimeLimit = &TimeLimit
	return params
}
func (params *CreateCallParams) SetUrl(Url string) *CreateCallParams {
	params.Url = &Url
	return params
}
func (params *CreateCallParams) SetTwiml(Twiml string) *CreateCallParams {
	params.Twiml = &Twiml
	return params
}
func (params *CreateCallParams) SetApplicationSid(ApplicationSid string) *CreateCallParams {
	params.ApplicationSid = &ApplicationSid
	return params
}

// Create a new outgoing call to phones, SIP-enabled endpoints or Twilio Client connections
func (c *ApiService) CreateCall(params *CreateCallParams) (*ApiV2010Call, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Calls.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.To != nil {
		data.Set("To", *params.To)
	}
	if params != nil && params.From != nil {
		data.Set("From", *params.From)
	}
	if params != nil && params.Method != nil {
		data.Set("Method", *params.Method)
	}
	if params != nil && params.FallbackUrl != nil {
		data.Set("FallbackUrl", *params.FallbackUrl)
	}
	if params != nil && params.FallbackMethod != nil {
		data.Set("FallbackMethod", *params.FallbackMethod)
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	if params != nil && params.StatusCallbackEvent != nil {
		for _, item := range *params.StatusCallbackEvent {
			data.Add("StatusCallbackEvent", item)
		}
	}
	if params != nil && params.StatusCallbackMethod != nil {
		data.Set("StatusCallbackMethod", *params.StatusCallbackMethod)
	}
	if params != nil && params.SendDigits != nil {
		data.Set("SendDigits", *params.SendDigits)
	}
	if params != nil && params.Timeout != nil {
		data.Set("Timeout", fmt.Sprint(*params.Timeout))
	}
	if params != nil && params.Record != nil {
		data.Set("Record", fmt.Sprint(*params.Record))
	}
	if params != nil && params.RecordingChannels != nil {
		data.Set("RecordingChannels", *params.RecordingChannels)
	}
	if params != nil && params.RecordingStatusCallback != nil {
		data.Set("RecordingStatusCallback", *params.RecordingStatusCallback)
	}
	if params != nil && params.RecordingStatusCallbackMethod != nil {
		data.Set("RecordingStatusCallbackMethod", *params.RecordingStatusCallbackMethod)
	}
	if params != nil && params.SipAuthUsername != nil {
		data.Set("SipAuthUsername", *params.SipAuthUsername)
	}
	if params != nil && params.SipAuthPassword != nil {
		data.Set("SipAuthPassword", *params.SipAuthPassword)
	}
	if params != nil && params.MachineDetection != nil {
		data.Set("MachineDetection", *params.MachineDetection)
	}
	if params != nil && params.MachineDetectionTimeout != nil {
		data.Set("MachineDetectionTimeout", fmt.Sprint(*params.MachineDetectionTimeout))
	}
	if params != nil && params.RecordingStatusCallbackEvent != nil {
		for _, item := range *params.RecordingStatusCallbackEvent {
			data.Add("RecordingStatusCallbackEvent", item)
		}
	}
	if params != nil && params.Trim != nil {
		data.Set("Trim", *params.Trim)
	}
	if params != nil && params.CallerId != nil {
		data.Set("CallerId", *params.CallerId)
	}
	if params != nil && params.MachineDetectionSpeechThreshold != nil {
		data.Set("MachineDetectionSpeechThreshold", fmt.Sprint(*params.MachineDetectionSpeechThreshold))
	}
	if params != nil && params.MachineDetectionSpeechEndThreshold != nil {
		data.Set("MachineDetectionSpeechEndThreshold", fmt.Sprint(*params.MachineDetectionSpeechEndThreshold))
	}
	if params != nil && params.MachineDetectionSilenceTimeout != nil {
		data.Set("MachineDetectionSilenceTimeout", fmt.Sprint(*params.MachineDetectionSilenceTimeout))
	}
	if params != nil && params.AsyncAmd != nil {
		data.Set("AsyncAmd", *params.AsyncAmd)
	}
	if params != nil && params.AsyncAmdStatusCallback != nil {
		data.Set("AsyncAmdStatusCallback", *params.AsyncAmdStatusCallback)
	}
	if params != nil && params.AsyncAmdStatusCallbackMethod != nil {
		data.Set("AsyncAmdStatusCallbackMethod", *params.AsyncAmdStatusCallbackMethod)
	}
	if params != nil && params.Byoc != nil {
		data.Set("Byoc", *params.Byoc)
	}
	if params != nil && params.CallReason != nil {
		data.Set("CallReason", *params.CallReason)
	}
	if params != nil && params.CallToken != nil {
		data.Set("CallToken", *params.CallToken)
	}
	if params != nil && params.RecordingTrack != nil {
		data.Set("RecordingTrack", *params.RecordingTrack)
	}
	if params != nil && params.TimeLimit != nil {
		data.Set("TimeLimit", fmt.Sprint(*params.TimeLimit))
	}
	if params != nil && params.Url != nil {
		data.Set("Url", *params.Url)
	}
	if params != nil && params.Twiml != nil {
		data.Set("Twiml", *params.Twiml)
	}
	if params != nil && params.ApplicationSid != nil {
		data.Set("ApplicationSid", *params.ApplicationSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010Call{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteCall'
type DeleteCallParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call resource(s) to delete.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteCallParams) SetPathAccountSid(PathAccountSid string) *DeleteCallParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Delete a Call record from your account. Once the record is deleted, it will no longer appear in the API and Account Portal logs.
func (c *ApiService) DeleteCall(Sid string, params *DeleteCallParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/Calls/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
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

// Optional parameters for the method 'FetchCall'
type FetchCallParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call resource(s) to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchCallParams) SetPathAccountSid(PathAccountSid string) *FetchCallParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch the call specified by the provided Call SID
func (c *ApiService) FetchCall(Sid string, params *FetchCallParams) (*ApiV2010Call, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Calls/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
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

	ps := &ApiV2010Call{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListCall'
type ListCallParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call resource(s) to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// Only show calls made to this phone number, SIP address, Client identifier or SIM SID.
	To *string `json:"To,omitempty"`
	// Only include calls from this phone number, SIP address, Client identifier or SIM SID.
	From *string `json:"From,omitempty"`
	// Only include calls spawned by calls with this SID.
	ParentCallSid *string `json:"ParentCallSid,omitempty"`
	// The status of the calls to include. Can be: `queued`, `ringing`, `in-progress`, `canceled`, `completed`, `failed`, `busy`, or `no-answer`.
	Status *string `json:"Status,omitempty"`
	// Only include calls that started on this date. Specify a date as `YYYY-MM-DD` in UTC, for example: `2009-07-06`, to read only calls that started on this date. You can also specify an inequality, such as `StartTime<=YYYY-MM-DD`, to read calls that started on or before midnight of this date, and `StartTime>=YYYY-MM-DD` to read calls that started on or after midnight of this date.
	StartTime *time.Time `json:"StartTime,omitempty"`
	// Only include calls that started on this date. Specify a date as `YYYY-MM-DD` in UTC, for example: `2009-07-06`, to read only calls that started on this date. You can also specify an inequality, such as `StartTime<=YYYY-MM-DD`, to read calls that started on or before midnight of this date, and `StartTime>=YYYY-MM-DD` to read calls that started on or after midnight of this date.
	StartTimeBefore *time.Time `json:"StartTime&lt;,omitempty"`
	// Only include calls that started on this date. Specify a date as `YYYY-MM-DD` in UTC, for example: `2009-07-06`, to read only calls that started on this date. You can also specify an inequality, such as `StartTime<=YYYY-MM-DD`, to read calls that started on or before midnight of this date, and `StartTime>=YYYY-MM-DD` to read calls that started on or after midnight of this date.
	StartTimeAfter *time.Time `json:"StartTime&gt;,omitempty"`
	// Only include calls that ended on this date. Specify a date as `YYYY-MM-DD` in UTC, for example: `2009-07-06`, to read only calls that ended on this date. You can also specify an inequality, such as `EndTime<=YYYY-MM-DD`, to read calls that ended on or before midnight of this date, and `EndTime>=YYYY-MM-DD` to read calls that ended on or after midnight of this date.
	EndTime *time.Time `json:"EndTime,omitempty"`
	// Only include calls that ended on this date. Specify a date as `YYYY-MM-DD` in UTC, for example: `2009-07-06`, to read only calls that ended on this date. You can also specify an inequality, such as `EndTime<=YYYY-MM-DD`, to read calls that ended on or before midnight of this date, and `EndTime>=YYYY-MM-DD` to read calls that ended on or after midnight of this date.
	EndTimeBefore *time.Time `json:"EndTime&lt;,omitempty"`
	// Only include calls that ended on this date. Specify a date as `YYYY-MM-DD` in UTC, for example: `2009-07-06`, to read only calls that ended on this date. You can also specify an inequality, such as `EndTime<=YYYY-MM-DD`, to read calls that ended on or before midnight of this date, and `EndTime>=YYYY-MM-DD` to read calls that ended on or after midnight of this date.
	EndTimeAfter *time.Time `json:"EndTime&gt;,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListCallParams) SetPathAccountSid(PathAccountSid string) *ListCallParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListCallParams) SetTo(To string) *ListCallParams {
	params.To = &To
	return params
}
func (params *ListCallParams) SetFrom(From string) *ListCallParams {
	params.From = &From
	return params
}
func (params *ListCallParams) SetParentCallSid(ParentCallSid string) *ListCallParams {
	params.ParentCallSid = &ParentCallSid
	return params
}
func (params *ListCallParams) SetStatus(Status string) *ListCallParams {
	params.Status = &Status
	return params
}
func (params *ListCallParams) SetStartTime(StartTime time.Time) *ListCallParams {
	params.StartTime = &StartTime
	return params
}
func (params *ListCallParams) SetStartTimeBefore(StartTimeBefore time.Time) *ListCallParams {
	params.StartTimeBefore = &StartTimeBefore
	return params
}
func (params *ListCallParams) SetStartTimeAfter(StartTimeAfter time.Time) *ListCallParams {
	params.StartTimeAfter = &StartTimeAfter
	return params
}
func (params *ListCallParams) SetEndTime(EndTime time.Time) *ListCallParams {
	params.EndTime = &EndTime
	return params
}
func (params *ListCallParams) SetEndTimeBefore(EndTimeBefore time.Time) *ListCallParams {
	params.EndTimeBefore = &EndTimeBefore
	return params
}
func (params *ListCallParams) SetEndTimeAfter(EndTimeAfter time.Time) *ListCallParams {
	params.EndTimeAfter = &EndTimeAfter
	return params
}
func (params *ListCallParams) SetPageSize(PageSize int) *ListCallParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListCallParams) SetLimit(Limit int) *ListCallParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Call records from the API. Request is executed immediately.
func (c *ApiService) PageCall(params *ListCallParams, pageToken, pageNumber string) (*ListCallResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Calls.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.To != nil {
		data.Set("To", *params.To)
	}
	if params != nil && params.From != nil {
		data.Set("From", *params.From)
	}
	if params != nil && params.ParentCallSid != nil {
		data.Set("ParentCallSid", *params.ParentCallSid)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", fmt.Sprint(*params.Status))
	}
	if params != nil && params.StartTime != nil {
		data.Set("StartTime", fmt.Sprint((*params.StartTime).Format(time.RFC3339)))
	}
	if params != nil && params.StartTimeBefore != nil {
		data.Set("StartTime<", fmt.Sprint((*params.StartTimeBefore).Format(time.RFC3339)))
	}
	if params != nil && params.StartTimeAfter != nil {
		data.Set("StartTime>", fmt.Sprint((*params.StartTimeAfter).Format(time.RFC3339)))
	}
	if params != nil && params.EndTime != nil {
		data.Set("EndTime", fmt.Sprint((*params.EndTime).Format(time.RFC3339)))
	}
	if params != nil && params.EndTimeBefore != nil {
		data.Set("EndTime<", fmt.Sprint((*params.EndTimeBefore).Format(time.RFC3339)))
	}
	if params != nil && params.EndTimeAfter != nil {
		data.Set("EndTime>", fmt.Sprint((*params.EndTimeAfter).Format(time.RFC3339)))
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

	ps := &ListCallResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Call records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListCall(params *ListCallParams) ([]ApiV2010Call, error) {
	response, errors := c.StreamCall(params)

	records := make([]ApiV2010Call, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Call records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamCall(params *ListCallParams) (chan ApiV2010Call, chan error) {
	if params == nil {
		params = &ListCallParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ApiV2010Call, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageCall(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamCall(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamCall(response *ListCallResponse, params *ListCallParams, recordChannel chan ApiV2010Call, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Calls
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListCallResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListCallResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListCallResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCallResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateCall'
type UpdateCallParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call resource(s) to update.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The absolute URL that returns the TwiML instructions for the call. We will call this URL using the `method` when the call connects. For more information, see the [Url Parameter](https://www.twilio.com/docs/voice/make-calls#specify-a-url-parameter) section in [Making Calls](https://www.twilio.com/docs/voice/make-calls).
	Url *string `json:"Url,omitempty"`
	// The HTTP method we should use when calling the `url`. Can be: `GET` or `POST` and the default is `POST`. If an `application_sid` parameter is present, this parameter is ignored.
	Method *string `json:"Method,omitempty"`
	//
	Status *string `json:"Status,omitempty"`
	// The URL that we call using the `fallback_method` if an error occurs when requesting or executing the TwiML at `url`. If an `application_sid` parameter is present, this parameter is ignored.
	FallbackUrl *string `json:"FallbackUrl,omitempty"`
	// The HTTP method that we should use to request the `fallback_url`. Can be: `GET` or `POST` and the default is `POST`. If an `application_sid` parameter is present, this parameter is ignored.
	FallbackMethod *string `json:"FallbackMethod,omitempty"`
	// The URL we should call using the `status_callback_method` to send status information to your application. If no `status_callback_event` is specified, we will send the `completed` status. If an `application_sid` parameter is present, this parameter is ignored. URLs must contain a valid hostname (underscores are not permitted).
	StatusCallback *string `json:"StatusCallback,omitempty"`
	// The HTTP method we should use when requesting the `status_callback` URL. Can be: `GET` or `POST` and the default is `POST`. If an `application_sid` parameter is present, this parameter is ignored.
	StatusCallbackMethod *string `json:"StatusCallbackMethod,omitempty"`
	// TwiML instructions for the call Twilio will use without fetching Twiml from url. Twiml and url parameters are mutually exclusive
	Twiml *string `json:"Twiml,omitempty"`
	// The maximum duration of the call in seconds. Constraints depend on account and configuration.
	TimeLimit *int `json:"TimeLimit,omitempty"`
}

func (params *UpdateCallParams) SetPathAccountSid(PathAccountSid string) *UpdateCallParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateCallParams) SetUrl(Url string) *UpdateCallParams {
	params.Url = &Url
	return params
}
func (params *UpdateCallParams) SetMethod(Method string) *UpdateCallParams {
	params.Method = &Method
	return params
}
func (params *UpdateCallParams) SetStatus(Status string) *UpdateCallParams {
	params.Status = &Status
	return params
}
func (params *UpdateCallParams) SetFallbackUrl(FallbackUrl string) *UpdateCallParams {
	params.FallbackUrl = &FallbackUrl
	return params
}
func (params *UpdateCallParams) SetFallbackMethod(FallbackMethod string) *UpdateCallParams {
	params.FallbackMethod = &FallbackMethod
	return params
}
func (params *UpdateCallParams) SetStatusCallback(StatusCallback string) *UpdateCallParams {
	params.StatusCallback = &StatusCallback
	return params
}
func (params *UpdateCallParams) SetStatusCallbackMethod(StatusCallbackMethod string) *UpdateCallParams {
	params.StatusCallbackMethod = &StatusCallbackMethod
	return params
}
func (params *UpdateCallParams) SetTwiml(Twiml string) *UpdateCallParams {
	params.Twiml = &Twiml
	return params
}
func (params *UpdateCallParams) SetTimeLimit(TimeLimit int) *UpdateCallParams {
	params.TimeLimit = &TimeLimit
	return params
}

// Initiates a call redirect or terminates a call
func (c *ApiService) UpdateCall(Sid string, params *UpdateCallParams) (*ApiV2010Call, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Calls/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Url != nil {
		data.Set("Url", *params.Url)
	}
	if params != nil && params.Method != nil {
		data.Set("Method", *params.Method)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", fmt.Sprint(*params.Status))
	}
	if params != nil && params.FallbackUrl != nil {
		data.Set("FallbackUrl", *params.FallbackUrl)
	}
	if params != nil && params.FallbackMethod != nil {
		data.Set("FallbackMethod", *params.FallbackMethod)
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	if params != nil && params.StatusCallbackMethod != nil {
		data.Set("StatusCallbackMethod", *params.StatusCallbackMethod)
	}
	if params != nil && params.Twiml != nil {
		data.Set("Twiml", *params.Twiml)
	}
	if params != nil && params.TimeLimit != nil {
		data.Set("TimeLimit", fmt.Sprint(*params.TimeLimit))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010Call{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
