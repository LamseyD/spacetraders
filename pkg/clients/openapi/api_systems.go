/*
SpaceTraders API

SpaceTraders is an open-universe game and learning platform that offers a set of HTTP endpoints to control a fleet of ships and explore a multiplayer universe.  The API is documented using [OpenAPI](https://github.com/SpaceTradersAPI/api-docs). You can send your first request right here in your browser to check the status of the game server.  ```json http {   \"method\": \"GET\",   \"url\": \"https://api.spacetraders.io/v2\", } ```  Unlike a traditional game, SpaceTraders does not have a first-party client or app to play the game. Instead, you can use the API to build your own client, write a script to automate your ships, or try an app built by the community.  We have a [Discord channel](https://discord.com/invite/jh6zurdWk5) where you can share your projects, ask questions, and get help from other players.

API version: 2.0.0
Contact: joel@spacetraders.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// SystemsAPIService SystemsAPI service
type SystemsAPIService service

type ApiGetJumpGateRequest struct {
	ctx            context.Context
	ApiService     *SystemsAPIService
	systemSymbol   string
	waypointSymbol string
}

func (r ApiGetJumpGateRequest) Execute() (*GetJumpGate200Response, *http.Response, error) {
	return r.ApiService.GetJumpGateExecute(r)
}

/*
GetJumpGate Get Jump Gate

Get jump gate details for a waypoint. Requires a waypoint of type `JUMP_GATE` to use.

The response will return all systems that are have a Jump Gate in range of this Jump Gate. Those systems can be jumped to from this Jump Gate.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param systemSymbol The system symbol
	@param waypointSymbol The waypoint symbol
	@return ApiGetJumpGateRequest
*/
func (a *SystemsAPIService) GetJumpGate(ctx context.Context, systemSymbol string, waypointSymbol string) ApiGetJumpGateRequest {
	return ApiGetJumpGateRequest{
		ApiService:     a,
		ctx:            ctx,
		systemSymbol:   systemSymbol,
		waypointSymbol: waypointSymbol,
	}
}

// Execute executes the request
//
//	@return GetJumpGate200Response
func (a *SystemsAPIService) GetJumpGateExecute(r ApiGetJumpGateRequest) (*GetJumpGate200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetJumpGate200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemsAPIService.GetJumpGate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systems/{systemSymbol}/waypoints/{waypointSymbol}/jump-gate"
	localVarPath = strings.Replace(localVarPath, "{"+"systemSymbol"+"}", url.PathEscape(parameterValueToString(r.systemSymbol, "systemSymbol")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"waypointSymbol"+"}", url.PathEscape(parameterValueToString(r.waypointSymbol, "waypointSymbol")), -1)

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetMarketRequest struct {
	ctx            context.Context
	ApiService     *SystemsAPIService
	systemSymbol   string
	waypointSymbol string
}

func (r ApiGetMarketRequest) Execute() (*GetMarket200Response, *http.Response, error) {
	return r.ApiService.GetMarketExecute(r)
}

/*
GetMarket Get Market

Retrieve imports, exports and exchange data from a marketplace. Requires a waypoint that has the `Marketplace` trait to use.

Send a ship to the waypoint to access trade good prices and recent transactions. Refer to the [Market Overview page](https://docs.spacetraders.io/game-concepts/markets) to gain better a understanding of the market in the game.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param systemSymbol The system symbol
	@param waypointSymbol The waypoint symbol
	@return ApiGetMarketRequest
*/
func (a *SystemsAPIService) GetMarket(ctx context.Context, systemSymbol string, waypointSymbol string) ApiGetMarketRequest {
	return ApiGetMarketRequest{
		ApiService:     a,
		ctx:            ctx,
		systemSymbol:   systemSymbol,
		waypointSymbol: waypointSymbol,
	}
}

// Execute executes the request
//
//	@return GetMarket200Response
func (a *SystemsAPIService) GetMarketExecute(r ApiGetMarketRequest) (*GetMarket200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetMarket200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemsAPIService.GetMarket")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systems/{systemSymbol}/waypoints/{waypointSymbol}/market"
	localVarPath = strings.Replace(localVarPath, "{"+"systemSymbol"+"}", url.PathEscape(parameterValueToString(r.systemSymbol, "systemSymbol")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"waypointSymbol"+"}", url.PathEscape(parameterValueToString(r.waypointSymbol, "waypointSymbol")), -1)

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetShipyardRequest struct {
	ctx            context.Context
	ApiService     *SystemsAPIService
	systemSymbol   string
	waypointSymbol string
}

func (r ApiGetShipyardRequest) Execute() (*GetShipyard200Response, *http.Response, error) {
	return r.ApiService.GetShipyardExecute(r)
}

/*
GetShipyard Get Shipyard

Get the shipyard for a waypoint. Requires a waypoint that has the `Shipyard` trait to use. Send a ship to the waypoint to access data on ships that are currently available for purchase and recent transactions.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param systemSymbol The system symbol
	@param waypointSymbol The waypoint symbol
	@return ApiGetShipyardRequest
*/
func (a *SystemsAPIService) GetShipyard(ctx context.Context, systemSymbol string, waypointSymbol string) ApiGetShipyardRequest {
	return ApiGetShipyardRequest{
		ApiService:     a,
		ctx:            ctx,
		systemSymbol:   systemSymbol,
		waypointSymbol: waypointSymbol,
	}
}

// Execute executes the request
//
//	@return GetShipyard200Response
func (a *SystemsAPIService) GetShipyardExecute(r ApiGetShipyardRequest) (*GetShipyard200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetShipyard200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemsAPIService.GetShipyard")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systems/{systemSymbol}/waypoints/{waypointSymbol}/shipyard"
	localVarPath = strings.Replace(localVarPath, "{"+"systemSymbol"+"}", url.PathEscape(parameterValueToString(r.systemSymbol, "systemSymbol")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"waypointSymbol"+"}", url.PathEscape(parameterValueToString(r.waypointSymbol, "waypointSymbol")), -1)

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetSystemRequest struct {
	ctx          context.Context
	ApiService   *SystemsAPIService
	systemSymbol string
}

func (r ApiGetSystemRequest) Execute() (*GetSystem200Response, *http.Response, error) {
	return r.ApiService.GetSystemExecute(r)
}

/*
GetSystem Get System

Get the details of a system.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param systemSymbol The system symbol
	@return ApiGetSystemRequest
*/
func (a *SystemsAPIService) GetSystem(ctx context.Context, systemSymbol string) ApiGetSystemRequest {
	return ApiGetSystemRequest{
		ApiService:   a,
		ctx:          ctx,
		systemSymbol: systemSymbol,
	}
}

// Execute executes the request
//
//	@return GetSystem200Response
func (a *SystemsAPIService) GetSystemExecute(r ApiGetSystemRequest) (*GetSystem200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetSystem200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemsAPIService.GetSystem")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systems/{systemSymbol}"
	localVarPath = strings.Replace(localVarPath, "{"+"systemSymbol"+"}", url.PathEscape(parameterValueToString(r.systemSymbol, "systemSymbol")), -1)

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetSystemWaypointsRequest struct {
	ctx          context.Context
	ApiService   *SystemsAPIService
	systemSymbol string
	page         *int32
	limit        *int32
}

// What entry offset to request
func (r ApiGetSystemWaypointsRequest) Page(page int32) ApiGetSystemWaypointsRequest {
	r.page = &page
	return r
}

// How many entries to return per page
func (r ApiGetSystemWaypointsRequest) Limit(limit int32) ApiGetSystemWaypointsRequest {
	r.limit = &limit
	return r
}

func (r ApiGetSystemWaypointsRequest) Execute() (*GetSystemWaypoints200Response, *http.Response, error) {
	return r.ApiService.GetSystemWaypointsExecute(r)
}

/*
GetSystemWaypoints List Waypoints in System

Return a paginated list of all of the waypoints for a given system.

If a waypoint is uncharted, it will return the `Uncharted` trait instead of its actual traits.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param systemSymbol The system symbol
	@return ApiGetSystemWaypointsRequest
*/
func (a *SystemsAPIService) GetSystemWaypoints(ctx context.Context, systemSymbol string) ApiGetSystemWaypointsRequest {
	return ApiGetSystemWaypointsRequest{
		ApiService:   a,
		ctx:          ctx,
		systemSymbol: systemSymbol,
	}
}

// Execute executes the request
//
//	@return GetSystemWaypoints200Response
func (a *SystemsAPIService) GetSystemWaypointsExecute(r ApiGetSystemWaypointsRequest) (*GetSystemWaypoints200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetSystemWaypoints200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemsAPIService.GetSystemWaypoints")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systems/{systemSymbol}/waypoints"
	localVarPath = strings.Replace(localVarPath, "{"+"systemSymbol"+"}", url.PathEscape(parameterValueToString(r.systemSymbol, "systemSymbol")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetSystemsRequest struct {
	ctx        context.Context
	ApiService *SystemsAPIService
	page       *int32
	limit      *int32
}

// What entry offset to request
func (r ApiGetSystemsRequest) Page(page int32) ApiGetSystemsRequest {
	r.page = &page
	return r
}

// How many entries to return per page
func (r ApiGetSystemsRequest) Limit(limit int32) ApiGetSystemsRequest {
	r.limit = &limit
	return r
}

func (r ApiGetSystemsRequest) Execute() (*GetSystems200Response, *http.Response, error) {
	return r.ApiService.GetSystemsExecute(r)
}

/*
GetSystems List Systems

Return a paginated list of all systems.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetSystemsRequest
*/
func (a *SystemsAPIService) GetSystems(ctx context.Context) ApiGetSystemsRequest {
	return ApiGetSystemsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetSystems200Response
func (a *SystemsAPIService) GetSystemsExecute(r ApiGetSystemsRequest) (*GetSystems200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetSystems200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemsAPIService.GetSystems")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systems"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetWaypointRequest struct {
	ctx            context.Context
	ApiService     *SystemsAPIService
	systemSymbol   string
	waypointSymbol string
}

func (r ApiGetWaypointRequest) Execute() (*GetWaypoint200Response, *http.Response, error) {
	return r.ApiService.GetWaypointExecute(r)
}

/*
GetWaypoint Get Waypoint

View the details of a waypoint.

If the waypoint is uncharted, it will return the 'Uncharted' trait instead of its actual traits.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param systemSymbol The system symbol
	@param waypointSymbol The waypoint symbol
	@return ApiGetWaypointRequest
*/
func (a *SystemsAPIService) GetWaypoint(ctx context.Context, systemSymbol string, waypointSymbol string) ApiGetWaypointRequest {
	return ApiGetWaypointRequest{
		ApiService:     a,
		ctx:            ctx,
		systemSymbol:   systemSymbol,
		waypointSymbol: waypointSymbol,
	}
}

// Execute executes the request
//
//	@return GetWaypoint200Response
func (a *SystemsAPIService) GetWaypointExecute(r ApiGetWaypointRequest) (*GetWaypoint200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetWaypoint200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemsAPIService.GetWaypoint")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systems/{systemSymbol}/waypoints/{waypointSymbol}"
	localVarPath = strings.Replace(localVarPath, "{"+"systemSymbol"+"}", url.PathEscape(parameterValueToString(r.systemSymbol, "systemSymbol")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"waypointSymbol"+"}", url.PathEscape(parameterValueToString(r.waypointSymbol, "waypointSymbol")), -1)

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
