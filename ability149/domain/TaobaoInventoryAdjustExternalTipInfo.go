package domain


type TaobaoInventoryAdjustExternalTipInfo struct {
    /*
        商品id     */
    ScItemId  *string `json:"sc_item_id,omitempty" `

    /*
        返回信息     */
    Info  *string `json:"info,omitempty" `

}

func (s *TaobaoInventoryAdjustExternalTipInfo) SetScItemId(v string) *TaobaoInventoryAdjustExternalTipInfo {
    s.ScItemId = &v
    return s
}
func (s *TaobaoInventoryAdjustExternalTipInfo) SetInfo(v string) *TaobaoInventoryAdjustExternalTipInfo {
    s.Info = &v
    return s
}
