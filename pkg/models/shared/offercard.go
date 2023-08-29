// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type OfferCardCardOffer struct {
	// Bank Name of Card.
	BankName   string   `json:"bank_name"`
	SchemeName []string `json:"scheme_name"`
	Type       []string `json:"type"`
}

func (o *OfferCardCardOffer) GetBankName() string {
	if o == nil {
		return ""
	}
	return o.BankName
}

func (o *OfferCardCardOffer) GetSchemeName() []string {
	if o == nil {
		return []string{}
	}
	return o.SchemeName
}

func (o *OfferCardCardOffer) GetType() []string {
	if o == nil {
		return []string{}
	}
	return o.Type
}

type OfferCard struct {
	Card OfferCardCardOffer `json:"card"`
}

func (o *OfferCard) GetCard() OfferCardCardOffer {
	if o == nil {
		return OfferCardCardOffer{}
	}
	return o.Card
}
