// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// AuthorizationRequestAction - Type of authorization to run. Can be one of 'CAPTURE' , 'VOID'
type AuthorizationRequestAction string

const (
	AuthorizationRequestActionCapture AuthorizationRequestAction = "CAPTURE"
	AuthorizationRequestActionVoid    AuthorizationRequestAction = "VOID"
)

func (e AuthorizationRequestAction) ToPointer() *AuthorizationRequestAction {
	return &e
}

func (e *AuthorizationRequestAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CAPTURE":
		fallthrough
	case "VOID":
		*e = AuthorizationRequestAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AuthorizationRequestAction: %v", v)
	}
}

type AuthorizationRequest struct {
	// Type of authorization to run. Can be one of 'CAPTURE' , 'VOID'
	Action *AuthorizationRequestAction `json:"action,omitempty"`
	// The amount if you are running a 'CAPTURE'
	Amount *float64 `json:"amount,omitempty"`
}

func (o *AuthorizationRequest) GetAction() *AuthorizationRequestAction {
	if o == nil {
		return nil
	}
	return o.Action
}

func (o *AuthorizationRequest) GetAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.Amount
}
