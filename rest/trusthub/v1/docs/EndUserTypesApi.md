# EndUserTypesApi

All URIs are relative to *https://trusthub.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchEndUserType**](EndUserTypesApi.md#FetchEndUserType) | **Get** /v1/EndUserTypes/{Sid} | Fetch a specific End-User Type Instance.
[**ListEndUserType**](EndUserTypesApi.md#ListEndUserType) | **Get** /v1/EndUserTypes | Retrieve a list of all End-User Types.



## FetchEndUserType

> TrusthubV1EndUserType FetchEndUserType(ctx, Sid)

Fetch a specific End-User Type Instance.

Fetch a specific End-User Type Instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that identifies the End-User Type resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchEndUserTypeParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TrusthubV1EndUserType**](TrusthubV1EndUserType.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEndUserType

> []TrusthubV1EndUserType ListEndUserType(ctx, optional)

Retrieve a list of all End-User Types.

Retrieve a list of all End-User Types.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListEndUserTypeParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]TrusthubV1EndUserType**](TrusthubV1EndUserType.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

