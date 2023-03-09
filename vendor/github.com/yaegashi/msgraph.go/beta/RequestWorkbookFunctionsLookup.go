// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsLookupRequestBuilder struct{ BaseRequestBuilder }

// Lookup action undocumented
func (b *WorkbookFunctionsRequestBuilder) Lookup(reqObj *WorkbookFunctionsLookupRequestParameter) *WorkbookFunctionsLookupRequestBuilder {
	bb := &WorkbookFunctionsLookupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/lookup"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsLookupRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsLookupRequestBuilder) Request() *WorkbookFunctionsLookupRequest {
	return &WorkbookFunctionsLookupRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsLookupRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}