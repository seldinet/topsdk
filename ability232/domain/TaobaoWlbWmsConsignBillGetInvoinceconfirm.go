package domain


type TaobaoWlbWmsConsignBillGetInvoinceconfirm struct {
    /*
        发票代码     */
    InvoiceCode  *string `json:"invoice_code,omitempty" `

    /*
        发票号     */
    InvoiceNumber  *string `json:"invoice_number,omitempty" `

    /*
        Erp发票ID     */
    BillId  *string `json:"bill_id,omitempty" `

}

func (s *TaobaoWlbWmsConsignBillGetInvoinceconfirm) SetInvoiceCode(v string) *TaobaoWlbWmsConsignBillGetInvoinceconfirm {
    s.InvoiceCode = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetInvoinceconfirm) SetInvoiceNumber(v string) *TaobaoWlbWmsConsignBillGetInvoinceconfirm {
    s.InvoiceNumber = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetInvoinceconfirm) SetBillId(v string) *TaobaoWlbWmsConsignBillGetInvoinceconfirm {
    s.BillId = &v
    return s
}
