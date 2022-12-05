package domain


type TmallChannelProductsGetProductSkuDo struct {
    /*
        关联的前端宝贝skuid     */
    AuctionSkuId  *int64 `json:"auction_sku_id,omitempty" `

    /*
        代销采购价:单位分     */
    CostPriceFen  *int64 `json:"cost_price_fen,omitempty" `

    /*
        经销采购价:单位分     */
    PriceCostDealerFen  *int64 `json:"price_cost_dealer_fen,omitempty" `

    /*
        sku的销售属性组合字符串。格式:pid:vid;pid:vid,如:1627207:3232483;1630696:3284570,表示:机身颜色:军绿色;手机套餐:一电一充。     */
    Properties  *string `json:"properties,omitempty" `

    /*
        库存     */
    Quantity  *int64 `json:"quantity,omitempty" `

    /*
        配额可用库存     */
    QuotaQuantity  *int64 `json:"quota_quantity,omitempty" `

    /*
        预扣库存     */
    ReservedQuantity  *int64 `json:"reserved_quantity,omitempty" `

    /*
        关联的后端商品id     */
    ScitemId  *int64 `json:"scitem_id,omitempty" `

    /*
        skuId     */
    SkuId  *int64 `json:"sku_id,omitempty" `

    /*
        商家编码     */
    OuterId  *string `json:"outer_id,omitempty" `

    /*
        市场价单位分     */
    StandardPriceFen  *int64 `json:"standard_price_fen,omitempty" `

}

func (s *TmallChannelProductsGetProductSkuDo) SetAuctionSkuId(v int64) *TmallChannelProductsGetProductSkuDo {
    s.AuctionSkuId = &v
    return s
}
func (s *TmallChannelProductsGetProductSkuDo) SetCostPriceFen(v int64) *TmallChannelProductsGetProductSkuDo {
    s.CostPriceFen = &v
    return s
}
func (s *TmallChannelProductsGetProductSkuDo) SetPriceCostDealerFen(v int64) *TmallChannelProductsGetProductSkuDo {
    s.PriceCostDealerFen = &v
    return s
}
func (s *TmallChannelProductsGetProductSkuDo) SetProperties(v string) *TmallChannelProductsGetProductSkuDo {
    s.Properties = &v
    return s
}
func (s *TmallChannelProductsGetProductSkuDo) SetQuantity(v int64) *TmallChannelProductsGetProductSkuDo {
    s.Quantity = &v
    return s
}
func (s *TmallChannelProductsGetProductSkuDo) SetQuotaQuantity(v int64) *TmallChannelProductsGetProductSkuDo {
    s.QuotaQuantity = &v
    return s
}
func (s *TmallChannelProductsGetProductSkuDo) SetReservedQuantity(v int64) *TmallChannelProductsGetProductSkuDo {
    s.ReservedQuantity = &v
    return s
}
func (s *TmallChannelProductsGetProductSkuDo) SetScitemId(v int64) *TmallChannelProductsGetProductSkuDo {
    s.ScitemId = &v
    return s
}
func (s *TmallChannelProductsGetProductSkuDo) SetSkuId(v int64) *TmallChannelProductsGetProductSkuDo {
    s.SkuId = &v
    return s
}
func (s *TmallChannelProductsGetProductSkuDo) SetOuterId(v string) *TmallChannelProductsGetProductSkuDo {
    s.OuterId = &v
    return s
}
func (s *TmallChannelProductsGetProductSkuDo) SetStandardPriceFen(v int64) *TmallChannelProductsGetProductSkuDo {
    s.StandardPriceFen = &v
    return s
}
