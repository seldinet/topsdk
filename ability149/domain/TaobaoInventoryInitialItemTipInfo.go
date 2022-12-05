package domain


type TaobaoInventoryInitialItemTipInfo struct {
    /*
        商品ID     */
    ScItemId  *string `json:"sc_item_id,omitempty" `

    /*
        提示信息     */
    Info  *string `json:"info,omitempty" `

}

func (s *TaobaoInventoryInitialItemTipInfo) SetScItemId(v string) *TaobaoInventoryInitialItemTipInfo {
    s.ScItemId = &v
    return s
}
func (s *TaobaoInventoryInitialItemTipInfo) SetInfo(v string) *TaobaoInventoryInitialItemTipInfo {
    s.Info = &v
    return s
}
