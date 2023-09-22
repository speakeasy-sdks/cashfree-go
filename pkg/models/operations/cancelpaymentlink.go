// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
	"github.com/speakeasy-sdks/cashfree-go/pkg/utils"
	"net/http"
)

type CancelPaymentLinkRequest struct {
	// The payment link ID which you want to cancel.
	LinkID string `pathParam:"style=simple,explode=false,name=link_id"`
	// API version to be used. Format is in YYYY-MM-DD
	XAPIVersion string `default:"2022-09-01" header:"style=simple,explode=false,name=x-api-version"`
	// Idempotency works by saving the resulting status code and body of the first request made for any given idempotency key, regardless of whether it succeeded or failed. Subsequent requests with the same key return the same result, including 500 errors.
	//
	// Currently supported on all POST calls that uses x-client-id & x-client-secret. To use enable, pass x-idempotency-key in the request header. The value of this header must be unique to each operation you are trying to do. One example can be to use the same order_id that you pass while creating orders
	//
	XIdempotencyKey *string `header:"style=simple,explode=false,name=x-idempotency-key"`
	// Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree
	XRequestID *string `header:"style=simple,explode=false,name=x-request-id"`
}

func (c CancelPaymentLinkRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CancelPaymentLinkRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CancelPaymentLinkRequest) GetLinkID() string {
	if o == nil {
		return ""
	}
	return o.LinkID
}

func (o *CancelPaymentLinkRequest) GetXAPIVersion() string {
	if o == nil {
		return ""
	}
	return o.XAPIVersion
}

func (o *CancelPaymentLinkRequest) GetXIdempotencyKey() *string {
	if o == nil {
		return nil
	}
	return o.XIdempotencyKey
}

func (o *CancelPaymentLinkRequest) GetXRequestID() *string {
	if o == nil {
		return nil
	}
	return o.XRequestID
}

type CancelPaymentLinkResponse struct {
	ContentType string
	Headers     map[string][]string
	// Payment Link cancelled
	LinkCancelledResponse *shared.LinkCancelledResponse
	StatusCode            int
	RawResponse           *http.Response
}

func (o *CancelPaymentLinkResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CancelPaymentLinkResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CancelPaymentLinkResponse) GetLinkCancelledResponse() *shared.LinkCancelledResponse {
	if o == nil {
		return nil
	}
	return o.LinkCancelledResponse
}

func (o *CancelPaymentLinkResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CancelPaymentLinkResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
