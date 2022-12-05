package domain


type TaobaoWlbWmsConsignBillGetOrderitem struct {
    /*
        商品属性列表     */
    InventoryItemList  *[]TaobaoWlbWmsConsignBillGetInventoryitemlist `json:"inventory_item_list,omitempty" `

    /*
        商家编码     */
    ItemCode  *string `json:"item_code,omitempty" `

    /*
        商品ID     */
    ItemId  *string `json:"item_id,omitempty" `

    /*
        订单明细行编码     */
    OrderItemId  *string `json:"order_item_id,omitempty" `

    /*
        交易单号     */
    OrderSourceCode  *string `json:"order_source_code,omitempty" `

}

func (s *TaobaoWlbWmsConsignBillGetOrderitem) SetInventoryItemList(v []TaobaoWlbWmsConsignBillGetInventoryitemlist) *TaobaoWlbWmsConsignBillGetOrderitem {
    s.InventoryItemList = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetOrderitem) SetItemCode(v string) *TaobaoWlbWmsConsignBillGetOrderitem {
    s.ItemCode = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetOrderitem) SetItemId(v string) *TaobaoWlbWmsConsignBillGetOrderitem {
    s.ItemId = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetOrderitem) SetOrderItemId(v string) *TaobaoWlbWmsConsignBillGetOrderitem {
    s.OrderItemId = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetOrderitem) SetOrderSourceCode(v string) *TaobaoWlbWmsConsignBillGetOrderitem {
    s.OrderSourceCode = &v
    return s
}
