package domain


type TaobaoFenxiaoRefundGetRefundLogistics struct {
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

func (s *TaobaoFenxiaoRefundGetRefundLogistics) SetCompanyCode(v string) *TaobaoFenxiaoRefundGetRefundLogistics {
    s.CompanyCode = &v
    return s
}
func (s *TaobaoFenxiaoRefundGetRefundLogistics) SetCompanyName(v string) *TaobaoFenxiaoRefundGetRefundLogistics {
    s.CompanyName = &v
    return s
}
func (s *TaobaoFenxiaoRefundGetRefundLogistics) SetMailNo(v string) *TaobaoFenxiaoRefundGetRefundLogistics {
    s.MailNo = &v
    return s
}
