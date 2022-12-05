package domain


type TaobaoFenxiaoProductcatsGetProductCat struct {
    /*
        产品线ID     */
    Id  *int64 `json:"id,omitempty" `

    /*
        产品线名称     */
    Name  *string `json:"name,omitempty" `

    /*
        产品数量     */
    ProductNum  *int64 `json:"product_num,omitempty" `

    /*
        产品最低零售价     */
    RetailLowPercent  *string `json:"retail_low_percent,omitempty" `

    /*
        产品最高零售价     */
    RetailHighPercent  *string `json:"retail_high_percent,omitempty" `

    /*
        产品代销采购价     */
    CostPercentAgent  *string `json:"cost_percent_agent,omitempty" `

    /*
        产品经销采购价     */
    CostPercentDealer  *string `json:"cost_percent_dealer,omitempty" `

}

func (s *TaobaoFenxiaoProductcatsGetProductCat) SetId(v int64) *TaobaoFenxiaoProductcatsGetProductCat {
    s.Id = &v
    return s
}
func (s *TaobaoFenxiaoProductcatsGetProductCat) SetName(v string) *TaobaoFenxiaoProductcatsGetProductCat {
    s.Name = &v
    return s
}
func (s *TaobaoFenxiaoProductcatsGetProductCat) SetProductNum(v int64) *TaobaoFenxiaoProductcatsGetProductCat {
    s.ProductNum = &v
    return s
}
func (s *TaobaoFenxiaoProductcatsGetProductCat) SetRetailLowPercent(v string) *TaobaoFenxiaoProductcatsGetProductCat {
    s.RetailLowPercent = &v
    return s
}
func (s *TaobaoFenxiaoProductcatsGetProductCat) SetRetailHighPercent(v string) *TaobaoFenxiaoProductcatsGetProductCat {
    s.RetailHighPercent = &v
    return s
}
func (s *TaobaoFenxiaoProductcatsGetProductCat) SetCostPercentAgent(v string) *TaobaoFenxiaoProductcatsGetProductCat {
    s.CostPercentAgent = &v
    return s
}
func (s *TaobaoFenxiaoProductcatsGetProductCat) SetCostPercentDealer(v string) *TaobaoFenxiaoProductcatsGetProductCat {
    s.CostPercentDealer = &v
    return s
}
