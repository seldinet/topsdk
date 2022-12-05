package domain


type TaobaoFenxiaoDistributorProductsGetSkuList struct {
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
        商家编码     */
    OuterId  *string `json:"outer_id,omitempty" `

    /*
        sku的销售属性组合字符串。格式:pid:vid;pid:vid,如:1627207:3232483;1630696:3284570,表示:机身颜色:军绿色;手机套餐:一电一充。     */
    Properties  *string `json:"properties,omitempty" `

    /*
        关联的后端商品id     */
    ScitemId  *int64 `json:"scitem_id,omitempty" `

}

func (s *TaobaoFenxiaoDistributorProductsGetSkuList) SetId(v int64) *TaobaoFenxiaoDistributorProductsGetSkuList {
    s.Id = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetSkuList) SetName(v string) *TaobaoFenxiaoDistributorProductsGetSkuList {
    s.Name = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetSkuList) SetStandardPrice(v string) *TaobaoFenxiaoDistributorProductsGetSkuList {
    s.StandardPrice = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetSkuList) SetCostPrice(v string) *TaobaoFenxiaoDistributorProductsGetSkuList {
    s.CostPrice = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetSkuList) SetDealerCostPrice(v string) *TaobaoFenxiaoDistributorProductsGetSkuList {
    s.DealerCostPrice = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetSkuList) SetQuantity(v int64) *TaobaoFenxiaoDistributorProductsGetSkuList {
    s.Quantity = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetSkuList) SetOuterId(v string) *TaobaoFenxiaoDistributorProductsGetSkuList {
    s.OuterId = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetSkuList) SetProperties(v string) *TaobaoFenxiaoDistributorProductsGetSkuList {
    s.Properties = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetSkuList) SetScitemId(v int64) *TaobaoFenxiaoDistributorProductsGetSkuList {
    s.ScitemId = &v
    return s
}
