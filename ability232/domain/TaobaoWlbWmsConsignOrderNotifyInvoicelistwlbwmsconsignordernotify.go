package domain


type TaobaoWlbWmsConsignOrderNotifyInvoicelistwlbwmsconsignordernotify struct {
    /*
        发票信息     */
    InvoiceInfo  *TaobaoWlbWmsConsignOrderNotifyInvoicewlbwmsconsignordernotify `json:"invoice_info,omitempty" `

}

func (s *TaobaoWlbWmsConsignOrderNotifyInvoicelistwlbwmsconsignordernotify) SetInvoiceInfo(v TaobaoWlbWmsConsignOrderNotifyInvoicewlbwmsconsignordernotify) *TaobaoWlbWmsConsignOrderNotifyInvoicelistwlbwmsconsignordernotify {
    s.InvoiceInfo = &v
    return s
}
