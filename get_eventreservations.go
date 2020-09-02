package resengo

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/omniboost/go-fortnox/utils"
)

func (c *Client) NewGetEventReservationsRequest() GetEventReservationsRequest {
	r := GetEventReservationsRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewGetEventReservationsQueryParams()
	r.pathParams = r.NewGetEventReservationsPathParams()
	r.requestBody = r.NewGetEventReservationsRequestBody()
	return r
}

type GetEventReservationsRequest struct {
	client      *Client
	queryParams *GetEventReservationsQueryParams
	pathParams  *GetEventReservationsPathParams
	method      string
	headers     http.Header
	requestBody GetEventReservationsRequestBody
}

func (r GetEventReservationsRequest) NewGetEventReservationsQueryParams() *GetEventReservationsQueryParams {
	return &GetEventReservationsQueryParams{
		// Pagination: odata.NewPagination(),
	}
}

type GetEventReservationsQueryParams struct {
	Page      int `schema:"Page,omitempty"`
	PageSize  int `schema:"PageSize,omitempty"`
	GuestName string
}

func (p GetEventReservationsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetEventReservationsRequest) QueryParams() *GetEventReservationsQueryParams {
	return r.queryParams
}

func (r GetEventReservationsRequest) NewGetEventReservationsPathParams() *GetEventReservationsPathParams {
	return &GetEventReservationsPathParams{}
}

type GetEventReservationsPathParams struct {
}

func (p *GetEventReservationsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetEventReservationsRequest) PathParams() *GetEventReservationsPathParams {
	return r.pathParams
}

func (r *GetEventReservationsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetEventReservationsRequest) Method() string {
	return r.method
}

func (r GetEventReservationsRequest) NewGetEventReservationsRequestBody() GetEventReservationsRequestBody {
	return GetEventReservationsRequestBody{}
}

type GetEventReservationsRequestBody struct{}

func (r *GetEventReservationsRequest) RequestBody() *GetEventReservationsRequestBody {
	return &r.requestBody
}

func (r *GetEventReservationsRequest) SetRequestBody(body GetEventReservationsRequestBody) {
	r.requestBody = body
}

func (r *GetEventReservationsRequest) NewResponseBody() *GetEventReservationsResponseBody {
	return &GetEventReservationsResponseBody{}
}

type GetEventReservationsResponseBody EventReservations

func (r *GetEventReservationsRequest) URL() url.URL {
	return r.client.GetEndpointURL("/company/{{.company_id}}/eventreservation", r.PathParams())
}

func (r *GetEventReservationsRequest) Do() (GetEventReservationsResponseBody, error) {
	// Create http request
	u := r.URL()
	u.Host = strings.Replace(u.Host, "{{.api}}", "reservationapi", -1)
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
