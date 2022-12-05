package domain


type TmallChannelProductsQueryProductTopDto struct {
    /*
        产品线ID     */
    ProductLineId  *int64 `json:"product_line_id,omitempty" `

    /*
        产品编码     */
    ProductNumber  *string `json:"product_number,omitempty" `

    /*
        产品描述地址     */
    DescPath  *string `json:"desc_path,omitempty" `

    /*
        基准价     */
    StandardPrice  *int64 `json:"standard_price,omitempty" `

    /*
        标题     */
    Title  *string `json:"title,omitempty" `

    /*
        产品Id     */
    ProductId  *int64 `json:"product_id,omitempty" `

    /*
        没有sku的情况下，产品对应的后端商品id     */
    ScItemId  *int64 `json:"sc_item_id,omitempty" `

    /*
        spuId     */
    SpuId  *int64 `json:"spu_id,omitempty" `

    /*
        供应商Id     */
    SupplierId  *int64 `json:"supplier_id,omitempty" `

    /*
        类目Id     */
    CategoryId  *int64 `json:"category_id,omitempty" `

    /*
        sku列表     */
    SkuList  *[]TmallChannelProductsQueryProductSkuTopDto `json:"sku_list,omitempty" `

}

func (s *TmallChannelProductsQueryProductTopDto) SetProductLineId(v int64) *TmallChannelProductsQueryProductTopDto {
    s.ProductLineId = &v
    return s
}
func (s *TmallChannelProductsQueryProductTopDto) SetProductNumber(v string) *TmallChannelProductsQueryProductTopDto {
    s.ProductNumber = &v
    return s
}
func (s *TmallChannelProductsQueryProductTopDto) SetDescPath(v string) *TmallChannelProductsQueryProductTopDto {
    s.DescPath = &v
    return s
}
func (s *TmallChannelProductsQueryProductTopDto) SetStandardPrice(v int64) *TmallChannelProductsQueryProductTopDto {
    s.StandardPrice = &v
    return s
}
func (s *TmallChannelProductsQueryProductTopDto) SetTitle(v string) *TmallChannelProductsQueryProductTopDto {
    s.Title = &v
    return s
}
func (s *TmallChannelProductsQueryProductTopDto) SetProductId(v int64) *TmallChannelProductsQueryProductTopDto {
    s.ProductId = &v
    return s
}
func (s *TmallChannelProductsQueryProductTopDto) SetScItemId(v int64) *TmallChannelProductsQueryProductTopDto {
    s.ScItemId = &v
    return s
}
func (s *TmallChannelProductsQueryProductTopDto) SetSpuId(v int64) *TmallChannelProductsQueryProductTopDto {
    s.SpuId = &v
    return s
}
func (s *TmallChannelProductsQueryProductTopDto) SetSupplierId(v int64) *TmallChannelProductsQueryProductTopDto {
    s.SupplierId = &v
    return s
}
func (s *TmallChannelProductsQueryProductTopDto) SetCategoryId(v int64) *TmallChannelProductsQueryProductTopDto {
    s.CategoryId = &v
    return s
}
func (s *TmallChannelProductsQueryProductTopDto) SetSkuList(v []TmallChannelProductsQueryProductSkuTopDto) *TmallChannelProductsQueryProductTopDto {
    s.SkuList = &v
    return s
}
