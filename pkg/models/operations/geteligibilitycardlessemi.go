// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
	"net/http"
)

type GetEligibilityCardlessEMIRequest struct {
	// Request body to check for eligibility for cardlessemi
	EligibilityCardlessEMIRequest *shared.EligibilityCardlessEMIRequest `request:"mediaType=application/json"`
	// API version to be used. Format is in YYYY-MM-DD
	XAPIVersion string `header:"style=simple,explode=false,name=x-api-version"`
	// Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree
	XRequestID *string `header:"style=simple,explode=false,name=x-request-id"`
}

func (o *GetEligibilityCardlessEMIRequest) GetEligibilityCardlessEMIRequest() *shared.EligibilityCardlessEMIRequest {
	if o == nil {
		return nil
	}
	return o.EligibilityCardlessEMIRequest
}

func (o *GetEligibilityCardlessEMIRequest) GetXAPIVersion() string {
	if o == nil {
		return ""
	}
	return o.XAPIVersion
}

func (o *GetEligibilityCardlessEMIRequest) GetXRequestID() *string {
	if o == nil {
		return nil
	}
	return o.XRequestID
}

type GetEligibilityCardlessEMIResponse struct {
	ContentType string
	// OK
	EligibleCardlessEMIEntities []shared.EligibleCardlessEMIEntity
	Headers                     map[string][]string
	StatusCode                  int
	RawResponse                 *http.Response
}

func (o *GetEligibilityCardlessEMIResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetEligibilityCardlessEMIResponse) GetEligibleCardlessEMIEntities() []shared.EligibleCardlessEMIEntity {
	if o == nil {
		return nil
	}
	return o.EligibleCardlessEMIEntities
}

func (o *GetEligibilityCardlessEMIResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *GetEligibilityCardlessEMIResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetEligibilityCardlessEMIResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
