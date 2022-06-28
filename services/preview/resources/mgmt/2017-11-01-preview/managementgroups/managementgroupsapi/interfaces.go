package managementgroupsapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2017-11-01-preview/managementgroups"
	"github.com/Azure/go-autorest/autorest"
)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	CreateOrUpdate(ctx context.Context, groupID string, createManagementGroupRequest managementgroups.CreateManagementGroupRequest, cacheControl string) (result managementgroups.ManagementGroup, err error)
	Delete(ctx context.Context, groupID string, cacheControl string) (result autorest.Response, err error)
	Get(ctx context.Context, groupID string, expand string, recurse *bool, cacheControl string) (result managementgroups.ManagementGroup, err error)
	List(ctx context.Context, cacheControl string, skiptoken string) (result managementgroups.ListResultPage, err error)
	ListComplete(ctx context.Context, cacheControl string, skiptoken string) (result managementgroups.ListResultIterator, err error)
	Update(ctx context.Context, groupID string, createManagementGroupRequest managementgroups.CreateManagementGroupRequest, cacheControl string) (result managementgroups.ManagementGroup, err error)
}

var _ ClientAPI = (*managementgroups.Client)(nil)

// SubscriptionsClientAPI contains the set of methods on the SubscriptionsClient type.
type SubscriptionsClientAPI interface {
	Create(ctx context.Context, groupID string, subscriptionID string, cacheControl string) (result autorest.Response, err error)
	Delete(ctx context.Context, groupID string, subscriptionID string, cacheControl string) (result autorest.Response, err error)
}

var _ SubscriptionsClientAPI = (*managementgroups.SubscriptionsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result managementgroups.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result managementgroups.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*managementgroups.OperationsClient)(nil)