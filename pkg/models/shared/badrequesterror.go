// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type BadRequestErrorType string

const (
	BadRequestErrorTypeInvalidRequestError BadRequestErrorType = "invalid_request_error"
)

func (e BadRequestErrorType) ToPointer() *BadRequestErrorType {
	return &e
}

func (e *BadRequestErrorType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "invalid_request_error":
		*e = BadRequestErrorType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BadRequestErrorType: %v", v)
	}
}
