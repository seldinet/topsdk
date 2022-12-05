package domain


type TaobaoWlbWmsStockInOrderNotifyOrderitemlistwlbwmsstockinordernotifywl struct {
    /*
        系统自动生成     */
    OrderItem  *TaobaoWlbWmsStockInOrderNotifyOrderitemwlbwmsstockinordernotifywl `json:"order_item,omitempty" `

}

func (s *TaobaoWlbWmsStockInOrderNotifyOrderitemlistwlbwmsstockinordernotifywl) SetOrderItem(v TaobaoWlbWmsStockInOrderNotifyOrderitemwlbwmsstockinordernotifywl) *TaobaoWlbWmsStockInOrderNotifyOrderitemlistwlbwmsstockinordernotifywl {
    s.OrderItem = &v
    return s
}
