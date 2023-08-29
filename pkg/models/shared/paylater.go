// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// PaylaterProvider - One of ["kotak", "flexipay", "zestmoney", "lazypay", "olapostpaid","simpl", "freechargepaylater"]. Please note that Flexipay is offered by HDFC bank
type PaylaterProvider string

const (
	PaylaterProviderKotak              PaylaterProvider = "kotak"
	PaylaterProviderFlexipay           PaylaterProvider = "flexipay"
	PaylaterProviderZestmoney          PaylaterProvider = "zestmoney"
	PaylaterProviderLazypay            PaylaterProvider = "lazypay"
	PaylaterProviderOlapostpaid        PaylaterProvider = "olapostpaid"
	PaylaterProviderSimpl              PaylaterProvider = "simpl"
	PaylaterProviderFreechargepaylater PaylaterProvider = "freechargepaylater"
)

func (e PaylaterProvider) ToPointer() *PaylaterProvider {
	return &e
}

func (e *PaylaterProvider) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "kotak":
		fallthrough
	case "flexipay":
		fallthrough
	case "zestmoney":
		fallthrough
	case "lazypay":
		fallthrough
	case "olapostpaid":
		fallthrough
	case "simpl":
		fallthrough
	case "freechargepaylater":
		*e = PaylaterProvider(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaylaterProvider: %v", v)
	}
}

type Paylater struct {
	// The channel for cardless EMI is always `link`
	Channel *string `json:"channel,omitempty"`
	// Customers phone number for this payment instrument. If the customer is not eligible you will receive a 400 error with type as 'invalid_request_error' and code as 'invalid_request_error'
	Phone *string `json:"phone,omitempty"`
	// One of ["kotak", "flexipay", "zestmoney", "lazypay", "olapostpaid","simpl", "freechargepaylater"]. Please note that Flexipay is offered by HDFC bank
	Provider *PaylaterProvider `json:"provider,omitempty"`
}

func (o *Paylater) GetChannel() *string {
	if o == nil {
		return nil
	}
	return o.Channel
}

func (o *Paylater) GetPhone() *string {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *Paylater) GetProvider() *PaylaterProvider {
	if o == nil {
		return nil
	}
	return o.Provider
}