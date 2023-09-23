// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
	"github.com/speakeasy-sdks/cashfree-go/pkg/utils"
	"net/http"
)

type GetEligibilityOfferRequest struct {
	// API version to be used. Format is in YYYY-MM-DD
	XAPIVersion string `default:"2022-09-01" header:"style=simple,explode=false,name=x-api-version"`
	// Request body to check for eligibility for offers
	EligibilityOffersRequest *shared.EligibilityOffersRequest `request:"mediaType=application/json"`
	// Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree
	XRequestID *string `header:"style=simple,explode=false,name=x-request-id"`
}

func (g GetEligibilityOfferRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetEligibilityOfferRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetEligibilityOfferRequest) GetXAPIVersion() string {
	if o == nil {
		return ""
	}
	return o.XAPIVersion
}

func (o *GetEligibilityOfferRequest) GetEligibilityOffersRequest() *shared.EligibilityOffersRequest {
	if o == nil {
		return nil
	}
	return o.EligibilityOffersRequest
}

func (o *GetEligibilityOfferRequest) GetXRequestID() *string {
	if o == nil {
		return nil
	}
	return o.XRequestID
}

type GetEligibilityOfferResponse struct {
	ContentType string
	// OK
	EligibleOffersEntities []shared.EligibleOffersEntity
	Headers                map[string][]string
	StatusCode             int
	RawResponse            *http.Response
}

func (o *GetEligibilityOfferResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetEligibilityOfferResponse) GetEligibleOffersEntities() []shared.EligibleOffersEntity {
	if o == nil {
		return nil
	}
	return o.EligibleOffersEntities
}

func (o *GetEligibilityOfferResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *GetEligibilityOfferResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetEligibilityOfferResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
