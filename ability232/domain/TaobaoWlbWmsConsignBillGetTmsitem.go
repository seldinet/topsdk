package domain


type TaobaoWlbWmsConsignBillGetTmsitem struct {
    /*
        此运单里该商品的数量     */
    ItemQty  *int64 `json:"item_qty,omitempty" `

    /*
        商家编码     */
    ItemCode  *string `json:"item_code,omitempty" `

    /*
        ERP商品ID     */
    ItemId  *string `json:"item_id,omitempty" `

}

func (s *TaobaoWlbWmsConsignBillGetTmsitem) SetItemQty(v int64) *TaobaoWlbWmsConsignBillGetTmsitem {
    s.ItemQty = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetTmsitem) SetItemCode(v string) *TaobaoWlbWmsConsignBillGetTmsitem {
    s.ItemCode = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetTmsitem) SetItemId(v string) *TaobaoWlbWmsConsignBillGetTmsitem {
    s.ItemId = &v
    return s
}
