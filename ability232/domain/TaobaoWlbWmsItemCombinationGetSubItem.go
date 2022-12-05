package domain


type TaobaoWlbWmsItemCombinationGetSubItem struct {
    /*
        子货品ID     */
    ScItemId  *int64 `json:"sc_item_id,omitempty" `

    /*
        子货品数量     */
    Count  *int64 `json:"count,omitempty" `

}

func (s *TaobaoWlbWmsItemCombinationGetSubItem) SetScItemId(v int64) *TaobaoWlbWmsItemCombinationGetSubItem {
    s.ScItemId = &v
    return s
}
func (s *TaobaoWlbWmsItemCombinationGetSubItem) SetCount(v int64) *TaobaoWlbWmsItemCombinationGetSubItem {
    s.Count = &v
    return s
}
