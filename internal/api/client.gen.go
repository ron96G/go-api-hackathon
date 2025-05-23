// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// ListPets request
	ListPets(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AddPetWithBody request with any body
	AddPetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AddPet(ctx context.Context, body AddPetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeletePet request
	DeletePet(ctx context.Context, petId PetId, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetPetById request
	GetPetById(ctx context.Context, petId PetId, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdatePetWithBody request with any body
	UpdatePetWithBody(ctx context.Context, petId PetId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdatePet(ctx context.Context, petId PetId, body UpdatePetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) ListPets(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListPetsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AddPetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAddPetRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AddPet(ctx context.Context, body AddPetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAddPetRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeletePet(ctx context.Context, petId PetId, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeletePetRequest(c.Server, petId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetPetById(ctx context.Context, petId PetId, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetPetByIdRequest(c.Server, petId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdatePetWithBody(ctx context.Context, petId PetId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdatePetRequestWithBody(c.Server, petId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdatePet(ctx context.Context, petId PetId, body UpdatePetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdatePetRequest(c.Server, petId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewListPetsRequest generates requests for ListPets
func NewListPetsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/pets")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewAddPetRequest calls the generic AddPet builder with application/json body
func NewAddPetRequest(server string, body AddPetJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAddPetRequestWithBody(server, "application/json", bodyReader)
}

// NewAddPetRequestWithBody generates requests for AddPet with any type of body
func NewAddPetRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/pets")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewDeletePetRequest generates requests for DeletePet
func NewDeletePetRequest(server string, petId PetId) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "petId", runtime.ParamLocationPath, petId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/pets/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetPetByIdRequest generates requests for GetPetById
func NewGetPetByIdRequest(server string, petId PetId) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "petId", runtime.ParamLocationPath, petId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/pets/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewUpdatePetRequest calls the generic UpdatePet builder with application/json body
func NewUpdatePetRequest(server string, petId PetId, body UpdatePetJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdatePetRequestWithBody(server, petId, "application/json", bodyReader)
}

// NewUpdatePetRequestWithBody generates requests for UpdatePet with any type of body
func NewUpdatePetRequestWithBody(server string, petId PetId, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "petId", runtime.ParamLocationPath, petId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/pets/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// ListPetsWithResponse request
	ListPetsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListPetsResponse, error)

	// AddPetWithBodyWithResponse request with any body
	AddPetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AddPetResponse, error)

	AddPetWithResponse(ctx context.Context, body AddPetJSONRequestBody, reqEditors ...RequestEditorFn) (*AddPetResponse, error)

	// DeletePetWithResponse request
	DeletePetWithResponse(ctx context.Context, petId PetId, reqEditors ...RequestEditorFn) (*DeletePetResponse, error)

	// GetPetByIdWithResponse request
	GetPetByIdWithResponse(ctx context.Context, petId PetId, reqEditors ...RequestEditorFn) (*GetPetByIdResponse, error)

	// UpdatePetWithBodyWithResponse request with any body
	UpdatePetWithBodyWithResponse(ctx context.Context, petId PetId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdatePetResponse, error)

	UpdatePetWithResponse(ctx context.Context, petId PetId, body UpdatePetJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdatePetResponse, error)
}

type ListPetsResponse struct {
	Body                          []byte
	HTTPResponse                  *http.Response
	JSON200                       *[]Pet
	ApplicationproblemJSON400     *ErrorResponse
	ApplicationproblemJSON500     *ErrorResponse
	ApplicationproblemJSONDefault *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r ListPetsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListPetsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AddPetResponse struct {
	Body                          []byte
	HTTPResponse                  *http.Response
	JSON200                       *Pet
	ApplicationproblemJSON400     *ErrorResponse
	ApplicationproblemJSON500     *ErrorResponse
	ApplicationproblemJSONDefault *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r AddPetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AddPetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeletePetResponse struct {
	Body                          []byte
	HTTPResponse                  *http.Response
	ApplicationproblemJSON400     *ErrorResponse
	ApplicationproblemJSON500     *ErrorResponse
	ApplicationproblemJSONDefault *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r DeletePetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeletePetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetPetByIdResponse struct {
	Body                          []byte
	HTTPResponse                  *http.Response
	JSON200                       *Pet
	ApplicationproblemJSON400     *ErrorResponse
	ApplicationproblemJSON404     *ErrorResponse
	ApplicationproblemJSON500     *ErrorResponse
	ApplicationproblemJSONDefault *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r GetPetByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetPetByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdatePetResponse struct {
	Body                          []byte
	HTTPResponse                  *http.Response
	JSON200                       *Pet
	ApplicationproblemJSON400     *ErrorResponse
	ApplicationproblemJSON404     *ErrorResponse
	ApplicationproblemJSON500     *ErrorResponse
	ApplicationproblemJSONDefault *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r UpdatePetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdatePetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ListPetsWithResponse request returning *ListPetsResponse
func (c *ClientWithResponses) ListPetsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListPetsResponse, error) {
	rsp, err := c.ListPets(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListPetsResponse(rsp)
}

// AddPetWithBodyWithResponse request with arbitrary body returning *AddPetResponse
func (c *ClientWithResponses) AddPetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AddPetResponse, error) {
	rsp, err := c.AddPetWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAddPetResponse(rsp)
}

func (c *ClientWithResponses) AddPetWithResponse(ctx context.Context, body AddPetJSONRequestBody, reqEditors ...RequestEditorFn) (*AddPetResponse, error) {
	rsp, err := c.AddPet(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAddPetResponse(rsp)
}

// DeletePetWithResponse request returning *DeletePetResponse
func (c *ClientWithResponses) DeletePetWithResponse(ctx context.Context, petId PetId, reqEditors ...RequestEditorFn) (*DeletePetResponse, error) {
	rsp, err := c.DeletePet(ctx, petId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeletePetResponse(rsp)
}

// GetPetByIdWithResponse request returning *GetPetByIdResponse
func (c *ClientWithResponses) GetPetByIdWithResponse(ctx context.Context, petId PetId, reqEditors ...RequestEditorFn) (*GetPetByIdResponse, error) {
	rsp, err := c.GetPetById(ctx, petId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetPetByIdResponse(rsp)
}

// UpdatePetWithBodyWithResponse request with arbitrary body returning *UpdatePetResponse
func (c *ClientWithResponses) UpdatePetWithBodyWithResponse(ctx context.Context, petId PetId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdatePetResponse, error) {
	rsp, err := c.UpdatePetWithBody(ctx, petId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdatePetResponse(rsp)
}

func (c *ClientWithResponses) UpdatePetWithResponse(ctx context.Context, petId PetId, body UpdatePetJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdatePetResponse, error) {
	rsp, err := c.UpdatePet(ctx, petId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdatePetResponse(rsp)
}

// ParseListPetsResponse parses an HTTP response from a ListPetsWithResponse call
func ParseListPetsResponse(rsp *http.Response) (*ListPetsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListPetsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Pet
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSONDefault = &dest

	}

	return response, nil
}

// ParseAddPetResponse parses an HTTP response from a AddPetWithResponse call
func ParseAddPetResponse(rsp *http.Response) (*AddPetResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AddPetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Pet
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSONDefault = &dest

	}

	return response, nil
}

// ParseDeletePetResponse parses an HTTP response from a DeletePetWithResponse call
func ParseDeletePetResponse(rsp *http.Response) (*DeletePetResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeletePetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSONDefault = &dest

	}

	return response, nil
}

// ParseGetPetByIdResponse parses an HTTP response from a GetPetByIdWithResponse call
func ParseGetPetByIdResponse(rsp *http.Response) (*GetPetByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetPetByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Pet
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSONDefault = &dest

	}

	return response, nil
}

// ParseUpdatePetResponse parses an HTTP response from a UpdatePetWithResponse call
func ParseUpdatePetResponse(rsp *http.Response) (*UpdatePetResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdatePetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Pet
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSONDefault = &dest

	}

	return response, nil
}
