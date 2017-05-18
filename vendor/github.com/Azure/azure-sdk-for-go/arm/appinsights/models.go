package appinsights

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

// ApplicationType enumerates the values for application type.
type ApplicationType string

const (
	// Other specifies the other state for application type.
	Other ApplicationType = "other"
	// Web specifies the web state for application type.
	Web ApplicationType = "web"
)

// FlowType enumerates the values for flow type.
type FlowType string

const (
	// Bluefield specifies the bluefield state for flow type.
	Bluefield FlowType = "Bluefield"
)

// RequestSource enumerates the values for request source.
type RequestSource string

const (
	// Rest specifies the rest state for request source.
	Rest RequestSource = "rest"
)

// WebTestKind enumerates the values for web test kind.
type WebTestKind string

const (
	// Multistep specifies the multistep state for web test kind.
	Multistep WebTestKind = "multistep"
	// Ping specifies the ping state for web test kind.
	Ping WebTestKind = "ping"
)

// ApplicationInsightsComponent is an Application Insights component
// definition.
type ApplicationInsightsComponent struct {
	autorest.Response                       `json:"-"`
	ID                                      *string             `json:"id,omitempty"`
	Name                                    *string             `json:"name,omitempty"`
	Type                                    *string             `json:"type,omitempty"`
	Location                                *string             `json:"location,omitempty"`
	Tags                                    *map[string]*string `json:"tags,omitempty"`
	Kind                                    *string             `json:"kind,omitempty"`
	*ApplicationInsightsComponentProperties `json:"properties,omitempty"`
}

// ApplicationInsightsComponentListResult is describes the list of Application
// Insights Resources.
type ApplicationInsightsComponentListResult struct {
	autorest.Response `json:"-"`
	Value             *[]ApplicationInsightsComponent `json:"value,omitempty"`
	NextLink          *string                         `json:"nextLink,omitempty"`
}

// ApplicationInsightsComponentListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ApplicationInsightsComponentListResult) ApplicationInsightsComponentListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ApplicationInsightsComponentProperties is properties that define an
// Application Insights component resource.
type ApplicationInsightsComponentProperties struct {
	ApplicationID      *string         `json:"ApplicationId,omitempty"`
	AppID              *string         `json:"AppId,omitempty"`
	ApplicationType    ApplicationType `json:"Application_Type,omitempty"`
	FlowType           FlowType        `json:"Flow_Type,omitempty"`
	RequestSource      RequestSource   `json:"Request_Source,omitempty"`
	InstrumentationKey *string         `json:"InstrumentationKey,omitempty"`
	CreationDate       *date.Time      `json:"CreationDate,omitempty"`
	TenantID           *string         `json:"TenantId,omitempty"`
	HockeyAppID        *string         `json:"HockeyAppId,omitempty"`
	HockeyAppToken     *string         `json:"HockeyAppToken,omitempty"`
	ProvisioningState  *string         `json:"provisioningState,omitempty"`
	SamplingPercentage *float64        `json:"SamplingPercentage,omitempty"`
}

// ErrorResponse is error reponse indicates Insights service is not able to
// process the incoming request. The reason is provided in the error message.
type ErrorResponse struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// Operation is cDN REST API operation
type Operation struct {
	Name    *string           `json:"name,omitempty"`
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay is the object that represents the operation.
type OperationDisplay struct {
	Provider  *string `json:"provider,omitempty"`
	Resource  *string `json:"resource,omitempty"`
	Operation *string `json:"operation,omitempty"`
}

// OperationListResult is result of the request to list CDN operations. It
// contains a list of operations and a URL link to get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Operation `json:"value,omitempty"`
	NextLink          *string      `json:"nextLink,omitempty"`
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

// Resource is an azure resource object
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}

// TagsResource is a container holding only the Tags for a resource, allowing
// the user to update the tags on a WebTest instance.
type TagsResource struct {
	Tags *map[string]*string `json:"tags,omitempty"`
}

// WebTest is an Application Insights web test definition.
type WebTest struct {
	autorest.Response  `json:"-"`
	ID                 *string             `json:"id,omitempty"`
	Name               *string             `json:"name,omitempty"`
	Type               *string             `json:"type,omitempty"`
	Location           *string             `json:"location,omitempty"`
	Tags               *map[string]*string `json:"tags,omitempty"`
	Kind               WebTestKind         `json:"kind,omitempty"`
	*WebTestProperties `json:"properties,omitempty"`
}

// WebTestGeolocation is geo-physical location to run a web test from. You must
// specify one or more locations for the test to run from.
type WebTestGeolocation struct {
	Location *string `json:"Id,omitempty"`
}

// WebTestListResult is a list of 0 or more Application Insights web test
// definitions.
type WebTestListResult struct {
	autorest.Response `json:"-"`
	Value             *[]WebTest `json:"value,omitempty"`
	NextLink          *string    `json:"nextLink,omitempty"`
}

// WebTestListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client WebTestListResult) WebTestListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// WebTestProperties is metadata describing a web test for an Azure resource.
type WebTestProperties struct {
	SyntheticMonitorID *string                         `json:"SyntheticMonitorId,omitempty"`
	WebTestName        *string                         `json:"Name,omitempty"`
	Description        *string                         `json:"Description,omitempty"`
	Enabled            *bool                           `json:"Enabled,omitempty"`
	Frequency          *int32                          `json:"Frequency,omitempty"`
	Timeout            *int32                          `json:"Timeout,omitempty"`
	WebTestKind        WebTestKind                     `json:"Kind,omitempty"`
	RetryEnabled       *bool                           `json:"RetryEnabled,omitempty"`
	Locations          *[]WebTestGeolocation           `json:"Locations,omitempty"`
	Configuration      *WebTestPropertiesConfiguration `json:"Configuration,omitempty"`
	ProvisioningState  *string                         `json:"provisioningState,omitempty"`
}

// WebTestPropertiesConfiguration is an XML configuration specification for a
// WebTest.
type WebTestPropertiesConfiguration struct {
	WebTest *string `json:"WebTest,omitempty"`
}
