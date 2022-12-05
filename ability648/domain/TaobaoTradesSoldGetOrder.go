package domain

import (
        "topsdk/util"
    )

type TaobaoTradesSoldGetOrder struct {
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
        商家外部编码(可与商家外部系统对接)。外部商家自己定义的商品Item的id，可以通过taobao.items.custom.get获取商品的Item的信息     */
    OuterIid  *string `json:"outer_iid,omitempty" `

    /*
        外部网店自己定义的Sku编号     */
    OuterSkuId  *string `json:"outer_sku_id,omitempty" `

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
        套餐ID     */
    ItemMealId  *int64 `json:"item_meal_id,omitempty" `

    /*
        套餐的值。如：M8原装电池:便携支架:M8专用座充:莫凡保护袋     */
    ItemMealName  *string `json:"item_meal_name,omitempty" `

    /*
        购买数量。取值范围:大于零的整数     */
    Num  *int64 `json:"num,omitempty" `

    /*
        买家是否已评价。可选值：true(已评价)，false(未评价)     */
    BuyerRate  *bool `json:"buyer_rate,omitempty" `

    /*
        卖家是否已评价。可选值：true(已评价)，false(未评价)     */
    SellerRate  *bool `json:"seller_rate,omitempty" `

    /*
        最近退款ID     */
    RefundId  *string `json:"refund_id,omitempty" `

    /*
        卖家类型，可选值为：B（商城商家），C（普通卖家）     */
    SellerType  *string `json:"seller_type,omitempty" `

    /*
        交易商品对应的类目ID     */
    Cid  *int64 `json:"cid,omitempty" `

    /*
        子订单的交易结束时间说明：子订单有单独的结束时间，与主订单的结束时间可能有所不同，在有退款发起的时候或者是主订单分阶段付款的时候，子订单的结束时间会早于主订单的结束时间，所以开放这个字段便于订单结束状态的判断     */
    EndTime  *util.LocalTime `json:"end_time,omitempty" `

    /*
        子订单发货时间，当卖家对订单进行了多次发货，子订单的发货时间和主订单的发货时间可能不一样了，那么就需要以子订单的时间为准。（没有进行多次发货的订单，主订单的发货时间和子订单的发货时间都一样）     */
    ConsignTime  *string `json:"consign_time,omitempty" `

    /*
        子订单的运送方式（卖家对订单进行多次发货之后，一个主订单下的子订单的运送方式可能不同，用order.shipping_type来区分子订单的运送方式）     */
    ShippingType  *string `json:"shipping_type,omitempty" `

    /*
        子订单发货的快递公司名称     */
    LogisticsCompany  *string `json:"logistics_company,omitempty" `

    /*
        子订单所在包裹的运单号     */
    InvoiceNo  *string `json:"invoice_no,omitempty" `

    /*
        捆绑的子订单号，表示该子订单要和捆绑的子订单一起发货，用于卖家子订单捆绑发货     */
    BindOid  *int64 `json:"bind_oid,omitempty" `

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
        天猫国际官网直供子订单关税税费     */
    SubOrderTaxFee  *string `json:"sub_order_tax_fee,omitempty" `

    /*
        天猫国际官网直供子订单关税税率     */
    SubOrderTaxRate  *string `json:"sub_order_tax_rate,omitempty" `

    /*
        征集预售订单征集状态：1（征集中），2（征集成功），3（征集失败）     */
    ZhengjiStatus  *string `json:"zhengji_status,omitempty" `

    /*
        免单资格     */
    MdQualification  *string `json:"md_qualification,omitempty" `

    /*
        免单金额     */
    MdFee  *string `json:"md_fee,omitempty" `

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

    /*
        是否为闲鱼订单 1是0否     */
    IsIdle  *string `json:"is_idle,omitempty" `

    /*
        承诺/最晚揽收时间     */
    PromiseCollectTime  *string `json:"promise_collect_time,omitempty" `

    /*
        是否自动流转到菜鸟仓发货     */
    IsForceDc  *bool `json:"is_force_dc,omitempty" `

}

func (s *TaobaoTradesSoldGetOrder) SetTitle(v string) *TaobaoTradesSoldGetOrder {
    s.Title = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetPicPath(v string) *TaobaoTradesSoldGetOrder {
    s.PicPath = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetPrice(v string) *TaobaoTradesSoldGetOrder {
    s.Price = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetNumIid(v int64) *TaobaoTradesSoldGetOrder {
    s.NumIid = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetSkuId(v string) *TaobaoTradesSoldGetOrder {
    s.SkuId = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetOuterIid(v string) *TaobaoTradesSoldGetOrder {
    s.OuterIid = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetOuterSkuId(v string) *TaobaoTradesSoldGetOrder {
    s.OuterSkuId = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetRefundStatus(v string) *TaobaoTradesSoldGetOrder {
    s.RefundStatus = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetStatus(v string) *TaobaoTradesSoldGetOrder {
    s.Status = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetOid(v int64) *TaobaoTradesSoldGetOrder {
    s.Oid = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetTotalFee(v string) *TaobaoTradesSoldGetOrder {
    s.TotalFee = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetPayment(v string) *TaobaoTradesSoldGetOrder {
    s.Payment = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetDiscountFee(v string) *TaobaoTradesSoldGetOrder {
    s.DiscountFee = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetAdjustFee(v string) *TaobaoTradesSoldGetOrder {
    s.AdjustFee = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetDivideOrderFee(v string) *TaobaoTradesSoldGetOrder {
    s.DivideOrderFee = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetPartMjzDiscount(v string) *TaobaoTradesSoldGetOrder {
    s.PartMjzDiscount = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetSkuPropertiesName(v string) *TaobaoTradesSoldGetOrder {
    s.SkuPropertiesName = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetItemMealId(v int64) *TaobaoTradesSoldGetOrder {
    s.ItemMealId = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetItemMealName(v string) *TaobaoTradesSoldGetOrder {
    s.ItemMealName = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetNum(v int64) *TaobaoTradesSoldGetOrder {
    s.Num = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetBuyerRate(v bool) *TaobaoTradesSoldGetOrder {
    s.BuyerRate = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetSellerRate(v bool) *TaobaoTradesSoldGetOrder {
    s.SellerRate = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetRefundId(v string) *TaobaoTradesSoldGetOrder {
    s.RefundId = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetSellerType(v string) *TaobaoTradesSoldGetOrder {
    s.SellerType = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetCid(v int64) *TaobaoTradesSoldGetOrder {
    s.Cid = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetEndTime(v util.LocalTime) *TaobaoTradesSoldGetOrder {
    s.EndTime = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetConsignTime(v string) *TaobaoTradesSoldGetOrder {
    s.ConsignTime = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetShippingType(v string) *TaobaoTradesSoldGetOrder {
    s.ShippingType = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetLogisticsCompany(v string) *TaobaoTradesSoldGetOrder {
    s.LogisticsCompany = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetInvoiceNo(v string) *TaobaoTradesSoldGetOrder {
    s.InvoiceNo = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetBindOid(v int64) *TaobaoTradesSoldGetOrder {
    s.BindOid = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetIsDaixiao(v bool) *TaobaoTradesSoldGetOrder {
    s.IsDaixiao = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetTicketOuterId(v string) *TaobaoTradesSoldGetOrder {
    s.TicketOuterId = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetTicketExpdateKey(v string) *TaobaoTradesSoldGetOrder {
    s.TicketExpdateKey = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetIsWww(v bool) *TaobaoTradesSoldGetOrder {
    s.IsWww = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetStoreCode(v string) *TaobaoTradesSoldGetOrder {
    s.StoreCode = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetTmserSpuCode(v string) *TaobaoTradesSoldGetOrder {
    s.TmserSpuCode = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetSubOrderTaxFee(v string) *TaobaoTradesSoldGetOrder {
    s.SubOrderTaxFee = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetSubOrderTaxRate(v string) *TaobaoTradesSoldGetOrder {
    s.SubOrderTaxRate = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetZhengjiStatus(v string) *TaobaoTradesSoldGetOrder {
    s.ZhengjiStatus = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetMdQualification(v string) *TaobaoTradesSoldGetOrder {
    s.MdQualification = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetMdFee(v string) *TaobaoTradesSoldGetOrder {
    s.MdFee = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetSubOrderTaxPromotionFee(v string) *TaobaoTradesSoldGetOrder {
    s.SubOrderTaxPromotionFee = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetTaxFree(v bool) *TaobaoTradesSoldGetOrder {
    s.TaxFree = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetTaxCouponDiscount(v string) *TaobaoTradesSoldGetOrder {
    s.TaxCouponDiscount = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetSpecialRefundType(v string) *TaobaoTradesSoldGetOrder {
    s.SpecialRefundType = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetIsDevalueFee(v bool) *TaobaoTradesSoldGetOrder {
    s.IsDevalueFee = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetIsIdle(v string) *TaobaoTradesSoldGetOrder {
    s.IsIdle = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetPromiseCollectTime(v string) *TaobaoTradesSoldGetOrder {
    s.PromiseCollectTime = &v
    return s
}
func (s *TaobaoTradesSoldGetOrder) SetIsForceDc(v bool) *TaobaoTradesSoldGetOrder {
    s.IsForceDc = &v
    return s
}
