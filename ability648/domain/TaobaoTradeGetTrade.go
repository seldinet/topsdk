package domain

import (
        "topsdk/util"
    )

type TaobaoTradeGetTrade struct {
    /*
        卖家昵称     */
    SellerNick  *string `json:"seller_nick,omitempty" `

    /*
        买家昵称     */
    BuyerNick  *string `json:"buyer_nick,omitempty" `

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
        卖家发货时间。格式:yyyy-MM-dd HH:mm:ss     */
    ConsignTime  *util.LocalTime `json:"consign_time,omitempty" `

    /*
        卖家实际收到的支付宝打款金额（由于子订单可以部分确认收货，这个金额会随着子订单的确认收货而不断增加，交易成功后等于买家实付款减去退款金额）。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
    ReceivedPayment  *string `json:"received_payment,omitempty" `

    /*
        交易佣金。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
    CommissionFee  *string `json:"commission_fee,omitempty" `

    /*
        买家备注（与淘宝网上订单的买家备注对应，只有买家才能查看该字段）     */
    BuyerMemo  *string `json:"buyer_memo,omitempty" `

    /*
        卖家备注（与淘宝网上订单的卖家备注对应，只有卖家才能查看该字段）     */
    SellerMemo  *string `json:"seller_memo,omitempty" `

    /*
        支付宝交易号，如：2009112081173831     */
    AlipayNo  *string `json:"alipay_no,omitempty" `

    /*
        买家留言     */
    BuyerMessage  *string `json:"buyer_message,omitempty" `

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
        表示订单交易是否含有对应的代销采购单。如果该订单中存在一个对应的代销采购单，那么该值为true；反之，该值为false。     */
    IsDaixiao  *bool `json:"is_daixiao,omitempty" `

    /*
        表示订单交易是否网厅订单。 如果该订单是网厅订单，那么该值为true；反之，该值为false。     */
    IsWt  *bool `json:"is_wt,omitempty" `

    /*
        物流到货时效，单位天     */
    ArriveInterval  *int64 `json:"arrive_interval,omitempty" `

    /*
        物流到货时效截单时间，格式 HH:mm     */
    ArriveCutTime  *string `json:"arrive_cut_time,omitempty" `

    /*
        物流发货时效，单位小时     */
    ConsignInterval  *int64 `json:"consign_interval,omitempty" `

    /*
        物流标签     */
    ServiceTags  *[]TaobaoTradeGetLogisticsTag `json:"service_tags,omitempty" `

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
    Orders  *[]TaobaoTradeGetOrder `json:"orders,omitempty" `

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
        天猫国际官网直供主订单关税税费     */
    OrderTaxFee  *string `json:"order_tax_fee,omitempty" `

    /*
        电子凭证服务上门服务的安装地址     */
    EticketServiceAddr  *string `json:"eticket_service_addr,omitempty" `

    /*
        分阶段交易的特权定金订单ID     */
    O2oEtOrderId  *string `json:"o2o_et_order_id,omitempty" `

    /*
        天猫国际计税优惠金额     */
    OrderTaxPromotionFee  *string `json:"order_tax_promotion_fee,omitempty" `

    /*
        透出前置营销工具     */
    Pmtp  *string `json:"pmtp,omitempty" `

    /*
        买家OpenUid     */
    BuyerOpenUid  *string `json:"buyer_open_uid,omitempty" `

}

func (s *TaobaoTradeGetTrade) SetSellerNick(v string) *TaobaoTradeGetTrade {
    s.SellerNick = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetBuyerNick(v string) *TaobaoTradeGetTrade {
    s.BuyerNick = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetTitle(v string) *TaobaoTradeGetTrade {
    s.Title = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetType(v string) *TaobaoTradeGetTrade {
    s.Type = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetCreated(v util.LocalTime) *TaobaoTradeGetTrade {
    s.Created = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetSid(v string) *TaobaoTradeGetTrade {
    s.Sid = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetTid(v int64) *TaobaoTradeGetTrade {
    s.Tid = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetSellerRate(v bool) *TaobaoTradeGetTrade {
    s.SellerRate = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetBuyerRate(v bool) *TaobaoTradeGetTrade {
    s.BuyerRate = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetStatus(v string) *TaobaoTradeGetTrade {
    s.Status = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetPayment(v string) *TaobaoTradeGetTrade {
    s.Payment = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetDiscountFee(v string) *TaobaoTradeGetTrade {
    s.DiscountFee = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetAdjustFee(v string) *TaobaoTradeGetTrade {
    s.AdjustFee = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetPostFee(v string) *TaobaoTradeGetTrade {
    s.PostFee = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetTotalFee(v string) *TaobaoTradeGetTrade {
    s.TotalFee = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetPayTime(v util.LocalTime) *TaobaoTradeGetTrade {
    s.PayTime = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetEndTime(v util.LocalTime) *TaobaoTradeGetTrade {
    s.EndTime = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetModified(v util.LocalTime) *TaobaoTradeGetTrade {
    s.Modified = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetConsignTime(v util.LocalTime) *TaobaoTradeGetTrade {
    s.ConsignTime = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetReceivedPayment(v string) *TaobaoTradeGetTrade {
    s.ReceivedPayment = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetCommissionFee(v string) *TaobaoTradeGetTrade {
    s.CommissionFee = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetBuyerMemo(v string) *TaobaoTradeGetTrade {
    s.BuyerMemo = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetSellerMemo(v string) *TaobaoTradeGetTrade {
    s.SellerMemo = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetAlipayNo(v string) *TaobaoTradeGetTrade {
    s.AlipayNo = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetBuyerMessage(v string) *TaobaoTradeGetTrade {
    s.BuyerMessage = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetPicPath(v string) *TaobaoTradeGetTrade {
    s.PicPath = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetNumIid(v int64) *TaobaoTradeGetTrade {
    s.NumIid = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetPrice(v string) *TaobaoTradeGetTrade {
    s.Price = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetCodFee(v string) *TaobaoTradeGetTrade {
    s.CodFee = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetCodStatus(v string) *TaobaoTradeGetTrade {
    s.CodStatus = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetShippingType(v string) *TaobaoTradeGetTrade {
    s.ShippingType = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetNum(v int64) *TaobaoTradeGetTrade {
    s.Num = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetPointFee(v int64) *TaobaoTradeGetTrade {
    s.PointFee = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetRealPointFee(v int64) *TaobaoTradeGetTrade {
    s.RealPointFee = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetBuyerObtainPointFee(v int64) *TaobaoTradeGetTrade {
    s.BuyerObtainPointFee = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetIsDaixiao(v bool) *TaobaoTradeGetTrade {
    s.IsDaixiao = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetIsWt(v bool) *TaobaoTradeGetTrade {
    s.IsWt = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetArriveInterval(v int64) *TaobaoTradeGetTrade {
    s.ArriveInterval = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetArriveCutTime(v string) *TaobaoTradeGetTrade {
    s.ArriveCutTime = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetConsignInterval(v int64) *TaobaoTradeGetTrade {
    s.ConsignInterval = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetServiceTags(v []TaobaoTradeGetLogisticsTag) *TaobaoTradeGetTrade {
    s.ServiceTags = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetO2o(v string) *TaobaoTradeGetTrade {
    s.O2o = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetO2oGuideId(v string) *TaobaoTradeGetTrade {
    s.O2oGuideId = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetO2oGuideName(v string) *TaobaoTradeGetTrade {
    s.O2oGuideName = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetO2oShopId(v string) *TaobaoTradeGetTrade {
    s.O2oShopId = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetO2oShopName(v string) *TaobaoTradeGetTrade {
    s.O2oShopName = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetO2oDelivery(v string) *TaobaoTradeGetTrade {
    s.O2oDelivery = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetO2oOutTradeId(v string) *TaobaoTradeGetTrade {
    s.O2oOutTradeId = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetOrders(v []TaobaoTradeGetOrder) *TaobaoTradeGetTrade {
    s.Orders = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetEtSerTime(v string) *TaobaoTradeGetTrade {
    s.EtSerTime = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetEtShopName(v string) *TaobaoTradeGetTrade {
    s.EtShopName = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetEtVerifiedShopName(v string) *TaobaoTradeGetTrade {
    s.EtVerifiedShopName = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetEtPlateNumber(v string) *TaobaoTradeGetTrade {
    s.EtPlateNumber = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetOrderTaxFee(v string) *TaobaoTradeGetTrade {
    s.OrderTaxFee = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetEticketServiceAddr(v string) *TaobaoTradeGetTrade {
    s.EticketServiceAddr = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetO2oEtOrderId(v string) *TaobaoTradeGetTrade {
    s.O2oEtOrderId = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetOrderTaxPromotionFee(v string) *TaobaoTradeGetTrade {
    s.OrderTaxPromotionFee = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetPmtp(v string) *TaobaoTradeGetTrade {
    s.Pmtp = &v
    return s
}
func (s *TaobaoTradeGetTrade) SetBuyerOpenUid(v string) *TaobaoTradeGetTrade {
    s.BuyerOpenUid = &v
    return s
}
