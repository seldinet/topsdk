package domain


type TaobaoItemSkuAddSku struct {
    /*
        sku所属商品id(注意：iid近期即将废弃，请用num_iid参数)     */
    Iid  *string `json:"iid,omitempty" `

    /*
        sku所属商品数字id     */
    NumIid  *int64 `json:"num_iid,omitempty" `

    /*
        sku创建日期 时间格式：yyyy-MM-dd HH:mm:ss     */
    Created  *string `json:"created,omitempty" `

    /*
        sku的id     */
    SkuId  *int64 `json:"sku_id,omitempty" `

}

func (s *TaobaoItemSkuAddSku) SetIid(v string) *TaobaoItemSkuAddSku {
    s.Iid = &v
    return s
}
func (s *TaobaoItemSkuAddSku) SetNumIid(v int64) *TaobaoItemSkuAddSku {
    s.NumIid = &v
    return s
}
func (s *TaobaoItemSkuAddSku) SetCreated(v string) *TaobaoItemSkuAddSku {
    s.Created = &v
    return s
}
func (s *TaobaoItemSkuAddSku) SetSkuId(v int64) *TaobaoItemSkuAddSku {
    s.SkuId = &v
    return s
}
