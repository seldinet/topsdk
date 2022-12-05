package domain


type TaobaoFenxiaoDistributorProductsGetPduList struct {
    /*
        产品ID     */
    ProductId  *int64 `json:"product_id,omitempty" `

    /*
        分销商ID     */
    DistributorId  *int64 `json:"distributor_id,omitempty" `

    /*
        分销商用户名     */
    DistributorName  *string `json:"distributor_name,omitempty" `

    /*
        产品销售属性值     */
    SkuProperties  *string `json:"sku_properties,omitempty" `

    /*
        产品代销配额库存     */
    QuantityAgent  *int64 `json:"quantity_agent,omitempty" `

}

func (s *TaobaoFenxiaoDistributorProductsGetPduList) SetProductId(v int64) *TaobaoFenxiaoDistributorProductsGetPduList {
    s.ProductId = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetPduList) SetDistributorId(v int64) *TaobaoFenxiaoDistributorProductsGetPduList {
    s.DistributorId = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetPduList) SetDistributorName(v string) *TaobaoFenxiaoDistributorProductsGetPduList {
    s.DistributorName = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetPduList) SetSkuProperties(v string) *TaobaoFenxiaoDistributorProductsGetPduList {
    s.SkuProperties = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetPduList) SetQuantityAgent(v int64) *TaobaoFenxiaoDistributorProductsGetPduList {
    s.QuantityAgent = &v
    return s
}
