// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type OrderPayDataPayload struct {
}

type OrderPayData struct {
	ContentType *string              `json:"content_type,omitempty"`
	Method      *string              `json:"method,omitempty"`
	Payload     *OrderPayDataPayload `json:"payload,omitempty"`
	URL         *string              `json:"url,omitempty"`
}

func (o *OrderPayData) GetContentType() *string {
	if o == nil {
		return nil
	}
	return o.ContentType
}

func (o *OrderPayData) GetMethod() *string {
	if o == nil {
		return nil
	}
	return o.Method
}

func (o *OrderPayData) GetPayload() *OrderPayDataPayload {
	if o == nil {
		return nil
	}
	return o.Payload
}

func (o *OrderPayData) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}
