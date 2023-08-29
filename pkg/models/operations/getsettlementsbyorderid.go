// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
	"net/http"
)

type GetSettlementsByOrderIDRequest struct {
	// Order ID for which you want to view the settlements.
	OrderID string `pathParam:"style=simple,explode=false,name=order_id"`
	// API version to be used. Format is in YYYY-MM-DD
	XAPIVersion string `header:"style=simple,explode=false,name=x-api-version"`
	// Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree
	XRequestID *string `header:"style=simple,explode=false,name=x-request-id"`
}

func (o *GetSettlementsByOrderIDRequest) GetOrderID() string {
	if o == nil {
		return ""
	}
	return o.OrderID
}

func (o *GetSettlementsByOrderIDRequest) GetXAPIVersion() string {
	if o == nil {
		return ""
	}
	return o.XAPIVersion
}

func (o *GetSettlementsByOrderIDRequest) GetXRequestID() *string {
	if o == nil {
		return nil
	}
	return o.XRequestID
}

type GetSettlementsByOrderIDResponse struct {
	ContentType string
	Headers     map[string][]string
	// OK
	SettlementsEntity *shared.SettlementsEntity
	StatusCode        int
	RawResponse       *http.Response
}

func (o *GetSettlementsByOrderIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSettlementsByOrderIDResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *GetSettlementsByOrderIDResponse) GetSettlementsEntity() *shared.SettlementsEntity {
	if o == nil {
		return nil
	}
	return o.SettlementsEntity
}

func (o *GetSettlementsByOrderIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSettlementsByOrderIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}