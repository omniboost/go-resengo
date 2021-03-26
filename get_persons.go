package resengo

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/omniboost/go-fortnox/utils"
)

func (c *Client) NewGetPersonsRequest() GetPersonsRequest {
	r := GetPersonsRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewGetPersonsQueryParams()
	r.pathParams = r.NewGetPersonsPathParams()
	r.requestBody = r.NewGetPersonsRequestBody()
	return r
}

type GetPersonsRequest struct {
	client      *Client
	queryParams *GetPersonsQueryParams
	pathParams  *GetPersonsPathParams
	method      string
	headers     http.Header
	requestBody GetPersonsRequestBody
}

func (r GetPersonsRequest) NewGetPersonsQueryParams() *GetPersonsQueryParams {
	return &GetPersonsQueryParams{
		// Pagination: odata.NewPagination(),
	}
}

type GetPersonsQueryParams struct {
	Page      int    `schema:"Page,omitempty"`
	PageSize  int    `schema:"PageSize,omitempty"`
	GuestName string `schema:"GuestName,omitempty"`
}

func (p GetPersonsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetPersonsRequest) QueryParams() *GetPersonsQueryParams {
	return r.queryParams
}

func (r GetPersonsRequest) NewGetPersonsPathParams() *GetPersonsPathParams {
	return &GetPersonsPathParams{}
}

type GetPersonsPathParams struct {
}

func (p *GetPersonsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetPersonsRequest) PathParams() *GetPersonsPathParams {
	return r.pathParams
}

func (r *GetPersonsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetPersonsRequest) Method() string {
	return r.method
}

func (r GetPersonsRequest) NewGetPersonsRequestBody() GetPersonsRequestBody {
	return GetPersonsRequestBody{}
}

type GetPersonsRequestBody struct{}

func (r *GetPersonsRequest) RequestBody() *GetPersonsRequestBody {
	return &r.requestBody
}

func (r *GetPersonsRequest) SetRequestBody(body GetPersonsRequestBody) {
	r.requestBody = body
}

func (r *GetPersonsRequest) NewResponseBody() *GetPersonsResponseBody {
	return &GetPersonsResponseBody{}
}

type GetPersonsResponseBody Persons

func (r *GetPersonsRequest) URL() url.URL {
	return r.client.GetEndpointURL("/company/{{.company_id}}/person", r.PathParams())
}

func (r *GetPersonsRequest) Do() (GetPersonsResponseBody, error) {
	// Create http request
	u := r.URL()
	u.Host = strings.Replace(u.Host, "{{.api}}", "personapi", -1)
	req, err := r.client.NewRequest(nil, r.Method(), u, nil)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
