// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PaymentMethodCardInPaymentsEntity struct {
	CardBankName           *string `json:"card_bank_name,omitempty"`
	CardCountry            *string `json:"card_country,omitempty"`
	CardNetwork            *string `json:"card_network,omitempty"`
	CardNetworkReferenceID *string `json:"card_network_reference_id,omitempty"`
	CardNumber             *string `json:"card_number,omitempty"`
	CardType               *string `json:"card_type,omitempty"`
	Channel                *string `json:"channel,omitempty"`
}

func (o *PaymentMethodCardInPaymentsEntity) GetCardBankName() *string {
	if o == nil {
		return nil
	}
	return o.CardBankName
}

func (o *PaymentMethodCardInPaymentsEntity) GetCardCountry() *string {
	if o == nil {
		return nil
	}
	return o.CardCountry
}

func (o *PaymentMethodCardInPaymentsEntity) GetCardNetwork() *string {
	if o == nil {
		return nil
	}
	return o.CardNetwork
}

func (o *PaymentMethodCardInPaymentsEntity) GetCardNetworkReferenceID() *string {
	if o == nil {
		return nil
	}
	return o.CardNetworkReferenceID
}

func (o *PaymentMethodCardInPaymentsEntity) GetCardNumber() *string {
	if o == nil {
		return nil
	}
	return o.CardNumber
}

func (o *PaymentMethodCardInPaymentsEntity) GetCardType() *string {
	if o == nil {
		return nil
	}
	return o.CardType
}

func (o *PaymentMethodCardInPaymentsEntity) GetChannel() *string {
	if o == nil {
		return nil
	}
	return o.Channel
}