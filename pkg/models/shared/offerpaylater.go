// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PaylaterOffer struct {
	Provider *string `json:"provider,omitempty"`
}

func (o *PaylaterOffer) GetProvider() *string {
	if o == nil {
		return nil
	}
	return o.Provider
}

type OfferPaylater struct {
	Paylater PaylaterOffer `json:"paylater"`
}

func (o *OfferPaylater) GetPaylater() PaylaterOffer {
	if o == nil {
		return PaylaterOffer{}
	}
	return o.Paylater
}
