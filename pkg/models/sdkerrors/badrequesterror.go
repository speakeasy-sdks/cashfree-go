// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"fmt"
)

type SchemasBadRequestErrorType string

const (
	SchemasBadRequestErrorTypeInvalidRequestError SchemasBadRequestErrorType = "invalid_request_error"
)

func (e SchemasBadRequestErrorType) ToPointer() *SchemasBadRequestErrorType {
	return &e
}

func (e *SchemasBadRequestErrorType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "invalid_request_error":
		*e = SchemasBadRequestErrorType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SchemasBadRequestErrorType: %v", v)
	}
}

// BadRequestError - Bad request error
type BadRequestError struct {
	Code    *string                     `json:"code,omitempty"`
	Message *string                     `json:"message,omitempty"`
	Type    *SchemasBadRequestErrorType `json:"type,omitempty"`
}

var _ error = &BadRequestError{}

func (e *BadRequestError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
