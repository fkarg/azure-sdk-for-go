//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// TransparentDataEncryptionsClient contains the methods for the TransparentDataEncryptions group.
// Don't use this type directly, use NewTransparentDataEncryptionsClient() instead.
type TransparentDataEncryptionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewTransparentDataEncryptionsClient creates a new instance of TransparentDataEncryptionsClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewTransparentDataEncryptionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TransparentDataEncryptionsClient, error) {
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
	client := &TransparentDataEncryptionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Updates a logical database's transparent data encryption configuration.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-02-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// databaseName - The name of the logical database for which the security alert policy is defined.
// tdeName - The name of the transparent data encryption configuration.
// parameters - The database transparent data encryption.
// options - TransparentDataEncryptionsClientCreateOrUpdateOptions contains the optional parameters for the TransparentDataEncryptionsClient.CreateOrUpdate
// method.
func (client *TransparentDataEncryptionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, tdeName TransparentDataEncryptionName, parameters LogicalDatabaseTransparentDataEncryption, options *TransparentDataEncryptionsClientCreateOrUpdateOptions) (TransparentDataEncryptionsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, databaseName, tdeName, parameters, options)
	if err != nil {
		return TransparentDataEncryptionsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TransparentDataEncryptionsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return TransparentDataEncryptionsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *TransparentDataEncryptionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, tdeName TransparentDataEncryptionName, parameters LogicalDatabaseTransparentDataEncryption, options *TransparentDataEncryptionsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/transparentDataEncryption/{tdeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if tdeName == "" {
		return nil, errors.New("parameter tdeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tdeName}", url.PathEscape(string(tdeName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *TransparentDataEncryptionsClient) createOrUpdateHandleResponse(resp *http.Response) (TransparentDataEncryptionsClientCreateOrUpdateResponse, error) {
	result := TransparentDataEncryptionsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogicalDatabaseTransparentDataEncryption); err != nil {
		return TransparentDataEncryptionsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Get - Gets a logical database's transparent data encryption.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-02-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// databaseName - The name of the logical database for which the transparent data encryption is defined.
// tdeName - The name of the transparent data encryption configuration.
// options - TransparentDataEncryptionsClientGetOptions contains the optional parameters for the TransparentDataEncryptionsClient.Get
// method.
func (client *TransparentDataEncryptionsClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, tdeName TransparentDataEncryptionName, options *TransparentDataEncryptionsClientGetOptions) (TransparentDataEncryptionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, databaseName, tdeName, options)
	if err != nil {
		return TransparentDataEncryptionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TransparentDataEncryptionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TransparentDataEncryptionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TransparentDataEncryptionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, tdeName TransparentDataEncryptionName, options *TransparentDataEncryptionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/transparentDataEncryption/{tdeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if tdeName == "" {
		return nil, errors.New("parameter tdeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tdeName}", url.PathEscape(string(tdeName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TransparentDataEncryptionsClient) getHandleResponse(resp *http.Response) (TransparentDataEncryptionsClientGetResponse, error) {
	result := TransparentDataEncryptionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogicalDatabaseTransparentDataEncryption); err != nil {
		return TransparentDataEncryptionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDatabasePager - Gets a list of the logical database's transparent data encryption.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-02-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// databaseName - The name of the logical database for which the transparent data encryption is defined.
// options - TransparentDataEncryptionsClientListByDatabaseOptions contains the optional parameters for the TransparentDataEncryptionsClient.ListByDatabase
// method.
func (client *TransparentDataEncryptionsClient) NewListByDatabasePager(resourceGroupName string, serverName string, databaseName string, options *TransparentDataEncryptionsClientListByDatabaseOptions) *runtime.Pager[TransparentDataEncryptionsClientListByDatabaseResponse] {
	return runtime.NewPager(runtime.PagingHandler[TransparentDataEncryptionsClientListByDatabaseResponse]{
		More: func(page TransparentDataEncryptionsClientListByDatabaseResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TransparentDataEncryptionsClientListByDatabaseResponse) (TransparentDataEncryptionsClientListByDatabaseResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByDatabaseCreateRequest(ctx, resourceGroupName, serverName, databaseName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return TransparentDataEncryptionsClientListByDatabaseResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return TransparentDataEncryptionsClientListByDatabaseResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TransparentDataEncryptionsClientListByDatabaseResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDatabaseHandleResponse(resp)
		},
	})
}

// listByDatabaseCreateRequest creates the ListByDatabase request.
func (client *TransparentDataEncryptionsClient) listByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, options *TransparentDataEncryptionsClientListByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/transparentDataEncryption"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDatabaseHandleResponse handles the ListByDatabase response.
func (client *TransparentDataEncryptionsClient) listByDatabaseHandleResponse(resp *http.Response) (TransparentDataEncryptionsClientListByDatabaseResponse, error) {
	result := TransparentDataEncryptionsClientListByDatabaseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogicalDatabaseTransparentDataEncryptionListResult); err != nil {
		return TransparentDataEncryptionsClientListByDatabaseResponse{}, err
	}
	return result, nil
}