package domain


type TaobaoItemSkuDeleteSku struct {
    /*
        sku所属商品id(注意：iid近期即将废弃，请用num_iid参数)     */
    Iid  *string `json:"iid,omitempty" `

    /*
        sku所属商品数字id     */
    NumIid  *int64 `json:"num_iid,omitempty" `

    /*
        sku最后修改日期 时间格式：yyyy-MM-dd HH:mm:ss     */
    Modified  *string `json:"modified,omitempty" `

}

func (s *TaobaoItemSkuDeleteSku) SetIid(v string) *TaobaoItemSkuDeleteSku {
    s.Iid = &v
    return s
}
func (s *TaobaoItemSkuDeleteSku) SetNumIid(v int64) *TaobaoItemSkuDeleteSku {
    s.NumIid = &v
    return s
}
func (s *TaobaoItemSkuDeleteSku) SetModified(v string) *TaobaoItemSkuDeleteSku {
    s.Modified = &v
    return s
}
