//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsubscription

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// AliasCreatePollerResponse contains the response from method Alias.Create.
type AliasCreatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *AliasCreatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l AliasCreatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (AliasCreateResponse, error) {
	respType := AliasCreateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.SubscriptionAliasResponse)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a AliasCreatePollerResponse from the provided client and resume token.
func (l *AliasCreatePollerResponse) Resume(ctx context.Context, client *AliasClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("AliasClient.Create", token, client.pl, client.createHandleError)
	if err != nil {
		return err
	}
	poller := &AliasCreatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// AliasCreateResponse contains the response from method Alias.Create.
type AliasCreateResponse struct {
	AliasCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AliasCreateResult contains the result from method Alias.Create.
type AliasCreateResult struct {
	SubscriptionAliasResponse
}

// AliasDeleteResponse contains the response from method Alias.Delete.
type AliasDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AliasGetResponse contains the response from method Alias.Get.
type AliasGetResponse struct {
	AliasGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AliasGetResult contains the result from method Alias.Get.
type AliasGetResult struct {
	SubscriptionAliasResponse
}

// AliasListResponse contains the response from method Alias.List.
type AliasListResponse struct {
	AliasListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AliasListResult contains the result from method Alias.List.
type AliasListResult struct {
	SubscriptionAliasListResult
}

// BillingAccountGetPolicyResponse contains the response from method BillingAccount.GetPolicy.
type BillingAccountGetPolicyResponse struct {
	BillingAccountGetPolicyResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// BillingAccountGetPolicyResult contains the result from method BillingAccount.GetPolicy.
type BillingAccountGetPolicyResult struct {
	BillingAccountPoliciesResponse
}

// OperationsListResponse contains the response from method Operations.List.
type OperationsListResponse struct {
	OperationsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsListResult contains the result from method Operations.List.
type OperationsListResult struct {
	OperationListResult
}

// SubscriptionAcceptOwnershipPollerResponse contains the response from method Subscription.AcceptOwnership.
type SubscriptionAcceptOwnershipPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *SubscriptionAcceptOwnershipPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l SubscriptionAcceptOwnershipPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (SubscriptionAcceptOwnershipResponse, error) {
	respType := SubscriptionAcceptOwnershipResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a SubscriptionAcceptOwnershipPollerResponse from the provided client and resume token.
func (l *SubscriptionAcceptOwnershipPollerResponse) Resume(ctx context.Context, client *SubscriptionClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("SubscriptionClient.AcceptOwnership", token, client.pl, client.acceptOwnershipHandleError)
	if err != nil {
		return err
	}
	poller := &SubscriptionAcceptOwnershipPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// SubscriptionAcceptOwnershipResponse contains the response from method Subscription.AcceptOwnership.
type SubscriptionAcceptOwnershipResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionAcceptOwnershipStatusResponse contains the response from method Subscription.AcceptOwnershipStatus.
type SubscriptionAcceptOwnershipStatusResponse struct {
	SubscriptionAcceptOwnershipStatusResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionAcceptOwnershipStatusResult contains the result from method Subscription.AcceptOwnershipStatus.
type SubscriptionAcceptOwnershipStatusResult struct {
	AcceptOwnershipStatusResponse
}

// SubscriptionCancelResponse contains the response from method Subscription.Cancel.
type SubscriptionCancelResponse struct {
	SubscriptionCancelResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionCancelResult contains the result from method Subscription.Cancel.
type SubscriptionCancelResult struct {
	CanceledSubscriptionID
}

// SubscriptionEnableResponse contains the response from method Subscription.Enable.
type SubscriptionEnableResponse struct {
	SubscriptionEnableResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionEnableResult contains the result from method Subscription.Enable.
type SubscriptionEnableResult struct {
	EnabledSubscriptionID
}

// SubscriptionPolicyAddUpdatePolicyForTenantResponse contains the response from method SubscriptionPolicy.AddUpdatePolicyForTenant.
type SubscriptionPolicyAddUpdatePolicyForTenantResponse struct {
	SubscriptionPolicyAddUpdatePolicyForTenantResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionPolicyAddUpdatePolicyForTenantResult contains the result from method SubscriptionPolicy.AddUpdatePolicyForTenant.
type SubscriptionPolicyAddUpdatePolicyForTenantResult struct {
	GetTenantPolicyResponse
}

// SubscriptionPolicyGetPolicyForTenantResponse contains the response from method SubscriptionPolicy.GetPolicyForTenant.
type SubscriptionPolicyGetPolicyForTenantResponse struct {
	SubscriptionPolicyGetPolicyForTenantResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionPolicyGetPolicyForTenantResult contains the result from method SubscriptionPolicy.GetPolicyForTenant.
type SubscriptionPolicyGetPolicyForTenantResult struct {
	GetTenantPolicyResponse
}

// SubscriptionPolicyListPolicyForTenantResponse contains the response from method SubscriptionPolicy.ListPolicyForTenant.
type SubscriptionPolicyListPolicyForTenantResponse struct {
	SubscriptionPolicyListPolicyForTenantResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionPolicyListPolicyForTenantResult contains the result from method SubscriptionPolicy.ListPolicyForTenant.
type SubscriptionPolicyListPolicyForTenantResult struct {
	GetTenantPolicyListResponse
}

// SubscriptionRenameResponse contains the response from method Subscription.Rename.
type SubscriptionRenameResponse struct {
	SubscriptionRenameResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionRenameResult contains the result from method Subscription.Rename.
type SubscriptionRenameResult struct {
	RenamedSubscriptionID
}