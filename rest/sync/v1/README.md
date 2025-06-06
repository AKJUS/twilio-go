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

All URIs are relative to *https://sync.twilio.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ServicesApi* | [**CreateService**](docs/ServicesApi.md#createservice) | **Post** /v1/Services | 
*ServicesApi* | [**DeleteService**](docs/ServicesApi.md#deleteservice) | **Delete** /v1/Services/{Sid} | 
*ServicesApi* | [**FetchService**](docs/ServicesApi.md#fetchservice) | **Get** /v1/Services/{Sid} | 
*ServicesApi* | [**ListService**](docs/ServicesApi.md#listservice) | **Get** /v1/Services | 
*ServicesApi* | [**UpdateService**](docs/ServicesApi.md#updateservice) | **Post** /v1/Services/{Sid} | 
*ServicesDocumentsApi* | [**CreateDocument**](docs/ServicesDocumentsApi.md#createdocument) | **Post** /v1/Services/{ServiceSid}/Documents | 
*ServicesDocumentsApi* | [**DeleteDocument**](docs/ServicesDocumentsApi.md#deletedocument) | **Delete** /v1/Services/{ServiceSid}/Documents/{Sid} | 
*ServicesDocumentsApi* | [**FetchDocument**](docs/ServicesDocumentsApi.md#fetchdocument) | **Get** /v1/Services/{ServiceSid}/Documents/{Sid} | 
*ServicesDocumentsApi* | [**ListDocument**](docs/ServicesDocumentsApi.md#listdocument) | **Get** /v1/Services/{ServiceSid}/Documents | 
*ServicesDocumentsApi* | [**UpdateDocument**](docs/ServicesDocumentsApi.md#updatedocument) | **Post** /v1/Services/{ServiceSid}/Documents/{Sid} | 
*ServicesDocumentsPermissionsApi* | [**DeleteDocumentPermission**](docs/ServicesDocumentsPermissionsApi.md#deletedocumentpermission) | **Delete** /v1/Services/{ServiceSid}/Documents/{DocumentSid}/Permissions/{Identity} | Delete a specific Sync Document Permission.
*ServicesDocumentsPermissionsApi* | [**FetchDocumentPermission**](docs/ServicesDocumentsPermissionsApi.md#fetchdocumentpermission) | **Get** /v1/Services/{ServiceSid}/Documents/{DocumentSid}/Permissions/{Identity} | Fetch a specific Sync Document Permission.
*ServicesDocumentsPermissionsApi* | [**ListDocumentPermission**](docs/ServicesDocumentsPermissionsApi.md#listdocumentpermission) | **Get** /v1/Services/{ServiceSid}/Documents/{DocumentSid}/Permissions | Retrieve a list of all Permissions applying to a Sync Document.
*ServicesDocumentsPermissionsApi* | [**UpdateDocumentPermission**](docs/ServicesDocumentsPermissionsApi.md#updatedocumentpermission) | **Post** /v1/Services/{ServiceSid}/Documents/{DocumentSid}/Permissions/{Identity} | Update an identity&#39;s access to a specific Sync Document.
*ServicesListsApi* | [**CreateSyncList**](docs/ServicesListsApi.md#createsynclist) | **Post** /v1/Services/{ServiceSid}/Lists | 
*ServicesListsApi* | [**DeleteSyncList**](docs/ServicesListsApi.md#deletesynclist) | **Delete** /v1/Services/{ServiceSid}/Lists/{Sid} | 
*ServicesListsApi* | [**FetchSyncList**](docs/ServicesListsApi.md#fetchsynclist) | **Get** /v1/Services/{ServiceSid}/Lists/{Sid} | 
*ServicesListsApi* | [**ListSyncList**](docs/ServicesListsApi.md#listsynclist) | **Get** /v1/Services/{ServiceSid}/Lists | 
*ServicesListsApi* | [**UpdateSyncList**](docs/ServicesListsApi.md#updatesynclist) | **Post** /v1/Services/{ServiceSid}/Lists/{Sid} | 
*ServicesListsItemsApi* | [**CreateSyncListItem**](docs/ServicesListsItemsApi.md#createsynclistitem) | **Post** /v1/Services/{ServiceSid}/Lists/{ListSid}/Items | 
*ServicesListsItemsApi* | [**DeleteSyncListItem**](docs/ServicesListsItemsApi.md#deletesynclistitem) | **Delete** /v1/Services/{ServiceSid}/Lists/{ListSid}/Items/{Index} | 
*ServicesListsItemsApi* | [**FetchSyncListItem**](docs/ServicesListsItemsApi.md#fetchsynclistitem) | **Get** /v1/Services/{ServiceSid}/Lists/{ListSid}/Items/{Index} | 
*ServicesListsItemsApi* | [**ListSyncListItem**](docs/ServicesListsItemsApi.md#listsynclistitem) | **Get** /v1/Services/{ServiceSid}/Lists/{ListSid}/Items | 
*ServicesListsItemsApi* | [**UpdateSyncListItem**](docs/ServicesListsItemsApi.md#updatesynclistitem) | **Post** /v1/Services/{ServiceSid}/Lists/{ListSid}/Items/{Index} | 
*ServicesListsPermissionsApi* | [**DeleteSyncListPermission**](docs/ServicesListsPermissionsApi.md#deletesynclistpermission) | **Delete** /v1/Services/{ServiceSid}/Lists/{ListSid}/Permissions/{Identity} | Delete a specific Sync List Permission.
*ServicesListsPermissionsApi* | [**FetchSyncListPermission**](docs/ServicesListsPermissionsApi.md#fetchsynclistpermission) | **Get** /v1/Services/{ServiceSid}/Lists/{ListSid}/Permissions/{Identity} | Fetch a specific Sync List Permission.
*ServicesListsPermissionsApi* | [**ListSyncListPermission**](docs/ServicesListsPermissionsApi.md#listsynclistpermission) | **Get** /v1/Services/{ServiceSid}/Lists/{ListSid}/Permissions | Retrieve a list of all Permissions applying to a Sync List.
*ServicesListsPermissionsApi* | [**UpdateSyncListPermission**](docs/ServicesListsPermissionsApi.md#updatesynclistpermission) | **Post** /v1/Services/{ServiceSid}/Lists/{ListSid}/Permissions/{Identity} | Update an identity&#39;s access to a specific Sync List.
*ServicesMapsApi* | [**CreateSyncMap**](docs/ServicesMapsApi.md#createsyncmap) | **Post** /v1/Services/{ServiceSid}/Maps | 
*ServicesMapsApi* | [**DeleteSyncMap**](docs/ServicesMapsApi.md#deletesyncmap) | **Delete** /v1/Services/{ServiceSid}/Maps/{Sid} | 
*ServicesMapsApi* | [**FetchSyncMap**](docs/ServicesMapsApi.md#fetchsyncmap) | **Get** /v1/Services/{ServiceSid}/Maps/{Sid} | 
*ServicesMapsApi* | [**ListSyncMap**](docs/ServicesMapsApi.md#listsyncmap) | **Get** /v1/Services/{ServiceSid}/Maps | 
*ServicesMapsApi* | [**UpdateSyncMap**](docs/ServicesMapsApi.md#updatesyncmap) | **Post** /v1/Services/{ServiceSid}/Maps/{Sid} | 
*ServicesMapsItemsApi* | [**CreateSyncMapItem**](docs/ServicesMapsItemsApi.md#createsyncmapitem) | **Post** /v1/Services/{ServiceSid}/Maps/{MapSid}/Items | 
*ServicesMapsItemsApi* | [**DeleteSyncMapItem**](docs/ServicesMapsItemsApi.md#deletesyncmapitem) | **Delete** /v1/Services/{ServiceSid}/Maps/{MapSid}/Items/{Key} | 
*ServicesMapsItemsApi* | [**FetchSyncMapItem**](docs/ServicesMapsItemsApi.md#fetchsyncmapitem) | **Get** /v1/Services/{ServiceSid}/Maps/{MapSid}/Items/{Key} | 
*ServicesMapsItemsApi* | [**ListSyncMapItem**](docs/ServicesMapsItemsApi.md#listsyncmapitem) | **Get** /v1/Services/{ServiceSid}/Maps/{MapSid}/Items | 
*ServicesMapsItemsApi* | [**UpdateSyncMapItem**](docs/ServicesMapsItemsApi.md#updatesyncmapitem) | **Post** /v1/Services/{ServiceSid}/Maps/{MapSid}/Items/{Key} | 
*ServicesMapsPermissionsApi* | [**DeleteSyncMapPermission**](docs/ServicesMapsPermissionsApi.md#deletesyncmappermission) | **Delete** /v1/Services/{ServiceSid}/Maps/{MapSid}/Permissions/{Identity} | Delete a specific Sync Map Permission.
*ServicesMapsPermissionsApi* | [**FetchSyncMapPermission**](docs/ServicesMapsPermissionsApi.md#fetchsyncmappermission) | **Get** /v1/Services/{ServiceSid}/Maps/{MapSid}/Permissions/{Identity} | Fetch a specific Sync Map Permission.
*ServicesMapsPermissionsApi* | [**ListSyncMapPermission**](docs/ServicesMapsPermissionsApi.md#listsyncmappermission) | **Get** /v1/Services/{ServiceSid}/Maps/{MapSid}/Permissions | Retrieve a list of all Permissions applying to a Sync Map.
*ServicesMapsPermissionsApi* | [**UpdateSyncMapPermission**](docs/ServicesMapsPermissionsApi.md#updatesyncmappermission) | **Post** /v1/Services/{ServiceSid}/Maps/{MapSid}/Permissions/{Identity} | Update an identity&#39;s access to a specific Sync Map.
*ServicesStreamsApi* | [**CreateSyncStream**](docs/ServicesStreamsApi.md#createsyncstream) | **Post** /v1/Services/{ServiceSid}/Streams | Create a new Stream.
*ServicesStreamsApi* | [**DeleteSyncStream**](docs/ServicesStreamsApi.md#deletesyncstream) | **Delete** /v1/Services/{ServiceSid}/Streams/{Sid} | Delete a specific Stream.
*ServicesStreamsApi* | [**FetchSyncStream**](docs/ServicesStreamsApi.md#fetchsyncstream) | **Get** /v1/Services/{ServiceSid}/Streams/{Sid} | Fetch a specific Stream.
*ServicesStreamsApi* | [**ListSyncStream**](docs/ServicesStreamsApi.md#listsyncstream) | **Get** /v1/Services/{ServiceSid}/Streams | Retrieve a list of all Streams in a Service Instance.
*ServicesStreamsApi* | [**UpdateSyncStream**](docs/ServicesStreamsApi.md#updatesyncstream) | **Post** /v1/Services/{ServiceSid}/Streams/{Sid} | Update a specific Stream.
*ServicesStreamsMessagesApi* | [**CreateStreamMessage**](docs/ServicesStreamsMessagesApi.md#createstreammessage) | **Post** /v1/Services/{ServiceSid}/Streams/{StreamSid}/Messages | Create a new Stream Message.


## Documentation For Models

 - [ListSyncStreamResponse](docs/ListSyncStreamResponse.md)
 - [SyncV1StreamMessage](docs/SyncV1StreamMessage.md)
 - [ListDocumentResponse](docs/ListDocumentResponse.md)
 - [SyncV1SyncMap](docs/SyncV1SyncMap.md)
 - [ListDocumentResponseMeta](docs/ListDocumentResponseMeta.md)
 - [ListSyncListPermissionResponse](docs/ListSyncListPermissionResponse.md)
 - [SyncV1SyncListItem](docs/SyncV1SyncListItem.md)
 - [ListSyncMapPermissionResponse](docs/ListSyncMapPermissionResponse.md)
 - [ListSyncMapResponse](docs/ListSyncMapResponse.md)
 - [SyncV1SyncListPermission](docs/SyncV1SyncListPermission.md)
 - [ListServiceResponse](docs/ListServiceResponse.md)
 - [SyncV1SyncMapPermission](docs/SyncV1SyncMapPermission.md)
 - [SyncV1DocumentPermission](docs/SyncV1DocumentPermission.md)
 - [SyncV1SyncMapItem](docs/SyncV1SyncMapItem.md)
 - [ListDocumentPermissionResponse](docs/ListDocumentPermissionResponse.md)
 - [SyncV1Document](docs/SyncV1Document.md)
 - [SyncV1SyncList](docs/SyncV1SyncList.md)
 - [ListSyncListResponse](docs/ListSyncListResponse.md)
 - [ListSyncListItemResponse](docs/ListSyncListItemResponse.md)
 - [SyncV1SyncStream](docs/SyncV1SyncStream.md)
 - [SyncV1Service](docs/SyncV1Service.md)
 - [ListSyncMapItemResponse](docs/ListSyncMapItemResponse.md)


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

