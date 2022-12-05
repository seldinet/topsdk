package domain


type TaobaoWlbWmsConsignOrderNotifyConsignorderitem struct {
    /*
        ERP订单明细行号ID     */
    OrderItemId  *string `json:"order_item_id,omitempty" `

    /*
        商品ID     */
    ItemId  *string `json:"item_id,omitempty" `

    /*
        商品编码     */
    ItemCode  *string `json:"item_code,omitempty" `

    /*
        商品数量     */
    ItemQuantity  *int64 `json:"item_quantity,omitempty" `

}

func (s *TaobaoWlbWmsConsignOrderNotifyConsignorderitem) SetOrderItemId(v string) *TaobaoWlbWmsConsignOrderNotifyConsignorderitem {
    s.OrderItemId = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyConsignorderitem) SetItemId(v string) *TaobaoWlbWmsConsignOrderNotifyConsignorderitem {
    s.ItemId = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyConsignorderitem) SetItemCode(v string) *TaobaoWlbWmsConsignOrderNotifyConsignorderitem {
    s.ItemCode = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyConsignorderitem) SetItemQuantity(v int64) *TaobaoWlbWmsConsignOrderNotifyConsignorderitem {
    s.ItemQuantity = &v
    return s
}
