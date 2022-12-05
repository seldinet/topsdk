package domain


type TaobaoLogisticsPartnersGetPartnerDetail struct {
    /*
        物流公司支付宝账号     */
    AccountNo  *string `json:"account_no,omitempty" `

    /*
        物流公司代码     */
    CompanyCode  *string `json:"company_code,omitempty" `

    /*
        物流公司id     */
    CompanyId  *int64 `json:"company_id,omitempty" `

    /*
        物流公司全名     */
    FullName  *string `json:"full_name,omitempty" `

    /*
        运单号验证正则表达式     */
    RegMailNo  *string `json:"reg_mail_no,omitempty" `

    /*
        旺旺id     */
    WangwangId  *string `json:"wangwang_id,omitempty" `

    /*
        物流公司简称     */
    CompanyName  *string `json:"company_name,omitempty" `

}

func (s *TaobaoLogisticsPartnersGetPartnerDetail) SetAccountNo(v string) *TaobaoLogisticsPartnersGetPartnerDetail {
    s.AccountNo = &v
    return s
}
func (s *TaobaoLogisticsPartnersGetPartnerDetail) SetCompanyCode(v string) *TaobaoLogisticsPartnersGetPartnerDetail {
    s.CompanyCode = &v
    return s
}
func (s *TaobaoLogisticsPartnersGetPartnerDetail) SetCompanyId(v int64) *TaobaoLogisticsPartnersGetPartnerDetail {
    s.CompanyId = &v
    return s
}
func (s *TaobaoLogisticsPartnersGetPartnerDetail) SetFullName(v string) *TaobaoLogisticsPartnersGetPartnerDetail {
    s.FullName = &v
    return s
}
func (s *TaobaoLogisticsPartnersGetPartnerDetail) SetRegMailNo(v string) *TaobaoLogisticsPartnersGetPartnerDetail {
    s.RegMailNo = &v
    return s
}
func (s *TaobaoLogisticsPartnersGetPartnerDetail) SetWangwangId(v string) *TaobaoLogisticsPartnersGetPartnerDetail {
    s.WangwangId = &v
    return s
}
func (s *TaobaoLogisticsPartnersGetPartnerDetail) SetCompanyName(v string) *TaobaoLogisticsPartnersGetPartnerDetail {
    s.CompanyName = &v
    return s
}
