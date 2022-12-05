package domain


type TaobaoWlbWmsConsignOrderNotifyItemdetailwlbwmsconsignordernotify struct {
    /*
        金额     */
    Amount  *string `json:"amount,omitempty" `

    /*
        商品数量     */
    Quantity  *int64 `json:"quantity,omitempty" `

    /*
        商品价格     */
    Price  *string `json:"price,omitempty" `

    /*
        单位     */
    Unit  *string `json:"unit,omitempty" `

    /*
        商品名称     */
    ItemName  *string `json:"item_name,omitempty" `

}

func (s *TaobaoWlbWmsConsignOrderNotifyItemdetailwlbwmsconsignordernotify) SetAmount(v string) *TaobaoWlbWmsConsignOrderNotifyItemdetailwlbwmsconsignordernotify {
    s.Amount = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyItemdetailwlbwmsconsignordernotify) SetQuantity(v int64) *TaobaoWlbWmsConsignOrderNotifyItemdetailwlbwmsconsignordernotify {
    s.Quantity = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyItemdetailwlbwmsconsignordernotify) SetPrice(v string) *TaobaoWlbWmsConsignOrderNotifyItemdetailwlbwmsconsignordernotify {
    s.Price = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyItemdetailwlbwmsconsignordernotify) SetUnit(v string) *TaobaoWlbWmsConsignOrderNotifyItemdetailwlbwmsconsignordernotify {
    s.Unit = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyItemdetailwlbwmsconsignordernotify) SetItemName(v string) *TaobaoWlbWmsConsignOrderNotifyItemdetailwlbwmsconsignordernotify {
    s.ItemName = &v
    return s
}
