// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
	"github.com/speakeasy-sdks/cashfree-go/pkg/utils"
	"net/http"
)

type CreateRefundRequest struct {
	// Request body to fetch refunds for an order
	CreateRefundRequest *shared.CreateRefundRequest `request:"mediaType=application/json"`
	// Order or the invoice ID for which refund must be initiated.
	OrderID string `pathParam:"style=simple,explode=false,name=order_id"`
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

func (c CreateRefundRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateRefundRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateRefundRequest) GetCreateRefundRequest() *shared.CreateRefundRequest {
	if o == nil {
		return nil
	}
	return o.CreateRefundRequest
}

func (o *CreateRefundRequest) GetOrderID() string {
	if o == nil {
		return ""
	}
	return o.OrderID
}

func (o *CreateRefundRequest) GetXAPIVersion() string {
	if o == nil {
		return ""
	}
	return o.XAPIVersion
}

func (o *CreateRefundRequest) GetXIdempotencyKey() *string {
	if o == nil {
		return nil
	}
	return o.XIdempotencyKey
}

func (o *CreateRefundRequest) GetXRequestID() *string {
	if o == nil {
		return nil
	}
	return o.XRequestID
}

type CreateRefundResponse struct {
	ContentType string
	Headers     map[string][]string
	// Refund created
	RefundsEntity *shared.RefundsEntity
	StatusCode    int
	RawResponse   *http.Response
}

func (o *CreateRefundResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateRefundResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CreateRefundResponse) GetRefundsEntity() *shared.RefundsEntity {
	if o == nil {
		return nil
	}
	return o.RefundsEntity
}

func (o *CreateRefundResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateRefundResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
