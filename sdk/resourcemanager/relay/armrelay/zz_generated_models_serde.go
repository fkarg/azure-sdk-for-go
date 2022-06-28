//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrelay

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type AuthorizationRuleProperties.
func (a AuthorizationRuleProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "rights", a.Rights)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type HybridConnectionProperties.
func (h HybridConnectionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", h.CreatedAt)
	populate(objectMap, "listenerCount", h.ListenerCount)
	populate(objectMap, "requiresClientAuthorization", h.RequiresClientAuthorization)
	populateTimeRFC3339(objectMap, "updatedAt", h.UpdatedAt)
	populate(objectMap, "userMetadata", h.UserMetadata)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type HybridConnectionProperties.
func (h *HybridConnectionProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", h, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, "CreatedAt", &h.CreatedAt)
			delete(rawMsg, key)
		case "listenerCount":
			err = unpopulate(val, "ListenerCount", &h.ListenerCount)
			delete(rawMsg, key)
		case "requiresClientAuthorization":
			err = unpopulate(val, "RequiresClientAuthorization", &h.RequiresClientAuthorization)
			delete(rawMsg, key)
		case "updatedAt":
			err = unpopulateTimeRFC3339(val, "UpdatedAt", &h.UpdatedAt)
			delete(rawMsg, key)
		case "userMetadata":
			err = unpopulate(val, "UserMetadata", &h.UserMetadata)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", h, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Namespace.
func (n Namespace) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", n.ID)
	populate(objectMap, "location", n.Location)
	populate(objectMap, "name", n.Name)
	populate(objectMap, "properties", n.Properties)
	populate(objectMap, "sku", n.SKU)
	populate(objectMap, "systemData", n.SystemData)
	populate(objectMap, "tags", n.Tags)
	populate(objectMap, "type", n.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type NamespaceProperties.
func (n NamespaceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", n.CreatedAt)
	populate(objectMap, "metricId", n.MetricID)
	populate(objectMap, "privateEndpointConnections", n.PrivateEndpointConnections)
	populate(objectMap, "provisioningState", n.ProvisioningState)
	populate(objectMap, "publicNetworkAccess", n.PublicNetworkAccess)
	populate(objectMap, "serviceBusEndpoint", n.ServiceBusEndpoint)
	populate(objectMap, "status", n.Status)
	populateTimeRFC3339(objectMap, "updatedAt", n.UpdatedAt)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type NamespaceProperties.
func (n *NamespaceProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", n, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, "CreatedAt", &n.CreatedAt)
			delete(rawMsg, key)
		case "metricId":
			err = unpopulate(val, "MetricID", &n.MetricID)
			delete(rawMsg, key)
		case "privateEndpointConnections":
			err = unpopulate(val, "PrivateEndpointConnections", &n.PrivateEndpointConnections)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &n.ProvisioningState)
			delete(rawMsg, key)
		case "publicNetworkAccess":
			err = unpopulate(val, "PublicNetworkAccess", &n.PublicNetworkAccess)
			delete(rawMsg, key)
		case "serviceBusEndpoint":
			err = unpopulate(val, "ServiceBusEndpoint", &n.ServiceBusEndpoint)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &n.Status)
			delete(rawMsg, key)
		case "updatedAt":
			err = unpopulateTimeRFC3339(val, "UpdatedAt", &n.UpdatedAt)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", n, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type NetworkRuleSetProperties.
func (n NetworkRuleSetProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "defaultAction", n.DefaultAction)
	populate(objectMap, "ipRules", n.IPRules)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceNamespacePatch.
func (r ResourceNamespacePatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "tags", r.Tags)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, "CreatedAt", &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, "CreatedBy", &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, "CreatedByType", &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, "LastModifiedAt", &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, "LastModifiedBy", &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, "LastModifiedByType", &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TrackedResource.
func (t TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", t.ID)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "tags", t.Tags)
	populate(objectMap, "type", t.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type UpdateParameters.
func (u UpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", u.ID)
	populate(objectMap, "name", u.Name)
	populate(objectMap, "properties", u.Properties)
	populate(objectMap, "sku", u.SKU)
	populate(objectMap, "tags", u.Tags)
	populate(objectMap, "type", u.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type WcfRelayProperties.
func (w WcfRelayProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", w.CreatedAt)
	populate(objectMap, "isDynamic", w.IsDynamic)
	populate(objectMap, "listenerCount", w.ListenerCount)
	populate(objectMap, "relayType", w.RelayType)
	populate(objectMap, "requiresClientAuthorization", w.RequiresClientAuthorization)
	populate(objectMap, "requiresTransportSecurity", w.RequiresTransportSecurity)
	populateTimeRFC3339(objectMap, "updatedAt", w.UpdatedAt)
	populate(objectMap, "userMetadata", w.UserMetadata)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type WcfRelayProperties.
func (w *WcfRelayProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", w, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, "CreatedAt", &w.CreatedAt)
			delete(rawMsg, key)
		case "isDynamic":
			err = unpopulate(val, "IsDynamic", &w.IsDynamic)
			delete(rawMsg, key)
		case "listenerCount":
			err = unpopulate(val, "ListenerCount", &w.ListenerCount)
			delete(rawMsg, key)
		case "relayType":
			err = unpopulate(val, "RelayType", &w.RelayType)
			delete(rawMsg, key)
		case "requiresClientAuthorization":
			err = unpopulate(val, "RequiresClientAuthorization", &w.RequiresClientAuthorization)
			delete(rawMsg, key)
		case "requiresTransportSecurity":
			err = unpopulate(val, "RequiresTransportSecurity", &w.RequiresTransportSecurity)
			delete(rawMsg, key)
		case "updatedAt":
			err = unpopulateTimeRFC3339(val, "UpdatedAt", &w.UpdatedAt)
			delete(rawMsg, key)
		case "userMetadata":
			err = unpopulate(val, "UserMetadata", &w.UserMetadata)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", w, err)
		}
	}
	return nil
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