package domain


type TaobaoTradeFullinfoGetPromotionDetail struct {
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

    /*
        赠品的宝贝id     */
    GiftItemId  *string `json:"gift_item_id,omitempty" `

    /*
        满就送礼物的礼物数量     */
    GiftItemNum  *string `json:"gift_item_num,omitempty" `

    /*
        优惠活动的描述     */
    PromotionDesc  *string `json:"promotion_desc,omitempty" `

    /*
        优惠id，(由营销工具id、优惠活动id和优惠详情id组成，结构为：营销工具id-优惠活动id_优惠详情id，如mjs-123024_211143）     */
    PromotionId  *string `json:"promotion_id,omitempty" `

    /*
        分摊优惠金额（免运费、限时打折时为空）,单位：元     */
    KdDiscountFee  *string `json:"kd_discount_fee,omitempty" `

    /*
        若优惠项在主订单上，返回子订单的分摊信息     */
    KdChildDiscountFee  *string `json:"kd_child_discount_fee,omitempty" `

}

func (s *TaobaoTradeFullinfoGetPromotionDetail) SetId(v int64) *TaobaoTradeFullinfoGetPromotionDetail {
    s.Id = &v
    return s
}
func (s *TaobaoTradeFullinfoGetPromotionDetail) SetPromotionName(v string) *TaobaoTradeFullinfoGetPromotionDetail {
    s.PromotionName = &v
    return s
}
func (s *TaobaoTradeFullinfoGetPromotionDetail) SetDiscountFee(v string) *TaobaoTradeFullinfoGetPromotionDetail {
    s.DiscountFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetPromotionDetail) SetGiftItemName(v string) *TaobaoTradeFullinfoGetPromotionDetail {
    s.GiftItemName = &v
    return s
}
func (s *TaobaoTradeFullinfoGetPromotionDetail) SetGiftItemId(v string) *TaobaoTradeFullinfoGetPromotionDetail {
    s.GiftItemId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetPromotionDetail) SetGiftItemNum(v string) *TaobaoTradeFullinfoGetPromotionDetail {
    s.GiftItemNum = &v
    return s
}
func (s *TaobaoTradeFullinfoGetPromotionDetail) SetPromotionDesc(v string) *TaobaoTradeFullinfoGetPromotionDetail {
    s.PromotionDesc = &v
    return s
}
func (s *TaobaoTradeFullinfoGetPromotionDetail) SetPromotionId(v string) *TaobaoTradeFullinfoGetPromotionDetail {
    s.PromotionId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetPromotionDetail) SetKdDiscountFee(v string) *TaobaoTradeFullinfoGetPromotionDetail {
    s.KdDiscountFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetPromotionDetail) SetKdChildDiscountFee(v string) *TaobaoTradeFullinfoGetPromotionDetail {
    s.KdChildDiscountFee = &v
    return s
}
