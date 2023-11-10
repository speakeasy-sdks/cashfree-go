// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type LinkCancelledResponse struct {
	CfLinkID                 *int64                     `json:"cf_link_id,omitempty"`
	CustomerDetails          *LinkCustomerDetailsEntity `json:"customer_details,omitempty"`
	LinkAmount               *float64                   `json:"link_amount,omitempty"`
	LinkAmountPaid           *float64                   `json:"link_amount_paid,omitempty"`
	LinkAutoReminders        *bool                      `json:"link_auto_reminders,omitempty"`
	LinkCreatedAt            *string                    `json:"link_created_at,omitempty"`
	LinkCurrency             *string                    `json:"link_currency,omitempty"`
	LinkExpiryTime           *string                    `json:"link_expiry_time,omitempty"`
	LinkID                   *string                    `json:"link_id,omitempty"`
	LinkMeta                 *LinkMetaEntity            `json:"link_meta,omitempty"`
	LinkMinimumPartialAmount *float64                   `json:"link_minimum_partial_amount,omitempty"`
	// Key-value pair that can be used to store additional information about the entity. Maximum 5 key-value pairs
	LinkNotes           map[string]string `json:"link_notes,omitempty"`
	LinkNotify          *LinkNotifyEntity `json:"link_notify,omitempty"`
	LinkPartialPayments *bool             `json:"link_partial_payments,omitempty"`
	LinkPurpose         *string           `json:"link_purpose,omitempty"`
	LinkStatus          *string           `json:"link_status,omitempty"`
	LinkURL             *string           `json:"link_url,omitempty"`
}

func (o *LinkCancelledResponse) GetCfLinkID() *int64 {
	if o == nil {
		return nil
	}
	return o.CfLinkID
}

func (o *LinkCancelledResponse) GetCustomerDetails() *LinkCustomerDetailsEntity {
	if o == nil {
		return nil
	}
	return o.CustomerDetails
}

func (o *LinkCancelledResponse) GetLinkAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.LinkAmount
}

func (o *LinkCancelledResponse) GetLinkAmountPaid() *float64 {
	if o == nil {
		return nil
	}
	return o.LinkAmountPaid
}

func (o *LinkCancelledResponse) GetLinkAutoReminders() *bool {
	if o == nil {
		return nil
	}
	return o.LinkAutoReminders
}

func (o *LinkCancelledResponse) GetLinkCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.LinkCreatedAt
}

func (o *LinkCancelledResponse) GetLinkCurrency() *string {
	if o == nil {
		return nil
	}
	return o.LinkCurrency
}

func (o *LinkCancelledResponse) GetLinkExpiryTime() *string {
	if o == nil {
		return nil
	}
	return o.LinkExpiryTime
}

func (o *LinkCancelledResponse) GetLinkID() *string {
	if o == nil {
		return nil
	}
	return o.LinkID
}

func (o *LinkCancelledResponse) GetLinkMeta() *LinkMetaEntity {
	if o == nil {
		return nil
	}
	return o.LinkMeta
}

func (o *LinkCancelledResponse) GetLinkMinimumPartialAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.LinkMinimumPartialAmount
}

func (o *LinkCancelledResponse) GetLinkNotes() map[string]string {
	if o == nil {
		return nil
	}
	return o.LinkNotes
}

func (o *LinkCancelledResponse) GetLinkNotify() *LinkNotifyEntity {
	if o == nil {
		return nil
	}
	return o.LinkNotify
}

func (o *LinkCancelledResponse) GetLinkPartialPayments() *bool {
	if o == nil {
		return nil
	}
	return o.LinkPartialPayments
}

func (o *LinkCancelledResponse) GetLinkPurpose() *string {
	if o == nil {
		return nil
	}
	return o.LinkPurpose
}

func (o *LinkCancelledResponse) GetLinkStatus() *string {
	if o == nil {
		return nil
	}
	return o.LinkStatus
}

func (o *LinkCancelledResponse) GetLinkURL() *string {
	if o == nil {
		return nil
	}
	return o.LinkURL
}
