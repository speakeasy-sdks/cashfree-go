// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SecurityOption1 struct {
	XClientID     string `security:"scheme,type=apiKey,subtype=header,name=x-client-id"`
	XClientSecret string `security:"scheme,type=apiKey,subtype=header,name=x-client-secret"`
}

func (o *SecurityOption1) GetXClientID() string {
	if o == nil {
		return ""
	}
	return o.XClientID
}

func (o *SecurityOption1) GetXClientSecret() string {
	if o == nil {
		return ""
	}
	return o.XClientSecret
}

type SecurityOption2 struct {
	XClientID      string `security:"scheme,type=apiKey,subtype=header,name=x-client-id"`
	XPartnerAPIKey string `security:"scheme,type=apiKey,subtype=header,name=x-partner-apikey"`
}

func (o *SecurityOption2) GetXClientID() string {
	if o == nil {
		return ""
	}
	return o.XClientID
}

func (o *SecurityOption2) GetXPartnerAPIKey() string {
	if o == nil {
		return ""
	}
	return o.XPartnerAPIKey
}

type SecurityOption3 struct {
	XClientID              string `security:"scheme,type=apiKey,subtype=header,name=x-client-id"`
	XClientSignatureHeader string `security:"scheme,type=apiKey,subtype=header,name=x-client-signature"`
}

func (o *SecurityOption3) GetXClientID() string {
	if o == nil {
		return ""
	}
	return o.XClientID
}

func (o *SecurityOption3) GetXClientSignatureHeader() string {
	if o == nil {
		return ""
	}
	return o.XClientSignatureHeader
}

type SecurityOption4 struct {
	XPartnerAPIKey     string `security:"scheme,type=apiKey,subtype=header,name=x-partner-apikey"`
	XPartnerMerchantID string `security:"scheme,type=apiKey,subtype=header,name=x-partner-merchantid"`
}

func (o *SecurityOption4) GetXPartnerAPIKey() string {
	if o == nil {
		return ""
	}
	return o.XPartnerAPIKey
}

func (o *SecurityOption4) GetXPartnerMerchantID() string {
	if o == nil {
		return ""
	}
	return o.XPartnerMerchantID
}

type Security struct {
	Option1 *SecurityOption1 `security:"option"`
	Option2 *SecurityOption2 `security:"option"`
	Option3 *SecurityOption3 `security:"option"`
	Option4 *SecurityOption4 `security:"option"`
}

func (o *Security) GetOption1() *SecurityOption1 {
	if o == nil {
		return nil
	}
	return o.Option1
}

func (o *Security) GetOption2() *SecurityOption2 {
	if o == nil {
		return nil
	}
	return o.Option2
}

func (o *Security) GetOption3() *SecurityOption3 {
	if o == nil {
		return nil
	}
	return o.Option3
}

func (o *Security) GetOption4() *SecurityOption4 {
	if o == nil {
		return nil
	}
	return o.Option4
}
