package domain

import (
        "topsdk/util"
    )

type TaobaoTradesSoldIncrementvGetTrade struct {
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
        物流运单号     */
    Sid  *string `json:"sid,omitempty" `

    /*
        交易编号 (父订单的交易编号)     */
    Tid  *int64 `json:"tid,omitempty" `

    /*
        Acookie订单唯一ID。     */
    AcookieId  *string `json:"acookie_id,omitempty" `

    /*
        卖家是否已评价。可选值:true(已评价),false(未评价)     */
    SellerRate  *bool `json:"seller_rate,omitempty" `

    /*
        买家是否已评价。可选值:true(已评价),false(未评价)。如买家只评价未打分，此字段仍返回false     */
    BuyerRate  *bool `json:"buyer_rate,omitempty" `

    /*
        交易状态。可选值:    * TRADE_NO_CREATE_PAY(没有创建支付宝交易)    * WAIT_BUYER_PAY(等待买家付款)    * SELLER_CONSIGNED_PART(卖家部分发货)    * WAIT_SELLER_SEND_GOODS(等待卖家发货,即:买家已付款)    * WAIT_BUYER_CONFIRM_GOODS(等待买家确认收货,即:卖家已发货)    * TRADE_BUYER_SIGNED(买家已签收,货到付款专用)    * TRADE_FINISHED(交易成功)    * TRADE_CLOSED(付款以后用户退款成功，交易自动关闭)    * TRADE_CLOSED_BY_TAOBAO(付款以前，卖家或买家主动关闭交易)    * PAY_PENDING(国际信用卡支付付款确认中)    * WAIT_PRE_AUTH_CONFIRM(0元购合约中)     */
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
        分阶段付款的订单状态（例如万人团订单等），目前有三返回状态FRONT_NOPAID_FINAL_NOPAID(定金未付尾款未付)，FRONT_PAID_FINAL_NOPAID(定金已付尾款未付)，FRONT_PAID_FINAL_PAID(定金和尾款都付)     */
    StepTradeStatus  *string `json:"step_trade_status,omitempty" `

    /*
        分阶段付款的已付金额（万人团订单已付金额）     */
    StepPaidFee  *string `json:"step_paid_fee,omitempty" `

    /*
        订单出现异常问题的时候，给予用户的描述,没有异常的时候，此值为空     */
    MarkDesc  *string `json:"mark_desc,omitempty" `

    /*
        是否是多次发货的订单如果卖家对订单进行多次发货，则为true否则为false     */
    IsPartConsign  *bool `json:"is_part_consign,omitempty" `

    /*
        表示订单交易是否含有对应的代销采购单。如果该订单中存在一个对应的代销采购单，那么该值为true；反之，该值为false。     */
    IsDaixiao  *bool `json:"is_daixiao,omitempty" `

    /*
        表示订单交易是否网厅订单。 如果该订单是网厅订单，那么该值为true；反之，该值为false。     */
    IsWt  *bool `json:"is_wt,omitempty" `

    /*
        订单列表     */
    Orders  *[]TaobaoTradesSoldIncrementvGetOrder `json:"orders,omitempty" `

    /*
        服务子订单列表     */
    ServiceOrders  *[]TaobaoTradesSoldIncrementvGetServiceOrder `json:"service_orders,omitempty" `

    /*
        收货人国籍     */
    ReceiverCountry  *string `json:"receiver_country,omitempty" `

    /*
        天猫国际官网直供主订单关税税费     */
    OrderTaxFee  *string `json:"order_tax_fee,omitempty" `

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
        无物流信息返回true，平台属性，业务不要依赖     */
    NoShipping  *bool `json:"no_shipping,omitempty" `

    /*
        aid     */
    Aid  *string `json:"aid,omitempty" `

    /*
        （收货人+手机号+座机+收货地址+create）5个字段组合成oaid，原始订单上座机为空也满足条件，否则生成不了oaid     */
    Oaid  *string `json:"oaid,omitempty" `

}

func (s *TaobaoTradesSoldIncrementvGetTrade) SetSellerNick(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.SellerNick = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetBuyerNick(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.BuyerNick = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetBuyerOpenUid(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.BuyerOpenUid = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetTitle(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.Title = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetType(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.Type = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetCreated(v util.LocalTime) *TaobaoTradesSoldIncrementvGetTrade {
    s.Created = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetSid(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.Sid = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetTid(v int64) *TaobaoTradesSoldIncrementvGetTrade {
    s.Tid = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetAcookieId(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.AcookieId = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetSellerRate(v bool) *TaobaoTradesSoldIncrementvGetTrade {
    s.SellerRate = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetBuyerRate(v bool) *TaobaoTradesSoldIncrementvGetTrade {
    s.BuyerRate = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetStatus(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.Status = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetPayment(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.Payment = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetDiscountFee(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.DiscountFee = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetAdjustFee(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.AdjustFee = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetPostFee(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.PostFee = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetTotalFee(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.TotalFee = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetPayTime(v util.LocalTime) *TaobaoTradesSoldIncrementvGetTrade {
    s.PayTime = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetEndTime(v util.LocalTime) *TaobaoTradesSoldIncrementvGetTrade {
    s.EndTime = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetModified(v util.LocalTime) *TaobaoTradesSoldIncrementvGetTrade {
    s.Modified = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetAsyncModified(v util.LocalTime) *TaobaoTradesSoldIncrementvGetTrade {
    s.AsyncModified = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetConsignTime(v util.LocalTime) *TaobaoTradesSoldIncrementvGetTrade {
    s.ConsignTime = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetReceivedPayment(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.ReceivedPayment = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetCommissionFee(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.CommissionFee = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetAlipayNo(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.AlipayNo = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetAlipayId(v int64) *TaobaoTradesSoldIncrementvGetTrade {
    s.AlipayId = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetBuyerArea(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.BuyerArea = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetNutFeature(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.NutFeature = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetPicPath(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.PicPath = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetNumIid(v int64) *TaobaoTradesSoldIncrementvGetTrade {
    s.NumIid = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetPrice(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.Price = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetCodFee(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.CodFee = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetCodStatus(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.CodStatus = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetShippingType(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.ShippingType = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetNum(v int64) *TaobaoTradesSoldIncrementvGetTrade {
    s.Num = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetPointFee(v int64) *TaobaoTradesSoldIncrementvGetTrade {
    s.PointFee = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetRealPointFee(v int64) *TaobaoTradesSoldIncrementvGetTrade {
    s.RealPointFee = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetBuyerObtainPointFee(v int64) *TaobaoTradesSoldIncrementvGetTrade {
    s.BuyerObtainPointFee = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetReceiverName(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.ReceiverName = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetReceiverState(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.ReceiverState = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetReceiverCity(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.ReceiverCity = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetReceiverDistrict(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.ReceiverDistrict = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetReceiverTown(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.ReceiverTown = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetReceiverAddress(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.ReceiverAddress = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetReceiverZip(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.ReceiverZip = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetReceiverMobile(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.ReceiverMobile = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetReceiverPhone(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.ReceiverPhone = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetSellerFlag(v int64) *TaobaoTradesSoldIncrementvGetTrade {
    s.SellerFlag = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetIsLgtype(v bool) *TaobaoTradesSoldIncrementvGetTrade {
    s.IsLgtype = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetLgAgingType(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.LgAgingType = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetLgAging(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.LgAging = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetIsBrandSale(v bool) *TaobaoTradesSoldIncrementvGetTrade {
    s.IsBrandSale = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetIsForceWlb(v bool) *TaobaoTradesSoldIncrementvGetTrade {
    s.IsForceWlb = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetHasBuyerMessage(v bool) *TaobaoTradesSoldIncrementvGetTrade {
    s.HasBuyerMessage = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetCreditCardFee(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.CreditCardFee = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetStepTradeStatus(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.StepTradeStatus = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetStepPaidFee(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.StepPaidFee = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetMarkDesc(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.MarkDesc = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetIsPartConsign(v bool) *TaobaoTradesSoldIncrementvGetTrade {
    s.IsPartConsign = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetIsDaixiao(v bool) *TaobaoTradesSoldIncrementvGetTrade {
    s.IsDaixiao = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetIsWt(v bool) *TaobaoTradesSoldIncrementvGetTrade {
    s.IsWt = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetOrders(v []TaobaoTradesSoldIncrementvGetOrder) *TaobaoTradesSoldIncrementvGetTrade {
    s.Orders = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetServiceOrders(v []TaobaoTradesSoldIncrementvGetServiceOrder) *TaobaoTradesSoldIncrementvGetTrade {
    s.ServiceOrders = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetReceiverCountry(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.ReceiverCountry = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetOrderTaxFee(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.OrderTaxFee = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetPostGateDeclare(v bool) *TaobaoTradesSoldIncrementvGetTrade {
    s.PostGateDeclare = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetCrossBondedDeclare(v bool) *TaobaoTradesSoldIncrementvGetTrade {
    s.CrossBondedDeclare = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetOrderTaxPromotionFee(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.OrderTaxPromotionFee = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetNoShipping(v bool) *TaobaoTradesSoldIncrementvGetTrade {
    s.NoShipping = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetAid(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.Aid = &v
    return s
}
func (s *TaobaoTradesSoldIncrementvGetTrade) SetOaid(v string) *TaobaoTradesSoldIncrementvGetTrade {
    s.Oaid = &v
    return s
}
