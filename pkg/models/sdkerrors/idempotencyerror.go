// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"fmt"
)

// IdempotencyErrorType - idempotency_error
type IdempotencyErrorType string

const (
	IdempotencyErrorTypeIdempotencyError IdempotencyErrorType = "idempotency_error"
)

func (e IdempotencyErrorType) ToPointer() *IdempotencyErrorType {
	return &e
}

func (e *IdempotencyErrorType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "idempotency_error":
		*e = IdempotencyErrorType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for IdempotencyErrorType: %v", v)
	}
}

type IdempotencyError struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	// idempotency_error
	Type *IdempotencyErrorType `json:"type,omitempty"`
}

var _ error = &IdempotencyError{}

func (e *IdempotencyError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
