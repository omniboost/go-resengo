package resengo

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/omniboost/go-fortnox/utils"
)

func (c *Client) NewPostEventReservationsRequest() PostEventReservationsRequest {
	r := PostEventReservationsRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewPostEventReservationsQueryParams()
	r.pathParams = r.NewPostEventReservationsPathParams()
	r.requestBody = r.NewPostEventReservationsRequestBody()
	return r
}

type PostEventReservationsRequest struct {
	client      *Client
	queryParams *PostEventReservationsQueryParams
	pathParams  *PostEventReservationsPathParams
	method      string
	headers     http.Header
	requestBody PostEventReservationsRequestBody
}

func (r PostEventReservationsRequest) NewPostEventReservationsQueryParams() *PostEventReservationsQueryParams {
	return &PostEventReservationsQueryParams{
		// Pagination: odata.NewPagination(),
	}
}

type PostEventReservationsQueryParams struct {
}

func (p PostEventReservationsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PostEventReservationsRequest) QueryParams() *PostEventReservationsQueryParams {
	return r.queryParams
}

func (r PostEventReservationsRequest) NewPostEventReservationsPathParams() *PostEventReservationsPathParams {
	return &PostEventReservationsPathParams{}
}

type PostEventReservationsPathParams struct {
}

func (p *PostEventReservationsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostEventReservationsRequest) PathParams() *PostEventReservationsPathParams {
	return r.pathParams
}

func (r *PostEventReservationsRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostEventReservationsRequest) Method() string {
	return r.method
}

func (r PostEventReservationsRequest) NewPostEventReservationsRequestBody() PostEventReservationsRequestBody {
	return PostEventReservationsRequestBody{}
}

type PostEventReservationsRequestBody EventReservation

func (r *PostEventReservationsRequest) RequestBody() *PostEventReservationsRequestBody {
	return &r.requestBody
}

func (r *PostEventReservationsRequest) SetRequestBody(body PostEventReservationsRequestBody) {
	r.requestBody = body
}

func (r *PostEventReservationsRequest) NewResponseBody() *PostEventReservationsResponseBody {
	return &PostEventReservationsResponseBody{}
}

type PostEventReservationsResponseBody EventReservation

func (r *PostEventReservationsRequest) URL() url.URL {
	return r.client.GetEndpointURL("/company/{{.company_id}}/eventreservation", r.PathParams())
}

func (r *PostEventReservationsRequest) Do() (PostEventReservationsResponseBody, error) {
	// Create http request
	u := r.URL()
	u.Host = strings.Replace(u.Host, "{{.api}}", "reservationapi", -1)
	req, err := r.client.NewRequest(nil, r.Method(), u, r.RequestBody())
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
