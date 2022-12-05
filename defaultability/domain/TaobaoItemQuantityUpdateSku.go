package domain


type TaobaoItemQuantityUpdateSku struct {
    /*
        sku的id     */
    SkuId  *int64 `json:"sku_id,omitempty" `

    /*
        sku最后修改日期 时间格式：yyyy-MM-dd HH:mm:ss     */
    Modified  *string `json:"modified,omitempty" `

    /*
        属于这个sku的商品的数量，     */
    Quantity  *int64 `json:"quantity,omitempty" `

}

func (s *TaobaoItemQuantityUpdateSku) SetSkuId(v int64) *TaobaoItemQuantityUpdateSku {
    s.SkuId = &v
    return s
}
func (s *TaobaoItemQuantityUpdateSku) SetModified(v string) *TaobaoItemQuantityUpdateSku {
    s.Modified = &v
    return s
}
func (s *TaobaoItemQuantityUpdateSku) SetQuantity(v int64) *TaobaoItemQuantityUpdateSku {
    s.Quantity = &v
    return s
}
