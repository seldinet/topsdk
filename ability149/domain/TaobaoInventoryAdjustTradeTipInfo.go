package domain


type TaobaoInventoryAdjustTradeTipInfo struct {
    /*
        商品id     */
    ScItemId  *string `json:"sc_item_id,omitempty" `

    /*
        返回信息     */
    Info  *string `json:"info,omitempty" `

}

func (s *TaobaoInventoryAdjustTradeTipInfo) SetScItemId(v string) *TaobaoInventoryAdjustTradeTipInfo {
    s.ScItemId = &v
    return s
}
func (s *TaobaoInventoryAdjustTradeTipInfo) SetInfo(v string) *TaobaoInventoryAdjustTradeTipInfo {
    s.Info = &v
    return s
}
