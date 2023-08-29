// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
	"net/http"
)

type CapturePreauthorizationRequest struct {
	// Request body to capture or void a transaction
	AuthorizationRequest *shared.AuthorizationRequest `request:"mediaType=application/json"`
	// The order or invoice ID for which you want to view the payment details.
	OrderID string `pathParam:"style=simple,explode=false,name=order_id"`
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

func (o *CapturePreauthorizationRequest) GetAuthorizationRequest() *shared.AuthorizationRequest {
	if o == nil {
		return nil
	}
	return o.AuthorizationRequest
}

func (o *CapturePreauthorizationRequest) GetOrderID() string {
	if o == nil {
		return ""
	}
	return o.OrderID
}

func (o *CapturePreauthorizationRequest) GetXAPIVersion() string {
	if o == nil {
		return ""
	}
	return o.XAPIVersion
}

func (o *CapturePreauthorizationRequest) GetXIdempotencyKey() *string {
	if o == nil {
		return nil
	}
	return o.XIdempotencyKey
}

func (o *CapturePreauthorizationRequest) GetXRequestID() *string {
	if o == nil {
		return nil
	}
	return o.XRequestID
}

type CapturePreauthorizationResponse struct {
	ContentType string
	Headers     map[string][]string
	// OK
	PaymentsEntity *shared.PaymentsEntity
	StatusCode     int
	RawResponse    *http.Response
}

func (o *CapturePreauthorizationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CapturePreauthorizationResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CapturePreauthorizationResponse) GetPaymentsEntity() *shared.PaymentsEntity {
	if o == nil {
		return nil
	}
	return o.PaymentsEntity
}

func (o *CapturePreauthorizationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CapturePreauthorizationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}