// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Cryptogram - OK
type Cryptogram struct {
	// last 4 digits of original card number
	CardDisplay *string `json:"card_display,omitempty"`
	// token pan expiry month
	CardExpiryMm *string `json:"card_expiry_mm,omitempty"`
	// token pan expiry year
	CardExpiryYy *string `json:"card_expiry_yy,omitempty"`
	// token pan number
	CardNumber *string `json:"card_number,omitempty"`
	// cryptogram
	Cryptogram *string `json:"cryptogram,omitempty"`
	// instrument_id of saved instrument
	InstrumentID *string `json:"instrument_id,omitempty"`
	// TRID issued by card networks
	TokenRequestorID *string `json:"token_requestor_id,omitempty"`
}

func (o *Cryptogram) GetCardDisplay() *string {
	if o == nil {
		return nil
	}
	return o.CardDisplay
}

func (o *Cryptogram) GetCardExpiryMm() *string {
	if o == nil {
		return nil
	}
	return o.CardExpiryMm
}

func (o *Cryptogram) GetCardExpiryYy() *string {
	if o == nil {
		return nil
	}
	return o.CardExpiryYy
}

func (o *Cryptogram) GetCardNumber() *string {
	if o == nil {
		return nil
	}
	return o.CardNumber
}

func (o *Cryptogram) GetCryptogram() *string {
	if o == nil {
		return nil
	}
	return o.Cryptogram
}

func (o *Cryptogram) GetInstrumentID() *string {
	if o == nil {
		return nil
	}
	return o.InstrumentID
}

func (o *Cryptogram) GetTokenRequestorID() *string {
	if o == nil {
		return nil
	}
	return o.TokenRequestorID
}