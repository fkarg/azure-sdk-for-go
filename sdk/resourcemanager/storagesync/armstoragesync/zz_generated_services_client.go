//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragesync

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ServicesClient contains the methods for the StorageSyncServices group.
// Don't use this type directly, use NewServicesClient() instead.
type ServicesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewServicesClient creates a new instance of ServicesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServicesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ServicesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CheckNameAvailability - Check the give namespace name availability.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-09-01
// locationName - The desired region for the name check.
// parameters - Parameters to check availability of the given namespace name
// options - ServicesClientCheckNameAvailabilityOptions contains the optional parameters for the ServicesClient.CheckNameAvailability
// method.
func (client *ServicesClient) CheckNameAvailability(ctx context.Context, locationName string, parameters CheckNameAvailabilityParameters, options *ServicesClientCheckNameAvailabilityOptions) (ServicesClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, locationName, parameters, options)
	if err != nil {
		return ServicesClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServicesClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServicesClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *ServicesClient) checkNameAvailabilityCreateRequest(ctx context.Context, locationName string, parameters CheckNameAvailabilityParameters, options *ServicesClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.StorageSync/locations/{locationName}/checkNameAvailability"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *ServicesClient) checkNameAvailabilityHandleResponse(resp *http.Response) (ServicesClientCheckNameAvailabilityResponse, error) {
	result := ServicesClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResult); err != nil {
		return ServicesClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreate - Create a new StorageSyncService.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-09-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageSyncServiceName - Name of Storage Sync Service resource.
// parameters - Storage Sync Service resource name.
// options - ServicesClientBeginCreateOptions contains the optional parameters for the ServicesClient.BeginCreate method.
func (client *ServicesClient) BeginCreate(ctx context.Context, resourceGroupName string, storageSyncServiceName string, parameters ServiceCreateParameters, options *ServicesClientBeginCreateOptions) (*runtime.Poller[ServicesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, storageSyncServiceName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ServicesClientCreateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ServicesClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Create a new StorageSyncService.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-09-01
func (client *ServicesClient) create(ctx context.Context, resourceGroupName string, storageSyncServiceName string, parameters ServiceCreateParameters, options *ServicesClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, storageSyncServiceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *ServicesClient) createCreateRequest(ctx context.Context, resourceGroupName string, storageSyncServiceName string, parameters ServiceCreateParameters, options *ServicesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageSyncServiceName == "" {
		return nil, errors.New("parameter storageSyncServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageSyncServiceName}", url.PathEscape(storageSyncServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Delete a given StorageSyncService.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-09-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageSyncServiceName - Name of Storage Sync Service resource.
// options - ServicesClientBeginDeleteOptions contains the optional parameters for the ServicesClient.BeginDelete method.
func (client *ServicesClient) BeginDelete(ctx context.Context, resourceGroupName string, storageSyncServiceName string, options *ServicesClientBeginDeleteOptions) (*runtime.Poller[ServicesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, storageSyncServiceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ServicesClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ServicesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete a given StorageSyncService.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-09-01
func (client *ServicesClient) deleteOperation(ctx context.Context, resourceGroupName string, storageSyncServiceName string, options *ServicesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, storageSyncServiceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ServicesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, storageSyncServiceName string, options *ServicesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageSyncServiceName == "" {
		return nil, errors.New("parameter storageSyncServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageSyncServiceName}", url.PathEscape(storageSyncServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a given StorageSyncService.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-09-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageSyncServiceName - Name of Storage Sync Service resource.
// options - ServicesClientGetOptions contains the optional parameters for the ServicesClient.Get method.
func (client *ServicesClient) Get(ctx context.Context, resourceGroupName string, storageSyncServiceName string, options *ServicesClientGetOptions) (ServicesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, storageSyncServiceName, options)
	if err != nil {
		return ServicesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServicesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServicesClient) getCreateRequest(ctx context.Context, resourceGroupName string, storageSyncServiceName string, options *ServicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageSyncServiceName == "" {
		return nil, errors.New("parameter storageSyncServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageSyncServiceName}", url.PathEscape(storageSyncServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServicesClient) getHandleResponse(resp *http.Response) (ServicesClientGetResponse, error) {
	result := ServicesClientGetResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if val := resp.Header.Get("x-ms-correlation-request-id"); val != "" {
		result.XMSCorrelationRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Service); err != nil {
		return ServicesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Get a StorageSyncService list by Resource group name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-09-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - ServicesClientListByResourceGroupOptions contains the optional parameters for the ServicesClient.ListByResourceGroup
// method.
func (client *ServicesClient) NewListByResourceGroupPager(resourceGroupName string, options *ServicesClientListByResourceGroupOptions) *runtime.Pager[ServicesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServicesClientListByResourceGroupResponse]{
		More: func(page ServicesClientListByResourceGroupResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ServicesClientListByResourceGroupResponse) (ServicesClientListByResourceGroupResponse, error) {
			req, err := client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			if err != nil {
				return ServicesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ServicesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServicesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ServicesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ServicesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ServicesClient) listByResourceGroupHandleResponse(resp *http.Response) (ServicesClientListByResourceGroupResponse, error) {
	result := ServicesClientListByResourceGroupResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if val := resp.Header.Get("x-ms-correlation-request-id"); val != "" {
		result.XMSCorrelationRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceArray); err != nil {
		return ServicesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Get a StorageSyncService list by subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-09-01
// options - ServicesClientListBySubscriptionOptions contains the optional parameters for the ServicesClient.ListBySubscription
// method.
func (client *ServicesClient) NewListBySubscriptionPager(options *ServicesClientListBySubscriptionOptions) *runtime.Pager[ServicesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServicesClientListBySubscriptionResponse]{
		More: func(page ServicesClientListBySubscriptionResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ServicesClientListBySubscriptionResponse) (ServicesClientListBySubscriptionResponse, error) {
			req, err := client.listBySubscriptionCreateRequest(ctx, options)
			if err != nil {
				return ServicesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ServicesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServicesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ServicesClient) listBySubscriptionCreateRequest(ctx context.Context, options *ServicesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.StorageSync/storageSyncServices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ServicesClient) listBySubscriptionHandleResponse(resp *http.Response) (ServicesClientListBySubscriptionResponse, error) {
	result := ServicesClientListBySubscriptionResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if val := resp.Header.Get("x-ms-correlation-request-id"); val != "" {
		result.XMSCorrelationRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceArray); err != nil {
		return ServicesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Patch a given StorageSyncService.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-09-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageSyncServiceName - Name of Storage Sync Service resource.
// options - ServicesClientBeginUpdateOptions contains the optional parameters for the ServicesClient.BeginUpdate method.
func (client *ServicesClient) BeginUpdate(ctx context.Context, resourceGroupName string, storageSyncServiceName string, options *ServicesClientBeginUpdateOptions) (*runtime.Poller[ServicesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, storageSyncServiceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ServicesClientUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ServicesClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Patch a given StorageSyncService.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-09-01
func (client *ServicesClient) update(ctx context.Context, resourceGroupName string, storageSyncServiceName string, options *ServicesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, storageSyncServiceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ServicesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, storageSyncServiceName string, options *ServicesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageSyncServiceName == "" {
		return nil, errors.New("parameter storageSyncServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageSyncServiceName}", url.PathEscape(storageSyncServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Parameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.Parameters)
	}
	return req, nil
}