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

// ServerAzureADAdministratorsClient contains the methods for the ServerAzureADAdministrators group.
// Don't use this type directly, use NewServerAzureADAdministratorsClient() instead.
type ServerAzureADAdministratorsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewServerAzureADAdministratorsClient creates a new instance of ServerAzureADAdministratorsClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewServerAzureADAdministratorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServerAzureADAdministratorsClient, error) {
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
	client := &ServerAzureADAdministratorsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates an existing Azure Active Directory administrator.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// administratorName - The name of server active directory administrator.
// parameters - The requested Azure Active Directory administrator Resource state.
// options - ServerAzureADAdministratorsClientBeginCreateOrUpdateOptions contains the optional parameters for the ServerAzureADAdministratorsClient.BeginCreateOrUpdate
// method.
func (client *ServerAzureADAdministratorsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, administratorName AdministratorName, parameters ServerAzureADAdministrator, options *ServerAzureADAdministratorsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ServerAzureADAdministratorsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, administratorName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ServerAzureADAdministratorsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ServerAzureADAdministratorsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates an existing Azure Active Directory administrator.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
func (client *ServerAzureADAdministratorsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, administratorName AdministratorName, parameters ServerAzureADAdministrator, options *ServerAzureADAdministratorsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, administratorName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServerAzureADAdministratorsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, administratorName AdministratorName, parameters ServerAzureADAdministrator, options *ServerAzureADAdministratorsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/administrators/{administratorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if administratorName == "" {
		return nil, errors.New("parameter administratorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{administratorName}", url.PathEscape(string(administratorName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the Azure Active Directory administrator with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// administratorName - The name of server active directory administrator.
// options - ServerAzureADAdministratorsClientBeginDeleteOptions contains the optional parameters for the ServerAzureADAdministratorsClient.BeginDelete
// method.
func (client *ServerAzureADAdministratorsClient) BeginDelete(ctx context.Context, resourceGroupName string, serverName string, administratorName AdministratorName, options *ServerAzureADAdministratorsClientBeginDeleteOptions) (*runtime.Poller[ServerAzureADAdministratorsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serverName, administratorName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ServerAzureADAdministratorsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ServerAzureADAdministratorsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the Azure Active Directory administrator with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
func (client *ServerAzureADAdministratorsClient) deleteOperation(ctx context.Context, resourceGroupName string, serverName string, administratorName AdministratorName, options *ServerAzureADAdministratorsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, administratorName, options)
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
func (client *ServerAzureADAdministratorsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, administratorName AdministratorName, options *ServerAzureADAdministratorsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/administrators/{administratorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if administratorName == "" {
		return nil, errors.New("parameter administratorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{administratorName}", url.PathEscape(string(administratorName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a Azure Active Directory administrator.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// administratorName - The name of server active directory administrator.
// options - ServerAzureADAdministratorsClientGetOptions contains the optional parameters for the ServerAzureADAdministratorsClient.Get
// method.
func (client *ServerAzureADAdministratorsClient) Get(ctx context.Context, resourceGroupName string, serverName string, administratorName AdministratorName, options *ServerAzureADAdministratorsClientGetOptions) (ServerAzureADAdministratorsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, administratorName, options)
	if err != nil {
		return ServerAzureADAdministratorsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServerAzureADAdministratorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerAzureADAdministratorsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServerAzureADAdministratorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, administratorName AdministratorName, options *ServerAzureADAdministratorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/administrators/{administratorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if administratorName == "" {
		return nil, errors.New("parameter administratorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{administratorName}", url.PathEscape(string(administratorName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServerAzureADAdministratorsClient) getHandleResponse(resp *http.Response) (ServerAzureADAdministratorsClientGetResponse, error) {
	result := ServerAzureADAdministratorsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerAzureADAdministrator); err != nil {
		return ServerAzureADAdministratorsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - Gets a list of Azure Active Directory administrators in a server.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// options - ServerAzureADAdministratorsClientListByServerOptions contains the optional parameters for the ServerAzureADAdministratorsClient.ListByServer
// method.
func (client *ServerAzureADAdministratorsClient) NewListByServerPager(resourceGroupName string, serverName string, options *ServerAzureADAdministratorsClientListByServerOptions) *runtime.Pager[ServerAzureADAdministratorsClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServerAzureADAdministratorsClientListByServerResponse]{
		More: func(page ServerAzureADAdministratorsClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServerAzureADAdministratorsClientListByServerResponse) (ServerAzureADAdministratorsClientListByServerResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ServerAzureADAdministratorsClientListByServerResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ServerAzureADAdministratorsClientListByServerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServerAzureADAdministratorsClientListByServerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServerHandleResponse(resp)
		},
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *ServerAzureADAdministratorsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerAzureADAdministratorsClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/administrators"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *ServerAzureADAdministratorsClient) listByServerHandleResponse(resp *http.Response) (ServerAzureADAdministratorsClientListByServerResponse, error) {
	result := ServerAzureADAdministratorsClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AdministratorListResult); err != nil {
		return ServerAzureADAdministratorsClientListByServerResponse{}, err
	}
	return result, nil
}