// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TerminalResponse - Terminal created
type TerminalResponse struct {
	AddedOn         *string `json:"added_on,omitempty"`
	CfTerminalID    *string `json:"cf_terminal_id,omitempty"`
	LastUpdatedOn   *string `json:"last_updated_on,omitempty"`
	TeminalID       *string `json:"teminal_id,omitempty"`
	TerminalArea    *string `json:"terminal_area,omitempty"`
	TerminalName    *string `json:"terminal_name,omitempty"`
	TerminalNote    *string `json:"terminal_note,omitempty"`
	TerminalPhoneNo *string `json:"terminal_phone_no,omitempty"`
	TerminalStatus  *string `json:"terminal_status,omitempty"`
}

func (o *TerminalResponse) GetAddedOn() *string {
	if o == nil {
		return nil
	}
	return o.AddedOn
}

func (o *TerminalResponse) GetCfTerminalID() *string {
	if o == nil {
		return nil
	}
	return o.CfTerminalID
}

func (o *TerminalResponse) GetLastUpdatedOn() *string {
	if o == nil {
		return nil
	}
	return o.LastUpdatedOn
}

func (o *TerminalResponse) GetTeminalID() *string {
	if o == nil {
		return nil
	}
	return o.TeminalID
}

func (o *TerminalResponse) GetTerminalArea() *string {
	if o == nil {
		return nil
	}
	return o.TerminalArea
}

func (o *TerminalResponse) GetTerminalName() *string {
	if o == nil {
		return nil
	}
	return o.TerminalName
}

func (o *TerminalResponse) GetTerminalNote() *string {
	if o == nil {
		return nil
	}
	return o.TerminalNote
}

func (o *TerminalResponse) GetTerminalPhoneNo() *string {
	if o == nil {
		return nil
	}
	return o.TerminalPhoneNo
}

func (o *TerminalResponse) GetTerminalStatus() *string {
	if o == nil {
		return nil
	}
	return o.TerminalStatus
}