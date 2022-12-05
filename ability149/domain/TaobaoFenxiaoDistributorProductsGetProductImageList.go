package domain


type TaobaoFenxiaoDistributorProductsGetProductImageList struct {
    /*
        图片id     */
    ImageId  *int64 `json:"image_id,omitempty" `

    /*
        图片对应的url     */
    ImageUrl  *string `json:"image_url,omitempty" `

    /*
        图片顺序     */
    ImagePosition  *int64 `json:"image_position,omitempty" `

    /*
        当图片类型为属性图片时，表示图片对应的属性pv对。     */
    Properties  *string `json:"properties,omitempty" `

    /*
        图片类型（PRODUCT：产品图片  EXTPRODUCT：产品辅图  PROPERTIES：产品属性图片 ）     */
    Type  *string `json:"type,omitempty" `

}

func (s *TaobaoFenxiaoDistributorProductsGetProductImageList) SetImageId(v int64) *TaobaoFenxiaoDistributorProductsGetProductImageList {
    s.ImageId = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetProductImageList) SetImageUrl(v string) *TaobaoFenxiaoDistributorProductsGetProductImageList {
    s.ImageUrl = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetProductImageList) SetImagePosition(v int64) *TaobaoFenxiaoDistributorProductsGetProductImageList {
    s.ImagePosition = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetProductImageList) SetProperties(v string) *TaobaoFenxiaoDistributorProductsGetProductImageList {
    s.Properties = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetProductImageList) SetType(v string) *TaobaoFenxiaoDistributorProductsGetProductImageList {
    s.Type = &v
    return s
}
