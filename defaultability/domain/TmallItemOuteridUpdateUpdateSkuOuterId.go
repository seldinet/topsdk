package domain


type TmallItemOuteridUpdateUpdateSkuOuterId struct {
    /*
        被更新的Sku的商家外部id     */
    OuterId  *string `json:"outer_id,omitempty" `

    /*
        Sku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充；如果填写将以属性对形式查找被更新SKU     */
    Properties  *string `json:"properties,omitempty" `

    /*
        SkuID，如果填写，将以SKUID查找被更新的SKU     */
    SkuId  *int64 `json:"sku_id,omitempty" `

}

func (s *TmallItemOuteridUpdateUpdateSkuOuterId) SetOuterId(v string) *TmallItemOuteridUpdateUpdateSkuOuterId {
    s.OuterId = &v
    return s
}
func (s *TmallItemOuteridUpdateUpdateSkuOuterId) SetProperties(v string) *TmallItemOuteridUpdateUpdateSkuOuterId {
    s.Properties = &v
    return s
}
func (s *TmallItemOuteridUpdateUpdateSkuOuterId) SetSkuId(v int64) *TmallItemOuteridUpdateUpdateSkuOuterId {
    s.SkuId = &v
    return s
}
