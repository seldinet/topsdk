package domain


type TmallChannelProductsGetTopQueryProductDo struct {
    /*
        要查询的产品id 列表     */
    Ids  *[]int64 `json:"ids,omitempty" `

    /*
        分页大小     */
    PageSize  *int64 `json:"page_size,omitempty" `

    /*
        产品线id     */
    ProductLineId  *int64 `json:"product_line_id,omitempty" `

    /*
        当前要查看的页数     */
    PageNum  *int64 `json:"page_num,omitempty" `

    /*
        产品商家编码     */
    ProductOuterId  *string `json:"product_outer_id,omitempty" `

    /*
        关联的前端宝贝id列表     */
    ItemIds  *[]int64 `json:"item_ids,omitempty" `

    /*
        sku商家编码     */
    SkuOuterId  *string `json:"sku_outer_id,omitempty" `

    /*
        渠道[21 零售plus]     */
    Channel  *int64 `json:"channel,omitempty" `

}

func (s *TmallChannelProductsGetTopQueryProductDo) SetIds(v []int64) *TmallChannelProductsGetTopQueryProductDo {
    s.Ids = &v
    return s
}
func (s *TmallChannelProductsGetTopQueryProductDo) SetPageSize(v int64) *TmallChannelProductsGetTopQueryProductDo {
    s.PageSize = &v
    return s
}
func (s *TmallChannelProductsGetTopQueryProductDo) SetProductLineId(v int64) *TmallChannelProductsGetTopQueryProductDo {
    s.ProductLineId = &v
    return s
}
func (s *TmallChannelProductsGetTopQueryProductDo) SetPageNum(v int64) *TmallChannelProductsGetTopQueryProductDo {
    s.PageNum = &v
    return s
}
func (s *TmallChannelProductsGetTopQueryProductDo) SetProductOuterId(v string) *TmallChannelProductsGetTopQueryProductDo {
    s.ProductOuterId = &v
    return s
}
func (s *TmallChannelProductsGetTopQueryProductDo) SetItemIds(v []int64) *TmallChannelProductsGetTopQueryProductDo {
    s.ItemIds = &v
    return s
}
func (s *TmallChannelProductsGetTopQueryProductDo) SetSkuOuterId(v string) *TmallChannelProductsGetTopQueryProductDo {
    s.SkuOuterId = &v
    return s
}
func (s *TmallChannelProductsGetTopQueryProductDo) SetChannel(v int64) *TmallChannelProductsGetTopQueryProductDo {
    s.Channel = &v
    return s
}
