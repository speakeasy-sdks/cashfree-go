// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
	"github.com/speakeasy-sdks/cashfree-go/pkg/utils"
	"net/http"
)

type GetEligibilityCardlessEMIRequest struct {
	// API version to be used. Format is in YYYY-MM-DD
	XAPIVersion string `default:"2022-09-01" header:"style=simple,explode=false,name=x-api-version"`
	// Request body to check for eligibility for cardlessemi
	EligibilityCardlessEMIRequest *shared.EligibilityCardlessEMIRequest `request:"mediaType=application/json"`
	// Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree
	XRequestID *string `header:"style=simple,explode=false,name=x-request-id"`
}

func (g GetEligibilityCardlessEMIRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetEligibilityCardlessEMIRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetEligibilityCardlessEMIRequest) GetXAPIVersion() string {
	if o == nil {
		return ""
	}
	return o.XAPIVersion
}

func (o *GetEligibilityCardlessEMIRequest) GetEligibilityCardlessEMIRequest() *shared.EligibilityCardlessEMIRequest {
	if o == nil {
		return nil
	}
	return o.EligibilityCardlessEMIRequest
}

func (o *GetEligibilityCardlessEMIRequest) GetXRequestID() *string {
	if o == nil {
		return nil
	}
	return o.XRequestID
}

type GetEligibilityCardlessEMIResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Classes []shared.EligibleCardlessEMIEntity
}

func (o *GetEligibilityCardlessEMIResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetEligibilityCardlessEMIResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
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

func (o *GetEligibilityCardlessEMIResponse) GetClasses() []shared.EligibleCardlessEMIEntity {
	if o == nil {
		return nil
	}
	return o.Classes
}
