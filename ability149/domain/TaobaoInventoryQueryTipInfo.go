package domain


type TaobaoInventoryQueryTipInfo struct {
    /*
        商品id     */
    ScItemId  *string `json:"sc_item_id,omitempty" `

    /*
        返回信息     */
    Info  *string `json:"info,omitempty" `

}

func (s *TaobaoInventoryQueryTipInfo) SetScItemId(v string) *TaobaoInventoryQueryTipInfo {
    s.ScItemId = &v
    return s
}
func (s *TaobaoInventoryQueryTipInfo) SetInfo(v string) *TaobaoInventoryQueryTipInfo {
    s.Info = &v
    return s
}
