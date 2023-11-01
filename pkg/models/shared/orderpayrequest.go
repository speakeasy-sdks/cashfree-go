// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"errors"
	"github.com/speakeasy-sdks/cashfree-go/pkg/utils"
)

type OrderPayRequestPaymentMethodType string

const (
	OrderPayRequestPaymentMethodTypeCardPaymentMethod        OrderPayRequestPaymentMethodType = "CardPaymentMethod"
	OrderPayRequestPaymentMethodTypeUPIPaymentMethod         OrderPayRequestPaymentMethodType = "UPIPaymentMethod"
	OrderPayRequestPaymentMethodTypeNetBankingPaymentMethod  OrderPayRequestPaymentMethodType = "NetBankingPaymentMethod"
	OrderPayRequestPaymentMethodTypeAppPaymentMethod         OrderPayRequestPaymentMethodType = "AppPaymentMethod"
	OrderPayRequestPaymentMethodTypeCardEMIPaymentMethod     OrderPayRequestPaymentMethodType = "CardEMIPaymentMethod"
	OrderPayRequestPaymentMethodTypeCardlessEMIPaymentMethod OrderPayRequestPaymentMethodType = "CardlessEMIPaymentMethod"
	OrderPayRequestPaymentMethodTypePaylaterPaymentMethod    OrderPayRequestPaymentMethodType = "PaylaterPaymentMethod"
)

type OrderPayRequestPaymentMethod struct {
	CardPaymentMethod        *CardPaymentMethod
	UPIPaymentMethod         *UPIPaymentMethod
	NetBankingPaymentMethod  *NetBankingPaymentMethod
	AppPaymentMethod         *AppPaymentMethod
	CardEMIPaymentMethod     *CardEMIPaymentMethod
	CardlessEMIPaymentMethod *CardlessEMIPaymentMethod
	PaylaterPaymentMethod    *PaylaterPaymentMethod

	Type OrderPayRequestPaymentMethodType
}

func CreateOrderPayRequestPaymentMethodCardPaymentMethod(cardPaymentMethod CardPaymentMethod) OrderPayRequestPaymentMethod {
	typ := OrderPayRequestPaymentMethodTypeCardPaymentMethod

	return OrderPayRequestPaymentMethod{
		CardPaymentMethod: &cardPaymentMethod,
		Type:              typ,
	}
}

func CreateOrderPayRequestPaymentMethodUPIPaymentMethod(upiPaymentMethod UPIPaymentMethod) OrderPayRequestPaymentMethod {
	typ := OrderPayRequestPaymentMethodTypeUPIPaymentMethod

	return OrderPayRequestPaymentMethod{
		UPIPaymentMethod: &upiPaymentMethod,
		Type:             typ,
	}
}

func CreateOrderPayRequestPaymentMethodNetBankingPaymentMethod(netBankingPaymentMethod NetBankingPaymentMethod) OrderPayRequestPaymentMethod {
	typ := OrderPayRequestPaymentMethodTypeNetBankingPaymentMethod

	return OrderPayRequestPaymentMethod{
		NetBankingPaymentMethod: &netBankingPaymentMethod,
		Type:                    typ,
	}
}

func CreateOrderPayRequestPaymentMethodAppPaymentMethod(appPaymentMethod AppPaymentMethod) OrderPayRequestPaymentMethod {
	typ := OrderPayRequestPaymentMethodTypeAppPaymentMethod

	return OrderPayRequestPaymentMethod{
		AppPaymentMethod: &appPaymentMethod,
		Type:             typ,
	}
}

func CreateOrderPayRequestPaymentMethodCardEMIPaymentMethod(cardEMIPaymentMethod CardEMIPaymentMethod) OrderPayRequestPaymentMethod {
	typ := OrderPayRequestPaymentMethodTypeCardEMIPaymentMethod

	return OrderPayRequestPaymentMethod{
		CardEMIPaymentMethod: &cardEMIPaymentMethod,
		Type:                 typ,
	}
}

func CreateOrderPayRequestPaymentMethodCardlessEMIPaymentMethod(cardlessEMIPaymentMethod CardlessEMIPaymentMethod) OrderPayRequestPaymentMethod {
	typ := OrderPayRequestPaymentMethodTypeCardlessEMIPaymentMethod

	return OrderPayRequestPaymentMethod{
		CardlessEMIPaymentMethod: &cardlessEMIPaymentMethod,
		Type:                     typ,
	}
}

func CreateOrderPayRequestPaymentMethodPaylaterPaymentMethod(paylaterPaymentMethod PaylaterPaymentMethod) OrderPayRequestPaymentMethod {
	typ := OrderPayRequestPaymentMethodTypePaylaterPaymentMethod

	return OrderPayRequestPaymentMethod{
		PaylaterPaymentMethod: &paylaterPaymentMethod,
		Type:                  typ,
	}
}

func (u *OrderPayRequestPaymentMethod) UnmarshalJSON(data []byte) error {

	cardPaymentMethod := CardPaymentMethod{}
	if err := utils.UnmarshalJSON(data, &cardPaymentMethod, "", true, true); err == nil {
		u.CardPaymentMethod = &cardPaymentMethod
		u.Type = OrderPayRequestPaymentMethodTypeCardPaymentMethod
		return nil
	}

	upiPaymentMethod := UPIPaymentMethod{}
	if err := utils.UnmarshalJSON(data, &upiPaymentMethod, "", true, true); err == nil {
		u.UPIPaymentMethod = &upiPaymentMethod
		u.Type = OrderPayRequestPaymentMethodTypeUPIPaymentMethod
		return nil
	}

	netBankingPaymentMethod := NetBankingPaymentMethod{}
	if err := utils.UnmarshalJSON(data, &netBankingPaymentMethod, "", true, true); err == nil {
		u.NetBankingPaymentMethod = &netBankingPaymentMethod
		u.Type = OrderPayRequestPaymentMethodTypeNetBankingPaymentMethod
		return nil
	}

	appPaymentMethod := AppPaymentMethod{}
	if err := utils.UnmarshalJSON(data, &appPaymentMethod, "", true, true); err == nil {
		u.AppPaymentMethod = &appPaymentMethod
		u.Type = OrderPayRequestPaymentMethodTypeAppPaymentMethod
		return nil
	}

	cardEMIPaymentMethod := CardEMIPaymentMethod{}
	if err := utils.UnmarshalJSON(data, &cardEMIPaymentMethod, "", true, true); err == nil {
		u.CardEMIPaymentMethod = &cardEMIPaymentMethod
		u.Type = OrderPayRequestPaymentMethodTypeCardEMIPaymentMethod
		return nil
	}

	cardlessEMIPaymentMethod := CardlessEMIPaymentMethod{}
	if err := utils.UnmarshalJSON(data, &cardlessEMIPaymentMethod, "", true, true); err == nil {
		u.CardlessEMIPaymentMethod = &cardlessEMIPaymentMethod
		u.Type = OrderPayRequestPaymentMethodTypeCardlessEMIPaymentMethod
		return nil
	}

	paylaterPaymentMethod := PaylaterPaymentMethod{}
	if err := utils.UnmarshalJSON(data, &paylaterPaymentMethod, "", true, true); err == nil {
		u.PaylaterPaymentMethod = &paylaterPaymentMethod
		u.Type = OrderPayRequestPaymentMethodTypePaylaterPaymentMethod
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u OrderPayRequestPaymentMethod) MarshalJSON() ([]byte, error) {
	if u.CardPaymentMethod != nil {
		return utils.MarshalJSON(u.CardPaymentMethod, "", true)
	}

	if u.UPIPaymentMethod != nil {
		return utils.MarshalJSON(u.UPIPaymentMethod, "", true)
	}

	if u.NetBankingPaymentMethod != nil {
		return utils.MarshalJSON(u.NetBankingPaymentMethod, "", true)
	}

	if u.AppPaymentMethod != nil {
		return utils.MarshalJSON(u.AppPaymentMethod, "", true)
	}

	if u.CardEMIPaymentMethod != nil {
		return utils.MarshalJSON(u.CardEMIPaymentMethod, "", true)
	}

	if u.CardlessEMIPaymentMethod != nil {
		return utils.MarshalJSON(u.CardlessEMIPaymentMethod, "", true)
	}

	if u.PaylaterPaymentMethod != nil {
		return utils.MarshalJSON(u.PaylaterPaymentMethod, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type OrderPayRequest struct {
	// This is required if any offers needs to be applied to the order.
	OfferID          *string                      `json:"offer_id,omitempty"`
	PaymentMethod    OrderPayRequestPaymentMethod `json:"payment_method"`
	PaymentSessionID string                       `json:"payment_session_id"`
	SaveInstrument   *bool                        `json:"save_instrument,omitempty"`
}

func (o *OrderPayRequest) GetOfferID() *string {
	if o == nil {
		return nil
	}
	return o.OfferID
}

func (o *OrderPayRequest) GetPaymentMethod() OrderPayRequestPaymentMethod {
	if o == nil {
		return OrderPayRequestPaymentMethod{}
	}
	return o.PaymentMethod
}

func (o *OrderPayRequest) GetPaymentSessionID() string {
	if o == nil {
		return ""
	}
	return o.PaymentSessionID
}

func (o *OrderPayRequest) GetSaveInstrument() *bool {
	if o == nil {
		return nil
	}
	return o.SaveInstrument
}
