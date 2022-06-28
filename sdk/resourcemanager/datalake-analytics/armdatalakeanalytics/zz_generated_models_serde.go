//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatalakeanalytics

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// UnmarshalJSON implements the json.Unmarshaller interface for type AccountProperties.
func (a *AccountProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "accountId":
			err = unpopulate(val, "AccountID", &a.AccountID)
			delete(rawMsg, key)
		case "computePolicies":
			err = unpopulate(val, "ComputePolicies", &a.ComputePolicies)
			delete(rawMsg, key)
		case "creationTime":
			err = unpopulateTimeRFC3339(val, "CreationTime", &a.CreationTime)
			delete(rawMsg, key)
		case "currentTier":
			err = unpopulate(val, "CurrentTier", &a.CurrentTier)
			delete(rawMsg, key)
		case "dataLakeStoreAccounts":
			err = unpopulate(val, "DataLakeStoreAccounts", &a.DataLakeStoreAccounts)
			delete(rawMsg, key)
		case "debugDataAccessLevel":
			err = unpopulate(val, "DebugDataAccessLevel", &a.DebugDataAccessLevel)
			delete(rawMsg, key)
		case "defaultDataLakeStoreAccount":
			err = unpopulate(val, "DefaultDataLakeStoreAccount", &a.DefaultDataLakeStoreAccount)
			delete(rawMsg, key)
		case "defaultDataLakeStoreAccountType":
			err = unpopulate(val, "DefaultDataLakeStoreAccountType", &a.DefaultDataLakeStoreAccountType)
			delete(rawMsg, key)
		case "endpoint":
			err = unpopulate(val, "Endpoint", &a.Endpoint)
			delete(rawMsg, key)
		case "firewallAllowAzureIps":
			err = unpopulate(val, "FirewallAllowAzureIPs", &a.FirewallAllowAzureIPs)
			delete(rawMsg, key)
		case "firewallRules":
			err = unpopulate(val, "FirewallRules", &a.FirewallRules)
			delete(rawMsg, key)
		case "firewallState":
			err = unpopulate(val, "FirewallState", &a.FirewallState)
			delete(rawMsg, key)
		case "hiveMetastores":
			err = unpopulate(val, "HiveMetastores", &a.HiveMetastores)
			delete(rawMsg, key)
		case "lastModifiedTime":
			err = unpopulateTimeRFC3339(val, "LastModifiedTime", &a.LastModifiedTime)
			delete(rawMsg, key)
		case "maxActiveJobCountPerUser":
			err = unpopulate(val, "MaxActiveJobCountPerUser", &a.MaxActiveJobCountPerUser)
			delete(rawMsg, key)
		case "maxDegreeOfParallelism":
			err = unpopulate(val, "MaxDegreeOfParallelism", &a.MaxDegreeOfParallelism)
			delete(rawMsg, key)
		case "maxDegreeOfParallelismPerJob":
			err = unpopulate(val, "MaxDegreeOfParallelismPerJob", &a.MaxDegreeOfParallelismPerJob)
			delete(rawMsg, key)
		case "maxJobCount":
			err = unpopulate(val, "MaxJobCount", &a.MaxJobCount)
			delete(rawMsg, key)
		case "maxJobRunningTimeInMin":
			err = unpopulate(val, "MaxJobRunningTimeInMin", &a.MaxJobRunningTimeInMin)
			delete(rawMsg, key)
		case "maxQueuedJobCountPerUser":
			err = unpopulate(val, "MaxQueuedJobCountPerUser", &a.MaxQueuedJobCountPerUser)
			delete(rawMsg, key)
		case "minPriorityPerJob":
			err = unpopulate(val, "MinPriorityPerJob", &a.MinPriorityPerJob)
			delete(rawMsg, key)
		case "newTier":
			err = unpopulate(val, "NewTier", &a.NewTier)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &a.ProvisioningState)
			delete(rawMsg, key)
		case "publicDataLakeStoreAccounts":
			err = unpopulate(val, "PublicDataLakeStoreAccounts", &a.PublicDataLakeStoreAccounts)
			delete(rawMsg, key)
		case "queryStoreRetention":
			err = unpopulate(val, "QueryStoreRetention", &a.QueryStoreRetention)
			delete(rawMsg, key)
		case "state":
			err = unpopulate(val, "State", &a.State)
			delete(rawMsg, key)
		case "storageAccounts":
			err = unpopulate(val, "StorageAccounts", &a.StorageAccounts)
			delete(rawMsg, key)
		case "systemMaxDegreeOfParallelism":
			err = unpopulate(val, "SystemMaxDegreeOfParallelism", &a.SystemMaxDegreeOfParallelism)
			delete(rawMsg, key)
		case "systemMaxJobCount":
			err = unpopulate(val, "SystemMaxJobCount", &a.SystemMaxJobCount)
			delete(rawMsg, key)
		case "virtualNetworkRules":
			err = unpopulate(val, "VirtualNetworkRules", &a.VirtualNetworkRules)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AccountPropertiesBasic.
func (a *AccountPropertiesBasic) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "accountId":
			err = unpopulate(val, "AccountID", &a.AccountID)
			delete(rawMsg, key)
		case "creationTime":
			err = unpopulateTimeRFC3339(val, "CreationTime", &a.CreationTime)
			delete(rawMsg, key)
		case "endpoint":
			err = unpopulate(val, "Endpoint", &a.Endpoint)
			delete(rawMsg, key)
		case "lastModifiedTime":
			err = unpopulateTimeRFC3339(val, "LastModifiedTime", &a.LastModifiedTime)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &a.ProvisioningState)
			delete(rawMsg, key)
		case "state":
			err = unpopulate(val, "State", &a.State)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CreateDataLakeAnalyticsAccountParameters.
func (c CreateDataLakeAnalyticsAccountParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "location", c.Location)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "tags", c.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type CreateDataLakeAnalyticsAccountProperties.
func (c CreateDataLakeAnalyticsAccountProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "computePolicies", c.ComputePolicies)
	populate(objectMap, "dataLakeStoreAccounts", c.DataLakeStoreAccounts)
	populate(objectMap, "defaultDataLakeStoreAccount", c.DefaultDataLakeStoreAccount)
	populate(objectMap, "firewallAllowAzureIps", c.FirewallAllowAzureIPs)
	populate(objectMap, "firewallRules", c.FirewallRules)
	populate(objectMap, "firewallState", c.FirewallState)
	populate(objectMap, "maxDegreeOfParallelism", c.MaxDegreeOfParallelism)
	populate(objectMap, "maxDegreeOfParallelismPerJob", c.MaxDegreeOfParallelismPerJob)
	populate(objectMap, "maxJobCount", c.MaxJobCount)
	populate(objectMap, "minPriorityPerJob", c.MinPriorityPerJob)
	populate(objectMap, "newTier", c.NewTier)
	populate(objectMap, "queryStoreRetention", c.QueryStoreRetention)
	populate(objectMap, "storageAccounts", c.StorageAccounts)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type StorageContainerProperties.
func (s *StorageContainerProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "lastModifiedTime":
			err = unpopulateTimeRFC3339(val, "LastModifiedTime", &s.LastModifiedTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type UpdateComputePolicyParameters.
func (u UpdateComputePolicyParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", u.Properties)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type UpdateDataLakeAnalyticsAccountParameters.
func (u UpdateDataLakeAnalyticsAccountParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", u.Properties)
	populate(objectMap, "tags", u.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type UpdateDataLakeAnalyticsAccountProperties.
func (u UpdateDataLakeAnalyticsAccountProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "computePolicies", u.ComputePolicies)
	populate(objectMap, "dataLakeStoreAccounts", u.DataLakeStoreAccounts)
	populate(objectMap, "firewallAllowAzureIps", u.FirewallAllowAzureIPs)
	populate(objectMap, "firewallRules", u.FirewallRules)
	populate(objectMap, "firewallState", u.FirewallState)
	populate(objectMap, "maxDegreeOfParallelism", u.MaxDegreeOfParallelism)
	populate(objectMap, "maxDegreeOfParallelismPerJob", u.MaxDegreeOfParallelismPerJob)
	populate(objectMap, "maxJobCount", u.MaxJobCount)
	populate(objectMap, "minPriorityPerJob", u.MinPriorityPerJob)
	populate(objectMap, "newTier", u.NewTier)
	populate(objectMap, "queryStoreRetention", u.QueryStoreRetention)
	populate(objectMap, "storageAccounts", u.StorageAccounts)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type UpdateFirewallRuleParameters.
func (u UpdateFirewallRuleParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", u.Properties)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type UpdateStorageAccountParameters.
func (u UpdateStorageAccountParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", u.Properties)
	return json.Marshal(objectMap)
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v interface{}) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}