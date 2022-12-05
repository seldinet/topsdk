package domain


type TaobaoWlbWmsConsignBillGetOrderitemlist struct {
    /*
        订单商品信息     */
    OrderItem  *TaobaoWlbWmsConsignBillGetOrderitem `json:"order_item,omitempty" `

}

func (s *TaobaoWlbWmsConsignBillGetOrderitemlist) SetOrderItem(v TaobaoWlbWmsConsignBillGetOrderitem) *TaobaoWlbWmsConsignBillGetOrderitemlist {
    s.OrderItem = &v
    return s
}
