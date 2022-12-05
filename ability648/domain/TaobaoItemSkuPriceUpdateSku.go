package domain


type TaobaoItemSkuPriceUpdateSku struct {
    /*
        sku所属商品id(注意：iid近期即将废弃，请用num_iid参数)     */
    Iid  *string `json:"iid,omitempty" `

    /*
        sku所属商品数字id     */
    NumIid  *int64 `json:"num_iid,omitempty" `

    /*
        sku最后修改日期 时间格式：yyyy-MM-dd HH:mm:ss     */
    Modified  *string `json:"modified,omitempty" `

    /*
        sku的id     */
    SkuId  *int64 `json:"sku_id,omitempty" `

}

func (s *TaobaoItemSkuPriceUpdateSku) SetIid(v string) *TaobaoItemSkuPriceUpdateSku {
    s.Iid = &v
    return s
}
func (s *TaobaoItemSkuPriceUpdateSku) SetNumIid(v int64) *TaobaoItemSkuPriceUpdateSku {
    s.NumIid = &v
    return s
}
func (s *TaobaoItemSkuPriceUpdateSku) SetModified(v string) *TaobaoItemSkuPriceUpdateSku {
    s.Modified = &v
    return s
}
func (s *TaobaoItemSkuPriceUpdateSku) SetSkuId(v int64) *TaobaoItemSkuPriceUpdateSku {
    s.SkuId = &v
    return s
}
