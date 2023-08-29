// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
	"net/http"
)

type FetchSpecificSavedInstrumentRequest struct {
	// The customer_id for which all the saved cards are queried
	CustomerID string `pathParam:"style=simple,explode=false,name=customer_id"`
	// The instrument_id of the saved instrument which needs to be queried
	InstrumentID string `pathParam:"style=simple,explode=false,name=instrument_id"`
	// API version to be used. Format is in YYYY-MM-DD
	XAPIVersion string `header:"style=simple,explode=false,name=x-api-version"`
	// Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree
	XRequestID *string `header:"style=simple,explode=false,name=x-request-id"`
}

func (o *FetchSpecificSavedInstrumentRequest) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

func (o *FetchSpecificSavedInstrumentRequest) GetInstrumentID() string {
	if o == nil {
		return ""
	}
	return o.InstrumentID
}

func (o *FetchSpecificSavedInstrumentRequest) GetXAPIVersion() string {
	if o == nil {
		return ""
	}
	return o.XAPIVersion
}

func (o *FetchSpecificSavedInstrumentRequest) GetXRequestID() *string {
	if o == nil {
		return nil
	}
	return o.XRequestID
}

type FetchSpecificSavedInstrumentResponse struct {
	ContentType string
	// OK
	FetchAllSavedInstruments *shared.FetchAllSavedInstruments
	Headers                  map[string][]string
	StatusCode               int
	RawResponse              *http.Response
}

func (o *FetchSpecificSavedInstrumentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *FetchSpecificSavedInstrumentResponse) GetFetchAllSavedInstruments() *shared.FetchAllSavedInstruments {
	if o == nil {
		return nil
	}
	return o.FetchAllSavedInstruments
}

func (o *FetchSpecificSavedInstrumentResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *FetchSpecificSavedInstrumentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *FetchSpecificSavedInstrumentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}