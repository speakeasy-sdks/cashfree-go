// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type OfferUPIUPIOffer struct {
}

type OfferUPI struct {
	Upi OfferUPIUPIOffer `json:"upi"`
}

func (o *OfferUPI) GetUpi() OfferUPIUPIOffer {
	if o == nil {
		return OfferUPIUPIOffer{}
	}
	return o.Upi
}
