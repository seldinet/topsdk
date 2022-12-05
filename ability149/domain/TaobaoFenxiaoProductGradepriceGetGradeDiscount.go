package domain


type TaobaoFenxiaoProductGradepriceGetGradeDiscount struct {
    /*
        采购价格     */
    Price  *string `json:"price,omitempty" `

    /*
        折扣类型：1-等级、2-分销商折扣     */
    DiscountType  *int64 `json:"discount_type,omitempty" `

    /*
        等级ID或分销商ID     */
    DiscountId  *int64 `json:"discount_id,omitempty" `

    /*
        skuId     */
    SkuId  *int64 `json:"sku_id,omitempty" `

    /*
        模式：1-代销、2-经销     */
    TradeType  *int64 `json:"trade_type,omitempty" `

}

func (s *TaobaoFenxiaoProductGradepriceGetGradeDiscount) SetPrice(v string) *TaobaoFenxiaoProductGradepriceGetGradeDiscount {
    s.Price = &v
    return s
}
func (s *TaobaoFenxiaoProductGradepriceGetGradeDiscount) SetDiscountType(v int64) *TaobaoFenxiaoProductGradepriceGetGradeDiscount {
    s.DiscountType = &v
    return s
}
func (s *TaobaoFenxiaoProductGradepriceGetGradeDiscount) SetDiscountId(v int64) *TaobaoFenxiaoProductGradepriceGetGradeDiscount {
    s.DiscountId = &v
    return s
}
func (s *TaobaoFenxiaoProductGradepriceGetGradeDiscount) SetSkuId(v int64) *TaobaoFenxiaoProductGradepriceGetGradeDiscount {
    s.SkuId = &v
    return s
}
func (s *TaobaoFenxiaoProductGradepriceGetGradeDiscount) SetTradeType(v int64) *TaobaoFenxiaoProductGradepriceGetGradeDiscount {
    s.TradeType = &v
    return s
}
