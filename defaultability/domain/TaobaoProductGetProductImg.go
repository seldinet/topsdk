package domain


type TaobaoProductGetProductImg struct {
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

func (s *TaobaoProductGetProductImg) SetId(v int64) *TaobaoProductGetProductImg {
    s.Id = &v
    return s
}
func (s *TaobaoProductGetProductImg) SetProductId(v int64) *TaobaoProductGetProductImg {
    s.ProductId = &v
    return s
}
func (s *TaobaoProductGetProductImg) SetUrl(v string) *TaobaoProductGetProductImg {
    s.Url = &v
    return s
}
func (s *TaobaoProductGetProductImg) SetPosition(v int64) *TaobaoProductGetProductImg {
    s.Position = &v
    return s
}
