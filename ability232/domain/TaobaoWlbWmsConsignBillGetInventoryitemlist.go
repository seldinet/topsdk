package domain


type TaobaoWlbWmsConsignBillGetInventoryitemlist struct {
    /*
        商品属性列表     */
    InventoryItem  *TaobaoWlbWmsConsignBillGetInventoryitem `json:"inventory_item,omitempty" `

}

func (s *TaobaoWlbWmsConsignBillGetInventoryitemlist) SetInventoryItem(v TaobaoWlbWmsConsignBillGetInventoryitem) *TaobaoWlbWmsConsignBillGetInventoryitemlist {
    s.InventoryItem = &v
    return s
}
