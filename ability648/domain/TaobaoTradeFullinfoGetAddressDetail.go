package domain


type TaobaoTradeFullinfoGetAddressDetail struct {
    /*
        详细地址     */
    DetailedAddress  *string `json:"detailed_address,omitempty" `

    /*
        行政区编码     */
    DivisionCode  *string `json:"division_code,omitempty" `

    /*
        区     */
    Area  *string `json:"area,omitempty" `

    /*
        电话号码     */
    Telephone  *string `json:"telephone,omitempty" `

    /*
        城市     */
    City  *string `json:"city,omitempty" `

    /*
        省份     */
    Prov  *string `json:"prov,omitempty" `

    /*
        国家     */
    Country  *string `json:"country,omitempty" `

    /*
        国家     */
    ContactName  *string `json:"contact_name,omitempty" `

    /*
        业务类型     */
    BizType  *string `json:"biz_type,omitempty" `

}

func (s *TaobaoTradeFullinfoGetAddressDetail) SetDetailedAddress(v string) *TaobaoTradeFullinfoGetAddressDetail {
    s.DetailedAddress = &v
    return s
}
func (s *TaobaoTradeFullinfoGetAddressDetail) SetDivisionCode(v string) *TaobaoTradeFullinfoGetAddressDetail {
    s.DivisionCode = &v
    return s
}
func (s *TaobaoTradeFullinfoGetAddressDetail) SetArea(v string) *TaobaoTradeFullinfoGetAddressDetail {
    s.Area = &v
    return s
}
func (s *TaobaoTradeFullinfoGetAddressDetail) SetTelephone(v string) *TaobaoTradeFullinfoGetAddressDetail {
    s.Telephone = &v
    return s
}
func (s *TaobaoTradeFullinfoGetAddressDetail) SetCity(v string) *TaobaoTradeFullinfoGetAddressDetail {
    s.City = &v
    return s
}
func (s *TaobaoTradeFullinfoGetAddressDetail) SetProv(v string) *TaobaoTradeFullinfoGetAddressDetail {
    s.Prov = &v
    return s
}
func (s *TaobaoTradeFullinfoGetAddressDetail) SetCountry(v string) *TaobaoTradeFullinfoGetAddressDetail {
    s.Country = &v
    return s
}
func (s *TaobaoTradeFullinfoGetAddressDetail) SetContactName(v string) *TaobaoTradeFullinfoGetAddressDetail {
    s.ContactName = &v
    return s
}
func (s *TaobaoTradeFullinfoGetAddressDetail) SetBizType(v string) *TaobaoTradeFullinfoGetAddressDetail {
    s.BizType = &v
    return s
}
