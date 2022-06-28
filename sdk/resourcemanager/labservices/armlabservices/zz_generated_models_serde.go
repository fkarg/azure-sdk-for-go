//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlabservices

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type ImageProperties.
func (i ImageProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "author", i.Author)
	populate(objectMap, "availableRegions", i.AvailableRegions)
	populate(objectMap, "description", i.Description)
	populate(objectMap, "displayName", i.DisplayName)
	populate(objectMap, "enabledState", i.EnabledState)
	populate(objectMap, "iconUrl", i.IconURL)
	populate(objectMap, "osState", i.OSState)
	populate(objectMap, "osType", i.OSType)
	populate(objectMap, "offer", i.Offer)
	populate(objectMap, "plan", i.Plan)
	populate(objectMap, "provisioningState", i.ProvisioningState)
	populate(objectMap, "publisher", i.Publisher)
	populate(objectMap, "sku", i.SKU)
	populate(objectMap, "sharedGalleryId", i.SharedGalleryID)
	populate(objectMap, "termsStatus", i.TermsStatus)
	populate(objectMap, "version", i.Version)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ImageUpdate.
func (i ImageUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", i.Properties)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Lab.
func (l Lab) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", l.ID)
	populate(objectMap, "location", l.Location)
	populate(objectMap, "name", l.Name)
	populate(objectMap, "properties", l.Properties)
	populate(objectMap, "systemData", l.SystemData)
	populate(objectMap, "tags", l.Tags)
	populate(objectMap, "type", l.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type LabPlan.
func (l LabPlan) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", l.ID)
	populate(objectMap, "location", l.Location)
	populate(objectMap, "name", l.Name)
	populate(objectMap, "properties", l.Properties)
	populate(objectMap, "systemData", l.SystemData)
	populate(objectMap, "tags", l.Tags)
	populate(objectMap, "type", l.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type LabPlanProperties.
func (l LabPlanProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "allowedRegions", l.AllowedRegions)
	populate(objectMap, "defaultAutoShutdownProfile", l.DefaultAutoShutdownProfile)
	populate(objectMap, "defaultConnectionProfile", l.DefaultConnectionProfile)
	populate(objectMap, "defaultNetworkProfile", l.DefaultNetworkProfile)
	populate(objectMap, "linkedLmsInstance", l.LinkedLmsInstance)
	populate(objectMap, "provisioningState", l.ProvisioningState)
	populate(objectMap, "sharedGalleryId", l.SharedGalleryID)
	populate(objectMap, "supportInfo", l.SupportInfo)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type LabPlanUpdate.
func (l LabPlanUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", l.Properties)
	populate(objectMap, "tags", l.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type LabPlanUpdateProperties.
func (l LabPlanUpdateProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "allowedRegions", l.AllowedRegions)
	populate(objectMap, "defaultAutoShutdownProfile", l.DefaultAutoShutdownProfile)
	populate(objectMap, "defaultConnectionProfile", l.DefaultConnectionProfile)
	populate(objectMap, "defaultNetworkProfile", l.DefaultNetworkProfile)
	populate(objectMap, "linkedLmsInstance", l.LinkedLmsInstance)
	populate(objectMap, "sharedGalleryId", l.SharedGalleryID)
	populate(objectMap, "supportInfo", l.SupportInfo)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type LabUpdate.
func (l LabUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", l.Properties)
	populate(objectMap, "tags", l.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationResult.
func (o *OperationResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "endTime":
			err = unpopulateTimeRFC3339(val, "EndTime", &o.EndTime)
			delete(rawMsg, key)
		case "error":
			err = unpopulate(val, "Error", &o.Error)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &o.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &o.Name)
			delete(rawMsg, key)
		case "percentComplete":
			err = unpopulate(val, "PercentComplete", &o.PercentComplete)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeRFC3339(val, "StartTime", &o.StartTime)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &o.Status)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RecurrencePattern.
func (r RecurrencePattern) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateDateType(objectMap, "expirationDate", r.ExpirationDate)
	populate(objectMap, "frequency", r.Frequency)
	populate(objectMap, "interval", r.Interval)
	populate(objectMap, "weekDays", r.WeekDays)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RecurrencePattern.
func (r *RecurrencePattern) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "expirationDate":
			err = unpopulateDateType(val, "ExpirationDate", &r.ExpirationDate)
			delete(rawMsg, key)
		case "frequency":
			err = unpopulate(val, "Frequency", &r.Frequency)
			delete(rawMsg, key)
		case "interval":
			err = unpopulate(val, "Interval", &r.Interval)
			delete(rawMsg, key)
		case "weekDays":
			err = unpopulate(val, "WeekDays", &r.WeekDays)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ScheduleProperties.
func (s ScheduleProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "notes", s.Notes)
	populate(objectMap, "provisioningState", s.ProvisioningState)
	populate(objectMap, "recurrencePattern", s.RecurrencePattern)
	populateTimeRFC3339(objectMap, "startAt", s.StartAt)
	populateTimeRFC3339(objectMap, "stopAt", s.StopAt)
	populate(objectMap, "timeZoneId", s.TimeZoneID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ScheduleProperties.
func (s *ScheduleProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "notes":
			err = unpopulate(val, "Notes", &s.Notes)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &s.ProvisioningState)
			delete(rawMsg, key)
		case "recurrencePattern":
			err = unpopulate(val, "RecurrencePattern", &s.RecurrencePattern)
			delete(rawMsg, key)
		case "startAt":
			err = unpopulateTimeRFC3339(val, "StartAt", &s.StartAt)
			delete(rawMsg, key)
		case "stopAt":
			err = unpopulateTimeRFC3339(val, "StopAt", &s.StopAt)
			delete(rawMsg, key)
		case "timeZoneId":
			err = unpopulate(val, "TimeZoneID", &s.TimeZoneID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ScheduleUpdate.
func (s ScheduleUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", s.Properties)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ScheduleUpdateProperties.
func (s ScheduleUpdateProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "notes", s.Notes)
	populate(objectMap, "recurrencePattern", s.RecurrencePattern)
	populateTimeRFC3339(objectMap, "startAt", s.StartAt)
	populateTimeRFC3339(objectMap, "stopAt", s.StopAt)
	populate(objectMap, "timeZoneId", s.TimeZoneID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ScheduleUpdateProperties.
func (s *ScheduleUpdateProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "notes":
			err = unpopulate(val, "Notes", &s.Notes)
			delete(rawMsg, key)
		case "recurrencePattern":
			err = unpopulate(val, "RecurrencePattern", &s.RecurrencePattern)
			delete(rawMsg, key)
		case "startAt":
			err = unpopulateTimeRFC3339(val, "StartAt", &s.StartAt)
			delete(rawMsg, key)
		case "stopAt":
			err = unpopulateTimeRFC3339(val, "StopAt", &s.StopAt)
			delete(rawMsg, key)
		case "timeZoneId":
			err = unpopulate(val, "TimeZoneID", &s.TimeZoneID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
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

// MarshalJSON implements the json.Marshaller interface for type TrackedResourceUpdate.
func (t TrackedResourceUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", t.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type UserProperties.
func (u UserProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalUsageQuota", u.AdditionalUsageQuota)
	populate(objectMap, "displayName", u.DisplayName)
	populate(objectMap, "email", u.Email)
	populateTimeRFC3339(objectMap, "invitationSent", u.InvitationSent)
	populate(objectMap, "invitationState", u.InvitationState)
	populate(objectMap, "provisioningState", u.ProvisioningState)
	populate(objectMap, "registrationState", u.RegistrationState)
	populate(objectMap, "totalUsage", u.TotalUsage)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UserProperties.
func (u *UserProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "additionalUsageQuota":
			err = unpopulate(val, "AdditionalUsageQuota", &u.AdditionalUsageQuota)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, "DisplayName", &u.DisplayName)
			delete(rawMsg, key)
		case "email":
			err = unpopulate(val, "Email", &u.Email)
			delete(rawMsg, key)
		case "invitationSent":
			err = unpopulateTimeRFC3339(val, "InvitationSent", &u.InvitationSent)
			delete(rawMsg, key)
		case "invitationState":
			err = unpopulate(val, "InvitationState", &u.InvitationState)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &u.ProvisioningState)
			delete(rawMsg, key)
		case "registrationState":
			err = unpopulate(val, "RegistrationState", &u.RegistrationState)
			delete(rawMsg, key)
		case "totalUsage":
			err = unpopulate(val, "TotalUsage", &u.TotalUsage)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", u, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type UserUpdate.
func (u UserUpdate) MarshalJSON() ([]byte, error) {
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