package containerregistry

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// PasswordName enumerates the values for password name.
type PasswordName string

const (
	// Password specifies the password state for password name.
	Password PasswordName = "password"
	// Password2 specifies the password 2 state for password name.
	Password2 PasswordName = "password2"
)

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Creating specifies the creating state for provisioning state.
	Creating ProvisioningState = "Creating"
	// Succeeded specifies the succeeded state for provisioning state.
	Succeeded ProvisioningState = "Succeeded"
)

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// Basic specifies the basic state for sku tier.
	Basic SkuTier = "Basic"
)

// OperationDefinition is the definition of a container registry operation.
type OperationDefinition struct {
	Name    *string                     `json:"name,omitempty"`
	Display *OperationDisplayDefinition `json:"display,omitempty"`
}

// OperationDisplayDefinition is the display information for a container
// registry operation.
type OperationDisplayDefinition struct {
	Provider    *string `json:"provider,omitempty"`
	Resource    *string `json:"resource,omitempty"`
	Operation   *string `json:"operation,omitempty"`
	Description *string `json:"description,omitempty"`
}

// OperationListResult is the result of a request to list container registry
// operations.
type OperationListResult struct {
	autorest.Response `json:"-"`
	Value             *[]OperationDefinition `json:"value,omitempty"`
	NextLink          *string                `json:"nextLink,omitempty"`
}

// OperationListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client OperationListResult) OperationListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// RegenerateCredentialParameters is the parameters used to regenerate the
// login credential.
type RegenerateCredentialParameters struct {
	Name PasswordName `json:"name,omitempty"`
}

// Registry is an object that represents a container registry.
type Registry struct {
	autorest.Response   `json:"-"`
	ID                  *string             `json:"id,omitempty"`
	Name                *string             `json:"name,omitempty"`
	Type                *string             `json:"type,omitempty"`
	Location            *string             `json:"location,omitempty"`
	Tags                *map[string]*string `json:"tags,omitempty"`
	Sku                 *Sku                `json:"sku,omitempty"`
	*RegistryProperties `json:"properties,omitempty"`
}

// RegistryCreateParameters is the parameters for creating a container
// registry.
type RegistryCreateParameters struct {
	Tags                                *map[string]*string `json:"tags,omitempty"`
	Location                            *string             `json:"location,omitempty"`
	Sku                                 *Sku                `json:"sku,omitempty"`
	*RegistryPropertiesCreateParameters `json:"properties,omitempty"`
}

// RegistryListCredentialsResult is the response from the ListCredentials
// operation.
type RegistryListCredentialsResult struct {
	autorest.Response `json:"-"`
	Username          *string             `json:"username,omitempty"`
	Passwords         *[]RegistryPassword `json:"passwords,omitempty"`
}

// RegistryListResult is the result of a request to list container registries.
type RegistryListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Registry `json:"value,omitempty"`
	NextLink          *string     `json:"nextLink,omitempty"`
}

// RegistryListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client RegistryListResult) RegistryListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// RegistryNameCheckRequest is a request to check whether a container registry
// name is available.
type RegistryNameCheckRequest struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// RegistryNameStatus is the result of a request to check the availability of a
// container registry name.
type RegistryNameStatus struct {
	autorest.Response `json:"-"`
	NameAvailable     *bool   `json:"nameAvailable,omitempty"`
	Reason            *string `json:"reason,omitempty"`
	Message           *string `json:"message,omitempty"`
}

// RegistryPassword is the login password for the container registry.
type RegistryPassword struct {
	Name  PasswordName `json:"name,omitempty"`
	Value *string      `json:"value,omitempty"`
}

// RegistryProperties is the properties of a container registry.
type RegistryProperties struct {
	LoginServer       *string                   `json:"loginServer,omitempty"`
	CreationDate      *date.Time                `json:"creationDate,omitempty"`
	ProvisioningState ProvisioningState         `json:"provisioningState,omitempty"`
	AdminUserEnabled  *bool                     `json:"adminUserEnabled,omitempty"`
	StorageAccount    *StorageAccountProperties `json:"storageAccount,omitempty"`
}

// RegistryPropertiesCreateParameters is the parameters for creating the
// properties of a container registry.
type RegistryPropertiesCreateParameters struct {
	AdminUserEnabled *bool                     `json:"adminUserEnabled,omitempty"`
	StorageAccount   *StorageAccountParameters `json:"storageAccount,omitempty"`
}

// RegistryPropertiesUpdateParameters is the parameters for updating the
// properties of a container registry.
type RegistryPropertiesUpdateParameters struct {
	AdminUserEnabled *bool                     `json:"adminUserEnabled,omitempty"`
	StorageAccount   *StorageAccountParameters `json:"storageAccount,omitempty"`
}

// RegistryUpdateParameters is the parameters for updating a container
// registry.
type RegistryUpdateParameters struct {
	Tags                                *map[string]*string `json:"tags,omitempty"`
	*RegistryPropertiesUpdateParameters `json:"properties,omitempty"`
}

// Resource is an Azure resource.
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}

// Sku is the SKU of a container registry.
type Sku struct {
	Name *string `json:"name,omitempty"`
	Tier SkuTier `json:"tier,omitempty"`
}

// StorageAccountParameters is the parameters of a storage account for a
// container registry.
type StorageAccountParameters struct {
	Name      *string `json:"name,omitempty"`
	AccessKey *string `json:"accessKey,omitempty"`
}

// StorageAccountProperties is the properties of a storage account for a
// container registry.
type StorageAccountProperties struct {
	Name *string `json:"name,omitempty"`
}
