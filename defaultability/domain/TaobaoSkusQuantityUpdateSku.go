package domain


type TaobaoSkusQuantityUpdateSku struct {
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

func (s *TaobaoSkusQuantityUpdateSku) SetSkuId(v int64) *TaobaoSkusQuantityUpdateSku {
    s.SkuId = &v
    return s
}
func (s *TaobaoSkusQuantityUpdateSku) SetModified(v string) *TaobaoSkusQuantityUpdateSku {
    s.Modified = &v
    return s
}
func (s *TaobaoSkusQuantityUpdateSku) SetQuantity(v int64) *TaobaoSkusQuantityUpdateSku {
    s.Quantity = &v
    return s
}
