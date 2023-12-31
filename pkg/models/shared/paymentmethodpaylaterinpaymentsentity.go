// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PaymentMethodPaylaterInPaymentsEntity struct {
	Channel  *string `json:"channel,omitempty"`
	Phone    *string `json:"phone,omitempty"`
	Provider *string `json:"provider,omitempty"`
}

func (o *PaymentMethodPaylaterInPaymentsEntity) GetChannel() *string {
	if o == nil {
		return nil
	}
	return o.Channel
}

func (o *PaymentMethodPaylaterInPaymentsEntity) GetPhone() *string {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *PaymentMethodPaylaterInPaymentsEntity) GetProvider() *string {
	if o == nil {
		return nil
	}
	return o.Provider
}
