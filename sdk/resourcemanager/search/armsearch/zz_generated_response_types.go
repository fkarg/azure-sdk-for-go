//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsearch

// AdminKeysClientGetResponse contains the response from method AdminKeysClient.Get.
type AdminKeysClientGetResponse struct {
	AdminKeyResult
}

// AdminKeysClientRegenerateResponse contains the response from method AdminKeysClient.Regenerate.
type AdminKeysClientRegenerateResponse struct {
	AdminKeyResult
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.Delete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListByServiceResponse contains the response from method PrivateEndpointConnectionsClient.ListByService.
type PrivateEndpointConnectionsClientListByServiceResponse struct {
	PrivateEndpointConnectionListResult
}

// PrivateEndpointConnectionsClientUpdateResponse contains the response from method PrivateEndpointConnectionsClient.Update.
type PrivateEndpointConnectionsClientUpdateResponse struct {
	PrivateEndpointConnection
}

// PrivateLinkResourcesClientListSupportedResponse contains the response from method PrivateLinkResourcesClient.ListSupported.
type PrivateLinkResourcesClientListSupportedResponse struct {
	PrivateLinkResourcesResult
}

// QueryKeysClientCreateResponse contains the response from method QueryKeysClient.Create.
type QueryKeysClientCreateResponse struct {
	QueryKey
}

// QueryKeysClientDeleteResponse contains the response from method QueryKeysClient.Delete.
type QueryKeysClientDeleteResponse struct {
	// placeholder for future response values
}

// QueryKeysClientListBySearchServiceResponse contains the response from method QueryKeysClient.ListBySearchService.
type QueryKeysClientListBySearchServiceResponse struct {
	ListQueryKeysResult
}

// ServicesClientCheckNameAvailabilityResponse contains the response from method ServicesClient.CheckNameAvailability.
type ServicesClientCheckNameAvailabilityResponse struct {
	CheckNameAvailabilityOutput
}

// ServicesClientCreateOrUpdateResponse contains the response from method ServicesClient.CreateOrUpdate.
type ServicesClientCreateOrUpdateResponse struct {
	Service
}

// ServicesClientDeleteResponse contains the response from method ServicesClient.Delete.
type ServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// ServicesClientGetResponse contains the response from method ServicesClient.Get.
type ServicesClientGetResponse struct {
	Service
}

// ServicesClientListByResourceGroupResponse contains the response from method ServicesClient.ListByResourceGroup.
type ServicesClientListByResourceGroupResponse struct {
	ServiceListResult
}

// ServicesClientListBySubscriptionResponse contains the response from method ServicesClient.ListBySubscription.
type ServicesClientListBySubscriptionResponse struct {
	ServiceListResult
}

// ServicesClientUpdateResponse contains the response from method ServicesClient.Update.
type ServicesClientUpdateResponse struct {
	Service
}

// SharedPrivateLinkResourcesClientCreateOrUpdateResponse contains the response from method SharedPrivateLinkResourcesClient.CreateOrUpdate.
type SharedPrivateLinkResourcesClientCreateOrUpdateResponse struct {
	SharedPrivateLinkResource
}

// SharedPrivateLinkResourcesClientDeleteResponse contains the response from method SharedPrivateLinkResourcesClient.Delete.
type SharedPrivateLinkResourcesClientDeleteResponse struct {
	// placeholder for future response values
}

// SharedPrivateLinkResourcesClientGetResponse contains the response from method SharedPrivateLinkResourcesClient.Get.
type SharedPrivateLinkResourcesClientGetResponse struct {
	SharedPrivateLinkResource
}

// SharedPrivateLinkResourcesClientListByServiceResponse contains the response from method SharedPrivateLinkResourcesClient.ListByService.
type SharedPrivateLinkResourcesClientListByServiceResponse struct {
	SharedPrivateLinkResourceListResult
}