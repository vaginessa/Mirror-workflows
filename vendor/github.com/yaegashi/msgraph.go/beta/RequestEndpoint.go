// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// EndpointRequestBuilder is request builder for Endpoint
type EndpointRequestBuilder struct{ BaseRequestBuilder }

// Request returns EndpointRequest
func (b *EndpointRequestBuilder) Request() *EndpointRequest {
	return &EndpointRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EndpointRequest is request for Endpoint
type EndpointRequest struct{ BaseRequest }

// Get performs GET request for Endpoint
func (r *EndpointRequest) Get(ctx context.Context) (resObj *Endpoint, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Endpoint
func (r *EndpointRequest) Update(ctx context.Context, reqObj *Endpoint) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Endpoint
func (r *EndpointRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}