package domain


type TaobaoWlbWmsConsignOrderNotifyOrderitemlistwlbwmsconsignordernotify struct {
    /*
        订单商品信息     */
    OrderItem  *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify `json:"order_item,omitempty" `

}

func (s *TaobaoWlbWmsConsignOrderNotifyOrderitemlistwlbwmsconsignordernotify) SetOrderItem(v TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify) *TaobaoWlbWmsConsignOrderNotifyOrderitemlistwlbwmsconsignordernotify {
    s.OrderItem = &v
    return s
}
