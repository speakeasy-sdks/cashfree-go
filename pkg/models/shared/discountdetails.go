// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// DiscountDetailsDiscountType - Type of discount
type DiscountDetailsDiscountType string

const (
	DiscountDetailsDiscountTypeFlat       DiscountDetailsDiscountType = "flat"
	DiscountDetailsDiscountTypePercentage DiscountDetailsDiscountType = "percentage"
)

func (e DiscountDetailsDiscountType) ToPointer() *DiscountDetailsDiscountType {
	return &e
}

func (e *DiscountDetailsDiscountType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "flat":
		fallthrough
	case "percentage":
		*e = DiscountDetailsDiscountType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DiscountDetailsDiscountType: %v", v)
	}
}

type DiscountDetails struct {
	// Type of discount
	DiscountType DiscountDetailsDiscountType `json:"discount_type"`
	// Value of Discount.
	DiscountValue string `json:"discount_value"`
	// Maximum Value of Discount allowed.
	MaxDiscountAmount string `json:"max_discount_amount"`
}

func (o *DiscountDetails) GetDiscountType() DiscountDetailsDiscountType {
	if o == nil {
		return DiscountDetailsDiscountType("")
	}
	return o.DiscountType
}

func (o *DiscountDetails) GetDiscountValue() string {
	if o == nil {
		return ""
	}
	return o.DiscountValue
}

func (o *DiscountDetails) GetMaxDiscountAmount() string {
	if o == nil {
		return ""
	}
	return o.MaxDiscountAmount
}