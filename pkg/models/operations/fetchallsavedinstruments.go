// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/cashfree-go/pkg/models/shared"
	"net/http"
)

// FetchAllSavedInstrumentsInstrumentType - type to instrument to query
type FetchAllSavedInstrumentsInstrumentType string

const (
	FetchAllSavedInstrumentsInstrumentTypeCard FetchAllSavedInstrumentsInstrumentType = "card"
)

func (e FetchAllSavedInstrumentsInstrumentType) ToPointer() *FetchAllSavedInstrumentsInstrumentType {
	return &e
}

func (e *FetchAllSavedInstrumentsInstrumentType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "card":
		*e = FetchAllSavedInstrumentsInstrumentType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FetchAllSavedInstrumentsInstrumentType: %v", v)
	}
}

type FetchAllSavedInstrumentsRequest struct {
	// The customer_id for which all the saved cards are queried
	CustomerID string `pathParam:"style=simple,explode=false,name=customer_id"`
	// type to instrument to query
	InstrumentType FetchAllSavedInstrumentsInstrumentType `queryParam:"style=form,explode=true,name=instrument_type"`
	// API version to be used. Format is in YYYY-MM-DD
	XAPIVersion string `header:"style=simple,explode=false,name=x-api-version"`
	// Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree
	XRequestID *string `header:"style=simple,explode=false,name=x-request-id"`
}

func (o *FetchAllSavedInstrumentsRequest) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

func (o *FetchAllSavedInstrumentsRequest) GetInstrumentType() FetchAllSavedInstrumentsInstrumentType {
	if o == nil {
		return FetchAllSavedInstrumentsInstrumentType("")
	}
	return o.InstrumentType
}

func (o *FetchAllSavedInstrumentsRequest) GetXAPIVersion() string {
	if o == nil {
		return ""
	}
	return o.XAPIVersion
}

func (o *FetchAllSavedInstrumentsRequest) GetXRequestID() *string {
	if o == nil {
		return nil
	}
	return o.XRequestID
}

type FetchAllSavedInstrumentsResponse struct {
	ContentType string
	// OK
	FetchAllSavedInstruments []shared.FetchAllSavedInstruments
	Headers                  map[string][]string
	StatusCode               int
	RawResponse              *http.Response
}

func (o *FetchAllSavedInstrumentsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *FetchAllSavedInstrumentsResponse) GetFetchAllSavedInstruments() []shared.FetchAllSavedInstruments {
	if o == nil {
		return nil
	}
	return o.FetchAllSavedInstruments
}

func (o *FetchAllSavedInstrumentsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *FetchAllSavedInstrumentsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *FetchAllSavedInstrumentsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}