package domain


type TaobaoWlbWmsInventoryQueryWmsInventoryQueryItemlist struct {
    /*
        商品详情     */
    Item  *TaobaoWlbWmsInventoryQueryWmsInventoryQueryItem `json:"item,omitempty" `

}

func (s *TaobaoWlbWmsInventoryQueryWmsInventoryQueryItemlist) SetItem(v TaobaoWlbWmsInventoryQueryWmsInventoryQueryItem) *TaobaoWlbWmsInventoryQueryWmsInventoryQueryItemlist {
    s.Item = &v
    return s
}
