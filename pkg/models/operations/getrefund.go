// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
	"github.com/speakeasy-sdks/cashfree-go/pkg/utils"
	"net/http"
)

type GetRefundRequest struct {
	// Order or the invoice ID for which you want to view the refund details.
	OrderID string `pathParam:"style=simple,explode=false,name=order_id"`
	// Refund Id of the refund you want to fetch.
	RefundID string `pathParam:"style=simple,explode=false,name=refund_id"`
	// API version to be used. Format is in YYYY-MM-DD
	XAPIVersion string `default:"2022-09-01" header:"style=simple,explode=false,name=x-api-version"`
	// Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree
	XRequestID *string `header:"style=simple,explode=false,name=x-request-id"`
}

func (g GetRefundRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetRefundRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetRefundRequest) GetOrderID() string {
	if o == nil {
		return ""
	}
	return o.OrderID
}

func (o *GetRefundRequest) GetRefundID() string {
	if o == nil {
		return ""
	}
	return o.RefundID
}

func (o *GetRefundRequest) GetXAPIVersion() string {
	if o == nil {
		return ""
	}
	return o.XAPIVersion
}

func (o *GetRefundRequest) GetXRequestID() *string {
	if o == nil {
		return nil
	}
	return o.XRequestID
}

type GetRefundResponse struct {
	ContentType string
	Headers     map[string][]string
	// OK
	RefundsEntity *shared.RefundsEntity
	StatusCode    int
	RawResponse   *http.Response
}

func (o *GetRefundResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRefundResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *GetRefundResponse) GetRefundsEntity() *shared.RefundsEntity {
	if o == nil {
		return nil
	}
	return o.RefundsEntity
}

func (o *GetRefundResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRefundResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
