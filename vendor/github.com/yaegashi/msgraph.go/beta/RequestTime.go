// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// TimeOffRequestBuilder is request builder for TimeOff
type TimeOffRequestBuilder struct{ BaseRequestBuilder }

// Request returns TimeOffRequest
func (b *TimeOffRequestBuilder) Request() *TimeOffRequest {
	return &TimeOffRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TimeOffRequest is request for TimeOff
type TimeOffRequest struct{ BaseRequest }

// Get performs GET request for TimeOff
func (r *TimeOffRequest) Get(ctx context.Context) (resObj *TimeOff, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TimeOff
func (r *TimeOffRequest) Update(ctx context.Context, reqObj *TimeOff) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TimeOff
func (r *TimeOffRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TimeOffReasonRequestBuilder is request builder for TimeOffReason
type TimeOffReasonRequestBuilder struct{ BaseRequestBuilder }

// Request returns TimeOffReasonRequest
func (b *TimeOffReasonRequestBuilder) Request() *TimeOffReasonRequest {
	return &TimeOffReasonRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TimeOffReasonRequest is request for TimeOffReason
type TimeOffReasonRequest struct{ BaseRequest }

// Get performs GET request for TimeOffReason
func (r *TimeOffReasonRequest) Get(ctx context.Context) (resObj *TimeOffReason, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TimeOffReason
func (r *TimeOffReasonRequest) Update(ctx context.Context, reqObj *TimeOffReason) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TimeOffReason
func (r *TimeOffReasonRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TimeOffRequestObjectRequestBuilder is request builder for TimeOffRequestObject
type TimeOffRequestObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns TimeOffRequestObjectRequest
func (b *TimeOffRequestObjectRequestBuilder) Request() *TimeOffRequestObjectRequest {
	return &TimeOffRequestObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TimeOffRequestObjectRequest is request for TimeOffRequestObject
type TimeOffRequestObjectRequest struct{ BaseRequest }

// Get performs GET request for TimeOffRequestObject
func (r *TimeOffRequestObjectRequest) Get(ctx context.Context) (resObj *TimeOffRequestObject, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TimeOffRequestObject
func (r *TimeOffRequestObjectRequest) Update(ctx context.Context, reqObj *TimeOffRequestObject) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TimeOffRequestObject
func (r *TimeOffRequestObjectRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}