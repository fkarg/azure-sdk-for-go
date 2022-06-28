//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridnetwork

import "encoding/json"

func unmarshalDevicePropertiesFormatClassification(rawMsg json.RawMessage) (DevicePropertiesFormatClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DevicePropertiesFormatClassification
	switch m["deviceType"] {
	case string(DeviceTypeAzureStackEdge):
		b = &AzureStackEdgeFormat{}
	default:
		b = &DevicePropertiesFormat{}
	}
	return b, json.Unmarshal(rawMsg, b)
}