// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// CountryRegionRequestBuilder is request builder for CountryRegion
type CountryRegionRequestBuilder struct{ BaseRequestBuilder }

// Request returns CountryRegionRequest
func (b *CountryRegionRequestBuilder) Request() *CountryRegionRequest {
	return &CountryRegionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CountryRegionRequest is request for CountryRegion
type CountryRegionRequest struct{ BaseRequest }

// Get performs GET request for CountryRegion
func (r *CountryRegionRequest) Get(ctx context.Context) (resObj *CountryRegion, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CountryRegion
func (r *CountryRegionRequest) Update(ctx context.Context, reqObj *CountryRegion) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CountryRegion
func (r *CountryRegionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}