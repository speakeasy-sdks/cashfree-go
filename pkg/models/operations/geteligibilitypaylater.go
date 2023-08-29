// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
	"net/http"
)

type GetEligibilityPaylaterRequest struct {
	// Request body to check for eligibility for paylater
	EligibilityCardlessEMIRequest *shared.EligibilityCardlessEMIRequest `request:"mediaType=application/json"`
	// API version to be used. Format is in YYYY-MM-DD
	XAPIVersion string `header:"style=simple,explode=false,name=x-api-version"`
	// Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree
	XRequestID *string `header:"style=simple,explode=false,name=x-request-id"`
}

func (o *GetEligibilityPaylaterRequest) GetEligibilityCardlessEMIRequest() *shared.EligibilityCardlessEMIRequest {
	if o == nil {
		return nil
	}
	return o.EligibilityCardlessEMIRequest
}

func (o *GetEligibilityPaylaterRequest) GetXAPIVersion() string {
	if o == nil {
		return ""
	}
	return o.XAPIVersion
}

func (o *GetEligibilityPaylaterRequest) GetXRequestID() *string {
	if o == nil {
		return nil
	}
	return o.XRequestID
}

type GetEligibilityPaylaterResponse struct {
	ContentType string
	// OK
	EligiblePaylaters []shared.EligiblePaylater
	Headers           map[string][]string
	StatusCode        int
	RawResponse       *http.Response
}

func (o *GetEligibilityPaylaterResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetEligibilityPaylaterResponse) GetEligiblePaylaters() []shared.EligiblePaylater {
	if o == nil {
		return nil
	}
	return o.EligiblePaylaters
}

func (o *GetEligibilityPaylaterResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *GetEligibilityPaylaterResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetEligibilityPaylaterResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
