package domain


type TaobaoWlbWmsConsignOrderNotifyInvoicewlbwmsconsignordernotify struct {
    /*
        发票金额     */
    BillAccount  *string `json:"bill_account,omitempty" `

    /*
        发票抬头     */
    BillTitle  *string `json:"bill_title,omitempty" `

    /*
        发票内容     */
    BillContent  *string `json:"bill_content,omitempty" `

    /*
        Erp发票ID     */
    BillId  *int64 `json:"bill_id,omitempty" `

    /*
        发票类型(VINVOICE - 增值税普通发票， EVINVOICE - 电子增票，电子发票仓库不打印)     */
    BillType  *string `json:"bill_type,omitempty" `

    /*
        发票明细列表     */
    DetailList  *[]TaobaoWlbWmsConsignOrderNotifyDetaillistwlbwmsconsignordernotify `json:"detail_list,omitempty" `

}

func (s *TaobaoWlbWmsConsignOrderNotifyInvoicewlbwmsconsignordernotify) SetBillAccount(v string) *TaobaoWlbWmsConsignOrderNotifyInvoicewlbwmsconsignordernotify {
    s.BillAccount = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyInvoicewlbwmsconsignordernotify) SetBillTitle(v string) *TaobaoWlbWmsConsignOrderNotifyInvoicewlbwmsconsignordernotify {
    s.BillTitle = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyInvoicewlbwmsconsignordernotify) SetBillContent(v string) *TaobaoWlbWmsConsignOrderNotifyInvoicewlbwmsconsignordernotify {
    s.BillContent = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyInvoicewlbwmsconsignordernotify) SetBillId(v int64) *TaobaoWlbWmsConsignOrderNotifyInvoicewlbwmsconsignordernotify {
    s.BillId = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyInvoicewlbwmsconsignordernotify) SetBillType(v string) *TaobaoWlbWmsConsignOrderNotifyInvoicewlbwmsconsignordernotify {
    s.BillType = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyInvoicewlbwmsconsignordernotify) SetDetailList(v []TaobaoWlbWmsConsignOrderNotifyDetaillistwlbwmsconsignordernotify) *TaobaoWlbWmsConsignOrderNotifyInvoicewlbwmsconsignordernotify {
    s.DetailList = &v
    return s
}
