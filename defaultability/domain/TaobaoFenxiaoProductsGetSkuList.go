package domain


type TaobaoFenxiaoProductsGetSkuList struct {
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

}

func (s *TaobaoFenxiaoProductsGetSkuList) SetId(v int64) *TaobaoFenxiaoProductsGetSkuList {
    s.Id = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetSkuList) SetName(v string) *TaobaoFenxiaoProductsGetSkuList {
    s.Name = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetSkuList) SetStandardPrice(v string) *TaobaoFenxiaoProductsGetSkuList {
    s.StandardPrice = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetSkuList) SetCostPrice(v string) *TaobaoFenxiaoProductsGetSkuList {
    s.CostPrice = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetSkuList) SetDealerCostPrice(v string) *TaobaoFenxiaoProductsGetSkuList {
    s.DealerCostPrice = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetSkuList) SetQuantity(v int64) *TaobaoFenxiaoProductsGetSkuList {
    s.Quantity = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetSkuList) SetReservedQuantity(v int64) *TaobaoFenxiaoProductsGetSkuList {
    s.ReservedQuantity = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetSkuList) SetQuotaQuantity(v int64) *TaobaoFenxiaoProductsGetSkuList {
    s.QuotaQuantity = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetSkuList) SetOuterId(v string) *TaobaoFenxiaoProductsGetSkuList {
    s.OuterId = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetSkuList) SetProperties(v string) *TaobaoFenxiaoProductsGetSkuList {
    s.Properties = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetSkuList) SetScitemId(v int64) *TaobaoFenxiaoProductsGetSkuList {
    s.ScitemId = &v
    return s
}
