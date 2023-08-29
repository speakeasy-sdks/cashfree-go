// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
	"net/http"
)

type GetSettlementsRequest struct {
	// application/json
	AcceptMedia *string `header:"style=simple,explode=false,name=Accept-Media"`
	// Request body to fetch settlements
	FetchSettlementReconRequest *shared.FetchSettlementReconRequest `request:"mediaType=application/json"`
	// API version to be used. Format is in YYYY-MM-DD
	XAPIVersion string `header:"style=simple,explode=false,name=x-api-version"`
	// Idempotency works by saving the resulting status code and body of the first request made for any given idempotency key, regardless of whether it succeeded or failed. Subsequent requests with the same key return the same result, including 500 errors.
	//
	// Currently supported on all POST calls that uses x-client-id & x-client-secret. To use enable, pass x-idempotency-key in the request header. The value of this header must be unique to each operation you are trying to do. One example can be to use the same order_id that you pass while creating orders
	//
	XIdempotencyKey *string `header:"style=simple,explode=false,name=x-idempotency-key"`
	// Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree
	XRequestID *string `header:"style=simple,explode=false,name=x-request-id"`
}

func (o *GetSettlementsRequest) GetAcceptMedia() *string {
	if o == nil {
		return nil
	}
	return o.AcceptMedia
}

func (o *GetSettlementsRequest) GetFetchSettlementReconRequest() *shared.FetchSettlementReconRequest {
	if o == nil {
		return nil
	}
	return o.FetchSettlementReconRequest
}

func (o *GetSettlementsRequest) GetXAPIVersion() string {
	if o == nil {
		return ""
	}
	return o.XAPIVersion
}

func (o *GetSettlementsRequest) GetXIdempotencyKey() *string {
	if o == nil {
		return nil
	}
	return o.XIdempotencyKey
}

func (o *GetSettlementsRequest) GetXRequestID() *string {
	if o == nil {
		return nil
	}
	return o.XRequestID
}

type GetSettlementsResponse struct {
	ContentType string
	// OK
	FetchSettlement interface{}
	Headers         map[string][]string
	StatusCode      int
	RawResponse     *http.Response
}

func (o *GetSettlementsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSettlementsResponse) GetFetchSettlement() interface{} {
	if o == nil {
		return nil
	}
	return o.FetchSettlement
}

func (o *GetSettlementsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *GetSettlementsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSettlementsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
