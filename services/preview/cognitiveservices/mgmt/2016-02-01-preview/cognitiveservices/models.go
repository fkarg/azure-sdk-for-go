package cognitiveservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/cognitiveservices/mgmt/2016-02-01-preview/cognitiveservices"

// Account cognitive Services Account is an Azure resource representing the provisioned account, its type,
// location and SKU.
type Account struct {
	autorest.Response `json:"-"`
	// Etag - Entity Tag
	Etag *string `json:"etag,omitempty"`
	// ID - The id of the created account
	ID *string `json:"id,omitempty"`
	// Kind - Type of cognitive service account.
	Kind *string `json:"kind,omitempty"`
	// Location - The location of the resource
	Location *string `json:"location,omitempty"`
	// Name - The name of the created account
	Name               *string `json:"name,omitempty"`
	*AccountProperties `json:"properties,omitempty"`
	Sku                *Sku `json:"sku,omitempty"`
	// Tags - Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags map[string]*string `json:"tags"`
	// Type - Resource type
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Account.
func (a Account) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if a.Etag != nil {
		objectMap["etag"] = a.Etag
	}
	if a.ID != nil {
		objectMap["id"] = a.ID
	}
	if a.Kind != nil {
		objectMap["kind"] = a.Kind
	}
	if a.Location != nil {
		objectMap["location"] = a.Location
	}
	if a.Name != nil {
		objectMap["name"] = a.Name
	}
	if a.AccountProperties != nil {
		objectMap["properties"] = a.AccountProperties
	}
	if a.Sku != nil {
		objectMap["sku"] = a.Sku
	}
	if a.Tags != nil {
		objectMap["tags"] = a.Tags
	}
	if a.Type != nil {
		objectMap["type"] = a.Type
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Account struct.
func (a *Account) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "etag":
			if v != nil {
				var etag string
				err = json.Unmarshal(*v, &etag)
				if err != nil {
					return err
				}
				a.Etag = &etag
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				a.ID = &ID
			}
		case "kind":
			if v != nil {
				var kind string
				err = json.Unmarshal(*v, &kind)
				if err != nil {
					return err
				}
				a.Kind = &kind
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				a.Location = &location
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				a.Name = &name
			}
		case "properties":
			if v != nil {
				var accountProperties AccountProperties
				err = json.Unmarshal(*v, &accountProperties)
				if err != nil {
					return err
				}
				a.AccountProperties = &accountProperties
			}
		case "sku":
			if v != nil {
				var sku Sku
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				a.Sku = &sku
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				a.Tags = tags
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				a.Type = &typeVar
			}
		}
	}

	return nil
}

// AccountCreateParameters the parameters to provide for the account.
type AccountCreateParameters struct {
	Sku *Sku `json:"sku,omitempty"`
	// Kind - Required. Indicates the type of cognitive service account. Possible values include: 'Academic', 'BingAutosuggest', 'BingSearch', 'BingSpeech', 'BingSpellCheck', 'ComputerVision', 'ContentModerator', 'Emotion', 'Face', 'LUIS', 'Recommendations', 'SpeakerRecognition', 'Speech', 'SpeechTranslation', 'TextAnalytics', 'TextTranslation', 'WebLM'
	Kind Kind `json:"kind,omitempty"`
	// Location - Required. Gets or sets the location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US, East US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is created, but if an identical geo region is specified on update the request will succeed.
	Location *string `json:"location,omitempty"`
	// Tags - Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags map[string]*string `json:"tags"`
	// Properties - Must exist in the request. Must not be null.
	Properties interface{} `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for AccountCreateParameters.
func (acp AccountCreateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if acp.Sku != nil {
		objectMap["sku"] = acp.Sku
	}
	if acp.Kind != "" {
		objectMap["kind"] = acp.Kind
	}
	if acp.Location != nil {
		objectMap["location"] = acp.Location
	}
	if acp.Tags != nil {
		objectMap["tags"] = acp.Tags
	}
	if acp.Properties != nil {
		objectMap["properties"] = acp.Properties
	}
	return json.Marshal(objectMap)
}

// AccountEnumerateSkusResult the list of cognitive services accounts operation response.
type AccountEnumerateSkusResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; Gets the list of Cognitive Services accounts and their properties.
	Value *[]ResourceAndSku `json:"value,omitempty"`
}

// MarshalJSON is the custom marshaler for AccountEnumerateSkusResult.
func (aesr AccountEnumerateSkusResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// AccountKeys the access keys for the cognitive services account.
type AccountKeys struct {
	autorest.Response `json:"-"`
	// Key1 - Gets the value of key 1.
	Key1 *string `json:"key1,omitempty"`
	// Key2 - Gets the value of key 2.
	Key2 *string `json:"key2,omitempty"`
}

// AccountListResult the list of cognitive services accounts operation response.
type AccountListResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; Gets the list of Cognitive Services accounts and their properties.
	Value *[]Account `json:"value,omitempty"`
}

// MarshalJSON is the custom marshaler for AccountListResult.
func (alr AccountListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// AccountProperties ...
type AccountProperties struct {
	// ProvisioningState - READ-ONLY; Gets the status of the cognitive services account at the time the operation was called. Possible values include: 'Creating', 'ResolvingDNS', 'Succeeded', 'Failed'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	// Endpoint - Endpoint of the created account
	Endpoint *string `json:"endpoint,omitempty"`
}

// MarshalJSON is the custom marshaler for AccountProperties.
func (ap AccountProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ap.Endpoint != nil {
		objectMap["endpoint"] = ap.Endpoint
	}
	return json.Marshal(objectMap)
}

// AccountUpdateParameters the parameters to provide for the account.
type AccountUpdateParameters struct {
	Sku *Sku `json:"sku,omitempty"`
	// Tags - Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for AccountUpdateParameters.
func (aup AccountUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if aup.Sku != nil {
		objectMap["sku"] = aup.Sku
	}
	if aup.Tags != nil {
		objectMap["tags"] = aup.Tags
	}
	return json.Marshal(objectMap)
}

// Error ...
type Error struct {
	Error *ErrorBody `json:"error,omitempty"`
}

// ErrorBody ...
type ErrorBody struct {
	// Code - error code
	Code *string `json:"code,omitempty"`
	// Message - error message
	Message *string `json:"message,omitempty"`
}

// RegenerateKeyParameters regenerate key parameters.
type RegenerateKeyParameters struct {
	// KeyName - key name to generate (Key1|Key2). Possible values include: 'Key1', 'Key2'
	KeyName KeyName `json:"keyName,omitempty"`
}

// ResourceAndSku ...
type ResourceAndSku struct {
	// ResourceType - Resource Namespace and Type
	ResourceType *string `json:"resourceType,omitempty"`
	Sku          *Sku    `json:"sku,omitempty"`
}

// Sku the SKU of the cognitive services account.
type Sku struct {
	// Name - Gets or sets the sku name. Required for account creation, optional for update. Possible values include: 'F0', 'P0', 'P1', 'P2', 'S0', 'S1', 'S2', 'S3', 'S4', 'S5', 'S6'
	Name SkuName `json:"name,omitempty"`
	// Tier - READ-ONLY; Gets the sku tier. This is based on the SKU name. Possible values include: 'Free', 'Standard', 'Premium'
	Tier SkuTier `json:"tier,omitempty"`
}

// MarshalJSON is the custom marshaler for Sku.
func (s Sku) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if s.Name != "" {
		objectMap["name"] = s.Name
	}
	return json.Marshal(objectMap)
}