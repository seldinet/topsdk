package domain


type TaobaoFenxiaoProductsGetPduList struct {
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

func (s *TaobaoFenxiaoProductsGetPduList) SetProductId(v int64) *TaobaoFenxiaoProductsGetPduList {
    s.ProductId = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetPduList) SetDistributorId(v int64) *TaobaoFenxiaoProductsGetPduList {
    s.DistributorId = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetPduList) SetDistributorName(v string) *TaobaoFenxiaoProductsGetPduList {
    s.DistributorName = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetPduList) SetSkuProperties(v string) *TaobaoFenxiaoProductsGetPduList {
    s.SkuProperties = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetPduList) SetQuantityAgent(v int64) *TaobaoFenxiaoProductsGetPduList {
    s.QuantityAgent = &v
    return s
}
