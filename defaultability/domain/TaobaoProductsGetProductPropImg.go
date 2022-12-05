package domain


type TaobaoProductsGetProductPropImg struct {
    /*
        产品属性图片ID     */
    Id  *int64 `json:"id,omitempty" `

    /*
        图片所属产品的ID     */
    ProductId  *int64 `json:"product_id,omitempty" `

    /*
        属性串(pid:vid),目前只有颜色属性.如:颜色:红色表示为　1627207:28326     */
    Props  *string `json:"props,omitempty" `

    /*
        图片地址.(绝对地址,格式:http://host/image_path)     */
    Url  *string `json:"url,omitempty" `

    /*
        图片序号。产品里的图片展示顺序，数据越小越靠前。要求是正整数。     */
    Position  *int64 `json:"position,omitempty" `

}

func (s *TaobaoProductsGetProductPropImg) SetId(v int64) *TaobaoProductsGetProductPropImg {
    s.Id = &v
    return s
}
func (s *TaobaoProductsGetProductPropImg) SetProductId(v int64) *TaobaoProductsGetProductPropImg {
    s.ProductId = &v
    return s
}
func (s *TaobaoProductsGetProductPropImg) SetProps(v string) *TaobaoProductsGetProductPropImg {
    s.Props = &v
    return s
}
func (s *TaobaoProductsGetProductPropImg) SetUrl(v string) *TaobaoProductsGetProductPropImg {
    s.Url = &v
    return s
}
func (s *TaobaoProductsGetProductPropImg) SetPosition(v int64) *TaobaoProductsGetProductPropImg {
    s.Position = &v
    return s
}
