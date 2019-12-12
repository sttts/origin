package datashare

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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// SharesClient is the creates a Microsoft.DataShare management client.
type SharesClient struct {
	BaseClient
}

// NewSharesClient creates an instance of the SharesClient client.
func NewSharesClient(subscriptionID string) SharesClient {
	return NewSharesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSharesClientWithBaseURI creates an instance of the SharesClient client.
func NewSharesClientWithBaseURI(baseURI string, subscriptionID string) SharesClient {
	return SharesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create create a share
// Parameters:
// resourceGroupName - the resource group name.
// accountName - the name of the share account.
// shareName - the name of the share.
// share - the share payload
func (client SharesClient) Create(ctx context.Context, resourceGroupName string, accountName string, shareName string, share Share) (result Share, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SharesClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, resourceGroupName, accountName, shareName, share)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.SharesClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datashare.SharesClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.SharesClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client SharesClient) CreatePreparer(ctx context.Context, resourceGroupName string, accountName string, shareName string, share Share) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"shareName":         autorest.Encode("path", shareName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}", pathParameters),
		autorest.WithJSON(share),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client SharesClient) CreateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client SharesClient) CreateResponder(resp *http.Response) (result Share, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a share
// Parameters:
// resourceGroupName - the resource group name.
// accountName - the name of the share account.
// shareName - the name of the share.
func (client SharesClient) Delete(ctx context.Context, resourceGroupName string, accountName string, shareName string) (result SharesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SharesClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, accountName, shareName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.SharesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.SharesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SharesClient) DeletePreparer(ctx context.Context, resourceGroupName string, accountName string, shareName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"shareName":         autorest.Encode("path", shareName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SharesClient) DeleteSender(req *http.Request) (future SharesDeleteFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SharesClient) DeleteResponder(resp *http.Response) (result OperationResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get a share
// Parameters:
// resourceGroupName - the resource group name.
// accountName - the name of the share account.
// shareName - the name of the share to retrieve.
func (client SharesClient) Get(ctx context.Context, resourceGroupName string, accountName string, shareName string) (result Share, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SharesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, accountName, shareName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.SharesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datashare.SharesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.SharesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SharesClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string, shareName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"shareName":         autorest.Encode("path", shareName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SharesClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SharesClient) GetResponder(resp *http.Response) (result Share, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByAccount list shares in an account
// Parameters:
// resourceGroupName - the resource group name.
// accountName - the name of the share account.
// skipToken - continuation Token
func (client SharesClient) ListByAccount(ctx context.Context, resourceGroupName string, accountName string, skipToken string) (result ShareListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SharesClient.ListByAccount")
		defer func() {
			sc := -1
			if result.sl.Response.Response != nil {
				sc = result.sl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByAccountNextResults
	req, err := client.ListByAccountPreparer(ctx, resourceGroupName, accountName, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.SharesClient", "ListByAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByAccountSender(req)
	if err != nil {
		result.sl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datashare.SharesClient", "ListByAccount", resp, "Failure sending request")
		return
	}

	result.sl, err = client.ListByAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.SharesClient", "ListByAccount", resp, "Failure responding to request")
	}

	return
}

// ListByAccountPreparer prepares the ListByAccount request.
func (client SharesClient) ListByAccountPreparer(ctx context.Context, resourceGroupName string, accountName string, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByAccountSender sends the ListByAccount request. The method will close the
// http.Response Body if it receives an error.
func (client SharesClient) ListByAccountSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByAccountResponder handles the response to the ListByAccount request. The method always
// closes the http.Response Body.
func (client SharesClient) ListByAccountResponder(resp *http.Response) (result ShareList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByAccountNextResults retrieves the next set of results, if any.
func (client SharesClient) listByAccountNextResults(ctx context.Context, lastResults ShareList) (result ShareList, err error) {
	req, err := lastResults.shareListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "datashare.SharesClient", "listByAccountNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "datashare.SharesClient", "listByAccountNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.SharesClient", "listByAccountNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByAccountComplete enumerates all values, automatically crossing page boundaries as required.
func (client SharesClient) ListByAccountComplete(ctx context.Context, resourceGroupName string, accountName string, skipToken string) (result ShareListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SharesClient.ListByAccount")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByAccount(ctx, resourceGroupName, accountName, skipToken)
	return
}

// ListSynchronizationDetails list synchronization details
// Parameters:
// resourceGroupName - the resource group name.
// accountName - the name of the share account.
// shareName - the name of the share.
// shareSynchronization - share Synchronization payload.
// skipToken - continuation token
func (client SharesClient) ListSynchronizationDetails(ctx context.Context, resourceGroupName string, accountName string, shareName string, shareSynchronization ShareSynchronization, skipToken string) (result SynchronizationDetailsListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SharesClient.ListSynchronizationDetails")
		defer func() {
			sc := -1
			if result.sdl.Response.Response != nil {
				sc = result.sdl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listSynchronizationDetailsNextResults
	req, err := client.ListSynchronizationDetailsPreparer(ctx, resourceGroupName, accountName, shareName, shareSynchronization, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.SharesClient", "ListSynchronizationDetails", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSynchronizationDetailsSender(req)
	if err != nil {
		result.sdl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datashare.SharesClient", "ListSynchronizationDetails", resp, "Failure sending request")
		return
	}

	result.sdl, err = client.ListSynchronizationDetailsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.SharesClient", "ListSynchronizationDetails", resp, "Failure responding to request")
	}

	return
}

// ListSynchronizationDetailsPreparer prepares the ListSynchronizationDetails request.
func (client SharesClient) ListSynchronizationDetailsPreparer(ctx context.Context, resourceGroupName string, accountName string, shareName string, shareSynchronization ShareSynchronization, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"shareName":         autorest.Encode("path", shareName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	shareSynchronization.SynchronizationMode = ""
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/listSynchronizationDetails", pathParameters),
		autorest.WithJSON(shareSynchronization),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSynchronizationDetailsSender sends the ListSynchronizationDetails request. The method will close the
// http.Response Body if it receives an error.
func (client SharesClient) ListSynchronizationDetailsSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListSynchronizationDetailsResponder handles the response to the ListSynchronizationDetails request. The method always
// closes the http.Response Body.
func (client SharesClient) ListSynchronizationDetailsResponder(resp *http.Response) (result SynchronizationDetailsList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listSynchronizationDetailsNextResults retrieves the next set of results, if any.
func (client SharesClient) listSynchronizationDetailsNextResults(ctx context.Context, lastResults SynchronizationDetailsList) (result SynchronizationDetailsList, err error) {
	req, err := lastResults.synchronizationDetailsListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "datashare.SharesClient", "listSynchronizationDetailsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSynchronizationDetailsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "datashare.SharesClient", "listSynchronizationDetailsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListSynchronizationDetailsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.SharesClient", "listSynchronizationDetailsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListSynchronizationDetailsComplete enumerates all values, automatically crossing page boundaries as required.
func (client SharesClient) ListSynchronizationDetailsComplete(ctx context.Context, resourceGroupName string, accountName string, shareName string, shareSynchronization ShareSynchronization, skipToken string) (result SynchronizationDetailsListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SharesClient.ListSynchronizationDetails")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListSynchronizationDetails(ctx, resourceGroupName, accountName, shareName, shareSynchronization, skipToken)
	return
}

// ListSynchronizations list synchronizations of a share
// Parameters:
// resourceGroupName - the resource group name.
// accountName - the name of the share account.
// shareName - the name of the share.
// skipToken - continuation token
func (client SharesClient) ListSynchronizations(ctx context.Context, resourceGroupName string, accountName string, shareName string, skipToken string) (result ShareSynchronizationListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SharesClient.ListSynchronizations")
		defer func() {
			sc := -1
			if result.ssl.Response.Response != nil {
				sc = result.ssl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listSynchronizationsNextResults
	req, err := client.ListSynchronizationsPreparer(ctx, resourceGroupName, accountName, shareName, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.SharesClient", "ListSynchronizations", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSynchronizationsSender(req)
	if err != nil {
		result.ssl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datashare.SharesClient", "ListSynchronizations", resp, "Failure sending request")
		return
	}

	result.ssl, err = client.ListSynchronizationsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.SharesClient", "ListSynchronizations", resp, "Failure responding to request")
	}

	return
}

// ListSynchronizationsPreparer prepares the ListSynchronizations request.
func (client SharesClient) ListSynchronizationsPreparer(ctx context.Context, resourceGroupName string, accountName string, shareName string, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"shareName":         autorest.Encode("path", shareName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/listSynchronizations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSynchronizationsSender sends the ListSynchronizations request. The method will close the
// http.Response Body if it receives an error.
func (client SharesClient) ListSynchronizationsSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListSynchronizationsResponder handles the response to the ListSynchronizations request. The method always
// closes the http.Response Body.
func (client SharesClient) ListSynchronizationsResponder(resp *http.Response) (result ShareSynchronizationList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listSynchronizationsNextResults retrieves the next set of results, if any.
func (client SharesClient) listSynchronizationsNextResults(ctx context.Context, lastResults ShareSynchronizationList) (result ShareSynchronizationList, err error) {
	req, err := lastResults.shareSynchronizationListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "datashare.SharesClient", "listSynchronizationsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSynchronizationsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "datashare.SharesClient", "listSynchronizationsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListSynchronizationsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.SharesClient", "listSynchronizationsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListSynchronizationsComplete enumerates all values, automatically crossing page boundaries as required.
func (client SharesClient) ListSynchronizationsComplete(ctx context.Context, resourceGroupName string, accountName string, shareName string, skipToken string) (result ShareSynchronizationListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SharesClient.ListSynchronizations")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListSynchronizations(ctx, resourceGroupName, accountName, shareName, skipToken)
	return
}
