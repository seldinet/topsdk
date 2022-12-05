package domain


type TaobaoInventoryInitialTipInfo struct {
    /*
        商品id     */
    ScItemId  *string `json:"sc_item_id,omitempty" `

    /*
        返回信息     */
    Info  *string `json:"info,omitempty" `

}

func (s *TaobaoInventoryInitialTipInfo) SetScItemId(v string) *TaobaoInventoryInitialTipInfo {
    s.ScItemId = &v
    return s
}
func (s *TaobaoInventoryInitialTipInfo) SetInfo(v string) *TaobaoInventoryInitialTipInfo {
    s.Info = &v
    return s
}
