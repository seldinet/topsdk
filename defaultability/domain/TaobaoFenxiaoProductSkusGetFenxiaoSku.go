package domain


type TaobaoFenxiaoProductSkusGetFenxiaoSku struct {
    /*
        SkuID     */
    Id  *int64 `json:"id,omitempty" `

    /*
        名称     */
    Name  *string `json:"name,omitempty" `

    /*
        市场价     */
    StandardPrice  *string `json:"standard_price,omitempty" `

    /*
        代销采购价，单位：元     */
    CostPrice  *string `json:"cost_price,omitempty" `

    /*
        经销采购价     */
    DealerCostPrice  *string `json:"dealer_cost_price,omitempty" `

    /*
        库存     */
    Quantity  *int64 `json:"quantity,omitempty" `

    /*
        预扣库存     */
    ReservedQuantity  *int64 `json:"reserved_quantity,omitempty" `

    /*
        配额可用库存     */
    QuotaQuantity  *int64 `json:"quota_quantity,omitempty" `

    /*
        商家编码     */
    OuterId  *string `json:"outer_id,omitempty" `

    /*
        sku的销售属性组合字符串。格式:pid:vid;pid:vid,如:1627207:3232483;1630696:3284570,表示:机身颜色:军绿色;手机套餐:一电一充。     */
    Properties  *string `json:"properties,omitempty" `

    /*
        关联的后端商品id     */
    ScitemId  *int64 `json:"scitem_id,omitempty" `

    /*
        关联的前端宝贝对应的skuid     */
    AuctionSkuId  *int64 `json:"auction_sku_id,omitempty" `

}

func (s *TaobaoFenxiaoProductSkusGetFenxiaoSku) SetId(v int64) *TaobaoFenxiaoProductSkusGetFenxiaoSku {
    s.Id = &v
    return s
}
func (s *TaobaoFenxiaoProductSkusGetFenxiaoSku) SetName(v string) *TaobaoFenxiaoProductSkusGetFenxiaoSku {
    s.Name = &v
    return s
}
func (s *TaobaoFenxiaoProductSkusGetFenxiaoSku) SetStandardPrice(v string) *TaobaoFenxiaoProductSkusGetFenxiaoSku {
    s.StandardPrice = &v
    return s
}
func (s *TaobaoFenxiaoProductSkusGetFenxiaoSku) SetCostPrice(v string) *TaobaoFenxiaoProductSkusGetFenxiaoSku {
    s.CostPrice = &v
    return s
}
func (s *TaobaoFenxiaoProductSkusGetFenxiaoSku) SetDealerCostPrice(v string) *TaobaoFenxiaoProductSkusGetFenxiaoSku {
    s.DealerCostPrice = &v
    return s
}
func (s *TaobaoFenxiaoProductSkusGetFenxiaoSku) SetQuantity(v int64) *TaobaoFenxiaoProductSkusGetFenxiaoSku {
    s.Quantity = &v
    return s
}
func (s *TaobaoFenxiaoProductSkusGetFenxiaoSku) SetReservedQuantity(v int64) *TaobaoFenxiaoProductSkusGetFenxiaoSku {
    s.ReservedQuantity = &v
    return s
}
func (s *TaobaoFenxiaoProductSkusGetFenxiaoSku) SetQuotaQuantity(v int64) *TaobaoFenxiaoProductSkusGetFenxiaoSku {
    s.QuotaQuantity = &v
    return s
}
func (s *TaobaoFenxiaoProductSkusGetFenxiaoSku) SetOuterId(v string) *TaobaoFenxiaoProductSkusGetFenxiaoSku {
    s.OuterId = &v
    return s
}
func (s *TaobaoFenxiaoProductSkusGetFenxiaoSku) SetProperties(v string) *TaobaoFenxiaoProductSkusGetFenxiaoSku {
    s.Properties = &v
    return s
}
func (s *TaobaoFenxiaoProductSkusGetFenxiaoSku) SetScitemId(v int64) *TaobaoFenxiaoProductSkusGetFenxiaoSku {
    s.ScitemId = &v
    return s
}
func (s *TaobaoFenxiaoProductSkusGetFenxiaoSku) SetAuctionSkuId(v int64) *TaobaoFenxiaoProductSkusGetFenxiaoSku {
    s.AuctionSkuId = &v
    return s
}
