package domain


type TaobaoWlbWmsItemCombinationGetSubItemList struct {
    /*
        子货品     */
    SubItem  *TaobaoWlbWmsItemCombinationGetSubItem `json:"sub_item,omitempty" `

}

func (s *TaobaoWlbWmsItemCombinationGetSubItemList) SetSubItem(v TaobaoWlbWmsItemCombinationGetSubItem) *TaobaoWlbWmsItemCombinationGetSubItemList {
    s.SubItem = &v
    return s
}
