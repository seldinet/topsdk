package domain


type TaobaoTradeAmountGetOrderAmount struct {
    /*
        子交易订单编号     */
    Oid  *int64 `json:"oid,omitempty" `

    /*
        商品标题     */
    Title  *string `json:"title,omitempty" `

    /*
        子订单对应的商品数字id     */
    NumIid  *int64 `json:"num_iid,omitempty" `

    /*
        子订单对应的商品的sku_id     */
    SkuId  *int64 `json:"sku_id,omitempty" `

    /*
        SKU的值。如：机身颜色:黑色;手机套餐:官方标配     */
    SkuPropertiesName  *string `json:"sku_properties_name,omitempty" `

    /*
        子交易订单中购买商品的数量     */
    Num  *int64 `json:"num,omitempty" `

    /*
        商品价格。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
    Price  *string `json:"price,omitempty" `

    /*
        子订单的系统优惠金额。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
    DiscountFee  *string `json:"discount_fee,omitempty" `

    /*
        子订单的系统优惠的名称，对应于discount_fee的名称     */
    PromotionName  *string `json:"promotion_name,omitempty" `

    /*
        卖家手工调整的子订单的优惠金额.格式为:1.01;单位:元;精确到小数点后两位.     */
    AdjustFee  *string `json:"adjust_fee,omitempty" `

    /*
        分摊之后的实付金额     */
    DivideOrderFee  *string `json:"divide_order_fee,omitempty" `

    /*
        优惠分摊     */
    PartMjzDiscount  *string `json:"part_mjz_discount,omitempty" `

    /*
        子订单实付金额。精确到2位小数，单位:元。如:200.07，表示:200元7分。计算公式如下：payment = price * num + adjust_fee - discount_fee + post_fee(邮费，单笔子订单时子订单实付金额包含邮费，多笔子订单时不包含邮费)；对于退款成功的子订单，由于主订单的优惠分摊金额，会造成该字段可能不为0.00元。建议使用退款前的实付金额减去退款单中的实际退款金额计算。     */
    Payment  *string `json:"payment,omitempty" `

}

func (s *TaobaoTradeAmountGetOrderAmount) SetOid(v int64) *TaobaoTradeAmountGetOrderAmount {
    s.Oid = &v
    return s
}
func (s *TaobaoTradeAmountGetOrderAmount) SetTitle(v string) *TaobaoTradeAmountGetOrderAmount {
    s.Title = &v
    return s
}
func (s *TaobaoTradeAmountGetOrderAmount) SetNumIid(v int64) *TaobaoTradeAmountGetOrderAmount {
    s.NumIid = &v
    return s
}
func (s *TaobaoTradeAmountGetOrderAmount) SetSkuId(v int64) *TaobaoTradeAmountGetOrderAmount {
    s.SkuId = &v
    return s
}
func (s *TaobaoTradeAmountGetOrderAmount) SetSkuPropertiesName(v string) *TaobaoTradeAmountGetOrderAmount {
    s.SkuPropertiesName = &v
    return s
}
func (s *TaobaoTradeAmountGetOrderAmount) SetNum(v int64) *TaobaoTradeAmountGetOrderAmount {
    s.Num = &v
    return s
}
func (s *TaobaoTradeAmountGetOrderAmount) SetPrice(v string) *TaobaoTradeAmountGetOrderAmount {
    s.Price = &v
    return s
}
func (s *TaobaoTradeAmountGetOrderAmount) SetDiscountFee(v string) *TaobaoTradeAmountGetOrderAmount {
    s.DiscountFee = &v
    return s
}
func (s *TaobaoTradeAmountGetOrderAmount) SetPromotionName(v string) *TaobaoTradeAmountGetOrderAmount {
    s.PromotionName = &v
    return s
}
func (s *TaobaoTradeAmountGetOrderAmount) SetAdjustFee(v string) *TaobaoTradeAmountGetOrderAmount {
    s.AdjustFee = &v
    return s
}
func (s *TaobaoTradeAmountGetOrderAmount) SetDivideOrderFee(v string) *TaobaoTradeAmountGetOrderAmount {
    s.DivideOrderFee = &v
    return s
}
func (s *TaobaoTradeAmountGetOrderAmount) SetPartMjzDiscount(v string) *TaobaoTradeAmountGetOrderAmount {
    s.PartMjzDiscount = &v
    return s
}
func (s *TaobaoTradeAmountGetOrderAmount) SetPayment(v string) *TaobaoTradeAmountGetOrderAmount {
    s.Payment = &v
    return s
}
