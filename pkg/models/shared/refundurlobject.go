// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type RefundURLObject struct {
	URL *string `json:"url,omitempty"`
}

func (o *RefundURLObject) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}
