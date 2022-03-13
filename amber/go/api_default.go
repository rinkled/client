/*
Amber Electric Public API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

// DefaultApiService DefaultApi service
type DefaultApiService service

type ApiSitesGetRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
}


func (r ApiSitesGetRequest) Execute() ([]InlineResponse200, *http.Response, error) {
	return r.ApiService.SitesGetExecute(r)
}

/*
SitesGet Method for SitesGet

Return all sites linked to your account

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSitesGetRequest
*/
func (a *DefaultApiService) SitesGet(ctx context.Context) ApiSitesGetRequest {
	return ApiSitesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []InlineResponse200
func (a *DefaultApiService) SitesGetExecute(r ApiSitesGetRequest) ([]InlineResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.SitesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sites"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSitesSiteIdUsageGetRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	siteId string
	startDate *interface{}
	endDate *interface{}
	resolution *float32
}

// Return all usage for each interval on and after this day.
func (r ApiSitesSiteIdUsageGetRequest) StartDate(startDate interface{}) ApiSitesSiteIdUsageGetRequest {
	r.startDate = &startDate
	return r
}
// Return all usage for each interval on and before this day.
func (r ApiSitesSiteIdUsageGetRequest) EndDate(endDate interface{}) ApiSitesSiteIdUsageGetRequest {
	r.endDate = &endDate
	return r
}
// Specify the required interval duration resolution. Valid options: 30. Default: 30
func (r ApiSitesSiteIdUsageGetRequest) Resolution(resolution float32) ApiSitesSiteIdUsageGetRequest {
	r.resolution = &resolution
	return r
}

func (r ApiSitesSiteIdUsageGetRequest) Execute() ([]interface{}, *http.Response, error) {
	return r.ApiService.SitesSiteIdUsageGetExecute(r)
}

/*
SitesSiteIdUsageGet Method for SitesSiteIdUsageGet

Returns all usage data between the start and end dates

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param siteId ID of the site you are fetching usage for. Can be found using the `/sites` enpoint
 @return ApiSitesSiteIdUsageGetRequest
*/
func (a *DefaultApiService) SitesSiteIdUsageGet(ctx context.Context, siteId string) ApiSitesSiteIdUsageGetRequest {
	return ApiSitesSiteIdUsageGetRequest{
		ApiService: a,
		ctx: ctx,
		siteId: siteId,
	}
}

// Execute executes the request
//  @return []interface{}
func (a *DefaultApiService) SitesSiteIdUsageGetExecute(r ApiSitesSiteIdUsageGetRequest) ([]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.SitesSiteIdUsageGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sites/{siteId}/usage"
	localVarPath = strings.Replace(localVarPath, "{"+"siteId"+"}", url.PathEscape(parameterToString(r.siteId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.startDate == nil {
		return localVarReturnValue, nil, reportError("startDate is required and must be specified")
	}
	if r.endDate == nil {
		return localVarReturnValue, nil, reportError("endDate is required and must be specified")
	}

	localVarQueryParams.Add("startDate", parameterToString(*r.startDate, ""))
	localVarQueryParams.Add("endDate", parameterToString(*r.endDate, ""))
	if r.resolution != nil {
		localVarQueryParams.Add("resolution", parameterToString(*r.resolution, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}