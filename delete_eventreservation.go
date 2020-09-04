package resengo

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/omniboost/go-fortnox/utils"
)

func (c *Client) NewDeleteEventReservationsRequest() DeleteEventReservationsRequest {
	r := DeleteEventReservationsRequest{
		client:  c,
		method:  http.MethodDelete,
		headers: http.Header{},
	}

	r.queryParams = r.NewDeleteEventReservationsQueryParams()
	r.pathParams = r.NewDeleteEventReservationsPathParams()
	r.requestBody = r.NewDeleteEventReservationsRequestBody()
	return r
}

type DeleteEventReservationsRequest struct {
	client      *Client
	queryParams *DeleteEventReservationsQueryParams
	pathParams  *DeleteEventReservationsPathParams
	method      string
	headers     http.Header
	requestBody DeleteEventReservationsRequestBody
}

func (r DeleteEventReservationsRequest) NewDeleteEventReservationsQueryParams() *DeleteEventReservationsQueryParams {
	return &DeleteEventReservationsQueryParams{
		// Pagination: odata.NewPagination(),
	}
}

type DeleteEventReservationsQueryParams struct{}

func (p DeleteEventReservationsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *DeleteEventReservationsRequest) QueryParams() *DeleteEventReservationsQueryParams {
	return r.queryParams
}

func (r DeleteEventReservationsRequest) NewDeleteEventReservationsPathParams() *DeleteEventReservationsPathParams {
	return &DeleteEventReservationsPathParams{}
}

type DeleteEventReservationsPathParams struct {
	EventReservationID int `schema:"eventreservation_id"`
}

func (p *DeleteEventReservationsPathParams) Params() map[string]string {
	return map[string]string{
		"eventreservation_id": strconv.Itoa(p.EventReservationID),
	}
}

func (r *DeleteEventReservationsRequest) PathParams() *DeleteEventReservationsPathParams {
	return r.pathParams
}

func (r *DeleteEventReservationsRequest) SetMethod(method string) {
	r.method = method
}

func (r *DeleteEventReservationsRequest) Method() string {
	return r.method
}

func (r DeleteEventReservationsRequest) NewDeleteEventReservationsRequestBody() DeleteEventReservationsRequestBody {
	return DeleteEventReservationsRequestBody{}
}

type DeleteEventReservationsRequestBody struct{}

func (r *DeleteEventReservationsRequest) RequestBody() *DeleteEventReservationsRequestBody {
	return &r.requestBody
}

func (r *DeleteEventReservationsRequest) SetRequestBody(body DeleteEventReservationsRequestBody) {
	r.requestBody = body
}

func (r *DeleteEventReservationsRequest) NewResponseBody() *DeleteEventReservationsResponseBody {
	return &DeleteEventReservationsResponseBody{}
}

type DeleteEventReservationsResponseBody struct{}

func (r *DeleteEventReservationsRequest) URL() url.URL {
	return r.client.GetEndpointURL("/company/{{.company_id}}/eventreservation/{{.eventreservation_id}}", r.PathParams())
}

func (r *DeleteEventReservationsRequest) Do() (DeleteEventReservationsResponseBody, error) {
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
