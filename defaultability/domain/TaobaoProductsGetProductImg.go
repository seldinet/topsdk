package domain


type TaobaoProductsGetProductImg struct {
    /*
        产品图片ID     */
    Id  *int64 `json:"id,omitempty" `

    /*
        图片所属产品的ID     */
    ProductId  *int64 `json:"product_id,omitempty" `

    /*
        图片地址.(绝对地址,格式:http://host/image_path)     */
    Url  *string `json:"url,omitempty" `

    /*
        图片序号。产品里的图片展示顺序，数据越小越靠前。要求是正整数。     */
    Position  *int64 `json:"position,omitempty" `

}

func (s *TaobaoProductsGetProductImg) SetId(v int64) *TaobaoProductsGetProductImg {
    s.Id = &v
    return s
}
func (s *TaobaoProductsGetProductImg) SetProductId(v int64) *TaobaoProductsGetProductImg {
    s.ProductId = &v
    return s
}
func (s *TaobaoProductsGetProductImg) SetUrl(v string) *TaobaoProductsGetProductImg {
    s.Url = &v
    return s
}
func (s *TaobaoProductsGetProductImg) SetPosition(v int64) *TaobaoProductsGetProductImg {
    s.Position = &v
    return s
}
