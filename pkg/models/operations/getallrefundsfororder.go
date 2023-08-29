// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
	"net/http"
)

type GetAllRefundsForOrderRequest struct {
	// Order or the invoice ID for which you want to view the refund details.
	OrderID string `pathParam:"style=simple,explode=false,name=order_id"`
	// API version to be used. Format is in YYYY-MM-DD
	XAPIVersion string `header:"style=simple,explode=false,name=x-api-version"`
	// Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree
	XRequestID *string `header:"style=simple,explode=false,name=x-request-id"`
}

func (o *GetAllRefundsForOrderRequest) GetOrderID() string {
	if o == nil {
		return ""
	}
	return o.OrderID
}

func (o *GetAllRefundsForOrderRequest) GetXAPIVersion() string {
	if o == nil {
		return ""
	}
	return o.XAPIVersion
}

func (o *GetAllRefundsForOrderRequest) GetXRequestID() *string {
	if o == nil {
		return nil
	}
	return o.XRequestID
}

type GetAllRefundsForOrderResponse struct {
	ContentType string
	Headers     map[string][]string
	// OK
	RefundsEntities []shared.RefundsEntity
	StatusCode      int
	RawResponse     *http.Response
}

func (o *GetAllRefundsForOrderResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAllRefundsForOrderResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *GetAllRefundsForOrderResponse) GetRefundsEntities() []shared.RefundsEntity {
	if o == nil {
		return nil
	}
	return o.RefundsEntities
}

func (o *GetAllRefundsForOrderResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAllRefundsForOrderResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}