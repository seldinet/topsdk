package domain


type TaobaoTradeAmountGetPromotionDetail struct {
    /*
        交易的主订单或子订单号     */
    Id  *int64 `json:"id,omitempty" `

    /*
        优惠信息的名称     */
    PromotionName  *string `json:"promotion_name,omitempty" `

    /*
        优惠金额（免运费、限时打折时为空）,单位：元     */
    DiscountFee  *string `json:"discount_fee,omitempty" `

    /*
        满就送商品时，所送商品的名称     */
    GiftItemName  *string `json:"gift_item_name,omitempty" `

}

func (s *TaobaoTradeAmountGetPromotionDetail) SetId(v int64) *TaobaoTradeAmountGetPromotionDetail {
    s.Id = &v
    return s
}
func (s *TaobaoTradeAmountGetPromotionDetail) SetPromotionName(v string) *TaobaoTradeAmountGetPromotionDetail {
    s.PromotionName = &v
    return s
}
func (s *TaobaoTradeAmountGetPromotionDetail) SetDiscountFee(v string) *TaobaoTradeAmountGetPromotionDetail {
    s.DiscountFee = &v
    return s
}
func (s *TaobaoTradeAmountGetPromotionDetail) SetGiftItemName(v string) *TaobaoTradeAmountGetPromotionDetail {
    s.GiftItemName = &v
    return s
}
