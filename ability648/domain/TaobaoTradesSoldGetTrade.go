package domain

import (
        "topsdk/util"
    )

type TaobaoTradesSoldGetTrade struct {
    /*
        卖家昵称     */
    SellerNick  *string `json:"seller_nick,omitempty" `

    /*
        买家昵称     */
    BuyerNick  *string `json:"buyer_nick,omitempty" `

    /*
        买家OpenUid     */
    BuyerOpenUid  *string `json:"buyer_open_uid,omitempty" `

    /*
        交易标题，以店铺名作为此标题的值。注:taobao.trades.get接口返回的Trade中的title是商品名称     */
    Title  *string `json:"title,omitempty" `

    /*
        交易类型列表，同时查询多种交易类型可用逗号分隔。默认同时查询guarantee_trade, auto_delivery, ec, cod的4种交易类型的数据 可选值 fixed(一口价) auction(拍卖) guarantee_trade(一口价、拍卖) auto_delivery(自动发货) independent_simple_trade(旺店入门版交易) independent_shop_trade(旺店标准版交易) ec(直冲) cod(货到付款) fenxiao(分销) game_equipment(游戏装备) shopex_trade(ShopEX交易) netcn_trade(万网交易) external_trade(统一外部交易)o2o_offlinetrade（O2O交易）step (万人团)nopaid(无付款订单)pre_auth_type(预授权0元购机交易)     */
    Type  *string `json:"type,omitempty" `

    /*
        交易创建时间。格式:yyyy-MM-dd HH:mm:ss     */
    Created  *util.LocalTime `json:"created,omitempty" `

    /*
        交易编号 (父订单的交易编号)     */
    Tid  *int64 `json:"tid,omitempty" `

    /*
        Acookie订单唯一ID。     */
    AcookieId  *string `json:"acookie_id,omitempty" `

    /*
        物流运单号     */
    Sid  *string `json:"sid,omitempty" `

    /*
        卖家是否已评价。可选值:true(已评价),false(未评价)     */
    SellerRate  *bool `json:"seller_rate,omitempty" `

    /*
        买家是否已评价。可选值:true(已评价),false(未评价)。如买家只评价未打分，此字段仍返回false     */
    BuyerRate  *bool `json:"buyer_rate,omitempty" `

    /*
        交易状态。可选值:    * TRADE_NO_CREATE_PAY(没有创建支付宝交易)    * WAIT_BUYER_PAY(等待买家付款)    * SELLER_CONSIGNED_PART(卖家部分发货)    * WAIT_SELLER_SEND_GOODS(等待卖家发货,即:买家已付款)    * WAIT_BUYER_CONFIRM_GOODS(等待买家确认收货,即:卖家已发货)    * TRADE_BUYER_SIGNED(买家已签收,货到付款专用)    * TRADE_FINISHED(交易成功)    * TRADE_CLOSED(付款以后用户退款成功，交易自动关闭)    * TRADE_CLOSED_BY_TAOBAO(付款以前，卖家或买家主动关闭交易)    * PAY_PENDING(国际信用卡支付付款确认中)    * WAIT_PRE_AUTH_CONFIRM(0元购合约中)	* PAID_FORBID_CONSIGN(拼团中订单或者发货强管控的订单，已付款但禁止发货)     */
    Status  *string `json:"status,omitempty" `

    /*
        实付金额。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
    Payment  *string `json:"payment,omitempty" `

    /*
        建议使用trade.promotion_details查询系统优惠系统优惠金额（如打折，VIP，满就送等），精确到2位小数，单位：元。如：200.07，表示：200元7分     */
    DiscountFee  *string `json:"discount_fee,omitempty" `

    /*
        卖家手工调整金额，精确到2位小数，单位：元。如：200.07，表示：200元7分。来源于订单价格修改，如果有多笔子订单的时候，这个为0，单笔的话则跟[order].adjust_fee一样     */
    AdjustFee  *string `json:"adjust_fee,omitempty" `

    /*
        邮费。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
    PostFee  *string `json:"post_fee,omitempty" `

    /*
        商品金额（商品价格乘以数量的总金额）。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
    TotalFee  *string `json:"total_fee,omitempty" `

    /*
        付款时间。格式:yyyy-MM-dd HH:mm:ss。订单的付款时间即为物流订单的创建时间。     */
    PayTime  *util.LocalTime `json:"pay_time,omitempty" `

    /*
        交易结束时间。交易成功时间(更新交易状态为成功的同时更新)/确认收货时间或者交易关闭时间 。格式:yyyy-MM-dd HH:mm:ss     */
    EndTime  *util.LocalTime `json:"end_time,omitempty" `

    /*
        交易修改时间(用户对订单的任何修改都会更新此字段)。格式:yyyy-MM-dd HH:mm:ss     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

    /*
        同步到卖家库的时间，taobao.trades.sold.incrementv.get接口返回此字段     */
    AsyncModified  *util.LocalTime `json:"async_modified,omitempty" `

    /*
        卖家发货时间。格式:yyyy-MM-dd HH:mm:ss     */
    ConsignTime  *util.LocalTime `json:"consign_time,omitempty" `

    /*
        卖家实际收到的支付宝打款金额（由于子订单可以部分确认收货，这个金额会随着子订单的确认收货而不断增加，交易成功后等于买家实付款减去退款金额）。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
    ReceivedPayment  *string `json:"received_payment,omitempty" `

    /*
        交易佣金。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
    CommissionFee  *string `json:"commission_fee,omitempty" `

    /*
        支付宝交易号，如：2009112081173831     */
    AlipayNo  *string `json:"alipay_no,omitempty" `

    /*
        买家的支付宝id号，在UIC中有记录，买家支付宝的唯一标示，不因为买家更换Email账号而改变。     */
    AlipayId  *int64 `json:"alipay_id,omitempty" `

    /*
        买家下单的地区     */
    BuyerArea  *string `json:"buyer_area,omitempty" `

    /*
        卡易售垂直表信息，去除下单ip之后的结果     */
    NutFeature  *string `json:"nut_feature,omitempty" `

    /*
        商品图片绝对途径     */
    PicPath  *string `json:"pic_path,omitempty" `

    /*
        商品数字编号     */
    NumIid  *int64 `json:"num_iid,omitempty" `

    /*
        商品价格。精确到2位小数；单位：元。如：200.07，表示：200元7分     */
    Price  *string `json:"price,omitempty" `

    /*
        货到付款服务费。精确到2位小数;单位:元。如:12.07，表示:12元7分。     */
    CodFee  *string `json:"cod_fee,omitempty" `

    /*
        货到付款物流状态。初始状态 NEW_CREATED,接单成功 ACCEPTED_BY_COMPANY,接单失败 REJECTED_BY_COMPANY,接单超时 RECIEVE_TIMEOUT,揽收成功 TAKEN_IN_SUCCESS,揽收失败 TAKEN_IN_FAILED,揽收超时 TAKEN_TIMEOUT,签收成功 SIGN_IN,签收失败 REJECTED_BY_OTHER_SIDE,订单等待发送给物流公司 WAITING_TO_BE_SENT,用户取消物流订单 CANCELED     */
    CodStatus  *string `json:"cod_status,omitempty" `

    /*
        创建交易时的物流方式（交易完成前，物流方式有可能改变，但系统里的这个字段一直不变）。可选值：free(卖家包邮),post(平邮),express(快递),ems(EMS),virtual(虚拟发货)，25(次日必达)，26(预约配送)。     */
    ShippingType  *string `json:"shipping_type,omitempty" `

    /*
        商品购买数量。取值范围：大于零的整数,对于一个trade对应多个order的时候（一笔主订单，对应多笔子订单），num=0，num是一个跟商品关联的属性，一笔订单对应多比子订单的时候，主订单上的num无意义。     */
    Num  *int64 `json:"num,omitempty" `

    /*
        买家使用积分,下单时生成，且一直不变。格式:100;单位:个.     */
    PointFee  *int64 `json:"point_fee,omitempty" `

    /*
        买家实际使用积分（扣除部分退款使用的积分），交易完成后生成（交易成功或关闭），交易未完成时该字段值为0。格式:100;单位:个     */
    RealPointFee  *int64 `json:"real_point_fee,omitempty" `

    /*
        买家获得积分,返点的积分。格式:100;单位:个。返点的积分要交易成功之后才能获得。     */
    BuyerObtainPointFee  *int64 `json:"buyer_obtain_point_fee,omitempty" `

    /*
        收货人的姓名     */
    ReceiverName  *string `json:"receiver_name,omitempty" `

    /*
        收货人的所在省份     */
    ReceiverState  *string `json:"receiver_state,omitempty" `

    /*
        收货人的所在城市<br>注：因为国家对于城市和地区的划分的有：省直辖市和省直辖县级行政区（区级别的）划分的，淘宝这边根据这个差异保存在不同字段里面比如：广东广州：广州属于一个直辖市是放在的receiver_city的字段里面；而河南济源：济源属于省直辖县级行政区划分，是区级别的，放在了receiver_district里面<br>建议：程序依赖于城市字段做物流等判断的操作，最好加一个判断逻辑：如果返回值里面只有receiver_district参数，该参数作为城市     */
    ReceiverCity  *string `json:"receiver_city,omitempty" `

    /*
        收货人的所在地区<br>注：因为国家对于城市和地区的划分的有：省直辖市和省直辖县级行政区（区级别的）划分的，淘宝这边根据这个差异保存在不同字段里面比如：广东广州：广州属于一个直辖市是放在的receiver_city的字段里面；而河南济源：济源属于省直辖县级行政区划分，是区级别的，放在了receiver_district里面<br>建议：程序依赖于城市字段做物流等判断的操作，最好加一个判断逻辑：如果返回值里面只有receiver_district参数，该参数作为城市     */
    ReceiverDistrict  *string `json:"receiver_district,omitempty" `

    /*
        收货人街道地址     */
    ReceiverTown  *string `json:"receiver_town,omitempty" `

    /*
        收货人的详细地址     */
    ReceiverAddress  *string `json:"receiver_address,omitempty" `

    /*
        收货人的邮编     */
    ReceiverZip  *string `json:"receiver_zip,omitempty" `

    /*
        收货人的手机号码     */
    ReceiverMobile  *string `json:"receiver_mobile,omitempty" `

    /*
        收货人的电话号码     */
    ReceiverPhone  *string `json:"receiver_phone,omitempty" `

    /*
        卖家备注旗帜（与淘宝网上订单的卖家备注旗帜对应，只有卖家才能查看该字段）红、黄、绿、蓝、紫 分别对应 1、2、3、4、5     */
    SellerFlag  *int64 `json:"seller_flag,omitempty" `

    /*
        是否保障速递，如果为true，则为保障速递订单，使用线下联系发货接口发货，如果未false，则该订单非保障速递，根据卖家设置的订单流转规则可使用物流宝或者常规物流发货。     */
    IsLgtype  *bool `json:"is_lgtype,omitempty" `

    /*
        次日达，三日达等送达类型     */
    LgAgingType  *string `json:"lg_aging_type,omitempty" `

    /*
        次日达订单送达时间     */
    LgAging  *string `json:"lg_aging,omitempty" `

    /*
        表示是否是品牌特卖（常规特卖，不包括特卖惠和特实惠）订单，如果是返回true，如果不是返回false。当此字段与is_force_wlb均为true时，订单强制物流宝发货。     */
    IsBrandSale  *bool `json:"is_brand_sale,omitempty" `

    /*
        订单是否强制使用物流宝发货。当此字段与is_brand_sale均为true时，订单强制物流宝发货。此字段为false时，该订单根据流转规则设置可以使用物流宝或者常规方式发货     */
    IsForceWlb  *bool `json:"is_force_wlb,omitempty" `

    /*
        判断订单是否有买家留言，有买家留言返回true，否则返回false     */
    HasBuyerMessage  *bool `json:"has_buyer_message,omitempty" `

    /*
        使用信用卡支付金额数     */
    CreditCardFee  *string `json:"credit_card_fee,omitempty" `

    /*
        交易内部来源。WAP(手机);HITAO(嗨淘);TOP(TOP平台);TAOBAO(普通淘宝);JHS(聚划算)一笔订单可能同时有以上多个标记，则以逗号分隔     */
    TradeFrom  *string `json:"trade_from,omitempty" `

    /*
        分阶段付款的订单状态（例如万人团订单等），目前有三返回状态FRONT_NOPAID_FINAL_NOPAID(定金未付尾款未付)，FRONT_PAID_FINAL_NOPAID(定金已付尾款未付)，FRONT_PAID_FINAL_PAID(定金和尾款都付)     */
    StepTradeStatus  *string `json:"step_trade_status,omitempty" `

    /*
        分阶段付款的已付金额（万人团订单已付金额）     */
    StepPaidFee  *string `json:"step_paid_fee,omitempty" `

    /*
        订单出现异常问题的时候，给予用户的描述,没有异常的时候，此值为空     */
    MarkDesc  *string `json:"mark_desc,omitempty" `

    /*
        订单中是否包含运费险订单，如果包含运费险订单返回true，不包含运费险订单，返回false     */
    HasYfx  *bool `json:"has_yfx,omitempty" `

    /*
        订单的运费险，单位为元     */
    YfxFee  *string `json:"yfx_fee,omitempty" `

    /*
        运费险支付号     */
    YfxId  *string `json:"yfx_id,omitempty" `

    /*
        运费险类型     */
    YfxType  *string `json:"yfx_type,omitempty" `

    /*
        交易外部来源：ownshop(商家官网)     */
    TradeSource  *string `json:"trade_source,omitempty" `

    /*
        订单将在此时间前发出，主要用于预售订单     */
    SendTime  *string `json:"send_time,omitempty" `

    /*
        卖家是否可以对订单进行评价     */
    SellerCanRate  *bool `json:"seller_can_rate,omitempty" `

    /*
        是否是多次发货的订单如果卖家对订单进行多次发货，则为true否则为false     */
    IsPartConsign  *bool `json:"is_part_consign,omitempty" `

    /*
        表示订单交易是否网厅订单。 如果该订单是网厅订单，那么该值为true；反之，该值为false。     */
    IsWt  *bool `json:"is_wt,omitempty" `

    /*
        表示订单交易是否含有对应的代销采购单。如果该订单中存在一个对应的代销采购单，那么该值为true；反之，该值为false。     */
    IsDaixiao  *bool `json:"is_daixiao,omitempty" `

    /*
        导购宝=crm     */
    O2o  *string `json:"o2o,omitempty" `

    /*
        导购员id     */
    O2oGuideId  *string `json:"o2o_guide_id,omitempty" `

    /*
        导购员名称     */
    O2oGuideName  *string `json:"o2o_guide_name,omitempty" `

    /*
        导购员门店id     */
    O2oShopId  *string `json:"o2o_shop_id,omitempty" `

    /*
        导购门店名称     */
    O2oShopName  *string `json:"o2o_shop_name,omitempty" `

    /*
        导购宝提货方式，inshop：店内提货，online：线上发货     */
    O2oDelivery  *string `json:"o2o_delivery,omitempty" `

    /*
        外部订单号     */
    O2oOutTradeId  *string `json:"o2o_out_trade_id,omitempty" `

    /*
        订单列表     */
    Orders  *[]TaobaoTradesSoldGetOrder `json:"orders,omitempty" `

    /*
        服务子订单列表     */
    ServiceOrders  *[]TaobaoTradesSoldGetServiceOrder `json:"service_orders,omitempty" `

    /*
        在返回的trade对象上专门增加一个字段zero_purchase来区分，这个为true的就是0元购机预授权交易     */
    ZeroPurchase  *bool `json:"zero_purchase,omitempty" `

    /*
        收货人国籍     */
    ReceiverCountry  *string `json:"receiver_country,omitempty" `

    /*
        天猫国际官网直供主订单关税税费     */
    OrderTaxFee  *string `json:"order_tax_fee,omitempty" `

    /*
        门店自提总店发货提货门店     */
    ShopPick  *string `json:"shop_pick,omitempty" `

    /*
        处方药未审核     */
    RxAuditStatus  *string `json:"rx_audit_status,omitempty" `

    /*
        邮关订单     */
    PostGateDeclare  *bool `json:"post_gate_declare,omitempty" `

    /*
        跨境订单     */
    CrossBondedDeclare  *bool `json:"cross_bonded_declare,omitempty" `

    /*
        天猫国际计税优惠金额     */
    OrderTaxPromotionFee  *string `json:"order_tax_promotion_fee,omitempty" `

    /*
        serviceType     */
    ServiceType  *string `json:"service_type,omitempty" `

    /*
        是否是智慧门店订单，只有true，或者 null 两种情况     */
    IsO2oPassport  *bool `json:"is_o2o_passport,omitempty" `

    /*
        tmallDelivery     */
    TmallDelivery  *bool `json:"tmall_delivery,omitempty" `

    /*
        threeplTiming     */
    ThreeplTiming  *bool `json:"threepl_timing,omitempty" `

    /*
        天猫直送服务     */
    CnService  *string `json:"cn_service,omitempty" `

    /*
        截单时间     */
    CutoffMinutes  *string `json:"cutoff_minutes,omitempty" `

    /*
        时效：天     */
    EsTime  *string `json:"es_time,omitempty" `

    /*
        发货时间     */
    DeliveryTime  *string `json:"delivery_time,omitempty" `

    /*
        揽收时间     */
    CollectTime  *string `json:"collect_time,omitempty" `

    /*
        派送时间     */
    DispatchTime  *string `json:"dispatch_time,omitempty" `

    /*
        签收时间     */
    SignTime  *string `json:"sign_time,omitempty" `

    /*
        派送CP     */
    DeliveryCps  *string `json:"delivery_cps,omitempty" `

    /*
        预计送达时间段     */
    EsRange  *string `json:"es_range,omitempty" `

    /*
        预计送达日期     */
    EsDate  *string `json:"es_date,omitempty" `

    /*
        预约送达时间段     */
    OsRange  *string `json:"os_range,omitempty" `

    /*
        预约送达日期     */
    OsDate  *string `json:"os_date,omitempty" `

    /*
        无物流信息返回true，平台属性，业务不要依赖     */
    NoShipping  *bool `json:"no_shipping,omitempty" `

    /*
        业务身份     */
    AsdpBizType  *string `json:"asdp_biz_type,omitempty" `

    /*
        送货上门标     */
    AsdpAds  *string `json:"asdp_ads,omitempty" `

    /*
        aid     */
    Aid  *string `json:"aid,omitempty" `

    /*
        是否疫情登记的订单。0=未登记，1=已登记     */
    DrugRegister  *string `json:"drug_register,omitempty" `

    /*
        承诺/最晚送达时间     */
    PromiseSignTime  *string `json:"promise_sign_time,omitempty" `

    /*
        （收货人+手机号+座机+收货地址+create）5个字段组合成oaid，原始订单上座机为空也满足条件，否则生成不了oaid     */
    Oaid  *string `json:"oaid,omitempty" `

    /*
        淘鲜达半日达     */
    ScenarioGroup  *string `json:"scenario_group,omitempty" `

    /*
        是否自动流转到菜鸟仓发货     */
    IsForceDc  *bool `json:"is_force_dc,omitempty" `

}

func (s *TaobaoTradesSoldGetTrade) SetSellerNick(v string) *TaobaoTradesSoldGetTrade {
    s.SellerNick = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetBuyerNick(v string) *TaobaoTradesSoldGetTrade {
    s.BuyerNick = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetBuyerOpenUid(v string) *TaobaoTradesSoldGetTrade {
    s.BuyerOpenUid = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetTitle(v string) *TaobaoTradesSoldGetTrade {
    s.Title = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetType(v string) *TaobaoTradesSoldGetTrade {
    s.Type = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetCreated(v util.LocalTime) *TaobaoTradesSoldGetTrade {
    s.Created = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetTid(v int64) *TaobaoTradesSoldGetTrade {
    s.Tid = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetAcookieId(v string) *TaobaoTradesSoldGetTrade {
    s.AcookieId = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetSid(v string) *TaobaoTradesSoldGetTrade {
    s.Sid = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetSellerRate(v bool) *TaobaoTradesSoldGetTrade {
    s.SellerRate = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetBuyerRate(v bool) *TaobaoTradesSoldGetTrade {
    s.BuyerRate = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetStatus(v string) *TaobaoTradesSoldGetTrade {
    s.Status = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetPayment(v string) *TaobaoTradesSoldGetTrade {
    s.Payment = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetDiscountFee(v string) *TaobaoTradesSoldGetTrade {
    s.DiscountFee = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetAdjustFee(v string) *TaobaoTradesSoldGetTrade {
    s.AdjustFee = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetPostFee(v string) *TaobaoTradesSoldGetTrade {
    s.PostFee = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetTotalFee(v string) *TaobaoTradesSoldGetTrade {
    s.TotalFee = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetPayTime(v util.LocalTime) *TaobaoTradesSoldGetTrade {
    s.PayTime = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetEndTime(v util.LocalTime) *TaobaoTradesSoldGetTrade {
    s.EndTime = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetModified(v util.LocalTime) *TaobaoTradesSoldGetTrade {
    s.Modified = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetAsyncModified(v util.LocalTime) *TaobaoTradesSoldGetTrade {
    s.AsyncModified = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetConsignTime(v util.LocalTime) *TaobaoTradesSoldGetTrade {
    s.ConsignTime = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetReceivedPayment(v string) *TaobaoTradesSoldGetTrade {
    s.ReceivedPayment = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetCommissionFee(v string) *TaobaoTradesSoldGetTrade {
    s.CommissionFee = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetAlipayNo(v string) *TaobaoTradesSoldGetTrade {
    s.AlipayNo = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetAlipayId(v int64) *TaobaoTradesSoldGetTrade {
    s.AlipayId = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetBuyerArea(v string) *TaobaoTradesSoldGetTrade {
    s.BuyerArea = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetNutFeature(v string) *TaobaoTradesSoldGetTrade {
    s.NutFeature = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetPicPath(v string) *TaobaoTradesSoldGetTrade {
    s.PicPath = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetNumIid(v int64) *TaobaoTradesSoldGetTrade {
    s.NumIid = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetPrice(v string) *TaobaoTradesSoldGetTrade {
    s.Price = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetCodFee(v string) *TaobaoTradesSoldGetTrade {
    s.CodFee = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetCodStatus(v string) *TaobaoTradesSoldGetTrade {
    s.CodStatus = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetShippingType(v string) *TaobaoTradesSoldGetTrade {
    s.ShippingType = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetNum(v int64) *TaobaoTradesSoldGetTrade {
    s.Num = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetPointFee(v int64) *TaobaoTradesSoldGetTrade {
    s.PointFee = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetRealPointFee(v int64) *TaobaoTradesSoldGetTrade {
    s.RealPointFee = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetBuyerObtainPointFee(v int64) *TaobaoTradesSoldGetTrade {
    s.BuyerObtainPointFee = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetReceiverName(v string) *TaobaoTradesSoldGetTrade {
    s.ReceiverName = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetReceiverState(v string) *TaobaoTradesSoldGetTrade {
    s.ReceiverState = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetReceiverCity(v string) *TaobaoTradesSoldGetTrade {
    s.ReceiverCity = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetReceiverDistrict(v string) *TaobaoTradesSoldGetTrade {
    s.ReceiverDistrict = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetReceiverTown(v string) *TaobaoTradesSoldGetTrade {
    s.ReceiverTown = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetReceiverAddress(v string) *TaobaoTradesSoldGetTrade {
    s.ReceiverAddress = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetReceiverZip(v string) *TaobaoTradesSoldGetTrade {
    s.ReceiverZip = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetReceiverMobile(v string) *TaobaoTradesSoldGetTrade {
    s.ReceiverMobile = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetReceiverPhone(v string) *TaobaoTradesSoldGetTrade {
    s.ReceiverPhone = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetSellerFlag(v int64) *TaobaoTradesSoldGetTrade {
    s.SellerFlag = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetIsLgtype(v bool) *TaobaoTradesSoldGetTrade {
    s.IsLgtype = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetLgAgingType(v string) *TaobaoTradesSoldGetTrade {
    s.LgAgingType = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetLgAging(v string) *TaobaoTradesSoldGetTrade {
    s.LgAging = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetIsBrandSale(v bool) *TaobaoTradesSoldGetTrade {
    s.IsBrandSale = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetIsForceWlb(v bool) *TaobaoTradesSoldGetTrade {
    s.IsForceWlb = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetHasBuyerMessage(v bool) *TaobaoTradesSoldGetTrade {
    s.HasBuyerMessage = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetCreditCardFee(v string) *TaobaoTradesSoldGetTrade {
    s.CreditCardFee = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetTradeFrom(v string) *TaobaoTradesSoldGetTrade {
    s.TradeFrom = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetStepTradeStatus(v string) *TaobaoTradesSoldGetTrade {
    s.StepTradeStatus = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetStepPaidFee(v string) *TaobaoTradesSoldGetTrade {
    s.StepPaidFee = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetMarkDesc(v string) *TaobaoTradesSoldGetTrade {
    s.MarkDesc = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetHasYfx(v bool) *TaobaoTradesSoldGetTrade {
    s.HasYfx = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetYfxFee(v string) *TaobaoTradesSoldGetTrade {
    s.YfxFee = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetYfxId(v string) *TaobaoTradesSoldGetTrade {
    s.YfxId = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetYfxType(v string) *TaobaoTradesSoldGetTrade {
    s.YfxType = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetTradeSource(v string) *TaobaoTradesSoldGetTrade {
    s.TradeSource = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetSendTime(v string) *TaobaoTradesSoldGetTrade {
    s.SendTime = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetSellerCanRate(v bool) *TaobaoTradesSoldGetTrade {
    s.SellerCanRate = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetIsPartConsign(v bool) *TaobaoTradesSoldGetTrade {
    s.IsPartConsign = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetIsWt(v bool) *TaobaoTradesSoldGetTrade {
    s.IsWt = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetIsDaixiao(v bool) *TaobaoTradesSoldGetTrade {
    s.IsDaixiao = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetO2o(v string) *TaobaoTradesSoldGetTrade {
    s.O2o = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetO2oGuideId(v string) *TaobaoTradesSoldGetTrade {
    s.O2oGuideId = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetO2oGuideName(v string) *TaobaoTradesSoldGetTrade {
    s.O2oGuideName = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetO2oShopId(v string) *TaobaoTradesSoldGetTrade {
    s.O2oShopId = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetO2oShopName(v string) *TaobaoTradesSoldGetTrade {
    s.O2oShopName = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetO2oDelivery(v string) *TaobaoTradesSoldGetTrade {
    s.O2oDelivery = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetO2oOutTradeId(v string) *TaobaoTradesSoldGetTrade {
    s.O2oOutTradeId = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetOrders(v []TaobaoTradesSoldGetOrder) *TaobaoTradesSoldGetTrade {
    s.Orders = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetServiceOrders(v []TaobaoTradesSoldGetServiceOrder) *TaobaoTradesSoldGetTrade {
    s.ServiceOrders = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetZeroPurchase(v bool) *TaobaoTradesSoldGetTrade {
    s.ZeroPurchase = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetReceiverCountry(v string) *TaobaoTradesSoldGetTrade {
    s.ReceiverCountry = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetOrderTaxFee(v string) *TaobaoTradesSoldGetTrade {
    s.OrderTaxFee = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetShopPick(v string) *TaobaoTradesSoldGetTrade {
    s.ShopPick = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetRxAuditStatus(v string) *TaobaoTradesSoldGetTrade {
    s.RxAuditStatus = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetPostGateDeclare(v bool) *TaobaoTradesSoldGetTrade {
    s.PostGateDeclare = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetCrossBondedDeclare(v bool) *TaobaoTradesSoldGetTrade {
    s.CrossBondedDeclare = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetOrderTaxPromotionFee(v string) *TaobaoTradesSoldGetTrade {
    s.OrderTaxPromotionFee = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetServiceType(v string) *TaobaoTradesSoldGetTrade {
    s.ServiceType = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetIsO2oPassport(v bool) *TaobaoTradesSoldGetTrade {
    s.IsO2oPassport = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetTmallDelivery(v bool) *TaobaoTradesSoldGetTrade {
    s.TmallDelivery = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetThreeplTiming(v bool) *TaobaoTradesSoldGetTrade {
    s.ThreeplTiming = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetCnService(v string) *TaobaoTradesSoldGetTrade {
    s.CnService = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetCutoffMinutes(v string) *TaobaoTradesSoldGetTrade {
    s.CutoffMinutes = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetEsTime(v string) *TaobaoTradesSoldGetTrade {
    s.EsTime = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetDeliveryTime(v string) *TaobaoTradesSoldGetTrade {
    s.DeliveryTime = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetCollectTime(v string) *TaobaoTradesSoldGetTrade {
    s.CollectTime = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetDispatchTime(v string) *TaobaoTradesSoldGetTrade {
    s.DispatchTime = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetSignTime(v string) *TaobaoTradesSoldGetTrade {
    s.SignTime = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetDeliveryCps(v string) *TaobaoTradesSoldGetTrade {
    s.DeliveryCps = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetEsRange(v string) *TaobaoTradesSoldGetTrade {
    s.EsRange = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetEsDate(v string) *TaobaoTradesSoldGetTrade {
    s.EsDate = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetOsRange(v string) *TaobaoTradesSoldGetTrade {
    s.OsRange = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetOsDate(v string) *TaobaoTradesSoldGetTrade {
    s.OsDate = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetNoShipping(v bool) *TaobaoTradesSoldGetTrade {
    s.NoShipping = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetAsdpBizType(v string) *TaobaoTradesSoldGetTrade {
    s.AsdpBizType = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetAsdpAds(v string) *TaobaoTradesSoldGetTrade {
    s.AsdpAds = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetAid(v string) *TaobaoTradesSoldGetTrade {
    s.Aid = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetDrugRegister(v string) *TaobaoTradesSoldGetTrade {
    s.DrugRegister = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetPromiseSignTime(v string) *TaobaoTradesSoldGetTrade {
    s.PromiseSignTime = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetOaid(v string) *TaobaoTradesSoldGetTrade {
    s.Oaid = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetScenarioGroup(v string) *TaobaoTradesSoldGetTrade {
    s.ScenarioGroup = &v
    return s
}
func (s *TaobaoTradesSoldGetTrade) SetIsForceDc(v bool) *TaobaoTradesSoldGetTrade {
    s.IsForceDc = &v
    return s
}
