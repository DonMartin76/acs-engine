package sql

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

// RecommendedElasticPoolsClient is the the Azure SQL Database management API
// provides a RESTful set of web services that interact with Azure SQL Database
// services to manage your databases. The API enables you to create, retrieve,
// update, and delete databases.
type RecommendedElasticPoolsClient struct {
	ManagementClient
}

// NewRecommendedElasticPoolsClient creates an instance of the
// RecommendedElasticPoolsClient client.
func NewRecommendedElasticPoolsClient(subscriptionID string) RecommendedElasticPoolsClient {
	return NewRecommendedElasticPoolsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRecommendedElasticPoolsClientWithBaseURI creates an instance of the
// RecommendedElasticPoolsClient client.
func NewRecommendedElasticPoolsClientWithBaseURI(baseURI string, subscriptionID string) RecommendedElasticPoolsClient {
	return RecommendedElasticPoolsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a recommented elastic pool.
//
// resourceGroupName is the name of the resource group that contains the
// resource. You can obtain this value from the Azure Resource Manager API or
// the portal. serverName is the name of the server. recommendedElasticPoolName
// is the name of the recommended elastic pool to be retrieved.
func (client RecommendedElasticPoolsClient) Get(resourceGroupName string, serverName string, recommendedElasticPoolName string) (result RecommendedElasticPool, err error) {
	req, err := client.GetPreparer(resourceGroupName, serverName, recommendedElasticPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.RecommendedElasticPoolsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.RecommendedElasticPoolsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.RecommendedElasticPoolsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RecommendedElasticPoolsClient) GetPreparer(resourceGroupName string, serverName string, recommendedElasticPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"recommendedElasticPoolName": autorest.Encode("path", recommendedElasticPoolName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"serverName":                 autorest.Encode("path", serverName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/recommendedElasticPools/{recommendedElasticPoolName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendedElasticPoolsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RecommendedElasticPoolsClient) GetResponder(resp *http.Response) (result RecommendedElasticPool, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDatabases gets a database inside of a recommented elastic pool.
//
// resourceGroupName is the name of the resource group that contains the
// resource. You can obtain this value from the Azure Resource Manager API or
// the portal. serverName is the name of the server. recommendedElasticPoolName
// is the name of the elastic pool to be retrieved. databaseName is the name of
// the database to be retrieved.
func (client RecommendedElasticPoolsClient) GetDatabases(resourceGroupName string, serverName string, recommendedElasticPoolName string, databaseName string) (result Database, err error) {
	req, err := client.GetDatabasesPreparer(resourceGroupName, serverName, recommendedElasticPoolName, databaseName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.RecommendedElasticPoolsClient", "GetDatabases", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDatabasesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.RecommendedElasticPoolsClient", "GetDatabases", resp, "Failure sending request")
		return
	}

	result, err = client.GetDatabasesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.RecommendedElasticPoolsClient", "GetDatabases", resp, "Failure responding to request")
	}

	return
}

// GetDatabasesPreparer prepares the GetDatabases request.
func (client RecommendedElasticPoolsClient) GetDatabasesPreparer(resourceGroupName string, serverName string, recommendedElasticPoolName string, databaseName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":               autorest.Encode("path", databaseName),
		"recommendedElasticPoolName": autorest.Encode("path", recommendedElasticPoolName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"serverName":                 autorest.Encode("path", serverName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/recommendedElasticPools/{recommendedElasticPoolName}/databases/{databaseName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetDatabasesSender sends the GetDatabases request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendedElasticPoolsClient) GetDatabasesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetDatabasesResponder handles the response to the GetDatabases request. The method always
// closes the http.Response Body.
func (client RecommendedElasticPoolsClient) GetDatabasesResponder(resp *http.Response) (result Database, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByServer returns recommended elastic pools.
//
// resourceGroupName is the name of the resource group that contains the
// resource. You can obtain this value from the Azure Resource Manager API or
// the portal. serverName is the name of the server.
func (client RecommendedElasticPoolsClient) ListByServer(resourceGroupName string, serverName string) (result RecommendedElasticPoolListResult, err error) {
	req, err := client.ListByServerPreparer(resourceGroupName, serverName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.RecommendedElasticPoolsClient", "ListByServer", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.RecommendedElasticPoolsClient", "ListByServer", resp, "Failure sending request")
		return
	}

	result, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.RecommendedElasticPoolsClient", "ListByServer", resp, "Failure responding to request")
	}

	return
}

// ListByServerPreparer prepares the ListByServer request.
func (client RecommendedElasticPoolsClient) ListByServerPreparer(resourceGroupName string, serverName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/recommendedElasticPools", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByServerSender sends the ListByServer request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendedElasticPoolsClient) ListByServerSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByServerResponder handles the response to the ListByServer request. The method always
// closes the http.Response Body.
func (client RecommendedElasticPoolsClient) ListByServerResponder(resp *http.Response) (result RecommendedElasticPoolListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListDatabases returns a list of databases inside a recommented elastic pool.
//
// resourceGroupName is the name of the resource group that contains the
// resource. You can obtain this value from the Azure Resource Manager API or
// the portal. serverName is the name of the server. recommendedElasticPoolName
// is the name of the recommended elastic pool to be retrieved.
func (client RecommendedElasticPoolsClient) ListDatabases(resourceGroupName string, serverName string, recommendedElasticPoolName string) (result DatabaseListResult, err error) {
	req, err := client.ListDatabasesPreparer(resourceGroupName, serverName, recommendedElasticPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.RecommendedElasticPoolsClient", "ListDatabases", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListDatabasesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.RecommendedElasticPoolsClient", "ListDatabases", resp, "Failure sending request")
		return
	}

	result, err = client.ListDatabasesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.RecommendedElasticPoolsClient", "ListDatabases", resp, "Failure responding to request")
	}

	return
}

// ListDatabasesPreparer prepares the ListDatabases request.
func (client RecommendedElasticPoolsClient) ListDatabasesPreparer(resourceGroupName string, serverName string, recommendedElasticPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"recommendedElasticPoolName": autorest.Encode("path", recommendedElasticPoolName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"serverName":                 autorest.Encode("path", serverName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/recommendedElasticPools/{recommendedElasticPoolName}/databases", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListDatabasesSender sends the ListDatabases request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendedElasticPoolsClient) ListDatabasesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListDatabasesResponder handles the response to the ListDatabases request. The method always
// closes the http.Response Body.
func (client RecommendedElasticPoolsClient) ListDatabasesResponder(resp *http.Response) (result DatabaseListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListMetrics returns a recommented elastic pool metrics.
//
// resourceGroupName is the name of the resource group that contains the
// resource. You can obtain this value from the Azure Resource Manager API or
// the portal. serverName is the name of the server. recommendedElasticPoolName
// is the name of the recommended elastic pool to be retrieved.
func (client RecommendedElasticPoolsClient) ListMetrics(resourceGroupName string, serverName string, recommendedElasticPoolName string) (result RecommendedElasticPoolListMetricsResult, err error) {
	req, err := client.ListMetricsPreparer(resourceGroupName, serverName, recommendedElasticPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.RecommendedElasticPoolsClient", "ListMetrics", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListMetricsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.RecommendedElasticPoolsClient", "ListMetrics", resp, "Failure sending request")
		return
	}

	result, err = client.ListMetricsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.RecommendedElasticPoolsClient", "ListMetrics", resp, "Failure responding to request")
	}

	return
}

// ListMetricsPreparer prepares the ListMetrics request.
func (client RecommendedElasticPoolsClient) ListMetricsPreparer(resourceGroupName string, serverName string, recommendedElasticPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"recommendedElasticPoolName": autorest.Encode("path", recommendedElasticPoolName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"serverName":                 autorest.Encode("path", serverName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/recommendedElasticPools/{recommendedElasticPoolName}/metrics", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListMetricsSender sends the ListMetrics request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendedElasticPoolsClient) ListMetricsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListMetricsResponder handles the response to the ListMetrics request. The method always
// closes the http.Response Body.
func (client RecommendedElasticPoolsClient) ListMetricsResponder(resp *http.Response) (result RecommendedElasticPoolListMetricsResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
