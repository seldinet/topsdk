package domain


type TmallChannelProductsQueryProductSkuTopDto struct {
    /*
        skuId     */
    SkuId  *int64 `json:"sku_id,omitempty" `

    /*
        sku商家编码     */
    SkuNumber  *string `json:"sku_number,omitempty" `

    /*
        sku后端货品     */
    SkuScItemId  *int64 `json:"sku_sc_item_id,omitempty" `

    /*
        基准价     */
    StandardPrice  *int64 `json:"standard_price,omitempty" `

    /*
        条形码     */
    BarCode  *string `json:"bar_code,omitempty" `

    /*
        图片链接     */
    PictureUrl  *string `json:"picture_url,omitempty" `

}

func (s *TmallChannelProductsQueryProductSkuTopDto) SetSkuId(v int64) *TmallChannelProductsQueryProductSkuTopDto {
    s.SkuId = &v
    return s
}
func (s *TmallChannelProductsQueryProductSkuTopDto) SetSkuNumber(v string) *TmallChannelProductsQueryProductSkuTopDto {
    s.SkuNumber = &v
    return s
}
func (s *TmallChannelProductsQueryProductSkuTopDto) SetSkuScItemId(v int64) *TmallChannelProductsQueryProductSkuTopDto {
    s.SkuScItemId = &v
    return s
}
func (s *TmallChannelProductsQueryProductSkuTopDto) SetStandardPrice(v int64) *TmallChannelProductsQueryProductSkuTopDto {
    s.StandardPrice = &v
    return s
}
func (s *TmallChannelProductsQueryProductSkuTopDto) SetBarCode(v string) *TmallChannelProductsQueryProductSkuTopDto {
    s.BarCode = &v
    return s
}
func (s *TmallChannelProductsQueryProductSkuTopDto) SetPictureUrl(v string) *TmallChannelProductsQueryProductSkuTopDto {
    s.PictureUrl = &v
    return s
}
