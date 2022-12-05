package domain


type TaobaoTradeFullinfoGetAgreeRefundCheck struct {
    /*
        提示文案     */
    DeliveryTips  *string `json:"delivery_tips,omitempty" `

    /*
        流程状态     */
    DeliveryProcess  *string `json:"delivery_process,omitempty" `

    /*
        订单id     */
    DetailOrderId  *string `json:"detail_order_id,omitempty" `

}

func (s *TaobaoTradeFullinfoGetAgreeRefundCheck) SetDeliveryTips(v string) *TaobaoTradeFullinfoGetAgreeRefundCheck {
    s.DeliveryTips = &v
    return s
}
func (s *TaobaoTradeFullinfoGetAgreeRefundCheck) SetDeliveryProcess(v string) *TaobaoTradeFullinfoGetAgreeRefundCheck {
    s.DeliveryProcess = &v
    return s
}
func (s *TaobaoTradeFullinfoGetAgreeRefundCheck) SetDetailOrderId(v string) *TaobaoTradeFullinfoGetAgreeRefundCheck {
    s.DetailOrderId = &v
    return s
}
