package domain


type TmallItemQuantityUpdateUpdateSkuQuantity struct {
    /*
        Sku的商家外部id，用于指定被修改库存的SKU     */
    OuterId  *string `json:"outer_id,omitempty" `

    /*
        Sku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充，用于指定被修改库存的SKU     */
    Properties  *string `json:"properties,omitempty" `

    /*
        SKU属于这个sku的商品的库存；增量编辑方式支持正数、负数     */
    Quantity  *int64 `json:"quantity,omitempty" `

    /*
        SkuID，用于指定被修改库存的     */
    SkuId  *int64 `json:"sku_id,omitempty" `

}

func (s *TmallItemQuantityUpdateUpdateSkuQuantity) SetOuterId(v string) *TmallItemQuantityUpdateUpdateSkuQuantity {
    s.OuterId = &v
    return s
}
func (s *TmallItemQuantityUpdateUpdateSkuQuantity) SetProperties(v string) *TmallItemQuantityUpdateUpdateSkuQuantity {
    s.Properties = &v
    return s
}
func (s *TmallItemQuantityUpdateUpdateSkuQuantity) SetQuantity(v int64) *TmallItemQuantityUpdateUpdateSkuQuantity {
    s.Quantity = &v
    return s
}
func (s *TmallItemQuantityUpdateUpdateSkuQuantity) SetSkuId(v int64) *TmallItemQuantityUpdateUpdateSkuQuantity {
    s.SkuId = &v
    return s
}
