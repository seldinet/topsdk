package domain

import (
        "topsdk/util"
    )

type TaobaoTradeFullinfoGetOrder struct {
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
        订单状态（请关注此状态，如果为TRADE_CLOSED_BY_TAOBAO状态，则不要对此订单进行发货，切记啊！）。可选值: <ul><li>TRADE_NO_CREATE_PAY(没有创建支付宝交易) </li><li>WAIT_BUYER_PAY(等待买家付款) </li><li>WAIT_SELLER_SEND_GOODS(等待卖家发货,即:买家已付款) </li><li>WAIT_BUYER_CONFIRM_GOODS(等待买家确认收货,即:卖家已发货) </li><li>TRADE_BUYER_SIGNED(买家已签收,货到付款专用) </li><li>TRADE_FINISHED(交易成功) </li><li>TRADE_CLOSED(付款以后用户退款成功，交易自动关闭) </li><li>TRADE_CLOSED_BY_TAOBAO(付款以前，卖家或买家主动关闭交易)</li><li>PAY_PENDING(国际信用卡支付付款确认中)</li></ul>     */
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
        订单快照URL     */
    SnapshotUrl  *string `json:"snapshot_url,omitempty" `

    /*
        订单超时到期时间。格式:yyyy-MM-dd HH:mm:ss     */
    TimeoutActionTime  *util.LocalTime `json:"timeout_action_time,omitempty" `

    /*
        商品备注     */
    ItemMemo  *string `json:"item_memo,omitempty" `

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
        是否超卖     */
    IsOversold  *bool `json:"is_oversold,omitempty" `

    /*
        子订单的交易结束时间说明：子订单有单独的结束时间，与主订单的结束时间可能有所不同，在有退款发起的时候或者是主订单分阶段付款的时候，子订单的结束时间会早于主订单的结束时间，所以开放这个字段便于订单结束状态的判断     */
    EndTime  *util.LocalTime `json:"end_time,omitempty" `

    /*
        子订单来源,如jhs(聚划算)、taobao(淘宝)、wap(无线)     */
    OrderFrom  *string `json:"order_from,omitempty" `

    /*
        是否是服务订单，是返回true，否返回false。     */
    IsServiceOrder  *bool `json:"is_service_order,omitempty" `

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
        子订单预计发货时间     */
    EstimateConTime  *string `json:"estimate_con_time,omitempty" `

    /*
        bind_oid字段的升级，支持返回绑定的多个子订单，多个子订单以半角逗号分隔     */
    BindOids  *string `json:"bind_oids,omitempty" `

    /*
        征集预售订单征集状态：1（征集中），2（征集成功），3（征集失败）     */
    ZhengjiStatus  *string `json:"zhengji_status,omitempty" `

    /*
        免单资格属性     */
    MdQualification  *string `json:"md_qualification,omitempty" `

    /*
        免单金额     */
    MdFee  *string `json:"md_fee,omitempty" `

    /*
        定制信息     */
    Customization  *string `json:"customization,omitempty" `

    /*
        库存类型：6为在途     */
    InvType  *string `json:"inv_type,omitempty" `

    /*
        是否发货     */
    IsShShip  *bool `json:"is_sh_ship,omitempty" `

    /*
        仓储信息     */
    Shipper  *string `json:"shipper,omitempty" `

    /*
        订单履行类型，如喵鲜生极速达（jsd）     */
    FType  *string `json:"f_type,omitempty" `

    /*
        订单履行状态，如喵鲜生极速达：分单完成     */
    FStatus  *string `json:"f_status,omitempty" `

    /*
        单履行内容，如喵鲜生极速达：storeId,phone     */
    FTerm  *string `json:"f_term,omitempty" `

    /*
        天猫搭配宝     */
    ComboId  *string `json:"combo_id,omitempty" `

    /*
        子订单扩展属性:<br/> is_free_down_payment:是否免首付：true：是，false：否，可选字段<br/> car_back_payment:返还免首付金额，单位：分，<br/> car_ref_activity_id:服务商传入活动ID，依赖外部服务商传入；<br/>     */
    OrderAttr  *string `json:"order_attr,omitempty" `

    /*
        assemblyRela     */
    AssemblyRela  *string `json:"assembly_rela,omitempty" `

    /*
        assemblyPrice     */
    AssemblyPrice  *string `json:"assembly_price,omitempty" `

    /*
        assemblyItem     */
    AssemblyItem  *string `json:"assembly_item,omitempty" `

    /*
        天猫国际子订单计税优惠金额     */
    SubOrderTaxPromotionFee  *string `json:"sub_order_tax_promotion_fee,omitempty" `

    /*
        downPayment     */
    DownPayment  *string `json:"down_payment,omitempty" `

    /*
        downPaymentRatio     */
    DownPaymentRatio  *string `json:"down_payment_ratio,omitempty" `

    /*
        monthPayment     */
    MonthPayment  *string `json:"month_payment,omitempty" `

    /*
        tailPayment     */
    TailPayment  *string `json:"tail_payment,omitempty" `

    /*
        installmentNum     */
    InstallmentNum  *string `json:"installment_num,omitempty" `

    /*
        penalty     */
    Penalty  *string `json:"penalty,omitempty" `

    /*
        serviceFee     */
    ServiceFee  *string `json:"service_fee,omitempty" `

    /*
        carTaker     */
    CarTaker  *string `json:"car_taker,omitempty" `

    /*
        carTakerPhone     */
    CarTakerPhone  *string `json:"car_taker_phone,omitempty" `

    /*
        carTakerID     */
    CarTakerId  *string `json:"car_taker_id,omitempty" `

    /*
        carStoreCode     */
    CarStoreCode  *string `json:"car_store_code,omitempty" `

    /*
        carStoreName     */
    CarStoreName  *string `json:"car_store_name,omitempty" `

    /*
        outUniqueId     */
    OutUniqueId  *string `json:"out_unique_id,omitempty" `

    /*
        wsBankApplyNo     */
    WsBankApplyNo  *string `json:"ws_bank_apply_no,omitempty" `

    /*
        oidStr     */
    OidStr  *string `json:"oid_str,omitempty" `

    /*
        花呗分期期数     */
    FqgNum  *int64 `json:"fqg_num,omitempty" `

    /*
        是否商家承担手续费     */
    IsFqgSFee  *bool `json:"is_fqg_s_fee,omitempty" `

    /*
        天猫国际订单是否包税     */
    TaxFree  *bool `json:"tax_free,omitempty" `

    /*
        天猫国际订单包税金额     */
    TaxCouponDiscount  *string `json:"tax_coupon_discount,omitempty" `

    /*
        nrReduceInvFail     */
    NrReduceInvFail  *string `json:"nr_reduce_inv_fail,omitempty" `

    /*
        新零售商家端商品唯一编号     */
    NrOuterIid  *string `json:"nr_outer_iid,omitempty" `

    /*
        bind_oids字段的升级，在交易成功和交易关闭状态下也能获取到，支持返回绑定的多个子订单，多个子订单以半角逗号分隔     */
    BindOidsAllStatus  *string `json:"bind_oids_all_status,omitempty" `

    /*
        sortInfo     */
    SortInfo  *TaobaoTradeFullinfoGetSortInfo `json:"sort_info,omitempty" `

    /*
        导购员Id     */
    O2oGuideId  *string `json:"o2o_guide_id,omitempty" `

    /*
        导购员名称     */
    O2oGuideName  *string `json:"o2o_guide_name,omitempty" `

    /*
        门店Id     */
    O2oShopId  *string `json:"o2o_shop_id,omitempty" `

    /*
        门店名称     */
    O2oShopName  *string `json:"o2o_shop_name,omitempty" `

    /*
        云店交易链路，为tmall.daogoubao.cloudstore时表示云店链路     */
    BizCode  *string `json:"biz_code,omitempty" `

    /*
        云店标记为1，且bizCode不为tmall.daogoubao.cloudstore时，为旗舰店订单     */
    CloudStore  *string `json:"cloud_store,omitempty" `

    /*
        云店是否扣拥     */
    HjSettleNoCommission  *string `json:"hj_settle_no_commission,omitempty" `

    /*
        云店接单标记     */
    OrderTaking  *string `json:"order_taking,omitempty" `

    /*
        云店改价用token     */
    CloudStoreToken  *string `json:"cloud_store_token,omitempty" `

    /*
        云店pos单号     */
    CloudStoreBindPos  *string `json:"cloud_store_bind_pos,omitempty" `

    /*
        天猫未来店线下店铺 ID     */
    RetailStoreId  *string `json:"retail_store_id,omitempty" `

    /*
        天猫未来店外部 ERP 商品 ID     */
    OutItemId  *string `json:"out_item_id,omitempty" `

    /*
        新零售全渠道订单：商家自有货品编码     */
    RtOmniOuterScId  *string `json:"rt_omni_outer_sc_id,omitempty" `

    /*
        新零售全渠道订单：后端货品ID     */
    RtOmniScId  *string `json:"rt_omni_sc_id,omitempty" `

    /*
        有值表示买家修改了地址；1表示付款后改地址；2表示付款前改地址     */
    ModifyAddress  *string `json:"modify_address,omitempty" `

    /*
        买家修改地址时间     */
    TiModifyAddressTime  *string `json:"ti_modify_address_time,omitempty" `

    /*
        有值表示信用购订单；1表示信用购一期；2表示信用购二期；3表示信用购三期     */
    CreditBuy  *string `json:"credit_buy,omitempty" `

    /*
        子订单商家代缴税费     */
    STariffFee  *string `json:"s_tariff_fee,omitempty" `

    /*
        时效服务身份，如tmallPromise代表天猫时效承诺     */
    TimingPromise  *string `json:"timing_promise,omitempty" `

    /*
        时效服务字段，服务字段，会有多个服务值，以英文半角逗号","切割     */
    PromiseService  *string `json:"promise_service,omitempty" `

    /*
        预计送达时间     */
    EsDate  *string `json:"es_date,omitempty" `

    /*
        预计配送时间段     */
    EsRange  *string `json:"es_range,omitempty" `

    /*
        预约配送，用户预约时间     */
    OsDate  *string `json:"os_date,omitempty" `

    /*
        预约配送，用户预约时间段     */
    OsRange  *string `json:"os_range,omitempty" `

    /*
        物流截单时间，分钟     */
    CutoffMinutes  *string `json:"cutoff_minutes,omitempty" `

    /*
        物流时效，相对时间，单位是天     */
    EsTime  *string `json:"es_time,omitempty" `

    /*
        最晚发货时间     */
    DeliveryTime  *string `json:"delivery_time,omitempty" `

    /*
        最晚揽收时间     */
    CollectTime  *string `json:"collect_time,omitempty" `

    /*
        最晚派送时间     */
    DispatchTime  *string `json:"dispatch_time,omitempty" `

    /*
        最晚签收时间     */
    SignTime  *string `json:"sign_time,omitempty" `

    /*
        信用购履约结束时间     */
    PromiseEndTime  *string `json:"promise_end_time,omitempty" `

    /*
        前N有礼活动id     */
    OsActivityId  *string `json:"os_activity_id,omitempty" `

    /*
        前N有礼奖品的商品id     */
    OsFgItemId  *string `json:"os_fg_item_id,omitempty" `

    /*
        前N有礼获得奖品的数量     */
    OsGiftCount  *string `json:"os_gift_count,omitempty" `

    /*
        前N有礼中奖名次，获得奖品的订单才会有该字段     */
    OsSortNum  *string `json:"os_sort_num,omitempty" `

    /*
        使用淘金币的数量，以分为单位，和订单标propoint中间那一段一样     */
    Propoint  *string `json:"propoint,omitempty" `

    /*
        是否是考拉商品订单     */
    IsKaola  *bool `json:"is_kaola,omitempty" `

    /*
        特殊的退款类型，比如直播返现、价保     */
    SpecialRefundType  *string `json:"special_refund_type,omitempty" `

    /*
        透出的额外信息     */
    ExtendInfo  *string `json:"extend_info,omitempty" `

    /*
        子订单优惠贬值     */
    IsDevalueFee  *bool `json:"is_devalue_fee,omitempty" `

    /*
        同城购渠道店id     */
    BrandLightShopStoreId  *string `json:"brand_light_shop_store_id,omitempty" `

    /*
        同城购订单来源     */
    BrandLightShopSource  *string `json:"brand_light_shop_source,omitempty" `

    /*
        服务供应链-服务订单类型,1:主子挂载；2：双主挂载；3：单独售卖     */
    ServiceOrderType  *string `json:"service_order_type,omitempty" `

    /*
        服务供应链-服务商外部编码     */
    ServiceOuterId  *string `json:"service_outer_id,omitempty" `

    /*
        购物金核销子订单权益金分摊金额（单位为元）     */
    ExpandCardExpandPriceUsedSuborder  *string `json:"expand_card_expand_price_used_suborder,omitempty" `

    /*
        购物金核销子订单本金分摊金额（单位为元）     */
    ExpandCardBasicPriceUsedSuborder  *string `json:"expand_card_basic_price_used_suborder,omitempty" `

    /*
        预售订单立减金额     */
    Lijian  *string `json:"lijian,omitempty" `

    /*
        优仓业务场景下 1（自动流转）/0（非自动流转）     */
    AutoFlow  *string `json:"auto_flow,omitempty" `

    /*
        订单履约类型     */
    TradeFulfillmentType  *int64 `json:"trade_fulfillment_type,omitempty" `

    /*
        订单是有代发订单，为空表示该订单暂无代发单据，distribute-该子订单有已分配代发单据，cancel-订单的代发单据都已取消，     */
    DistributeStatus  *string `json:"distribute_status,omitempty" `

    /*
        是否为闲鱼订单，1是，0否     */
    IsIdle  *string `json:"is_idle,omitempty" `

    /*
        赠品关联的id，主品订单下此字段表示赠品订单id，赠品订单表示主品订单id     */
    GiftMids  *string `json:"gift_mids,omitempty" `

    /*
        是否是赠品订单     */
    IsFreeGift  *bool `json:"is_free_gift,omitempty" `

    /*
        是否含有赠品     */
    HasGift  *bool `json:"has_gift,omitempty" `

    /*
        承诺/最晚揽收时间     */
    PromiseCollectTime  *string `json:"promise_collect_time,omitempty" `

    /*
        表示该子订单会ascp会自动流转到菜鸟仓发货     */
    IsForceDc  *bool `json:"is_force_dc,omitempty" `

}

func (s *TaobaoTradeFullinfoGetOrder) SetTitle(v string) *TaobaoTradeFullinfoGetOrder {
    s.Title = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetPicPath(v string) *TaobaoTradeFullinfoGetOrder {
    s.PicPath = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetPrice(v string) *TaobaoTradeFullinfoGetOrder {
    s.Price = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetNumIid(v int64) *TaobaoTradeFullinfoGetOrder {
    s.NumIid = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetSkuId(v string) *TaobaoTradeFullinfoGetOrder {
    s.SkuId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetOuterIid(v string) *TaobaoTradeFullinfoGetOrder {
    s.OuterIid = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetOuterSkuId(v string) *TaobaoTradeFullinfoGetOrder {
    s.OuterSkuId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetRefundStatus(v string) *TaobaoTradeFullinfoGetOrder {
    s.RefundStatus = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetStatus(v string) *TaobaoTradeFullinfoGetOrder {
    s.Status = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetOid(v int64) *TaobaoTradeFullinfoGetOrder {
    s.Oid = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetTotalFee(v string) *TaobaoTradeFullinfoGetOrder {
    s.TotalFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetPayment(v string) *TaobaoTradeFullinfoGetOrder {
    s.Payment = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetDiscountFee(v string) *TaobaoTradeFullinfoGetOrder {
    s.DiscountFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetAdjustFee(v string) *TaobaoTradeFullinfoGetOrder {
    s.AdjustFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetDivideOrderFee(v string) *TaobaoTradeFullinfoGetOrder {
    s.DivideOrderFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetPartMjzDiscount(v string) *TaobaoTradeFullinfoGetOrder {
    s.PartMjzDiscount = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetSkuPropertiesName(v string) *TaobaoTradeFullinfoGetOrder {
    s.SkuPropertiesName = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetItemMealId(v int64) *TaobaoTradeFullinfoGetOrder {
    s.ItemMealId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetItemMealName(v string) *TaobaoTradeFullinfoGetOrder {
    s.ItemMealName = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetNum(v int64) *TaobaoTradeFullinfoGetOrder {
    s.Num = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetSnapshotUrl(v string) *TaobaoTradeFullinfoGetOrder {
    s.SnapshotUrl = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetTimeoutActionTime(v util.LocalTime) *TaobaoTradeFullinfoGetOrder {
    s.TimeoutActionTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetItemMemo(v string) *TaobaoTradeFullinfoGetOrder {
    s.ItemMemo = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetBuyerRate(v bool) *TaobaoTradeFullinfoGetOrder {
    s.BuyerRate = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetSellerRate(v bool) *TaobaoTradeFullinfoGetOrder {
    s.SellerRate = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetRefundId(v string) *TaobaoTradeFullinfoGetOrder {
    s.RefundId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetSellerType(v string) *TaobaoTradeFullinfoGetOrder {
    s.SellerType = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetCid(v int64) *TaobaoTradeFullinfoGetOrder {
    s.Cid = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetIsOversold(v bool) *TaobaoTradeFullinfoGetOrder {
    s.IsOversold = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetEndTime(v util.LocalTime) *TaobaoTradeFullinfoGetOrder {
    s.EndTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetOrderFrom(v string) *TaobaoTradeFullinfoGetOrder {
    s.OrderFrom = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetIsServiceOrder(v bool) *TaobaoTradeFullinfoGetOrder {
    s.IsServiceOrder = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetConsignTime(v string) *TaobaoTradeFullinfoGetOrder {
    s.ConsignTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetShippingType(v string) *TaobaoTradeFullinfoGetOrder {
    s.ShippingType = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetLogisticsCompany(v string) *TaobaoTradeFullinfoGetOrder {
    s.LogisticsCompany = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetInvoiceNo(v string) *TaobaoTradeFullinfoGetOrder {
    s.InvoiceNo = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetBindOid(v int64) *TaobaoTradeFullinfoGetOrder {
    s.BindOid = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetIsDaixiao(v bool) *TaobaoTradeFullinfoGetOrder {
    s.IsDaixiao = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetTicketOuterId(v string) *TaobaoTradeFullinfoGetOrder {
    s.TicketOuterId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetTicketExpdateKey(v string) *TaobaoTradeFullinfoGetOrder {
    s.TicketExpdateKey = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetIsWww(v bool) *TaobaoTradeFullinfoGetOrder {
    s.IsWww = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetStoreCode(v string) *TaobaoTradeFullinfoGetOrder {
    s.StoreCode = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetTmserSpuCode(v string) *TaobaoTradeFullinfoGetOrder {
    s.TmserSpuCode = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetSubOrderTaxFee(v string) *TaobaoTradeFullinfoGetOrder {
    s.SubOrderTaxFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetSubOrderTaxRate(v string) *TaobaoTradeFullinfoGetOrder {
    s.SubOrderTaxRate = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetEtSerTime(v string) *TaobaoTradeFullinfoGetOrder {
    s.EtSerTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetEtShopName(v string) *TaobaoTradeFullinfoGetOrder {
    s.EtShopName = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetEtVerifiedShopName(v string) *TaobaoTradeFullinfoGetOrder {
    s.EtVerifiedShopName = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetEtPlateNumber(v string) *TaobaoTradeFullinfoGetOrder {
    s.EtPlateNumber = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetEstimateConTime(v string) *TaobaoTradeFullinfoGetOrder {
    s.EstimateConTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetBindOids(v string) *TaobaoTradeFullinfoGetOrder {
    s.BindOids = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetZhengjiStatus(v string) *TaobaoTradeFullinfoGetOrder {
    s.ZhengjiStatus = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetMdQualification(v string) *TaobaoTradeFullinfoGetOrder {
    s.MdQualification = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetMdFee(v string) *TaobaoTradeFullinfoGetOrder {
    s.MdFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetCustomization(v string) *TaobaoTradeFullinfoGetOrder {
    s.Customization = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetInvType(v string) *TaobaoTradeFullinfoGetOrder {
    s.InvType = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetIsShShip(v bool) *TaobaoTradeFullinfoGetOrder {
    s.IsShShip = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetShipper(v string) *TaobaoTradeFullinfoGetOrder {
    s.Shipper = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetFType(v string) *TaobaoTradeFullinfoGetOrder {
    s.FType = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetFStatus(v string) *TaobaoTradeFullinfoGetOrder {
    s.FStatus = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetFTerm(v string) *TaobaoTradeFullinfoGetOrder {
    s.FTerm = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetComboId(v string) *TaobaoTradeFullinfoGetOrder {
    s.ComboId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetOrderAttr(v string) *TaobaoTradeFullinfoGetOrder {
    s.OrderAttr = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetAssemblyRela(v string) *TaobaoTradeFullinfoGetOrder {
    s.AssemblyRela = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetAssemblyPrice(v string) *TaobaoTradeFullinfoGetOrder {
    s.AssemblyPrice = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetAssemblyItem(v string) *TaobaoTradeFullinfoGetOrder {
    s.AssemblyItem = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetSubOrderTaxPromotionFee(v string) *TaobaoTradeFullinfoGetOrder {
    s.SubOrderTaxPromotionFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetDownPayment(v string) *TaobaoTradeFullinfoGetOrder {
    s.DownPayment = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetDownPaymentRatio(v string) *TaobaoTradeFullinfoGetOrder {
    s.DownPaymentRatio = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetMonthPayment(v string) *TaobaoTradeFullinfoGetOrder {
    s.MonthPayment = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetTailPayment(v string) *TaobaoTradeFullinfoGetOrder {
    s.TailPayment = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetInstallmentNum(v string) *TaobaoTradeFullinfoGetOrder {
    s.InstallmentNum = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetPenalty(v string) *TaobaoTradeFullinfoGetOrder {
    s.Penalty = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetServiceFee(v string) *TaobaoTradeFullinfoGetOrder {
    s.ServiceFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetCarTaker(v string) *TaobaoTradeFullinfoGetOrder {
    s.CarTaker = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetCarTakerPhone(v string) *TaobaoTradeFullinfoGetOrder {
    s.CarTakerPhone = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetCarTakerId(v string) *TaobaoTradeFullinfoGetOrder {
    s.CarTakerId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetCarStoreCode(v string) *TaobaoTradeFullinfoGetOrder {
    s.CarStoreCode = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetCarStoreName(v string) *TaobaoTradeFullinfoGetOrder {
    s.CarStoreName = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetOutUniqueId(v string) *TaobaoTradeFullinfoGetOrder {
    s.OutUniqueId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetWsBankApplyNo(v string) *TaobaoTradeFullinfoGetOrder {
    s.WsBankApplyNo = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetOidStr(v string) *TaobaoTradeFullinfoGetOrder {
    s.OidStr = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetFqgNum(v int64) *TaobaoTradeFullinfoGetOrder {
    s.FqgNum = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetIsFqgSFee(v bool) *TaobaoTradeFullinfoGetOrder {
    s.IsFqgSFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetTaxFree(v bool) *TaobaoTradeFullinfoGetOrder {
    s.TaxFree = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetTaxCouponDiscount(v string) *TaobaoTradeFullinfoGetOrder {
    s.TaxCouponDiscount = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetNrReduceInvFail(v string) *TaobaoTradeFullinfoGetOrder {
    s.NrReduceInvFail = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetNrOuterIid(v string) *TaobaoTradeFullinfoGetOrder {
    s.NrOuterIid = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetBindOidsAllStatus(v string) *TaobaoTradeFullinfoGetOrder {
    s.BindOidsAllStatus = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetSortInfo(v TaobaoTradeFullinfoGetSortInfo) *TaobaoTradeFullinfoGetOrder {
    s.SortInfo = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetO2oGuideId(v string) *TaobaoTradeFullinfoGetOrder {
    s.O2oGuideId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetO2oGuideName(v string) *TaobaoTradeFullinfoGetOrder {
    s.O2oGuideName = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetO2oShopId(v string) *TaobaoTradeFullinfoGetOrder {
    s.O2oShopId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetO2oShopName(v string) *TaobaoTradeFullinfoGetOrder {
    s.O2oShopName = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetBizCode(v string) *TaobaoTradeFullinfoGetOrder {
    s.BizCode = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetCloudStore(v string) *TaobaoTradeFullinfoGetOrder {
    s.CloudStore = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetHjSettleNoCommission(v string) *TaobaoTradeFullinfoGetOrder {
    s.HjSettleNoCommission = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetOrderTaking(v string) *TaobaoTradeFullinfoGetOrder {
    s.OrderTaking = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetCloudStoreToken(v string) *TaobaoTradeFullinfoGetOrder {
    s.CloudStoreToken = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetCloudStoreBindPos(v string) *TaobaoTradeFullinfoGetOrder {
    s.CloudStoreBindPos = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetRetailStoreId(v string) *TaobaoTradeFullinfoGetOrder {
    s.RetailStoreId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetOutItemId(v string) *TaobaoTradeFullinfoGetOrder {
    s.OutItemId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetRtOmniOuterScId(v string) *TaobaoTradeFullinfoGetOrder {
    s.RtOmniOuterScId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetRtOmniScId(v string) *TaobaoTradeFullinfoGetOrder {
    s.RtOmniScId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetModifyAddress(v string) *TaobaoTradeFullinfoGetOrder {
    s.ModifyAddress = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetTiModifyAddressTime(v string) *TaobaoTradeFullinfoGetOrder {
    s.TiModifyAddressTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetCreditBuy(v string) *TaobaoTradeFullinfoGetOrder {
    s.CreditBuy = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetSTariffFee(v string) *TaobaoTradeFullinfoGetOrder {
    s.STariffFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetTimingPromise(v string) *TaobaoTradeFullinfoGetOrder {
    s.TimingPromise = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetPromiseService(v string) *TaobaoTradeFullinfoGetOrder {
    s.PromiseService = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetEsDate(v string) *TaobaoTradeFullinfoGetOrder {
    s.EsDate = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetEsRange(v string) *TaobaoTradeFullinfoGetOrder {
    s.EsRange = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetOsDate(v string) *TaobaoTradeFullinfoGetOrder {
    s.OsDate = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetOsRange(v string) *TaobaoTradeFullinfoGetOrder {
    s.OsRange = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetCutoffMinutes(v string) *TaobaoTradeFullinfoGetOrder {
    s.CutoffMinutes = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetEsTime(v string) *TaobaoTradeFullinfoGetOrder {
    s.EsTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetDeliveryTime(v string) *TaobaoTradeFullinfoGetOrder {
    s.DeliveryTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetCollectTime(v string) *TaobaoTradeFullinfoGetOrder {
    s.CollectTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetDispatchTime(v string) *TaobaoTradeFullinfoGetOrder {
    s.DispatchTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetSignTime(v string) *TaobaoTradeFullinfoGetOrder {
    s.SignTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetPromiseEndTime(v string) *TaobaoTradeFullinfoGetOrder {
    s.PromiseEndTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetOsActivityId(v string) *TaobaoTradeFullinfoGetOrder {
    s.OsActivityId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetOsFgItemId(v string) *TaobaoTradeFullinfoGetOrder {
    s.OsFgItemId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetOsGiftCount(v string) *TaobaoTradeFullinfoGetOrder {
    s.OsGiftCount = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetOsSortNum(v string) *TaobaoTradeFullinfoGetOrder {
    s.OsSortNum = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetPropoint(v string) *TaobaoTradeFullinfoGetOrder {
    s.Propoint = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetIsKaola(v bool) *TaobaoTradeFullinfoGetOrder {
    s.IsKaola = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetSpecialRefundType(v string) *TaobaoTradeFullinfoGetOrder {
    s.SpecialRefundType = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetExtendInfo(v string) *TaobaoTradeFullinfoGetOrder {
    s.ExtendInfo = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetIsDevalueFee(v bool) *TaobaoTradeFullinfoGetOrder {
    s.IsDevalueFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetBrandLightShopStoreId(v string) *TaobaoTradeFullinfoGetOrder {
    s.BrandLightShopStoreId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetBrandLightShopSource(v string) *TaobaoTradeFullinfoGetOrder {
    s.BrandLightShopSource = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetServiceOrderType(v string) *TaobaoTradeFullinfoGetOrder {
    s.ServiceOrderType = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetServiceOuterId(v string) *TaobaoTradeFullinfoGetOrder {
    s.ServiceOuterId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetExpandCardExpandPriceUsedSuborder(v string) *TaobaoTradeFullinfoGetOrder {
    s.ExpandCardExpandPriceUsedSuborder = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetExpandCardBasicPriceUsedSuborder(v string) *TaobaoTradeFullinfoGetOrder {
    s.ExpandCardBasicPriceUsedSuborder = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetLijian(v string) *TaobaoTradeFullinfoGetOrder {
    s.Lijian = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetAutoFlow(v string) *TaobaoTradeFullinfoGetOrder {
    s.AutoFlow = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetTradeFulfillmentType(v int64) *TaobaoTradeFullinfoGetOrder {
    s.TradeFulfillmentType = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetDistributeStatus(v string) *TaobaoTradeFullinfoGetOrder {
    s.DistributeStatus = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetIsIdle(v string) *TaobaoTradeFullinfoGetOrder {
    s.IsIdle = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetGiftMids(v string) *TaobaoTradeFullinfoGetOrder {
    s.GiftMids = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetIsFreeGift(v bool) *TaobaoTradeFullinfoGetOrder {
    s.IsFreeGift = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetHasGift(v bool) *TaobaoTradeFullinfoGetOrder {
    s.HasGift = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetPromiseCollectTime(v string) *TaobaoTradeFullinfoGetOrder {
    s.PromiseCollectTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetOrder) SetIsForceDc(v bool) *TaobaoTradeFullinfoGetOrder {
    s.IsForceDc = &v
    return s
}
