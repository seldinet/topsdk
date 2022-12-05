package domain


type TaobaoVmarketEticketCodesGetEticketCode struct {
    /*
        可用数量     */
    Num  *int64 `json:"num,omitempty" `

    /*
        电子凭证码     */
    Code  *string `json:"code,omitempty" `

    /*
        订单ID     */
    OrderId  *int64 `json:"order_id,omitempty" `

    /*
        码状态     */
    Status  *int64 `json:"status,omitempty" `

    /*
        二维码的图片地址     */
    QrcodeUrl  *string `json:"qrcode_url,omitempty" `

}

func (s *TaobaoVmarketEticketCodesGetEticketCode) SetNum(v int64) *TaobaoVmarketEticketCodesGetEticketCode {
    s.Num = &v
    return s
}
func (s *TaobaoVmarketEticketCodesGetEticketCode) SetCode(v string) *TaobaoVmarketEticketCodesGetEticketCode {
    s.Code = &v
    return s
}
func (s *TaobaoVmarketEticketCodesGetEticketCode) SetOrderId(v int64) *TaobaoVmarketEticketCodesGetEticketCode {
    s.OrderId = &v
    return s
}
func (s *TaobaoVmarketEticketCodesGetEticketCode) SetStatus(v int64) *TaobaoVmarketEticketCodesGetEticketCode {
    s.Status = &v
    return s
}
func (s *TaobaoVmarketEticketCodesGetEticketCode) SetQrcodeUrl(v string) *TaobaoVmarketEticketCodesGetEticketCode {
    s.QrcodeUrl = &v
    return s
}
