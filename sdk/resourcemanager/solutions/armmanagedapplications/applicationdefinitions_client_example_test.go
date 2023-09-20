//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmanagedapplications_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/solutions/armmanagedapplications/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5fb045bd44f143bae17da2e01552ae531f77d0ba/specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/getApplicationDefinition.json
func ExampleApplicationDefinitionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedapplications.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationDefinitionsClient().Get(ctx, "rg", "myManagedApplicationDef", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ApplicationDefinition = armmanagedapplications.ApplicationDefinition{
	// 	Name: to.Ptr("myManagedApplicationDef"),
	// 	Type: to.Ptr("Microsoft.Solutions/applicationDefinitions"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applicationDefinitions/myManagedApplicationDef"),
	// 	Properties: &armmanagedapplications.ApplicationDefinitionProperties{
	// 		Description: to.Ptr("myManagedApplicationDef description"),
	// 		Artifacts: []*armmanagedapplications.ApplicationDefinitionArtifact{
	// 			{
	// 				Name: to.Ptr(armmanagedapplications.ApplicationDefinitionArtifactNameCreateUIDefinition),
	// 				Type: to.Ptr(armmanagedapplications.ApplicationArtifactTypeCustom),
	// 				URI: to.Ptr("https://path/to/managedApplicationCreateUiDefinition.json"),
	// 			},
	// 			{
	// 				Name: to.Ptr(armmanagedapplications.ApplicationDefinitionArtifactNameApplicationResourceTemplate),
	// 				Type: to.Ptr(armmanagedapplications.ApplicationArtifactTypeTemplate),
	// 				URI: to.Ptr("https://path/to/mainTemplate.json"),
	// 		}},
	// 		Authorizations: []*armmanagedapplications.ApplicationAuthorization{
	// 			{
	// 				PrincipalID: to.Ptr("validprincipalguid"),
	// 				RoleDefinitionID: to.Ptr("validroleguid"),
	// 		}},
	// 		DisplayName: to.Ptr("myManagedApplicationDef"),
	// 		LockLevel: to.Ptr(armmanagedapplications.ApplicationLockLevelNone),
	// 		PackageFileURI: to.Ptr("https://path/to/packagezipfile"),
	// 		Policies: []*armmanagedapplications.ApplicationPolicy{
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5fb045bd44f143bae17da2e01552ae531f77d0ba/specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/deleteApplicationDefinition.json
func ExampleApplicationDefinitionsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedapplications.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewApplicationDefinitionsClient().Delete(ctx, "rg", "myManagedApplicationDef", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5fb045bd44f143bae17da2e01552ae531f77d0ba/specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/createOrUpdateApplicationDefinition.json
func ExampleApplicationDefinitionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedapplications.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationDefinitionsClient().CreateOrUpdate(ctx, "rg", "myManagedApplicationDef", armmanagedapplications.ApplicationDefinition{
		Properties: &armmanagedapplications.ApplicationDefinitionProperties{
			Description: to.Ptr("myManagedApplicationDef description"),
			Authorizations: []*armmanagedapplications.ApplicationAuthorization{
				{
					PrincipalID:      to.Ptr("validprincipalguid"),
					RoleDefinitionID: to.Ptr("validroleguid"),
				}},
			DisplayName:    to.Ptr("myManagedApplicationDef"),
			LockLevel:      to.Ptr(armmanagedapplications.ApplicationLockLevelNone),
			PackageFileURI: to.Ptr("https://path/to/packagezipfile"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ApplicationDefinition = armmanagedapplications.ApplicationDefinition{
	// 	Name: to.Ptr("myManagedApplicationDef"),
	// 	Type: to.Ptr("Microsoft.Solutions/applicationDefinitions"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applicationDefinitions/myManagedApplicationDef"),
	// 	Properties: &armmanagedapplications.ApplicationDefinitionProperties{
	// 		Description: to.Ptr("myManagedApplicationDef description"),
	// 		Artifacts: []*armmanagedapplications.ApplicationDefinitionArtifact{
	// 			{
	// 				Name: to.Ptr(armmanagedapplications.ApplicationDefinitionArtifactNameCreateUIDefinition),
	// 				Type: to.Ptr(armmanagedapplications.ApplicationArtifactTypeCustom),
	// 				URI: to.Ptr("https://path/to/managedApplicationCreateUiDefinition.json"),
	// 			},
	// 			{
	// 				Name: to.Ptr(armmanagedapplications.ApplicationDefinitionArtifactNameApplicationResourceTemplate),
	// 				Type: to.Ptr(armmanagedapplications.ApplicationArtifactTypeTemplate),
	// 				URI: to.Ptr("https://path/to/mainTemplate.json"),
	// 		}},
	// 		Authorizations: []*armmanagedapplications.ApplicationAuthorization{
	// 			{
	// 				PrincipalID: to.Ptr("validprincipalguid"),
	// 				RoleDefinitionID: to.Ptr("validroleguid"),
	// 		}},
	// 		DisplayName: to.Ptr("myManagedApplicationDef"),
	// 		LockLevel: to.Ptr(armmanagedapplications.ApplicationLockLevelNone),
	// 		PackageFileURI: to.Ptr("https://path/to/packagezipfile"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5fb045bd44f143bae17da2e01552ae531f77d0ba/specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/updateApplicationDefinition.json
func ExampleApplicationDefinitionsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedapplications.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationDefinitionsClient().Update(ctx, "rg", "myManagedApplicationDef", armmanagedapplications.ApplicationDefinitionPatchable{
		Tags: map[string]*string{
			"department": to.Ptr("Finance"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ApplicationDefinition = armmanagedapplications.ApplicationDefinition{
	// 	Name: to.Ptr("myManagedApplicationDef"),
	// 	Type: to.Ptr("Microsoft.Solutions/applicationDefinitions"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applicationDefinitions/myManagedApplicationDef"),
	// 	Properties: &armmanagedapplications.ApplicationDefinitionProperties{
	// 		Description: to.Ptr("myManagedApplicationDef description"),
	// 		Artifacts: []*armmanagedapplications.ApplicationDefinitionArtifact{
	// 			{
	// 				Name: to.Ptr(armmanagedapplications.ApplicationDefinitionArtifactNameCreateUIDefinition),
	// 				Type: to.Ptr(armmanagedapplications.ApplicationArtifactTypeCustom),
	// 				URI: to.Ptr("https://path/to/managedApplicationCreateUiDefinition.json"),
	// 			},
	// 			{
	// 				Name: to.Ptr(armmanagedapplications.ApplicationDefinitionArtifactNameApplicationResourceTemplate),
	// 				Type: to.Ptr(armmanagedapplications.ApplicationArtifactTypeTemplate),
	// 				URI: to.Ptr("https://path/to/mainTemplate.json"),
	// 		}},
	// 		Authorizations: []*armmanagedapplications.ApplicationAuthorization{
	// 			{
	// 				PrincipalID: to.Ptr("validprincipalguid"),
	// 				RoleDefinitionID: to.Ptr("validroleguid"),
	// 		}},
	// 		DisplayName: to.Ptr("myManagedApplicationDef"),
	// 		LockLevel: to.Ptr(armmanagedapplications.ApplicationLockLevelNone),
	// 		PackageFileURI: to.Ptr("https://path/to/packagezipfile"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5fb045bd44f143bae17da2e01552ae531f77d0ba/specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/listApplicationDefinitionsByResourceGroup.json
func ExampleApplicationDefinitionsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedapplications.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewApplicationDefinitionsClient().NewListByResourceGroupPager("rg", nil)
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
		// page.ApplicationDefinitionListResult = armmanagedapplications.ApplicationDefinitionListResult{
		// 	Value: []*armmanagedapplications.ApplicationDefinition{
		// 		{
		// 			Name: to.Ptr("myManagedApplicationDef"),
		// 			Type: to.Ptr("Microsoft.Solutions/applicationDefinitions"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applicationDefinitions/myManagedApplicationDef"),
		// 			Properties: &armmanagedapplications.ApplicationDefinitionProperties{
		// 				Description: to.Ptr("myManagedApplicationDef description"),
		// 				Artifacts: []*armmanagedapplications.ApplicationDefinitionArtifact{
		// 					{
		// 						Name: to.Ptr(armmanagedapplications.ApplicationDefinitionArtifactNameCreateUIDefinition),
		// 						Type: to.Ptr(armmanagedapplications.ApplicationArtifactTypeCustom),
		// 						URI: to.Ptr("https://path/to/managedApplicationCreateUiDefinition.json"),
		// 					},
		// 					{
		// 						Name: to.Ptr(armmanagedapplications.ApplicationDefinitionArtifactNameApplicationResourceTemplate),
		// 						Type: to.Ptr(armmanagedapplications.ApplicationArtifactTypeTemplate),
		// 						URI: to.Ptr("https://path/to/mainTemplate.json"),
		// 				}},
		// 				Authorizations: []*armmanagedapplications.ApplicationAuthorization{
		// 					{
		// 						PrincipalID: to.Ptr("validprincipalguid"),
		// 						RoleDefinitionID: to.Ptr("validroleguid"),
		// 				}},
		// 				DisplayName: to.Ptr("myManagedApplicationDef"),
		// 				LockLevel: to.Ptr(armmanagedapplications.ApplicationLockLevelNone),
		// 				PackageFileURI: to.Ptr("https://path/to/packagezipfile"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myManagedApplicationDef2"),
		// 			Type: to.Ptr("Microsoft.Solutions/applicationDefinitions"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applicationDefinitions/myManagedApplicationDef2"),
		// 			Properties: &armmanagedapplications.ApplicationDefinitionProperties{
		// 				Description: to.Ptr("myManagedApplicationDef2 description"),
		// 				Artifacts: []*armmanagedapplications.ApplicationDefinitionArtifact{
		// 					{
		// 						Name: to.Ptr(armmanagedapplications.ApplicationDefinitionArtifactNameCreateUIDefinition),
		// 						Type: to.Ptr(armmanagedapplications.ApplicationArtifactTypeCustom),
		// 						URI: to.Ptr("https://path/to/managedApplicationCreateUiDefinition.json"),
		// 					},
		// 					{
		// 						Name: to.Ptr(armmanagedapplications.ApplicationDefinitionArtifactNameApplicationResourceTemplate),
		// 						Type: to.Ptr(armmanagedapplications.ApplicationArtifactTypeTemplate),
		// 						URI: to.Ptr("https://path/to/mainTemplate.json"),
		// 				}},
		// 				Authorizations: []*armmanagedapplications.ApplicationAuthorization{
		// 					{
		// 						PrincipalID: to.Ptr("validprincipalguid"),
		// 						RoleDefinitionID: to.Ptr("validroleguid"),
		// 				}},
		// 				DisplayName: to.Ptr("myManagedApplicationDef2"),
		// 				LockLevel: to.Ptr(armmanagedapplications.ApplicationLockLevelNone),
		// 				PackageFileURI: to.Ptr("https://path/to/packagezipfile"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5fb045bd44f143bae17da2e01552ae531f77d0ba/specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/listApplicationDefinitionsBySubscription.json
func ExampleApplicationDefinitionsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedapplications.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewApplicationDefinitionsClient().NewListBySubscriptionPager(nil)
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
		// page.ApplicationDefinitionListResult = armmanagedapplications.ApplicationDefinitionListResult{
		// 	Value: []*armmanagedapplications.ApplicationDefinition{
		// 		{
		// 			Name: to.Ptr("myManagedApplicationDef"),
		// 			Type: to.Ptr("Microsoft.Solutions/applicationDefinitions"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applicationDefinitions/myManagedApplicationDef"),
		// 			Properties: &armmanagedapplications.ApplicationDefinitionProperties{
		// 				Description: to.Ptr("myManagedApplicationDef description"),
		// 				Artifacts: []*armmanagedapplications.ApplicationDefinitionArtifact{
		// 					{
		// 						Name: to.Ptr(armmanagedapplications.ApplicationDefinitionArtifactNameCreateUIDefinition),
		// 						Type: to.Ptr(armmanagedapplications.ApplicationArtifactTypeCustom),
		// 						URI: to.Ptr("https://path/to/managedApplicationCreateUiDefinition.json"),
		// 					},
		// 					{
		// 						Name: to.Ptr(armmanagedapplications.ApplicationDefinitionArtifactNameApplicationResourceTemplate),
		// 						Type: to.Ptr(armmanagedapplications.ApplicationArtifactTypeTemplate),
		// 						URI: to.Ptr("https://path/to/mainTemplate.json"),
		// 				}},
		// 				Authorizations: []*armmanagedapplications.ApplicationAuthorization{
		// 					{
		// 						PrincipalID: to.Ptr("validprincipalguid"),
		// 						RoleDefinitionID: to.Ptr("validroleguid"),
		// 				}},
		// 				DisplayName: to.Ptr("myManagedApplicationDef"),
		// 				LockLevel: to.Ptr(armmanagedapplications.ApplicationLockLevelNone),
		// 				PackageFileURI: to.Ptr("https://path/to/packagezipfile"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myManagedApplicationDef2"),
		// 			Type: to.Ptr("Microsoft.Solutions/applicationDefinitions"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applicationDefinitions/myManagedApplicationDef2"),
		// 			Properties: &armmanagedapplications.ApplicationDefinitionProperties{
		// 				Description: to.Ptr("myManagedApplicationDef2 description"),
		// 				Artifacts: []*armmanagedapplications.ApplicationDefinitionArtifact{
		// 					{
		// 						Name: to.Ptr(armmanagedapplications.ApplicationDefinitionArtifactNameCreateUIDefinition),
		// 						Type: to.Ptr(armmanagedapplications.ApplicationArtifactTypeCustom),
		// 						URI: to.Ptr("https://path/to/managedApplicationCreateUiDefinition.json"),
		// 					},
		// 					{
		// 						Name: to.Ptr(armmanagedapplications.ApplicationDefinitionArtifactNameApplicationResourceTemplate),
		// 						Type: to.Ptr(armmanagedapplications.ApplicationArtifactTypeTemplate),
		// 						URI: to.Ptr("https://path/to/mainTemplate.json"),
		// 				}},
		// 				Authorizations: []*armmanagedapplications.ApplicationAuthorization{
		// 					{
		// 						PrincipalID: to.Ptr("validprincipalguid"),
		// 						RoleDefinitionID: to.Ptr("validroleguid"),
		// 				}},
		// 				DisplayName: to.Ptr("myManagedApplicationDef2"),
		// 				LockLevel: to.Ptr(armmanagedapplications.ApplicationLockLevelNone),
		// 				PackageFileURI: to.Ptr("https://path/to/packagezipfile"),
		// 			},
		// 	}},
		// }
	}
}
