package domain


type TaobaoItemSkuUpdateSku struct {
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

func (s *TaobaoItemSkuUpdateSku) SetIid(v string) *TaobaoItemSkuUpdateSku {
    s.Iid = &v
    return s
}
func (s *TaobaoItemSkuUpdateSku) SetNumIid(v int64) *TaobaoItemSkuUpdateSku {
    s.NumIid = &v
    return s
}
func (s *TaobaoItemSkuUpdateSku) SetModified(v string) *TaobaoItemSkuUpdateSku {
    s.Modified = &v
    return s
}
func (s *TaobaoItemSkuUpdateSku) SetSkuId(v int64) *TaobaoItemSkuUpdateSku {
    s.SkuId = &v
    return s
}
