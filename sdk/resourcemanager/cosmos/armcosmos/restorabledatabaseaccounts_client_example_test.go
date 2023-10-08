//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c280892951a9e45c059132c05aace25a9c752d48/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-09-15/examples/CosmosDBRestorableDatabaseAccountList.json
func ExampleRestorableDatabaseAccountsClient_NewListByLocationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRestorableDatabaseAccountsClient().NewListByLocationPager("West US", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.RestorableDatabaseAccountsListResult = armcosmos.RestorableDatabaseAccountsListResult{
		// 	Value: []*armcosmos.RestorableDatabaseAccountGetResult{
		// 		{
		// 			Name: to.Ptr("d9b26648-2f53-4541-b3d8-3044f4f9810d"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/locations/restorableDatabaseAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.DocumentDB/locations/West US/restorableDatabaseAccounts/d9b26648-2f53-4541-b3d8-3044f4f9810d"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armcosmos.RestorableDatabaseAccountProperties{
		// 				AccountName: to.Ptr("ddb1"),
		// 				APIType: to.Ptr(armcosmos.APITypeSQL),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-11T21:56:15Z"); return t}()),
		// 				DeletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-12T22:05:09Z"); return t}()),
		// 				OldestRestorableTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-12T22:05:09Z"); return t}()),
		// 				RestorableLocations: []*armcosmos.RestorableLocationResource{
		// 					{
		// 						CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-30T21:13:10Z"); return t}()),
		// 						DeletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-30T21:13:35Z"); return t}()),
		// 						LocationName: to.Ptr("South Central US"),
		// 						RegionalDatabaseAccountInstanceID: to.Ptr("d7a01f78-606f-45c6-9dac-0df32f433bb5"),
		// 					},
		// 					{
		// 						CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-30T21:13:10Z"); return t}()),
		// 						LocationName: to.Ptr("West US"),
		// 						RegionalDatabaseAccountInstanceID: to.Ptr("fdb43d84-1572-4697-b6e7-2bcda0c51b2c"),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("4f9e6ace-ac7a-446c-98bc-194c502a06b4"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/locations/restorableDatabaseAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.DocumentDB/locations/West US/restorableDatabaseAccounts/4f9e6ace-ac7a-446c-98bc-194c502a06b4"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armcosmos.RestorableDatabaseAccountProperties{
		// 				AccountName: to.Ptr("ddb2"),
		// 				APIType: to.Ptr(armcosmos.APITypeSQL),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-01T08:05:18Z"); return t}()),
		// 				OldestRestorableTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-01T08:05:18Z"); return t}()),
		// 				RestorableLocations: []*armcosmos.RestorableLocationResource{
		// 					{
		// 						CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-30T21:13:10Z"); return t}()),
		// 						DeletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-30T21:13:35Z"); return t}()),
		// 						LocationName: to.Ptr("South Central US"),
		// 						RegionalDatabaseAccountInstanceID: to.Ptr("d7a01f78-606f-45c6-9dac-0df32f433bb5"),
		// 					},
		// 					{
		// 						CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-30T21:13:10Z"); return t}()),
		// 						LocationName: to.Ptr("West US"),
		// 						RegionalDatabaseAccountInstanceID: to.Ptr("fdb43d84-1572-4697-b6e7-2bcda0c51b2c"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c280892951a9e45c059132c05aace25a9c752d48/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-09-15/examples/CosmosDBRestorableDatabaseAccountNoLocationList.json
func ExampleRestorableDatabaseAccountsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRestorableDatabaseAccountsClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.RestorableDatabaseAccountsListResult = armcosmos.RestorableDatabaseAccountsListResult{
		// 	Value: []*armcosmos.RestorableDatabaseAccountGetResult{
		// 		{
		// 			Name: to.Ptr("d9b26648-2f53-4541-b3d8-3044f4f9810d"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/locations/restorableDatabaseAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.DocumentDB/locations/West US/restorableDatabaseAccounts/d9b26648-2f53-4541-b3d8-3044f4f9810d"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armcosmos.RestorableDatabaseAccountProperties{
		// 				AccountName: to.Ptr("ddb1"),
		// 				APIType: to.Ptr(armcosmos.APITypeSQL),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-11T21:56:15Z"); return t}()),
		// 				DeletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-12T22:05:09Z"); return t}()),
		// 				OldestRestorableTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-12T22:05:09Z"); return t}()),
		// 				RestorableLocations: []*armcosmos.RestorableLocationResource{
		// 					{
		// 						CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-30T21:13:10Z"); return t}()),
		// 						DeletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-30T21:13:35Z"); return t}()),
		// 						LocationName: to.Ptr("South Central US"),
		// 						RegionalDatabaseAccountInstanceID: to.Ptr("d7a01f78-606f-45c6-9dac-0df32f433bb5"),
		// 					},
		// 					{
		// 						CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-30T21:13:10Z"); return t}()),
		// 						LocationName: to.Ptr("West US"),
		// 						RegionalDatabaseAccountInstanceID: to.Ptr("fdb43d84-1572-4697-b6e7-2bcda0c51b2c"),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("4f9e6ace-ac7a-446c-98bc-194c502a06b4"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/locations/restorableDatabaseAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.DocumentDB/locations/West US/restorableDatabaseAccounts/4f9e6ace-ac7a-446c-98bc-194c502a06b4"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armcosmos.RestorableDatabaseAccountProperties{
		// 				AccountName: to.Ptr("ddb2"),
		// 				APIType: to.Ptr(armcosmos.APITypeSQL),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-01T08:05:18Z"); return t}()),
		// 				OldestRestorableTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-01T08:05:18Z"); return t}()),
		// 				RestorableLocations: []*armcosmos.RestorableLocationResource{
		// 					{
		// 						CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-30T21:13:10Z"); return t}()),
		// 						DeletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-30T21:13:35Z"); return t}()),
		// 						LocationName: to.Ptr("South Central US"),
		// 						RegionalDatabaseAccountInstanceID: to.Ptr("d7a01f78-606f-45c6-9dac-0df32f433bb5"),
		// 					},
		// 					{
		// 						CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-30T21:13:10Z"); return t}()),
		// 						LocationName: to.Ptr("West US"),
		// 						RegionalDatabaseAccountInstanceID: to.Ptr("fdb43d84-1572-4697-b6e7-2bcda0c51b2c"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c280892951a9e45c059132c05aace25a9c752d48/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-09-15/examples/CosmosDBRestorableDatabaseAccountGet.json
func ExampleRestorableDatabaseAccountsClient_GetByLocation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRestorableDatabaseAccountsClient().GetByLocation(ctx, "West US", "d9b26648-2f53-4541-b3d8-3044f4f9810d", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RestorableDatabaseAccountGetResult = armcosmos.RestorableDatabaseAccountGetResult{
	// 	Name: to.Ptr("d9b26648-2f53-4541-b3d8-3044f4f9810d"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/locations/restorableDatabaseAccounts"),
	// 	ID: to.Ptr("/subscriptions/subid/providers/Microsoft.DocumentDB/locations/West US/restorableDatabaseAccounts/d9b26648-2f53-4541-b3d8-3044f4f9810d"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcosmos.RestorableDatabaseAccountProperties{
	// 		AccountName: to.Ptr("ddb1"),
	// 		APIType: to.Ptr(armcosmos.APITypeSQL),
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-11T21:56:15Z"); return t}()),
	// 		DeletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-12T22:05:09Z"); return t}()),
	// 		OldestRestorableTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-01T22:05:09Z"); return t}()),
	// 		RestorableLocations: []*armcosmos.RestorableLocationResource{
	// 			{
	// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-30T21:13:10Z"); return t}()),
	// 				DeletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-30T21:13:35Z"); return t}()),
	// 				LocationName: to.Ptr("South Central US"),
	// 				RegionalDatabaseAccountInstanceID: to.Ptr("d7a01f78-606f-45c6-9dac-0df32f433bb5"),
	// 			},
	// 			{
	// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-30T21:13:10Z"); return t}()),
	// 				LocationName: to.Ptr("West US"),
	// 				RegionalDatabaseAccountInstanceID: to.Ptr("fdb43d84-1572-4697-b6e7-2bcda0c51b2c"),
	// 		}},
	// 	},
	// }
}
