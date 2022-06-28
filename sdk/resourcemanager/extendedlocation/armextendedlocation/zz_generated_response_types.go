//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armextendedlocation

// CustomLocationsClientCreateOrUpdateResponse contains the response from method CustomLocationsClient.CreateOrUpdate.
type CustomLocationsClientCreateOrUpdateResponse struct {
	CustomLocation
}

// CustomLocationsClientDeleteResponse contains the response from method CustomLocationsClient.Delete.
type CustomLocationsClientDeleteResponse struct {
	// placeholder for future response values
}

// CustomLocationsClientGetResponse contains the response from method CustomLocationsClient.Get.
type CustomLocationsClientGetResponse struct {
	CustomLocation
}

// CustomLocationsClientListByResourceGroupResponse contains the response from method CustomLocationsClient.ListByResourceGroup.
type CustomLocationsClientListByResourceGroupResponse struct {
	CustomLocationListResult
}

// CustomLocationsClientListBySubscriptionResponse contains the response from method CustomLocationsClient.ListBySubscription.
type CustomLocationsClientListBySubscriptionResponse struct {
	CustomLocationListResult
}

// CustomLocationsClientListEnabledResourceTypesResponse contains the response from method CustomLocationsClient.ListEnabledResourceTypes.
type CustomLocationsClientListEnabledResourceTypesResponse struct {
	EnabledResourceTypesListResult
}

// CustomLocationsClientListOperationsResponse contains the response from method CustomLocationsClient.ListOperations.
type CustomLocationsClientListOperationsResponse struct {
	CustomLocationOperationsList
}

// CustomLocationsClientUpdateResponse contains the response from method CustomLocationsClient.Update.
type CustomLocationsClientUpdateResponse struct {
	CustomLocation
}