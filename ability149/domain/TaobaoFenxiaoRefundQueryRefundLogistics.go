package domain


type TaobaoFenxiaoRefundQueryRefundLogistics struct {
    /*
        退货物流公司编码，如顺丰、韵达等     */
    CompanyCode  *string `json:"company_code,omitempty" `

    /*
        退货物流公司名称，如顺丰     */
    CompanyName  *string `json:"company_name,omitempty" `

    /*
        退货物流运单号     */
    MailNo  *string `json:"mail_no,omitempty" `

}

func (s *TaobaoFenxiaoRefundQueryRefundLogistics) SetCompanyCode(v string) *TaobaoFenxiaoRefundQueryRefundLogistics {
    s.CompanyCode = &v
    return s
}
func (s *TaobaoFenxiaoRefundQueryRefundLogistics) SetCompanyName(v string) *TaobaoFenxiaoRefundQueryRefundLogistics {
    s.CompanyName = &v
    return s
}
func (s *TaobaoFenxiaoRefundQueryRefundLogistics) SetMailNo(v string) *TaobaoFenxiaoRefundQueryRefundLogistics {
    s.MailNo = &v
    return s
}
