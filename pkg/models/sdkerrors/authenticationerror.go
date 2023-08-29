// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
)

// AuthenticationError - Authentication Error
type AuthenticationError struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	// authentication_error
	Type *string `json:"type,omitempty"`
}

var _ error = &AuthenticationError{}

func (e *AuthenticationError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}