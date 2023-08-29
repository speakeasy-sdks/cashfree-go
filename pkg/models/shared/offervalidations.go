// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
)

type OfferValidationsPaymentMethodType string

const (
	OfferValidationsPaymentMethodTypeOfferAll      OfferValidationsPaymentMethodType = "OfferAll"
	OfferValidationsPaymentMethodTypeOfferCard     OfferValidationsPaymentMethodType = "OfferCard"
	OfferValidationsPaymentMethodTypeOfferNB       OfferValidationsPaymentMethodType = "OfferNB"
	OfferValidationsPaymentMethodTypeOfferWallet   OfferValidationsPaymentMethodType = "OfferWallet"
	OfferValidationsPaymentMethodTypeOfferUPI      OfferValidationsPaymentMethodType = "OfferUPI"
	OfferValidationsPaymentMethodTypeOfferPaylater OfferValidationsPaymentMethodType = "OfferPaylater"
	OfferValidationsPaymentMethodTypeOfferEMI      OfferValidationsPaymentMethodType = "OfferEMI"
)

type OfferValidationsPaymentMethod struct {
	OfferAll      *OfferAll
	OfferCard     *OfferCard
	OfferNB       *OfferNB
	OfferWallet   *OfferWallet
	OfferUPI      *OfferUPI
	OfferPaylater *OfferPaylater
	OfferEMI      *OfferEMI

	Type OfferValidationsPaymentMethodType
}

func CreateOfferValidationsPaymentMethodOfferAll(offerAll OfferAll) OfferValidationsPaymentMethod {
	typ := OfferValidationsPaymentMethodTypeOfferAll

	return OfferValidationsPaymentMethod{
		OfferAll: &offerAll,
		Type:     typ,
	}
}

func CreateOfferValidationsPaymentMethodOfferCard(offerCard OfferCard) OfferValidationsPaymentMethod {
	typ := OfferValidationsPaymentMethodTypeOfferCard

	return OfferValidationsPaymentMethod{
		OfferCard: &offerCard,
		Type:      typ,
	}
}

func CreateOfferValidationsPaymentMethodOfferNB(offerNB OfferNB) OfferValidationsPaymentMethod {
	typ := OfferValidationsPaymentMethodTypeOfferNB

	return OfferValidationsPaymentMethod{
		OfferNB: &offerNB,
		Type:    typ,
	}
}

func CreateOfferValidationsPaymentMethodOfferWallet(offerWallet OfferWallet) OfferValidationsPaymentMethod {
	typ := OfferValidationsPaymentMethodTypeOfferWallet

	return OfferValidationsPaymentMethod{
		OfferWallet: &offerWallet,
		Type:        typ,
	}
}

func CreateOfferValidationsPaymentMethodOfferUPI(offerUPI OfferUPI) OfferValidationsPaymentMethod {
	typ := OfferValidationsPaymentMethodTypeOfferUPI

	return OfferValidationsPaymentMethod{
		OfferUPI: &offerUPI,
		Type:     typ,
	}
}

func CreateOfferValidationsPaymentMethodOfferPaylater(offerPaylater OfferPaylater) OfferValidationsPaymentMethod {
	typ := OfferValidationsPaymentMethodTypeOfferPaylater

	return OfferValidationsPaymentMethod{
		OfferPaylater: &offerPaylater,
		Type:          typ,
	}
}

func CreateOfferValidationsPaymentMethodOfferEMI(offerEMI OfferEMI) OfferValidationsPaymentMethod {
	typ := OfferValidationsPaymentMethodTypeOfferEMI

	return OfferValidationsPaymentMethod{
		OfferEMI: &offerEMI,
		Type:     typ,
	}
}

func (u *OfferValidationsPaymentMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	offerAll := new(OfferAll)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&offerAll); err == nil {
		u.OfferAll = offerAll
		u.Type = OfferValidationsPaymentMethodTypeOfferAll
		return nil
	}

	offerCard := new(OfferCard)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&offerCard); err == nil {
		u.OfferCard = offerCard
		u.Type = OfferValidationsPaymentMethodTypeOfferCard
		return nil
	}

	offerNB := new(OfferNB)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&offerNB); err == nil {
		u.OfferNB = offerNB
		u.Type = OfferValidationsPaymentMethodTypeOfferNB
		return nil
	}

	offerWallet := new(OfferWallet)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&offerWallet); err == nil {
		u.OfferWallet = offerWallet
		u.Type = OfferValidationsPaymentMethodTypeOfferWallet
		return nil
	}

	offerUPI := new(OfferUPI)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&offerUPI); err == nil {
		u.OfferUPI = offerUPI
		u.Type = OfferValidationsPaymentMethodTypeOfferUPI
		return nil
	}

	offerPaylater := new(OfferPaylater)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&offerPaylater); err == nil {
		u.OfferPaylater = offerPaylater
		u.Type = OfferValidationsPaymentMethodTypeOfferPaylater
		return nil
	}

	offerEMI := new(OfferEMI)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&offerEMI); err == nil {
		u.OfferEMI = offerEMI
		u.Type = OfferValidationsPaymentMethodTypeOfferEMI
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u OfferValidationsPaymentMethod) MarshalJSON() ([]byte, error) {
	if u.OfferAll != nil {
		return json.Marshal(u.OfferAll)
	}

	if u.OfferCard != nil {
		return json.Marshal(u.OfferCard)
	}

	if u.OfferNB != nil {
		return json.Marshal(u.OfferNB)
	}

	if u.OfferWallet != nil {
		return json.Marshal(u.OfferWallet)
	}

	if u.OfferUPI != nil {
		return json.Marshal(u.OfferUPI)
	}

	if u.OfferPaylater != nil {
		return json.Marshal(u.OfferPaylater)
	}

	if u.OfferEMI != nil {
		return json.Marshal(u.OfferEMI)
	}

	return nil, nil
}

type OfferValidations struct {
	// Maximum amount of Offer that can be availed.
	MaxAllowed float64 `json:"max_allowed"`
	// Minimum Amount for Offer to be Applicable
	MinAmount     *float64                      `json:"min_amount,omitempty"`
	PaymentMethod OfferValidationsPaymentMethod `json:"payment_method"`
}

func (o *OfferValidations) GetMaxAllowed() float64 {
	if o == nil {
		return 0.0
	}
	return o.MaxAllowed
}

func (o *OfferValidations) GetMinAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.MinAmount
}

func (o *OfferValidations) GetPaymentMethod() OfferValidationsPaymentMethod {
	if o == nil {
		return OfferValidationsPaymentMethod{}
	}
	return o.PaymentMethod
}
