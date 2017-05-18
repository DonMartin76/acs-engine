package devtestlabs

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
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// ServiceRunnersClient is the the DevTest Labs Client.
type ServiceRunnersClient struct {
	ManagementClient
}

// NewServiceRunnersClient creates an instance of the ServiceRunnersClient
// client.
func NewServiceRunnersClient(subscriptionID string) ServiceRunnersClient {
	return NewServiceRunnersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewServiceRunnersClientWithBaseURI creates an instance of the
// ServiceRunnersClient client.
func NewServiceRunnersClientWithBaseURI(baseURI string, subscriptionID string) ServiceRunnersClient {
	return ServiceRunnersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or replace an existing Service runner.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. name is the name of the service runner. serviceRunner is a
// container for a managed identity to execute DevTest lab services.
func (client ServiceRunnersClient) CreateOrUpdate(resourceGroupName string, labName string, name string, serviceRunner ServiceRunner) (result ServiceRunner, err error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, labName, name, serviceRunner)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.ServiceRunnersClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "devtestlabs.ServiceRunnersClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.ServiceRunnersClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ServiceRunnersClient) CreateOrUpdatePreparer(resourceGroupName string, labName string, name string, serviceRunner ServiceRunner) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/servicerunners/{name}", pathParameters),
		autorest.WithJSON(serviceRunner),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceRunnersClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ServiceRunnersClient) CreateOrUpdateResponder(resp *http.Response) (result ServiceRunner, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete service runner.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. name is the name of the service runner.
func (client ServiceRunnersClient) Delete(resourceGroupName string, labName string, name string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName, labName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.ServiceRunnersClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "devtestlabs.ServiceRunnersClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.ServiceRunnersClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ServiceRunnersClient) DeletePreparer(resourceGroupName string, labName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/servicerunners/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceRunnersClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ServiceRunnersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get service runner.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. name is the name of the service runner.
func (client ServiceRunnersClient) Get(resourceGroupName string, labName string, name string) (result ServiceRunner, err error) {
	req, err := client.GetPreparer(resourceGroupName, labName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.ServiceRunnersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "devtestlabs.ServiceRunnersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.ServiceRunnersClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ServiceRunnersClient) GetPreparer(resourceGroupName string, labName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/servicerunners/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceRunnersClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ServiceRunnersClient) GetResponder(resp *http.Response) (result ServiceRunner, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list service runners in a given lab.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. filter is the filter to apply to the operation. top is the maximum
// number of resources to return from the operation. orderby is the ordering
// expression for the results, using OData notation.
func (client ServiceRunnersClient) List(resourceGroupName string, labName string, filter string, top *int32, orderby string) (result ResponseWithContinuationServiceRunner, err error) {
	req, err := client.ListPreparer(resourceGroupName, labName, filter, top, orderby)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.ServiceRunnersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "devtestlabs.ServiceRunnersClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.ServiceRunnersClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ServiceRunnersClient) ListPreparer(resourceGroupName string, labName string, filter string, top *int32, orderby string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/servicerunners", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceRunnersClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ServiceRunnersClient) ListResponder(resp *http.Response) (result ResponseWithContinuationServiceRunner, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client ServiceRunnersClient) ListNextResults(lastResults ResponseWithContinuationServiceRunner) (result ResponseWithContinuationServiceRunner, err error) {
	req, err := lastResults.ResponseWithContinuationServiceRunnerPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.ServiceRunnersClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "devtestlabs.ServiceRunnersClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.ServiceRunnersClient", "List", resp, "Failure responding to next results request")
	}

	return
}
