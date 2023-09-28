// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"errors"
	"github.com/speakeasy-sdks/cashfree-go/pkg/utils"
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

	offerAll := new(OfferAll)
	if err := utils.UnmarshalJSON(data, &offerAll, "", true, true); err == nil {
		u.OfferAll = offerAll
		u.Type = OfferValidationsPaymentMethodTypeOfferAll
		return nil
	}

	offerCard := new(OfferCard)
	if err := utils.UnmarshalJSON(data, &offerCard, "", true, true); err == nil {
		u.OfferCard = offerCard
		u.Type = OfferValidationsPaymentMethodTypeOfferCard
		return nil
	}

	offerNB := new(OfferNB)
	if err := utils.UnmarshalJSON(data, &offerNB, "", true, true); err == nil {
		u.OfferNB = offerNB
		u.Type = OfferValidationsPaymentMethodTypeOfferNB
		return nil
	}

	offerWallet := new(OfferWallet)
	if err := utils.UnmarshalJSON(data, &offerWallet, "", true, true); err == nil {
		u.OfferWallet = offerWallet
		u.Type = OfferValidationsPaymentMethodTypeOfferWallet
		return nil
	}

	offerUPI := new(OfferUPI)
	if err := utils.UnmarshalJSON(data, &offerUPI, "", true, true); err == nil {
		u.OfferUPI = offerUPI
		u.Type = OfferValidationsPaymentMethodTypeOfferUPI
		return nil
	}

	offerPaylater := new(OfferPaylater)
	if err := utils.UnmarshalJSON(data, &offerPaylater, "", true, true); err == nil {
		u.OfferPaylater = offerPaylater
		u.Type = OfferValidationsPaymentMethodTypeOfferPaylater
		return nil
	}

	offerEMI := new(OfferEMI)
	if err := utils.UnmarshalJSON(data, &offerEMI, "", true, true); err == nil {
		u.OfferEMI = offerEMI
		u.Type = OfferValidationsPaymentMethodTypeOfferEMI
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u OfferValidationsPaymentMethod) MarshalJSON() ([]byte, error) {
	if u.OfferAll != nil {
		return utils.MarshalJSON(u.OfferAll, "", true)
	}

	if u.OfferCard != nil {
		return utils.MarshalJSON(u.OfferCard, "", true)
	}

	if u.OfferNB != nil {
		return utils.MarshalJSON(u.OfferNB, "", true)
	}

	if u.OfferWallet != nil {
		return utils.MarshalJSON(u.OfferWallet, "", true)
	}

	if u.OfferUPI != nil {
		return utils.MarshalJSON(u.OfferUPI, "", true)
	}

	if u.OfferPaylater != nil {
		return utils.MarshalJSON(u.OfferPaylater, "", true)
	}

	if u.OfferEMI != nil {
		return utils.MarshalJSON(u.OfferEMI, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
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
