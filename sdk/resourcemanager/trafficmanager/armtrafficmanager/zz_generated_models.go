//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtrafficmanager

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"time"
)

// CheckTrafficManagerRelativeDNSNameAvailabilityParameters - Parameters supplied to check Traffic Manager name operation.
type CheckTrafficManagerRelativeDNSNameAvailabilityParameters struct {
	// The name of the resource.
	Name *string `json:"name,omitempty"`

	// The type of the resource.
	Type *string `json:"type,omitempty"`
}

// CloudError - An error returned by the Azure Resource Manager
// Implements the error and azcore.HTTPResponse interfaces.
type CloudError struct {
	raw string
	// The content of the error.
	InnerError *CloudErrorBody `json:"error,omitempty"`
}

// Error implements the error interface for type CloudError.
// The contents of the error text are not contractual and subject to change.
func (e CloudError) Error() string {
	return e.raw
}

// CloudErrorBody - The content of an error returned by the Azure Resource Manager
type CloudErrorBody struct {
	// Error code
	Code *string `json:"code,omitempty"`

	// Error details
	Details []*CloudErrorBody `json:"details,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	// Error target
	Target *string `json:"target,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type CloudErrorBody.
func (c CloudErrorBody) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", c.Code)
	populate(objectMap, "details", c.Details)
	populate(objectMap, "message", c.Message)
	populate(objectMap, "target", c.Target)
	return json.Marshal(objectMap)
}

// DNSConfig - Class containing DNS settings in a Traffic Manager profile.
type DNSConfig struct {
	// The relative DNS name provided by this Traffic Manager profile. This value is combined with the DNS domain name used by Azure Traffic Manager to form
	// the fully-qualified domain name (FQDN) of the
	// profile.
	RelativeName *string `json:"relativeName,omitempty"`

	// The DNS Time-To-Live (TTL), in seconds. This informs the local DNS resolvers and DNS clients how long to cache DNS responses provided by this Traffic
	// Manager profile.
	TTL *int64 `json:"ttl,omitempty"`

	// READ-ONLY; The fully-qualified domain name (FQDN) of the Traffic Manager profile. This is formed from the concatenation of the RelativeName with the
	// DNS domain used by Azure Traffic Manager.
	Fqdn *string `json:"fqdn,omitempty" azure:"ro"`
}

// DeleteOperationResult - The result of the request or operation.
type DeleteOperationResult struct {
	// READ-ONLY; The result of the operation or request.
	OperationResult *bool `json:"boolean,omitempty" azure:"ro"`
}

// Endpoint - Class representing a Traffic Manager endpoint.
type Endpoint struct {
	ProxyResource
	// The properties of the Traffic Manager endpoint.
	Properties *EndpointProperties `json:"properties,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type Endpoint.
func (e Endpoint) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	e.ProxyResource.marshalInternal(objectMap)
	populate(objectMap, "properties", e.Properties)
	return json.Marshal(objectMap)
}

// EndpointProperties - Class representing a Traffic Manager endpoint properties.
type EndpointProperties struct {
	// List of custom headers.
	CustomHeaders []*EndpointPropertiesCustomHeadersItem `json:"customHeaders,omitempty"`

	// Specifies the location of the external or nested endpoints when using the 'Performance' traffic routing method.
	EndpointLocation *string `json:"endpointLocation,omitempty"`

	// The monitoring status of the endpoint.
	EndpointMonitorStatus *EndpointMonitorStatus `json:"endpointMonitorStatus,omitempty"`

	// The status of the endpoint. If the endpoint is Enabled, it is probed for endpoint health and is included in the traffic routing method.
	EndpointStatus *EndpointStatus `json:"endpointStatus,omitempty"`

	// The list of countries/regions mapped to this endpoint when using the 'Geographic' traffic routing method. Please consult Traffic Manager Geographic documentation
	// for a full list of accepted values.
	GeoMapping []*string `json:"geoMapping,omitempty"`

	// The minimum number of endpoints that must be available in the child profile in order for the parent profile to be considered available. Only applicable
	// to endpoint of type 'NestedEndpoints'.
	MinChildEndpoints *int64 `json:"minChildEndpoints,omitempty"`

	// The minimum number of IPv4 (DNS record type A) endpoints that must be available in the child profile in order for the parent profile to be considered
	// available. Only applicable to endpoint of type
	// 'NestedEndpoints'.
	MinChildEndpointsIPv4 *int64 `json:"minChildEndpointsIPv4,omitempty"`

	// The minimum number of IPv6 (DNS record type AAAA) endpoints that must be available in the child profile in order for the parent profile to be considered
	// available. Only applicable to endpoint of type
	// 'NestedEndpoints'.
	MinChildEndpointsIPv6 *int64 `json:"minChildEndpointsIPv6,omitempty"`

	// The priority of this endpoint when using the 'Priority' traffic routing method. Possible values are from 1 to 1000, lower values represent higher priority.
	// This is an optional parameter. If specified,
	// it must be specified on all endpoints, and no two endpoints can share the same priority value.
	Priority *int64 `json:"priority,omitempty"`

	// The list of subnets, IP addresses, and/or address ranges mapped to this endpoint when using the 'Subnet' traffic routing method. An empty list will match
	// all ranges not covered by other endpoints.
	Subnets []*EndpointPropertiesSubnetsItem `json:"subnets,omitempty"`

	// The fully-qualified DNS name or IP address of the endpoint. Traffic Manager returns this value in DNS responses to direct traffic to this endpoint.
	Target *string `json:"target,omitempty"`

	// The Azure Resource URI of the of the endpoint. Not applicable to endpoints of type 'ExternalEndpoints'.
	TargetResourceID *string `json:"targetResourceId,omitempty"`

	// The weight of this endpoint when using the 'Weighted' traffic routing method. Possible values are from 1 to 1000.
	Weight *int64 `json:"weight,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type EndpointProperties.
func (e EndpointProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "customHeaders", e.CustomHeaders)
	populate(objectMap, "endpointLocation", e.EndpointLocation)
	populate(objectMap, "endpointMonitorStatus", e.EndpointMonitorStatus)
	populate(objectMap, "endpointStatus", e.EndpointStatus)
	populate(objectMap, "geoMapping", e.GeoMapping)
	populate(objectMap, "minChildEndpoints", e.MinChildEndpoints)
	populate(objectMap, "minChildEndpointsIPv4", e.MinChildEndpointsIPv4)
	populate(objectMap, "minChildEndpointsIPv6", e.MinChildEndpointsIPv6)
	populate(objectMap, "priority", e.Priority)
	populate(objectMap, "subnets", e.Subnets)
	populate(objectMap, "target", e.Target)
	populate(objectMap, "targetResourceId", e.TargetResourceID)
	populate(objectMap, "weight", e.Weight)
	return json.Marshal(objectMap)
}

// EndpointPropertiesCustomHeadersItem - Custom header name and value.
type EndpointPropertiesCustomHeadersItem struct {
	// Header name.
	Name *string `json:"name,omitempty"`

	// Header value.
	Value *string `json:"value,omitempty"`
}

// EndpointPropertiesSubnetsItem - Subnet first address, scope, and/or last address.
type EndpointPropertiesSubnetsItem struct {
	// First address in the subnet.
	First *string `json:"first,omitempty"`

	// Last address in the subnet.
	Last *string `json:"last,omitempty"`

	// Block size (number of leading bits in the subnet mask).
	Scope *int32 `json:"scope,omitempty"`
}

// EndpointsCreateOrUpdateOptions contains the optional parameters for the Endpoints.CreateOrUpdate method.
type EndpointsCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// EndpointsDeleteOptions contains the optional parameters for the Endpoints.Delete method.
type EndpointsDeleteOptions struct {
	// placeholder for future optional parameters
}

// EndpointsGetOptions contains the optional parameters for the Endpoints.Get method.
type EndpointsGetOptions struct {
	// placeholder for future optional parameters
}

// EndpointsUpdateOptions contains the optional parameters for the Endpoints.Update method.
type EndpointsUpdateOptions struct {
	// placeholder for future optional parameters
}

// GeographicHierarchiesGetDefaultOptions contains the optional parameters for the GeographicHierarchies.GetDefault method.
type GeographicHierarchiesGetDefaultOptions struct {
	// placeholder for future optional parameters
}

// GeographicHierarchyProperties - Class representing the properties of the Geographic hierarchy used with the Geographic traffic routing method.
type GeographicHierarchyProperties struct {
	// The region at the root of the hierarchy from all the regions in the hierarchy can be retrieved.
	GeographicHierarchy *Region `json:"geographicHierarchy,omitempty"`
}

// HeatMapEndpoint - Class which is a sparse representation of a Traffic Manager endpoint.
type HeatMapEndpoint struct {
	// A number uniquely identifying this endpoint in query experiences.
	EndpointID *int32 `json:"endpointId,omitempty"`

	// The ARM Resource ID of this Traffic Manager endpoint.
	ResourceID *string `json:"resourceId,omitempty"`
}

// HeatMapGetOptions contains the optional parameters for the HeatMap.Get method.
type HeatMapGetOptions struct {
	// The bottom right latitude,longitude pair of the rectangular viewport to query for.
	BotRight []float64
	// The top left latitude,longitude pair of the rectangular viewport to query for.
	TopLeft []float64
}

// HeatMapModel - Class representing a Traffic Manager HeatMap.
type HeatMapModel struct {
	ProxyResource
	// The properties of the Traffic Manager HeatMap.
	Properties *HeatMapProperties `json:"properties,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type HeatMapModel.
func (h HeatMapModel) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	h.ProxyResource.marshalInternal(objectMap)
	populate(objectMap, "properties", h.Properties)
	return json.Marshal(objectMap)
}

// HeatMapProperties - Class representing a Traffic Manager HeatMap properties.
type HeatMapProperties struct {
	// The ending of the time window for this HeatMap, exclusive.
	EndTime *time.Time `json:"endTime,omitempty"`

	// The endpoints used in this HeatMap calculation.
	Endpoints []*HeatMapEndpoint `json:"endpoints,omitempty"`

	// The beginning of the time window for this HeatMap, inclusive.
	StartTime *time.Time `json:"startTime,omitempty"`

	// The traffic flows produced in this HeatMap calculation.
	TrafficFlows []*TrafficFlow `json:"trafficFlows,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type HeatMapProperties.
func (h HeatMapProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "endTime", h.EndTime)
	populate(objectMap, "endpoints", h.Endpoints)
	populateTimeRFC3339(objectMap, "startTime", h.StartTime)
	populate(objectMap, "trafficFlows", h.TrafficFlows)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type HeatMapProperties.
func (h *HeatMapProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "endTime":
			err = unpopulateTimeRFC3339(val, &h.EndTime)
			delete(rawMsg, key)
		case "endpoints":
			err = unpopulate(val, &h.Endpoints)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeRFC3339(val, &h.StartTime)
			delete(rawMsg, key)
		case "trafficFlows":
			err = unpopulate(val, &h.TrafficFlows)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MonitorConfig - Class containing endpoint monitoring settings in a Traffic Manager profile.
type MonitorConfig struct {
	// List of custom headers.
	CustomHeaders []*MonitorConfigCustomHeadersItem `json:"customHeaders,omitempty"`

	// List of expected status code ranges.
	ExpectedStatusCodeRanges []*MonitorConfigExpectedStatusCodeRangesItem `json:"expectedStatusCodeRanges,omitempty"`

	// The monitor interval for endpoints in this profile. This is the interval at which Traffic Manager will check the health of each endpoint in this profile.
	IntervalInSeconds *int64 `json:"intervalInSeconds,omitempty"`

	// The path relative to the endpoint domain name used to probe for endpoint health.
	Path *string `json:"path,omitempty"`

	// The TCP port used to probe for endpoint health.
	Port *int64 `json:"port,omitempty"`

	// The profile-level monitoring status of the Traffic Manager profile.
	ProfileMonitorStatus *ProfileMonitorStatus `json:"profileMonitorStatus,omitempty"`

	// The protocol (HTTP, HTTPS or TCP) used to probe for endpoint health.
	Protocol *MonitorProtocol `json:"protocol,omitempty"`

	// The monitor timeout for endpoints in this profile. This is the time that Traffic Manager allows endpoints in this profile to response to the health check.
	TimeoutInSeconds *int64 `json:"timeoutInSeconds,omitempty"`

	// The number of consecutive failed health check that Traffic Manager tolerates before declaring an endpoint in this profile Degraded after the next failed
	// health check.
	ToleratedNumberOfFailures *int64 `json:"toleratedNumberOfFailures,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type MonitorConfig.
func (m MonitorConfig) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "customHeaders", m.CustomHeaders)
	populate(objectMap, "expectedStatusCodeRanges", m.ExpectedStatusCodeRanges)
	populate(objectMap, "intervalInSeconds", m.IntervalInSeconds)
	populate(objectMap, "path", m.Path)
	populate(objectMap, "port", m.Port)
	populate(objectMap, "profileMonitorStatus", m.ProfileMonitorStatus)
	populate(objectMap, "protocol", m.Protocol)
	populate(objectMap, "timeoutInSeconds", m.TimeoutInSeconds)
	populate(objectMap, "toleratedNumberOfFailures", m.ToleratedNumberOfFailures)
	return json.Marshal(objectMap)
}

// MonitorConfigCustomHeadersItem - Custom header name and value.
type MonitorConfigCustomHeadersItem struct {
	// Header name.
	Name *string `json:"name,omitempty"`

	// Header value.
	Value *string `json:"value,omitempty"`
}

// MonitorConfigExpectedStatusCodeRangesItem - Min and max value of a status code range.
type MonitorConfigExpectedStatusCodeRangesItem struct {
	// Max status code.
	Max *int32 `json:"max,omitempty"`

	// Min status code.
	Min *int32 `json:"min,omitempty"`
}

// Profile - Class representing a Traffic Manager profile.
type Profile struct {
	TrackedResource
	// The properties of the Traffic Manager profile.
	Properties *ProfileProperties `json:"properties,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type Profile.
func (p Profile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	p.TrackedResource.marshalInternal(objectMap)
	populate(objectMap, "properties", p.Properties)
	return json.Marshal(objectMap)
}

// ProfileListResult - The list Traffic Manager profiles operation response.
type ProfileListResult struct {
	// Gets the list of Traffic manager profiles.
	Value []*Profile `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ProfileListResult.
func (p ProfileListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// ProfileProperties - Class representing the Traffic Manager profile properties.
type ProfileProperties struct {
	// The list of allowed endpoint record types.
	AllowedEndpointRecordTypes []*AllowedEndpointRecordType `json:"allowedEndpointRecordTypes,omitempty"`

	// The DNS settings of the Traffic Manager profile.
	DNSConfig *DNSConfig `json:"dnsConfig,omitempty"`

	// The list of endpoints in the Traffic Manager profile.
	Endpoints []*Endpoint `json:"endpoints,omitempty"`

	// Maximum number of endpoints to be returned for MultiValue routing type.
	MaxReturn *int64 `json:"maxReturn,omitempty"`

	// The endpoint monitoring settings of the Traffic Manager profile.
	MonitorConfig *MonitorConfig `json:"monitorConfig,omitempty"`

	// The status of the Traffic Manager profile.
	ProfileStatus *ProfileStatus `json:"profileStatus,omitempty"`

	// The traffic routing method of the Traffic Manager profile.
	TrafficRoutingMethod *TrafficRoutingMethod `json:"trafficRoutingMethod,omitempty"`

	// Indicates whether Traffic View is 'Enabled' or 'Disabled' for the Traffic Manager profile. Null, indicates 'Disabled'. Enabling this feature will increase
	// the cost of the Traffic Manage profile.
	TrafficViewEnrollmentStatus *TrafficViewEnrollmentStatus `json:"trafficViewEnrollmentStatus,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ProfileProperties.
func (p ProfileProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "allowedEndpointRecordTypes", p.AllowedEndpointRecordTypes)
	populate(objectMap, "dnsConfig", p.DNSConfig)
	populate(objectMap, "endpoints", p.Endpoints)
	populate(objectMap, "maxReturn", p.MaxReturn)
	populate(objectMap, "monitorConfig", p.MonitorConfig)
	populate(objectMap, "profileStatus", p.ProfileStatus)
	populate(objectMap, "trafficRoutingMethod", p.TrafficRoutingMethod)
	populate(objectMap, "trafficViewEnrollmentStatus", p.TrafficViewEnrollmentStatus)
	return json.Marshal(objectMap)
}

// ProfilesCheckTrafficManagerRelativeDNSNameAvailabilityOptions contains the optional parameters for the Profiles.CheckTrafficManagerRelativeDNSNameAvailability
// method.
type ProfilesCheckTrafficManagerRelativeDNSNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// ProfilesCreateOrUpdateOptions contains the optional parameters for the Profiles.CreateOrUpdate method.
type ProfilesCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ProfilesDeleteOptions contains the optional parameters for the Profiles.Delete method.
type ProfilesDeleteOptions struct {
	// placeholder for future optional parameters
}

// ProfilesGetOptions contains the optional parameters for the Profiles.Get method.
type ProfilesGetOptions struct {
	// placeholder for future optional parameters
}

// ProfilesListByResourceGroupOptions contains the optional parameters for the Profiles.ListByResourceGroup method.
type ProfilesListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// ProfilesListBySubscriptionOptions contains the optional parameters for the Profiles.ListBySubscription method.
type ProfilesListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// ProfilesUpdateOptions contains the optional parameters for the Profiles.Update method.
type ProfilesUpdateOptions struct {
	// placeholder for future optional parameters
}

// ProxyResource - The resource model definition for a ARM proxy resource. It will have everything other than required location and tags
type ProxyResource struct {
	Resource
}

func (p ProxyResource) marshalInternal(objectMap map[string]interface{}) {
	p.Resource.marshalInternal(objectMap)
}

// QueryExperience - Class representing a Traffic Manager HeatMap query experience properties.
type QueryExperience struct {
	// REQUIRED; The id of the endpoint from the 'endpoints' array which these queries were routed to.
	EndpointID *int32 `json:"endpointId,omitempty"`

	// REQUIRED; The number of queries originating from this location.
	QueryCount *int32 `json:"queryCount,omitempty"`

	// The latency experienced by queries originating from this location.
	Latency *float64 `json:"latency,omitempty"`
}

// Region - Class representing a region in the Geographic hierarchy used with the Geographic traffic routing method.
type Region struct {
	// The code of the region
	Code *string `json:"code,omitempty"`

	// The name of the region
	Name *string `json:"name,omitempty"`

	// The list of Regions grouped under this Region in the Geographic Hierarchy.
	Regions []*Region `json:"regions,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type Region.
func (r Region) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", r.Code)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "regions", r.Regions)
	return json.Marshal(objectMap)
}

// Resource - The core properties of ARM resources
type Resource struct {
	// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{resourceName}
	ID *string `json:"id,omitempty"`

	// The name of the resource
	Name *string `json:"name,omitempty"`

	// The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	r.marshalInternal(objectMap)
	return json.Marshal(objectMap)
}

func (r Resource) marshalInternal(objectMap map[string]interface{}) {
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "type", r.Type)
}

// TrackedResource - The resource model definition for a ARM tracked top level resource
type TrackedResource struct {
	Resource
	// The Azure Region where the resource lives
	Location *string `json:"location,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type TrackedResource.
func (t TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	t.marshalInternal(objectMap)
	return json.Marshal(objectMap)
}

func (t TrackedResource) marshalInternal(objectMap map[string]interface{}) {
	t.Resource.marshalInternal(objectMap)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "tags", t.Tags)
}

// TrafficFlow - Class representing a Traffic Manager HeatMap traffic flow properties.
type TrafficFlow struct {
	// The approximate latitude that these queries originated from.
	Latitude *float64 `json:"latitude,omitempty"`

	// The approximate longitude that these queries originated from.
	Longitude *float64 `json:"longitude,omitempty"`

	// The query experiences produced in this HeatMap calculation.
	QueryExperiences []*QueryExperience `json:"queryExperiences,omitempty"`

	// The IP address that this query experience originated from.
	SourceIP *string `json:"sourceIp,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type TrafficFlow.
func (t TrafficFlow) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "latitude", t.Latitude)
	populate(objectMap, "longitude", t.Longitude)
	populate(objectMap, "queryExperiences", t.QueryExperiences)
	populate(objectMap, "sourceIp", t.SourceIP)
	return json.Marshal(objectMap)
}

// TrafficManagerGeographicHierarchy - Class representing the Geographic hierarchy used with the Geographic traffic routing method.
type TrafficManagerGeographicHierarchy struct {
	ProxyResource
	// The properties of the Geographic Hierarchy resource.
	Properties *GeographicHierarchyProperties `json:"properties,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type TrafficManagerGeographicHierarchy.
func (t TrafficManagerGeographicHierarchy) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	t.ProxyResource.marshalInternal(objectMap)
	populate(objectMap, "properties", t.Properties)
	return json.Marshal(objectMap)
}

// TrafficManagerNameAvailability - Class representing a Traffic Manager Name Availability response.
type TrafficManagerNameAvailability struct {
	// Descriptive message that explains why the name is not available, when applicable.
	Message *string `json:"message,omitempty"`

	// The relative name.
	Name *string `json:"name,omitempty"`

	// Describes whether the relative name is available or not.
	NameAvailable *bool `json:"nameAvailable,omitempty"`

	// The reason why the name is not available, when applicable.
	Reason *string `json:"reason,omitempty"`

	// Traffic Manager profile resource type.
	Type *string `json:"type,omitempty"`
}

// TrafficManagerUserMetricsKeysCreateOrUpdateOptions contains the optional parameters for the TrafficManagerUserMetricsKeys.CreateOrUpdate method.
type TrafficManagerUserMetricsKeysCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// TrafficManagerUserMetricsKeysDeleteOptions contains the optional parameters for the TrafficManagerUserMetricsKeys.Delete method.
type TrafficManagerUserMetricsKeysDeleteOptions struct {
	// placeholder for future optional parameters
}

// TrafficManagerUserMetricsKeysGetOptions contains the optional parameters for the TrafficManagerUserMetricsKeys.Get method.
type TrafficManagerUserMetricsKeysGetOptions struct {
	// placeholder for future optional parameters
}

// UserMetricsModel - Class representing Traffic Manager User Metrics.
type UserMetricsModel struct {
	ProxyResource
	// The properties of the Traffic Manager User Metrics.
	Properties *UserMetricsProperties `json:"properties,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type UserMetricsModel.
func (u UserMetricsModel) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	u.ProxyResource.marshalInternal(objectMap)
	populate(objectMap, "properties", u.Properties)
	return json.Marshal(objectMap)
}

// UserMetricsProperties - Class representing a Traffic Manager Real User Metrics key response.
type UserMetricsProperties struct {
	// The key returned by the User Metrics operation.
	Key *string `json:"key,omitempty"`
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

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}