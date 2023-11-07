// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
	"github.com/speakeasy-sdks/cashfree-go/pkg/utils"
	"net/http"
)

type GetPaymentsforOrderRequest struct {
	// Order or the invoice ID for which you want to view the payment details.
	OrderID string `pathParam:"style=simple,explode=false,name=order_id"`
	// API version to be used. Format is in YYYY-MM-DD
	XAPIVersion string `default:"2022-09-01" header:"style=simple,explode=false,name=x-api-version"`
	// Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree
	XRequestID *string `header:"style=simple,explode=false,name=x-request-id"`
}

func (g GetPaymentsforOrderRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetPaymentsforOrderRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetPaymentsforOrderRequest) GetOrderID() string {
	if o == nil {
		return ""
	}
	return o.OrderID
}

func (o *GetPaymentsforOrderRequest) GetXAPIVersion() string {
	if o == nil {
		return ""
	}
	return o.XAPIVersion
}

func (o *GetPaymentsforOrderRequest) GetXRequestID() *string {
	if o == nil {
		return nil
	}
	return o.XRequestID
}

type GetPaymentsforOrderResponseBodyType string

const (
	GetPaymentsforOrderResponseBodyTypePaymentsEntity GetPaymentsforOrderResponseBodyType = "PaymentsEntity"
)

type GetPaymentsforOrderResponseBody struct {
	PaymentsEntity *shared.PaymentsEntity

	Type GetPaymentsforOrderResponseBodyType
}

func CreateGetPaymentsforOrderResponseBodyPaymentsEntity(paymentsEntity shared.PaymentsEntity) GetPaymentsforOrderResponseBody {
	typ := GetPaymentsforOrderResponseBodyTypePaymentsEntity

	return GetPaymentsforOrderResponseBody{
		PaymentsEntity: &paymentsEntity,
		Type:           typ,
	}
}

func (u *GetPaymentsforOrderResponseBody) UnmarshalJSON(data []byte) error {

	paymentsEntity := shared.PaymentsEntity{}
	if err := utils.UnmarshalJSON(data, &paymentsEntity, "", true, true); err == nil {
		u.PaymentsEntity = &paymentsEntity
		u.Type = GetPaymentsforOrderResponseBodyTypePaymentsEntity
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetPaymentsforOrderResponseBody) MarshalJSON() ([]byte, error) {
	if u.PaymentsEntity != nil {
		return utils.MarshalJSON(u.PaymentsEntity, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetPaymentsforOrderResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	OneOf *GetPaymentsforOrderResponseBody
}

func (o *GetPaymentsforOrderResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPaymentsforOrderResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *GetPaymentsforOrderResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPaymentsforOrderResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetPaymentsforOrderResponse) GetOneOf() *GetPaymentsforOrderResponseBody {
	if o == nil {
		return nil
	}
	return o.OneOf
}
