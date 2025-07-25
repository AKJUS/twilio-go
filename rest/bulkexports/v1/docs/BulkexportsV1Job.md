# BulkexportsV1Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | Pointer to **string** | The type of communication – Messages, Calls, Conferences, and Participants |
**FriendlyName** | Pointer to **string** | The friendly name specified when creating the job |
**Details** | Pointer to **[]interface{}** | The details of a job which is an object that contains an array of status grouped by `status` state.  Each `status` object has a `status` string, a count which is the number of days in that `status`, and list of days in that `status`. The day strings are in the format yyyy-MM-dd. As an example, a currently running job may have a status object for COMPLETED and a `status` object for SUBMITTED each with its own count and list of days. |
**StartDay** | Pointer to **string** | The start time for the export specified when creating the job |
**EndDay** | Pointer to **string** | The end time for the export specified when creating the job |
**JobSid** | Pointer to **string** | The job_sid returned when the export was created |
**WebhookUrl** | Pointer to **string** | The optional webhook url called on completion |
**WebhookMethod** | Pointer to **string** | This is the method used to call the webhook |
**Email** | Pointer to **string** | The optional email to send the completion notification to |
**Url** | Pointer to **string** |  |
**JobQueuePosition** | Pointer to **string** | This is the job position from the 1st in line. Your queue position will never increase. As jobs ahead of yours in the queue are processed, the queue position number will decrease |
**EstimatedCompletionTime** | Pointer to **string** | this is the time estimated until your job is complete. This is calculated each time you request the job list. The time is calculated based on the current rate of job completion (which may vary) and your job queue position |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


