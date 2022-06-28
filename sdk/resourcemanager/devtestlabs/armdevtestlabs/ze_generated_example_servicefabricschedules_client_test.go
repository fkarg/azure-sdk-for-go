//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ServiceFabricSchedules_List.json
func ExampleServiceFabricSchedulesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevtestlabs.NewServiceFabricSchedulesClient("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("resourceGroupName",
		"{labName}",
		"@me",
		"{serviceFrabicName}",
		&armdevtestlabs.ServiceFabricSchedulesClientListOptions{Expand: nil,
			Filter:  nil,
			Top:     nil,
			Orderby: nil,
		})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ServiceFabricSchedules_Get.json
func ExampleServiceFabricSchedulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevtestlabs.NewServiceFabricSchedulesClient("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"resourceGroupName",
		"{labName}",
		"@me",
		"{serviceFrabicName}",
		"{scheduleName}",
		&armdevtestlabs.ServiceFabricSchedulesClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ServiceFabricSchedules_CreateOrUpdate.json
func ExampleServiceFabricSchedulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevtestlabs.NewServiceFabricSchedulesClient("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"resourceGroupName",
		"{labName}",
		"@me",
		"{serviceFrabicName}",
		"{scheduleName}",
		armdevtestlabs.Schedule{
			Location: to.Ptr("{location}"),
			Tags: map[string]*string{
				"tagName1": to.Ptr("tagValue1"),
			},
			Properties: &armdevtestlabs.ScheduleProperties{
				DailyRecurrence: &armdevtestlabs.DayDetails{
					Time: to.Ptr("19:00"),
				},
				HourlyRecurrence: &armdevtestlabs.HourDetails{
					Minute: to.Ptr[int32](0),
				},
				NotificationSettings: &armdevtestlabs.NotificationSettings{
					EmailRecipient:     to.Ptr("{email}"),
					NotificationLocale: to.Ptr("EN"),
					Status:             to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
					TimeInMinutes:      to.Ptr[int32](15),
					WebhookURL:         to.Ptr("{webhoolUrl}"),
				},
				Status:           to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
				TargetResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.DevTestLab/labs/{labName}/users/{uniqueIdentifier}/servicefabrics/{serviceFrabicName}"),
				TaskType:         to.Ptr("{Unknown|LabVmsShutdownTask|LabVmsStartupTask|LabVmReclamationTask|ComputeVmShutdownTask}"),
				TimeZoneID:       to.Ptr("Pacific Standard Time"),
				WeeklyRecurrence: &armdevtestlabs.WeekDetails{
					Time: to.Ptr("19:00"),
					Weekdays: []*string{
						to.Ptr("Monday"),
						to.Ptr("Tuesday"),
						to.Ptr("Wednesday"),
						to.Ptr("Thursday"),
						to.Ptr("Friday"),
						to.Ptr("Saturday"),
						to.Ptr("Sunday")},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ServiceFabricSchedules_Delete.json
func ExampleServiceFabricSchedulesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevtestlabs.NewServiceFabricSchedulesClient("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"resourceGroupName",
		"{labName}",
		"@me",
		"{serviceFrabicName}",
		"{scheduleName}",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ServiceFabricSchedules_Update.json
func ExampleServiceFabricSchedulesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevtestlabs.NewServiceFabricSchedulesClient("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"resourceGroupName",
		"{labName}",
		"@me",
		"{serviceFrabicName}",
		"{scheduleName}",
		armdevtestlabs.ScheduleFragment{
			Tags: map[string]*string{
				"tagName1": to.Ptr("tagValue1"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ServiceFabricSchedules_Execute.json
func ExampleServiceFabricSchedulesClient_BeginExecute() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevtestlabs.NewServiceFabricSchedulesClient("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginExecute(ctx,
		"resourceGroupName",
		"{labName}",
		"@me",
		"{serviceFrabicName}",
		"{scheduleName}",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}