//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbilling

// AccountsClientGetResponse contains the response from method AccountsClient.Get.
type AccountsClientGetResponse struct {
	Account
}

// AccountsClientListInvoiceSectionsByCreateSubscriptionPermissionResponse contains the response from method AccountsClient.ListInvoiceSectionsByCreateSubscriptionPermission.
type AccountsClientListInvoiceSectionsByCreateSubscriptionPermissionResponse struct {
	InvoiceSectionListWithCreateSubPermissionResult
}

// AccountsClientListResponse contains the response from method AccountsClient.List.
type AccountsClientListResponse struct {
	AccountListResult
}

// AccountsClientUpdateResponse contains the response from method AccountsClient.Update.
type AccountsClientUpdateResponse struct {
	Account
}

// AddressClientValidateResponse contains the response from method AddressClient.Validate.
type AddressClientValidateResponse struct {
	ValidateAddressResponse
}

// AgreementsClientGetResponse contains the response from method AgreementsClient.Get.
type AgreementsClientGetResponse struct {
	Agreement
}

// AgreementsClientListByBillingAccountResponse contains the response from method AgreementsClient.ListByBillingAccount.
type AgreementsClientListByBillingAccountResponse struct {
	AgreementListResult
}

// AvailableBalancesClientGetResponse contains the response from method AvailableBalancesClient.Get.
type AvailableBalancesClientGetResponse struct {
	AvailableBalance
}

// CustomersClientGetResponse contains the response from method CustomersClient.Get.
type CustomersClientGetResponse struct {
	Customer
}

// CustomersClientListByBillingAccountResponse contains the response from method CustomersClient.ListByBillingAccount.
type CustomersClientListByBillingAccountResponse struct {
	CustomerListResult
}

// CustomersClientListByBillingProfileResponse contains the response from method CustomersClient.ListByBillingProfile.
type CustomersClientListByBillingProfileResponse struct {
	CustomerListResult
}

// EnrollmentAccountsClientGetResponse contains the response from method EnrollmentAccountsClient.Get.
type EnrollmentAccountsClientGetResponse struct {
	EnrollmentAccountSummary
}

// EnrollmentAccountsClientListResponse contains the response from method EnrollmentAccountsClient.List.
type EnrollmentAccountsClientListResponse struct {
	EnrollmentAccountListResult
}

// InstructionsClientGetResponse contains the response from method InstructionsClient.Get.
type InstructionsClientGetResponse struct {
	Instruction
}

// InstructionsClientListByBillingProfileResponse contains the response from method InstructionsClient.ListByBillingProfile.
type InstructionsClientListByBillingProfileResponse struct {
	InstructionListResult
}

// InstructionsClientPutResponse contains the response from method InstructionsClient.Put.
type InstructionsClientPutResponse struct {
	Instruction
}

// InvoiceSectionsClientCreateOrUpdateResponse contains the response from method InvoiceSectionsClient.CreateOrUpdate.
type InvoiceSectionsClientCreateOrUpdateResponse struct {
	InvoiceSection
}

// InvoiceSectionsClientGetResponse contains the response from method InvoiceSectionsClient.Get.
type InvoiceSectionsClientGetResponse struct {
	InvoiceSection
}

// InvoiceSectionsClientListByBillingProfileResponse contains the response from method InvoiceSectionsClient.ListByBillingProfile.
type InvoiceSectionsClientListByBillingProfileResponse struct {
	InvoiceSectionListResult
}

// InvoicesClientDownloadBillingSubscriptionInvoiceResponse contains the response from method InvoicesClient.DownloadBillingSubscriptionInvoice.
type InvoicesClientDownloadBillingSubscriptionInvoiceResponse struct {
	DownloadURL
}

// InvoicesClientDownloadInvoiceResponse contains the response from method InvoicesClient.DownloadInvoice.
type InvoicesClientDownloadInvoiceResponse struct {
	DownloadURL
}

// InvoicesClientDownloadMultipleBillingProfileInvoicesResponse contains the response from method InvoicesClient.DownloadMultipleBillingProfileInvoices.
type InvoicesClientDownloadMultipleBillingProfileInvoicesResponse struct {
	DownloadURL
}

// InvoicesClientDownloadMultipleBillingSubscriptionInvoicesResponse contains the response from method InvoicesClient.DownloadMultipleBillingSubscriptionInvoices.
type InvoicesClientDownloadMultipleBillingSubscriptionInvoicesResponse struct {
	DownloadURL
}

// InvoicesClientGetByIDResponse contains the response from method InvoicesClient.GetByID.
type InvoicesClientGetByIDResponse struct {
	Invoice
}

// InvoicesClientGetBySubscriptionAndInvoiceIDResponse contains the response from method InvoicesClient.GetBySubscriptionAndInvoiceID.
type InvoicesClientGetBySubscriptionAndInvoiceIDResponse struct {
	Invoice
}

// InvoicesClientGetResponse contains the response from method InvoicesClient.Get.
type InvoicesClientGetResponse struct {
	Invoice
}

// InvoicesClientListByBillingAccountResponse contains the response from method InvoicesClient.ListByBillingAccount.
type InvoicesClientListByBillingAccountResponse struct {
	InvoiceListResult
}

// InvoicesClientListByBillingProfileResponse contains the response from method InvoicesClient.ListByBillingProfile.
type InvoicesClientListByBillingProfileResponse struct {
	InvoiceListResult
}

// InvoicesClientListByBillingSubscriptionResponse contains the response from method InvoicesClient.ListByBillingSubscription.
type InvoicesClientListByBillingSubscriptionResponse struct {
	InvoiceListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// PeriodsClientGetResponse contains the response from method PeriodsClient.Get.
type PeriodsClientGetResponse struct {
	Period
}

// PeriodsClientListResponse contains the response from method PeriodsClient.List.
type PeriodsClientListResponse struct {
	PeriodsListResult
}

// PermissionsClientListByBillingAccountResponse contains the response from method PermissionsClient.ListByBillingAccount.
type PermissionsClientListByBillingAccountResponse struct {
	PermissionsListResult
}

// PermissionsClientListByBillingProfileResponse contains the response from method PermissionsClient.ListByBillingProfile.
type PermissionsClientListByBillingProfileResponse struct {
	PermissionsListResult
}

// PermissionsClientListByCustomerResponse contains the response from method PermissionsClient.ListByCustomer.
type PermissionsClientListByCustomerResponse struct {
	PermissionsListResult
}

// PermissionsClientListByInvoiceSectionsResponse contains the response from method PermissionsClient.ListByInvoiceSections.
type PermissionsClientListByInvoiceSectionsResponse struct {
	PermissionsListResult
}

// PoliciesClientGetByBillingProfileResponse contains the response from method PoliciesClient.GetByBillingProfile.
type PoliciesClientGetByBillingProfileResponse struct {
	Policy
}

// PoliciesClientGetByCustomerResponse contains the response from method PoliciesClient.GetByCustomer.
type PoliciesClientGetByCustomerResponse struct {
	CustomerPolicy
}

// PoliciesClientUpdateCustomerResponse contains the response from method PoliciesClient.UpdateCustomer.
type PoliciesClientUpdateCustomerResponse struct {
	CustomerPolicy
}

// PoliciesClientUpdateResponse contains the response from method PoliciesClient.Update.
type PoliciesClientUpdateResponse struct {
	Policy
}

// ProductsClientGetResponse contains the response from method ProductsClient.Get.
type ProductsClientGetResponse struct {
	Product
}

// ProductsClientListByBillingAccountResponse contains the response from method ProductsClient.ListByBillingAccount.
type ProductsClientListByBillingAccountResponse struct {
	ProductsListResult
}

// ProductsClientListByBillingProfileResponse contains the response from method ProductsClient.ListByBillingProfile.
type ProductsClientListByBillingProfileResponse struct {
	ProductsListResult
}

// ProductsClientListByCustomerResponse contains the response from method ProductsClient.ListByCustomer.
type ProductsClientListByCustomerResponse struct {
	ProductsListResult
}

// ProductsClientListByInvoiceSectionResponse contains the response from method ProductsClient.ListByInvoiceSection.
type ProductsClientListByInvoiceSectionResponse struct {
	ProductsListResult
}

// ProductsClientMoveResponse contains the response from method ProductsClient.Move.
type ProductsClientMoveResponse struct {
	Product
	// Location contains the information returned from the Location header response.
	Location *string

	// RetryAfter contains the information returned from the Retry-After header response.
	RetryAfter *int32
}

// ProductsClientUpdateResponse contains the response from method ProductsClient.Update.
type ProductsClientUpdateResponse struct {
	Product
}

// ProductsClientValidateMoveResponse contains the response from method ProductsClient.ValidateMove.
type ProductsClientValidateMoveResponse struct {
	ValidateProductTransferEligibilityResult
}

// ProfilesClientCreateOrUpdateResponse contains the response from method ProfilesClient.CreateOrUpdate.
type ProfilesClientCreateOrUpdateResponse struct {
	Profile
}

// ProfilesClientGetResponse contains the response from method ProfilesClient.Get.
type ProfilesClientGetResponse struct {
	Profile
}

// ProfilesClientListByBillingAccountResponse contains the response from method ProfilesClient.ListByBillingAccount.
type ProfilesClientListByBillingAccountResponse struct {
	ProfileListResult
}

// PropertyClientGetResponse contains the response from method PropertyClient.Get.
type PropertyClientGetResponse struct {
	Property
}

// PropertyClientUpdateResponse contains the response from method PropertyClient.Update.
type PropertyClientUpdateResponse struct {
	Property
}

// ReservationsClientListByBillingAccountResponse contains the response from method ReservationsClient.ListByBillingAccount.
type ReservationsClientListByBillingAccountResponse struct {
	ReservationsListResult
}

// ReservationsClientListByBillingProfileResponse contains the response from method ReservationsClient.ListByBillingProfile.
type ReservationsClientListByBillingProfileResponse struct {
	ReservationsListResult
}

// RoleAssignmentsClientDeleteByBillingAccountResponse contains the response from method RoleAssignmentsClient.DeleteByBillingAccount.
type RoleAssignmentsClientDeleteByBillingAccountResponse struct {
	RoleAssignment
}

// RoleAssignmentsClientDeleteByBillingProfileResponse contains the response from method RoleAssignmentsClient.DeleteByBillingProfile.
type RoleAssignmentsClientDeleteByBillingProfileResponse struct {
	RoleAssignment
}

// RoleAssignmentsClientDeleteByInvoiceSectionResponse contains the response from method RoleAssignmentsClient.DeleteByInvoiceSection.
type RoleAssignmentsClientDeleteByInvoiceSectionResponse struct {
	RoleAssignment
}

// RoleAssignmentsClientGetByBillingAccountResponse contains the response from method RoleAssignmentsClient.GetByBillingAccount.
type RoleAssignmentsClientGetByBillingAccountResponse struct {
	RoleAssignment
}

// RoleAssignmentsClientGetByBillingProfileResponse contains the response from method RoleAssignmentsClient.GetByBillingProfile.
type RoleAssignmentsClientGetByBillingProfileResponse struct {
	RoleAssignment
}

// RoleAssignmentsClientGetByInvoiceSectionResponse contains the response from method RoleAssignmentsClient.GetByInvoiceSection.
type RoleAssignmentsClientGetByInvoiceSectionResponse struct {
	RoleAssignment
}

// RoleAssignmentsClientListByBillingAccountResponse contains the response from method RoleAssignmentsClient.ListByBillingAccount.
type RoleAssignmentsClientListByBillingAccountResponse struct {
	RoleAssignmentListResult
}

// RoleAssignmentsClientListByBillingProfileResponse contains the response from method RoleAssignmentsClient.ListByBillingProfile.
type RoleAssignmentsClientListByBillingProfileResponse struct {
	RoleAssignmentListResult
}

// RoleAssignmentsClientListByInvoiceSectionResponse contains the response from method RoleAssignmentsClient.ListByInvoiceSection.
type RoleAssignmentsClientListByInvoiceSectionResponse struct {
	RoleAssignmentListResult
}

// RoleDefinitionsClientGetByBillingAccountResponse contains the response from method RoleDefinitionsClient.GetByBillingAccount.
type RoleDefinitionsClientGetByBillingAccountResponse struct {
	RoleDefinition
}

// RoleDefinitionsClientGetByBillingProfileResponse contains the response from method RoleDefinitionsClient.GetByBillingProfile.
type RoleDefinitionsClientGetByBillingProfileResponse struct {
	RoleDefinition
}

// RoleDefinitionsClientGetByInvoiceSectionResponse contains the response from method RoleDefinitionsClient.GetByInvoiceSection.
type RoleDefinitionsClientGetByInvoiceSectionResponse struct {
	RoleDefinition
}

// RoleDefinitionsClientListByBillingAccountResponse contains the response from method RoleDefinitionsClient.ListByBillingAccount.
type RoleDefinitionsClientListByBillingAccountResponse struct {
	RoleDefinitionListResult
}

// RoleDefinitionsClientListByBillingProfileResponse contains the response from method RoleDefinitionsClient.ListByBillingProfile.
type RoleDefinitionsClientListByBillingProfileResponse struct {
	RoleDefinitionListResult
}

// RoleDefinitionsClientListByInvoiceSectionResponse contains the response from method RoleDefinitionsClient.ListByInvoiceSection.
type RoleDefinitionsClientListByInvoiceSectionResponse struct {
	RoleDefinitionListResult
}

// SubscriptionsClientGetResponse contains the response from method SubscriptionsClient.Get.
type SubscriptionsClientGetResponse struct {
	Subscription
}

// SubscriptionsClientListByBillingAccountResponse contains the response from method SubscriptionsClient.ListByBillingAccount.
type SubscriptionsClientListByBillingAccountResponse struct {
	SubscriptionsListResult
}

// SubscriptionsClientListByBillingProfileResponse contains the response from method SubscriptionsClient.ListByBillingProfile.
type SubscriptionsClientListByBillingProfileResponse struct {
	SubscriptionsListResult
}

// SubscriptionsClientListByCustomerResponse contains the response from method SubscriptionsClient.ListByCustomer.
type SubscriptionsClientListByCustomerResponse struct {
	SubscriptionsListResult
}

// SubscriptionsClientListByInvoiceSectionResponse contains the response from method SubscriptionsClient.ListByInvoiceSection.
type SubscriptionsClientListByInvoiceSectionResponse struct {
	SubscriptionsListResult
}

// SubscriptionsClientMoveResponse contains the response from method SubscriptionsClient.Move.
type SubscriptionsClientMoveResponse struct {
	Subscription
}

// SubscriptionsClientUpdateResponse contains the response from method SubscriptionsClient.Update.
type SubscriptionsClientUpdateResponse struct {
	Subscription
}

// SubscriptionsClientValidateMoveResponse contains the response from method SubscriptionsClient.ValidateMove.
type SubscriptionsClientValidateMoveResponse struct {
	ValidateSubscriptionTransferEligibilityResult
}

// TransactionsClientListByInvoiceResponse contains the response from method TransactionsClient.ListByInvoice.
type TransactionsClientListByInvoiceResponse struct {
	TransactionListResult
}