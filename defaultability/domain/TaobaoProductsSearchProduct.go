package domain


type TaobaoProductsSearchProduct struct {
    /*
        产品ID     */
    ProductId  *int64 `json:"product_id,omitempty" `

    /*
        产品名称     */
    Name  *string `json:"name,omitempty" `

    /*
        产品的主图片地址.(绝对地址,格式:http://host/image_path)     */
    PicUrl  *string `json:"pic_url,omitempty" `

    /*
        商品类目ID.必须是叶子类目ID     */
    Cid  *int64 `json:"cid,omitempty" `

    /*
        产品的关键属性列表.格式：pid:vid;pid:vid     */
    Props  *string `json:"props,omitempty" `

    /*
        产品的市场价.单位为元.精确到2位小数;如:200.07     */
    Price  *string `json:"price,omitempty" `

    /*
        淘宝标准产品编码     */
    Tsc  *string `json:"tsc,omitempty" `

    /*
        当前状态(0 商家确认 1 屏蔽 3 小二确认 2 未确认 -1 删除)     */
    Status  *int64 `json:"status,omitempty" `

}

func (s *TaobaoProductsSearchProduct) SetProductId(v int64) *TaobaoProductsSearchProduct {
    s.ProductId = &v
    return s
}
func (s *TaobaoProductsSearchProduct) SetName(v string) *TaobaoProductsSearchProduct {
    s.Name = &v
    return s
}
func (s *TaobaoProductsSearchProduct) SetPicUrl(v string) *TaobaoProductsSearchProduct {
    s.PicUrl = &v
    return s
}
func (s *TaobaoProductsSearchProduct) SetCid(v int64) *TaobaoProductsSearchProduct {
    s.Cid = &v
    return s
}
func (s *TaobaoProductsSearchProduct) SetProps(v string) *TaobaoProductsSearchProduct {
    s.Props = &v
    return s
}
func (s *TaobaoProductsSearchProduct) SetPrice(v string) *TaobaoProductsSearchProduct {
    s.Price = &v
    return s
}
func (s *TaobaoProductsSearchProduct) SetTsc(v string) *TaobaoProductsSearchProduct {
    s.Tsc = &v
    return s
}
func (s *TaobaoProductsSearchProduct) SetStatus(v int64) *TaobaoProductsSearchProduct {
    s.Status = &v
    return s
}
