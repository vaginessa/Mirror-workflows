// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// PolicyRequestBuilder is request builder for Policy
type PolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns PolicyRequest
func (b *PolicyRequestBuilder) Request() *PolicyRequest {
	return &PolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PolicyRequest is request for Policy
type PolicyRequest struct{ BaseRequest }

// Get performs GET request for Policy
func (r *PolicyRequest) Get(ctx context.Context) (resObj *Policy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Policy
func (r *PolicyRequest) Update(ctx context.Context, reqObj *Policy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Policy
func (r *PolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PolicySetRequestBuilder is request builder for PolicySet
type PolicySetRequestBuilder struct{ BaseRequestBuilder }

// Request returns PolicySetRequest
func (b *PolicySetRequestBuilder) Request() *PolicySetRequest {
	return &PolicySetRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PolicySetRequest is request for PolicySet
type PolicySetRequest struct{ BaseRequest }

// Get performs GET request for PolicySet
func (r *PolicySetRequest) Get(ctx context.Context) (resObj *PolicySet, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PolicySet
func (r *PolicySetRequest) Update(ctx context.Context, reqObj *PolicySet) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PolicySet
func (r *PolicySetRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PolicySetAssignmentRequestBuilder is request builder for PolicySetAssignment
type PolicySetAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns PolicySetAssignmentRequest
func (b *PolicySetAssignmentRequestBuilder) Request() *PolicySetAssignmentRequest {
	return &PolicySetAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PolicySetAssignmentRequest is request for PolicySetAssignment
type PolicySetAssignmentRequest struct{ BaseRequest }

// Get performs GET request for PolicySetAssignment
func (r *PolicySetAssignmentRequest) Get(ctx context.Context) (resObj *PolicySetAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PolicySetAssignment
func (r *PolicySetAssignmentRequest) Update(ctx context.Context, reqObj *PolicySetAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PolicySetAssignment
func (r *PolicySetAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PolicySetItemRequestBuilder is request builder for PolicySetItem
type PolicySetItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns PolicySetItemRequest
func (b *PolicySetItemRequestBuilder) Request() *PolicySetItemRequest {
	return &PolicySetItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PolicySetItemRequest is request for PolicySetItem
type PolicySetItemRequest struct{ BaseRequest }

// Get performs GET request for PolicySetItem
func (r *PolicySetItemRequest) Get(ctx context.Context) (resObj *PolicySetItem, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PolicySetItem
func (r *PolicySetItemRequest) Update(ctx context.Context, reqObj *PolicySetItem) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PolicySetItem
func (r *PolicySetItemRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type PolicySetCollectionGetPolicySetsRequestBuilder struct{ BaseRequestBuilder }

// GetPolicySets action undocumented
func (b *DeviceAppManagementPolicySetsCollectionRequestBuilder) GetPolicySets(reqObj *PolicySetCollectionGetPolicySetsRequestParameter) *PolicySetCollectionGetPolicySetsRequestBuilder {
	bb := &PolicySetCollectionGetPolicySetsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getPolicySets"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type PolicySetCollectionGetPolicySetsRequest struct{ BaseRequest }

//
func (b *PolicySetCollectionGetPolicySetsRequestBuilder) Request() *PolicySetCollectionGetPolicySetsRequest {
	return &PolicySetCollectionGetPolicySetsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *PolicySetCollectionGetPolicySetsRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]PolicySet, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []PolicySet
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []PolicySet
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

//
func (r *PolicySetCollectionGetPolicySetsRequest) PostN(ctx context.Context, n int) ([]PolicySet, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

//
func (r *PolicySetCollectionGetPolicySetsRequest) Post(ctx context.Context) ([]PolicySet, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}

//
type PolicySetUpdateRequestBuilder struct{ BaseRequestBuilder }

// Update action undocumented
func (b *PolicySetRequestBuilder) Update(reqObj *PolicySetUpdateRequestParameter) *PolicySetUpdateRequestBuilder {
	bb := &PolicySetUpdateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/update"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type PolicySetUpdateRequest struct{ BaseRequest }

//
func (b *PolicySetUpdateRequestBuilder) Request() *PolicySetUpdateRequest {
	return &PolicySetUpdateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *PolicySetUpdateRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}