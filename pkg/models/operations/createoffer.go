// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
	"github.com/speakeasy-sdks/cashfree-go/pkg/utils"
	"net/http"
)

type CreateOfferRequest struct {
	// API version to be used. Format is in YYYY-MM-DD
	XAPIVersion string `default:"2022-09-01" header:"style=simple,explode=false,name=x-api-version"`
	// Request body to create an offer at Cashfree
	CreateOfferBackendRequest *shared.CreateOfferBackendRequest `request:"mediaType=application/json"`
	// Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree
	XRequestID *string `header:"style=simple,explode=false,name=x-request-id"`
}

func (c CreateOfferRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateOfferRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateOfferRequest) GetXAPIVersion() string {
	if o == nil {
		return ""
	}
	return o.XAPIVersion
}

func (o *CreateOfferRequest) GetCreateOfferBackendRequest() *shared.CreateOfferBackendRequest {
	if o == nil {
		return nil
	}
	return o.CreateOfferBackendRequest
}

func (o *CreateOfferRequest) GetXRequestID() *string {
	if o == nil {
		return nil
	}
	return o.XRequestID
}

type CreateOfferResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// OK
	OfferEntity *shared.OfferEntity
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateOfferResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateOfferResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CreateOfferResponse) GetOfferEntity() *shared.OfferEntity {
	if o == nil {
		return nil
	}
	return o.OfferEntity
}

func (o *CreateOfferResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateOfferResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
