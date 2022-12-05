package domain


type TaobaoItemSellerGetLocalityLife struct {
    /*
        表示是否使用邮寄 0: 代表不使用邮寄； 1：代表使用邮寄；如果不设置这个值，代表不使用邮寄     */
    ChooseLogis  *string `json:"choose_logis,omitempty" `

    /*
        电子凭证业务属性     */
    Eticket  *string `json:"eticket,omitempty" `

    /*
        电子交易凭证有效期，有三种格式：如果有效期为起止日期类型，此值为2012-08-06,2012-08-16 如果有效期为【购买成功日 至】类型则格式为2012-08-16如果有效期为天数类型则格式为15     */
    Expirydate  *string `json:"expirydate,omitempty" `

    /*
        格式为 码商id:nick     */
    Merchant  *string `json:"merchant,omitempty" `

    /*
        网点ID,在参数empty_fields里设置locality_life.network_id可删除网点ID     */
    NetworkId  *string `json:"network_id,omitempty" `

    /*
        电子凭证售中自动退款比例     */
    OnsaleAutoRefundRatio  *int64 `json:"onsale_auto_refund_ratio,omitempty" `

    /*
        退款比例，百分比%前的数字，1-100的正整数值；在参数empty_fields里设置locality_life.refund_ratio可删除退款比例     */
    RefundRatio  *int64 `json:"refund_ratio,omitempty" `

    /*
        退款码费承担方。发布电子凭证宝贝的时候会增加“退款码费承担方”配置项，可选填：(1)s（卖家承担） (2)b(买家承担)     */
    Refundmafee  *string `json:"refundmafee,omitempty" `

    /*
        核销打款:1代表核销打款,0代表非核销打款;在参数empty_fields里设置locality_life.verification可删除核销打款     */
    Verification  *string `json:"verification,omitempty" `

}

func (s *TaobaoItemSellerGetLocalityLife) SetChooseLogis(v string) *TaobaoItemSellerGetLocalityLife {
    s.ChooseLogis = &v
    return s
}
func (s *TaobaoItemSellerGetLocalityLife) SetEticket(v string) *TaobaoItemSellerGetLocalityLife {
    s.Eticket = &v
    return s
}
func (s *TaobaoItemSellerGetLocalityLife) SetExpirydate(v string) *TaobaoItemSellerGetLocalityLife {
    s.Expirydate = &v
    return s
}
func (s *TaobaoItemSellerGetLocalityLife) SetMerchant(v string) *TaobaoItemSellerGetLocalityLife {
    s.Merchant = &v
    return s
}
func (s *TaobaoItemSellerGetLocalityLife) SetNetworkId(v string) *TaobaoItemSellerGetLocalityLife {
    s.NetworkId = &v
    return s
}
func (s *TaobaoItemSellerGetLocalityLife) SetOnsaleAutoRefundRatio(v int64) *TaobaoItemSellerGetLocalityLife {
    s.OnsaleAutoRefundRatio = &v
    return s
}
func (s *TaobaoItemSellerGetLocalityLife) SetRefundRatio(v int64) *TaobaoItemSellerGetLocalityLife {
    s.RefundRatio = &v
    return s
}
func (s *TaobaoItemSellerGetLocalityLife) SetRefundmafee(v string) *TaobaoItemSellerGetLocalityLife {
    s.Refundmafee = &v
    return s
}
func (s *TaobaoItemSellerGetLocalityLife) SetVerification(v string) *TaobaoItemSellerGetLocalityLife {
    s.Verification = &v
    return s
}
