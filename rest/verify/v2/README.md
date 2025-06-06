# Go API client for openapi

This is the public Twilio REST API.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project from the OpenAPI specs located at [twilio/twilio-oai](https://github.com/twilio/twilio-oai/tree/main/spec).  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: com.twilio.oai.TwilioGoGenerator
For more information, please visit [https://support.twilio.com](https://support.twilio.com)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import "./openapi"
```

## Documentation for API Endpoints

All URIs are relative to *https://verify.twilio.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AttemptsApi* | [**FetchVerificationAttempt**](docs/AttemptsApi.md#fetchverificationattempt) | **Get** /v2/Attempts/{Sid} | Fetch a specific verification attempt.
*AttemptsApi* | [**ListVerificationAttempt**](docs/AttemptsApi.md#listverificationattempt) | **Get** /v2/Attempts | List all the verification attempts for a given Account.
*AttemptsSummaryApi* | [**FetchVerificationAttemptsSummary**](docs/AttemptsSummaryApi.md#fetchverificationattemptssummary) | **Get** /v2/Attempts/Summary | Get a summary of how many attempts were made and how many were converted.
*FormsApi* | [**FetchForm**](docs/FormsApi.md#fetchform) | **Get** /v2/Forms/{FormType} | Fetch the forms for a specific Form Type.
*SafeListNumbersApi* | [**CreateSafelist**](docs/SafeListNumbersApi.md#createsafelist) | **Post** /v2/SafeList/Numbers | Add a new phone number to SafeList.
*SafeListNumbersApi* | [**DeleteSafelist**](docs/SafeListNumbersApi.md#deletesafelist) | **Delete** /v2/SafeList/Numbers/{PhoneNumber} | Remove a phone number from SafeList.
*SafeListNumbersApi* | [**FetchSafelist**](docs/SafeListNumbersApi.md#fetchsafelist) | **Get** /v2/SafeList/Numbers/{PhoneNumber} | Check if a phone number exists in SafeList.
*ServicesApi* | [**CreateService**](docs/ServicesApi.md#createservice) | **Post** /v2/Services | Create a new Verification Service.
*ServicesApi* | [**DeleteService**](docs/ServicesApi.md#deleteservice) | **Delete** /v2/Services/{Sid} | Delete a specific Verification Service Instance.
*ServicesApi* | [**FetchService**](docs/ServicesApi.md#fetchservice) | **Get** /v2/Services/{Sid} | Fetch specific Verification Service Instance.
*ServicesApi* | [**ListService**](docs/ServicesApi.md#listservice) | **Get** /v2/Services | Retrieve a list of all Verification Services for an account.
*ServicesApi* | [**UpdateService**](docs/ServicesApi.md#updateservice) | **Post** /v2/Services/{Sid} | Update a specific Verification Service.
*ServicesAccessTokensApi* | [**CreateAccessToken**](docs/ServicesAccessTokensApi.md#createaccesstoken) | **Post** /v2/Services/{ServiceSid}/AccessTokens | Create a new enrollment Access Token for the Entity
*ServicesAccessTokensApi* | [**FetchAccessToken**](docs/ServicesAccessTokensApi.md#fetchaccesstoken) | **Get** /v2/Services/{ServiceSid}/AccessTokens/{Sid} | Fetch an Access Token for the Entity
*ServicesEntitiesApi* | [**CreateEntity**](docs/ServicesEntitiesApi.md#createentity) | **Post** /v2/Services/{ServiceSid}/Entities | Create a new Entity for the Service
*ServicesEntitiesApi* | [**DeleteEntity**](docs/ServicesEntitiesApi.md#deleteentity) | **Delete** /v2/Services/{ServiceSid}/Entities/{Identity} | Delete a specific Entity.
*ServicesEntitiesApi* | [**FetchEntity**](docs/ServicesEntitiesApi.md#fetchentity) | **Get** /v2/Services/{ServiceSid}/Entities/{Identity} | Fetch a specific Entity.
*ServicesEntitiesApi* | [**ListEntity**](docs/ServicesEntitiesApi.md#listentity) | **Get** /v2/Services/{ServiceSid}/Entities | Retrieve a list of all Entities for a Service.
*ServicesEntitiesChallengesApi* | [**CreateChallenge**](docs/ServicesEntitiesChallengesApi.md#createchallenge) | **Post** /v2/Services/{ServiceSid}/Entities/{Identity}/Challenges | Create a new Challenge for the Factor
*ServicesEntitiesChallengesApi* | [**FetchChallenge**](docs/ServicesEntitiesChallengesApi.md#fetchchallenge) | **Get** /v2/Services/{ServiceSid}/Entities/{Identity}/Challenges/{Sid} | Fetch a specific Challenge.
*ServicesEntitiesChallengesApi* | [**ListChallenge**](docs/ServicesEntitiesChallengesApi.md#listchallenge) | **Get** /v2/Services/{ServiceSid}/Entities/{Identity}/Challenges | Retrieve a list of all Challenges for a Factor.
*ServicesEntitiesChallengesApi* | [**UpdateChallenge**](docs/ServicesEntitiesChallengesApi.md#updatechallenge) | **Post** /v2/Services/{ServiceSid}/Entities/{Identity}/Challenges/{Sid} | Verify a specific Challenge.
*ServicesEntitiesChallengesNotificationsApi* | [**CreateNotification**](docs/ServicesEntitiesChallengesNotificationsApi.md#createnotification) | **Post** /v2/Services/{ServiceSid}/Entities/{Identity}/Challenges/{ChallengeSid}/Notifications | Create a new Notification for the corresponding Challenge
*ServicesEntitiesFactorsApi* | [**CreateNewFactor**](docs/ServicesEntitiesFactorsApi.md#createnewfactor) | **Post** /v2/Services/{ServiceSid}/Entities/{Identity}/Factors | Create a new Factor for the Entity
*ServicesEntitiesFactorsApi* | [**DeleteFactor**](docs/ServicesEntitiesFactorsApi.md#deletefactor) | **Delete** /v2/Services/{ServiceSid}/Entities/{Identity}/Factors/{Sid} | Delete a specific Factor.
*ServicesEntitiesFactorsApi* | [**FetchFactor**](docs/ServicesEntitiesFactorsApi.md#fetchfactor) | **Get** /v2/Services/{ServiceSid}/Entities/{Identity}/Factors/{Sid} | Fetch a specific Factor.
*ServicesEntitiesFactorsApi* | [**ListFactor**](docs/ServicesEntitiesFactorsApi.md#listfactor) | **Get** /v2/Services/{ServiceSid}/Entities/{Identity}/Factors | Retrieve a list of all Factors for an Entity.
*ServicesEntitiesFactorsApi* | [**UpdateFactor**](docs/ServicesEntitiesFactorsApi.md#updatefactor) | **Post** /v2/Services/{ServiceSid}/Entities/{Identity}/Factors/{Sid} | Update a specific Factor. This endpoint can be used to Verify a Factor if passed an &#x60;AuthPayload&#x60; param.
*ServicesMessagingConfigurationsApi* | [**CreateMessagingConfiguration**](docs/ServicesMessagingConfigurationsApi.md#createmessagingconfiguration) | **Post** /v2/Services/{ServiceSid}/MessagingConfigurations | Create a new MessagingConfiguration for a service.
*ServicesMessagingConfigurationsApi* | [**DeleteMessagingConfiguration**](docs/ServicesMessagingConfigurationsApi.md#deletemessagingconfiguration) | **Delete** /v2/Services/{ServiceSid}/MessagingConfigurations/{Country} | Delete a specific MessagingConfiguration.
*ServicesMessagingConfigurationsApi* | [**FetchMessagingConfiguration**](docs/ServicesMessagingConfigurationsApi.md#fetchmessagingconfiguration) | **Get** /v2/Services/{ServiceSid}/MessagingConfigurations/{Country} | Fetch a specific MessagingConfiguration.
*ServicesMessagingConfigurationsApi* | [**ListMessagingConfiguration**](docs/ServicesMessagingConfigurationsApi.md#listmessagingconfiguration) | **Get** /v2/Services/{ServiceSid}/MessagingConfigurations | Retrieve a list of all Messaging Configurations for a Service.
*ServicesMessagingConfigurationsApi* | [**UpdateMessagingConfiguration**](docs/ServicesMessagingConfigurationsApi.md#updatemessagingconfiguration) | **Post** /v2/Services/{ServiceSid}/MessagingConfigurations/{Country} | Update a specific MessagingConfiguration
*ServicesRateLimitsApi* | [**CreateRateLimit**](docs/ServicesRateLimitsApi.md#createratelimit) | **Post** /v2/Services/{ServiceSid}/RateLimits | Create a new Rate Limit for a Service
*ServicesRateLimitsApi* | [**DeleteRateLimit**](docs/ServicesRateLimitsApi.md#deleteratelimit) | **Delete** /v2/Services/{ServiceSid}/RateLimits/{Sid} | Delete a specific Rate Limit.
*ServicesRateLimitsApi* | [**FetchRateLimit**](docs/ServicesRateLimitsApi.md#fetchratelimit) | **Get** /v2/Services/{ServiceSid}/RateLimits/{Sid} | Fetch a specific Rate Limit.
*ServicesRateLimitsApi* | [**ListRateLimit**](docs/ServicesRateLimitsApi.md#listratelimit) | **Get** /v2/Services/{ServiceSid}/RateLimits | Retrieve a list of all Rate Limits for a service.
*ServicesRateLimitsApi* | [**UpdateRateLimit**](docs/ServicesRateLimitsApi.md#updateratelimit) | **Post** /v2/Services/{ServiceSid}/RateLimits/{Sid} | Update a specific Rate Limit.
*ServicesRateLimitsBucketsApi* | [**CreateBucket**](docs/ServicesRateLimitsBucketsApi.md#createbucket) | **Post** /v2/Services/{ServiceSid}/RateLimits/{RateLimitSid}/Buckets | Create a new Bucket for a Rate Limit
*ServicesRateLimitsBucketsApi* | [**DeleteBucket**](docs/ServicesRateLimitsBucketsApi.md#deletebucket) | **Delete** /v2/Services/{ServiceSid}/RateLimits/{RateLimitSid}/Buckets/{Sid} | Delete a specific Bucket.
*ServicesRateLimitsBucketsApi* | [**FetchBucket**](docs/ServicesRateLimitsBucketsApi.md#fetchbucket) | **Get** /v2/Services/{ServiceSid}/RateLimits/{RateLimitSid}/Buckets/{Sid} | Fetch a specific Bucket.
*ServicesRateLimitsBucketsApi* | [**ListBucket**](docs/ServicesRateLimitsBucketsApi.md#listbucket) | **Get** /v2/Services/{ServiceSid}/RateLimits/{RateLimitSid}/Buckets | Retrieve a list of all Buckets for a Rate Limit.
*ServicesRateLimitsBucketsApi* | [**UpdateBucket**](docs/ServicesRateLimitsBucketsApi.md#updatebucket) | **Post** /v2/Services/{ServiceSid}/RateLimits/{RateLimitSid}/Buckets/{Sid} | Update a specific Bucket.
*ServicesVerificationCheckApi* | [**CreateVerificationCheck**](docs/ServicesVerificationCheckApi.md#createverificationcheck) | **Post** /v2/Services/{ServiceSid}/VerificationCheck | challenge a specific Verification Check.
*ServicesVerificationsApi* | [**CreateVerification**](docs/ServicesVerificationsApi.md#createverification) | **Post** /v2/Services/{ServiceSid}/Verifications | Create a new Verification using a Service
*ServicesVerificationsApi* | [**FetchVerification**](docs/ServicesVerificationsApi.md#fetchverification) | **Get** /v2/Services/{ServiceSid}/Verifications/{Sid} | Fetch a specific Verification
*ServicesVerificationsApi* | [**UpdateVerification**](docs/ServicesVerificationsApi.md#updateverification) | **Post** /v2/Services/{ServiceSid}/Verifications/{Sid} | Update a Verification status
*ServicesWebhooksApi* | [**CreateWebhook**](docs/ServicesWebhooksApi.md#createwebhook) | **Post** /v2/Services/{ServiceSid}/Webhooks | Create a new Webhook for the Service
*ServicesWebhooksApi* | [**DeleteWebhook**](docs/ServicesWebhooksApi.md#deletewebhook) | **Delete** /v2/Services/{ServiceSid}/Webhooks/{Sid} | Delete a specific Webhook.
*ServicesWebhooksApi* | [**FetchWebhook**](docs/ServicesWebhooksApi.md#fetchwebhook) | **Get** /v2/Services/{ServiceSid}/Webhooks/{Sid} | Fetch a specific Webhook.
*ServicesWebhooksApi* | [**ListWebhook**](docs/ServicesWebhooksApi.md#listwebhook) | **Get** /v2/Services/{ServiceSid}/Webhooks | Retrieve a list of all Webhooks for a Service.
*ServicesWebhooksApi* | [**UpdateWebhook**](docs/ServicesWebhooksApi.md#updatewebhook) | **Post** /v2/Services/{ServiceSid}/Webhooks/{Sid} | 
*TemplatesApi* | [**ListVerificationTemplate**](docs/TemplatesApi.md#listverificationtemplate) | **Get** /v2/Templates | List all the available templates for a given Account.


## Documentation For Models

 - [ListMessagingConfigurationResponse](docs/ListMessagingConfigurationResponse.md)
 - [VerifyV2RateLimit](docs/VerifyV2RateLimit.md)
 - [VerifyV2AccessToken](docs/VerifyV2AccessToken.md)
 - [VerifyV2Webhook](docs/VerifyV2Webhook.md)
 - [ListServiceResponse](docs/ListServiceResponse.md)
 - [VerifyV2Notification](docs/VerifyV2Notification.md)
 - [VerifyV2Factor](docs/VerifyV2Factor.md)
 - [VerifyV2Form](docs/VerifyV2Form.md)
 - [VerifyV2VerificationAttempt](docs/VerifyV2VerificationAttempt.md)
 - [VerifyV2Safelist](docs/VerifyV2Safelist.md)
 - [VerifyV2Challenge](docs/VerifyV2Challenge.md)
 - [ListEntityResponse](docs/ListEntityResponse.md)
 - [ListWebhookResponse](docs/ListWebhookResponse.md)
 - [ListRateLimitResponse](docs/ListRateLimitResponse.md)
 - [ListFactorResponse](docs/ListFactorResponse.md)
 - [ListVerificationTemplateResponse](docs/ListVerificationTemplateResponse.md)
 - [VerifyV2MessagingConfiguration](docs/VerifyV2MessagingConfiguration.md)
 - [VerifyV2Service](docs/VerifyV2Service.md)
 - [ListBucketResponseMeta](docs/ListBucketResponseMeta.md)
 - [VerifyV2VerificationTemplate](docs/VerifyV2VerificationTemplate.md)
 - [VerifyV2VerificationCheck](docs/VerifyV2VerificationCheck.md)
 - [ListVerificationAttemptResponse](docs/ListVerificationAttemptResponse.md)
 - [VerifyV2NewFactor](docs/VerifyV2NewFactor.md)
 - [VerifyV2VerificationAttemptsSummary](docs/VerifyV2VerificationAttemptsSummary.md)
 - [VerifyV2Verification](docs/VerifyV2Verification.md)
 - [ListBucketResponse](docs/ListBucketResponse.md)
 - [ListChallengeResponse](docs/ListChallengeResponse.md)
 - [VerifyV2Bucket](docs/VerifyV2Bucket.md)
 - [VerifyV2Entity](docs/VerifyV2Entity.md)


## Documentation For Authorization



## accountSid_authToken

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```

