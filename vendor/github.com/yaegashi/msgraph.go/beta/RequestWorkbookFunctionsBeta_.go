// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsBeta_DistRequestBuilder struct{ BaseRequestBuilder }

// Beta_Dist action undocumented
func (b *WorkbookFunctionsRequestBuilder) Beta_Dist(reqObj *WorkbookFunctionsBeta_DistRequestParameter) *WorkbookFunctionsBeta_DistRequestBuilder {
	bb := &WorkbookFunctionsBeta_DistRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/beta_Dist"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsBeta_DistRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsBeta_DistRequestBuilder) Request() *WorkbookFunctionsBeta_DistRequest {
	return &WorkbookFunctionsBeta_DistRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsBeta_DistRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookFunctionsBeta_InvRequestBuilder struct{ BaseRequestBuilder }

// Beta_Inv action undocumented
func (b *WorkbookFunctionsRequestBuilder) Beta_Inv(reqObj *WorkbookFunctionsBeta_InvRequestParameter) *WorkbookFunctionsBeta_InvRequestBuilder {
	bb := &WorkbookFunctionsBeta_InvRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/beta_Inv"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsBeta_InvRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsBeta_InvRequestBuilder) Request() *WorkbookFunctionsBeta_InvRequest {
	return &WorkbookFunctionsBeta_InvRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsBeta_InvRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}