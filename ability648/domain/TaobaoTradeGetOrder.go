package domain


type TaobaoTradeGetOrder struct {
    /*
        商品标题     */
    Title  *string `json:"title,omitempty" `

    /*
        商品图片的绝对路径     */
    PicPath  *string `json:"pic_path,omitempty" `

    /*
        商品价格。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
    Price  *string `json:"price,omitempty" `

    /*
        商品数字ID     */
    NumIid  *int64 `json:"num_iid,omitempty" `

    /*
        商品的最小库存单位Sku的id.可以通过taobao.item.sku.get获取详细的Sku信息     */
    SkuId  *string `json:"sku_id,omitempty" `

    /*
        退款状态。退款状态。可选值 WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意) WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货) WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货) SELLER_REFUSE_BUYER(卖家拒绝退款) CLOSED(退款关闭) SUCCESS(退款成功)     */
    RefundStatus  *string `json:"refund_status,omitempty" `

    /*
        订单状态（请关注此状态，如果为TRADE_CLOSED_BY_TAOBAO状态，则不要对此订单进行发货，切记啊！）。可选值: <ul><li>TRADE_NO_CREATE_PAY(没有创建支付宝交易) <li>WAIT_BUYER_PAY(等待买家付款) <li>WAIT_SELLER_SEND_GOODS(等待卖家发货,即:买家已付款) <li>WAIT_BUYER_CONFIRM_GOODS(等待买家确认收货,即:卖家已发货) <li>TRADE_BUYER_SIGNED(买家已签收,货到付款专用) <li>TRADE_FINISHED(交易成功) <li>TRADE_CLOSED(付款以后用户退款成功，交易自动关闭) <li>TRADE_CLOSED_BY_TAOBAO(付款以前，卖家或买家主动关闭交易)<li>PAY_PENDING(国际信用卡支付付款确认中)     */
    Status  *string `json:"status,omitempty" `

    /*
        子订单编号     */
    Oid  *int64 `json:"oid,omitempty" `

    /*
        应付金额（商品价格 * 商品数量 + 手工调整金额 - 子订单级订单优惠金额）。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
    TotalFee  *string `json:"total_fee,omitempty" `

    /*
        子订单实付金额。精确到2位小数，单位:元。如:200.07，表示:200元7分。对于多子订单的交易，计算公式如下：payment = price * num + adjust_fee - discount_fee ；单子订单交易，payment与主订单的payment一致，对于退款成功的子订单，由于主订单的优惠分摊金额，会造成该字段可能不为0.00元。建议使用退款前的实付金额减去退款单中的实际退款金额计算。     */
    Payment  *string `json:"payment,omitempty" `

    /*
        子订单级订单优惠金额。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
    DiscountFee  *string `json:"discount_fee,omitempty" `

    /*
        手工调整金额.格式为:1.01;单位:元;精确到小数点后两位.     */
    AdjustFee  *string `json:"adjust_fee,omitempty" `

    /*
        分摊之后的实付金额     */
    DivideOrderFee  *string `json:"divide_order_fee,omitempty" `

    /*
        优惠分摊     */
    PartMjzDiscount  *string `json:"part_mjz_discount,omitempty" `

    /*
        SKU的值。如：机身颜色:黑色;手机套餐:官方标配     */
    SkuPropertiesName  *string `json:"sku_properties_name,omitempty" `

    /*
        套餐的值。如：M8原装电池:便携支架:M8专用座充:莫凡保护袋     */
    ItemMealName  *string `json:"item_meal_name,omitempty" `

    /*
        购买数量。取值范围:大于零的整数     */
    Num  *int64 `json:"num,omitempty" `

    /*
        外部网店自己定义的Sku编号     */
    OuterSkuId  *string `json:"outer_sku_id,omitempty" `

    /*
        商家外部编码(可与商家外部系统对接)。外部商家自己定义的商品Item的id，可以通过taobao.items.custom.get获取商品的Item的信息     */
    OuterIid  *string `json:"outer_iid,omitempty" `

    /*
        买家是否已评价。可选值：true(已评价)，false(未评价)     */
    BuyerRate  *bool `json:"buyer_rate,omitempty" `

    /*
        卖家是否已评价。可选值：true(已评价)，false(未评价)     */
    SellerRate  *bool `json:"seller_rate,omitempty" `

    /*
        卖家类型，可选值为：B（商城商家），C（普通卖家）     */
    SellerType  *string `json:"seller_type,omitempty" `

    /*
        表示订单交易是否含有对应的代销采购单。如果该订单中存在一个对应的代销采购单，那么该值为true；反之，该值为false。     */
    IsDaixiao  *bool `json:"is_daixiao,omitempty" `

    /*
        对应门票有效期的外部id     */
    TicketOuterId  *string `json:"ticket_outer_id,omitempty" `

    /*
        门票有效期的key     */
    TicketExpdateKey  *string `json:"ticket_expdate_key,omitempty" `

    /*
        子订单是否是www订单     */
    IsWww  *bool `json:"is_www,omitempty" `

    /*
        发货的仓库编码     */
    StoreCode  *string `json:"store_code,omitempty" `

    /*
        支持家装类物流的类型     */
    TmserSpuCode  *string `json:"tmser_spu_code,omitempty" `

    /*
        天猫汽车服务预约时间     */
    EtSerTime  *string `json:"et_ser_time,omitempty" `

    /*
        电子凭证预约门店地址     */
    EtShopName  *string `json:"et_shop_name,omitempty" `

    /*
        电子凭证核销门店地址     */
    EtVerifiedShopName  *string `json:"et_verified_shop_name,omitempty" `

    /*
        车牌号码     */
    EtPlateNumber  *string `json:"et_plate_number,omitempty" `

    /*
        天猫国际子订单关税税费     */
    SubOrderTaxFee  *string `json:"sub_order_tax_fee,omitempty" `

    /*
        天猫国际子订单关税税率     */
    SubOrderTaxRate  *string `json:"sub_order_tax_rate,omitempty" `

    /*
        征集预售订单征集状态：1（征集中），2（征集成功），3（征集失败）     */
    ZhengjiStatus  *string `json:"zhengji_status,omitempty" `

    /*
        天猫国际子订单计税优惠金额     */
    SubOrderTaxPromotionFee  *string `json:"sub_order_tax_promotion_fee,omitempty" `

    /*
        天猫国际子订单是否包税     */
    TaxFree  *bool `json:"tax_free,omitempty" `

    /*
        天猫国际子订单包税金额     */
    TaxCouponDiscount  *string `json:"tax_coupon_discount,omitempty" `

    /*
        特殊退款状态     */
    SpecialRefundType  *string `json:"special_refund_type,omitempty" `

    /*
        子订单是否优惠贬值     */
    IsDevalueFee  *bool `json:"is_devalue_fee,omitempty" `

}

func (s *TaobaoTradeGetOrder) SetTitle(v string) *TaobaoTradeGetOrder {
    s.Title = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetPicPath(v string) *TaobaoTradeGetOrder {
    s.PicPath = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetPrice(v string) *TaobaoTradeGetOrder {
    s.Price = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetNumIid(v int64) *TaobaoTradeGetOrder {
    s.NumIid = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetSkuId(v string) *TaobaoTradeGetOrder {
    s.SkuId = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetRefundStatus(v string) *TaobaoTradeGetOrder {
    s.RefundStatus = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetStatus(v string) *TaobaoTradeGetOrder {
    s.Status = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetOid(v int64) *TaobaoTradeGetOrder {
    s.Oid = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetTotalFee(v string) *TaobaoTradeGetOrder {
    s.TotalFee = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetPayment(v string) *TaobaoTradeGetOrder {
    s.Payment = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetDiscountFee(v string) *TaobaoTradeGetOrder {
    s.DiscountFee = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetAdjustFee(v string) *TaobaoTradeGetOrder {
    s.AdjustFee = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetDivideOrderFee(v string) *TaobaoTradeGetOrder {
    s.DivideOrderFee = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetPartMjzDiscount(v string) *TaobaoTradeGetOrder {
    s.PartMjzDiscount = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetSkuPropertiesName(v string) *TaobaoTradeGetOrder {
    s.SkuPropertiesName = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetItemMealName(v string) *TaobaoTradeGetOrder {
    s.ItemMealName = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetNum(v int64) *TaobaoTradeGetOrder {
    s.Num = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetOuterSkuId(v string) *TaobaoTradeGetOrder {
    s.OuterSkuId = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetOuterIid(v string) *TaobaoTradeGetOrder {
    s.OuterIid = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetBuyerRate(v bool) *TaobaoTradeGetOrder {
    s.BuyerRate = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetSellerRate(v bool) *TaobaoTradeGetOrder {
    s.SellerRate = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetSellerType(v string) *TaobaoTradeGetOrder {
    s.SellerType = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetIsDaixiao(v bool) *TaobaoTradeGetOrder {
    s.IsDaixiao = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetTicketOuterId(v string) *TaobaoTradeGetOrder {
    s.TicketOuterId = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetTicketExpdateKey(v string) *TaobaoTradeGetOrder {
    s.TicketExpdateKey = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetIsWww(v bool) *TaobaoTradeGetOrder {
    s.IsWww = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetStoreCode(v string) *TaobaoTradeGetOrder {
    s.StoreCode = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetTmserSpuCode(v string) *TaobaoTradeGetOrder {
    s.TmserSpuCode = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetEtSerTime(v string) *TaobaoTradeGetOrder {
    s.EtSerTime = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetEtShopName(v string) *TaobaoTradeGetOrder {
    s.EtShopName = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetEtVerifiedShopName(v string) *TaobaoTradeGetOrder {
    s.EtVerifiedShopName = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetEtPlateNumber(v string) *TaobaoTradeGetOrder {
    s.EtPlateNumber = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetSubOrderTaxFee(v string) *TaobaoTradeGetOrder {
    s.SubOrderTaxFee = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetSubOrderTaxRate(v string) *TaobaoTradeGetOrder {
    s.SubOrderTaxRate = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetZhengjiStatus(v string) *TaobaoTradeGetOrder {
    s.ZhengjiStatus = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetSubOrderTaxPromotionFee(v string) *TaobaoTradeGetOrder {
    s.SubOrderTaxPromotionFee = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetTaxFree(v bool) *TaobaoTradeGetOrder {
    s.TaxFree = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetTaxCouponDiscount(v string) *TaobaoTradeGetOrder {
    s.TaxCouponDiscount = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetSpecialRefundType(v string) *TaobaoTradeGetOrder {
    s.SpecialRefundType = &v
    return s
}
func (s *TaobaoTradeGetOrder) SetIsDevalueFee(v bool) *TaobaoTradeGetOrder {
    s.IsDevalueFee = &v
    return s
}
