// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type EligibilityOffersRequest struct {
	Filters *OfferFilters `json:"filters,omitempty"`
	Queries OfferQueries  `json:"queries"`
}

func (o *EligibilityOffersRequest) GetFilters() *OfferFilters {
	if o == nil {
		return nil
	}
	return o.Filters
}

func (o *EligibilityOffersRequest) GetQueries() OfferQueries {
	if o == nil {
		return OfferQueries{}
	}
	return o.Queries
}
