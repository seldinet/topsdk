package domain

import (
        "topsdk/util"
    )

type TaobaoTradeFullinfoGetTrade struct {
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
        买家下单的地区     */
    BuyerArea  *string `json:"buyer_area,omitempty" `

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
        买家货到付款服务费。精确到2位小数;单位:元。如:12.07，表示:12元7分     */
    BuyerCodFee  *string `json:"buyer_cod_fee,omitempty" `

    /*
        卖家货到付款服务费。精确到2位小数;单位:元。如:12.07，表示:12元7分。卖家不承担服务费的订单：未发货的订单获取服务费为0，发货后就能获取到正确值。     */
    SellerCodFee  *string `json:"seller_cod_fee,omitempty" `

    /*
        快递代收款。精确到2位小数;单位:元。如:212.07，表示:212元7分     */
    ExpressAgencyFee  *string `json:"express_agency_fee,omitempty" `

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
        买家支付宝账号     */
    BuyerAlipayNo  *string `json:"buyer_alipay_no,omitempty" `

    /*
        收货人的姓名     */
    ReceiverName  *string `json:"receiver_name,omitempty" `

    /*
        收货人国籍     */
    ReceiverCountry  *string `json:"receiver_country,omitempty" `

    /*
        收货人的所在省份     */
    ReceiverState  *string `json:"receiver_state,omitempty" `

    /*
        收货人的所在城市<br/>注：因为国家对于城市和地区的划分的有：省直辖市和省直辖县级行政区（区级别的）划分的，淘宝这边根据这个差异保存在不同字段里面比如：广东广州：广州属于一个直辖市是放在的receiver_city的字段里面；而河南济源：济源属于省直辖县级行政区划分，是区级别的，放在了receiver_district里面<br/>建议：程序依赖于城市字段做物流等判断的操作，最好加一个判断逻辑：如果返回值里面只有receiver_district参数，该参数作为城市     */
    ReceiverCity  *string `json:"receiver_city,omitempty" `

    /*
        收货人的所在地区<br/>注：因为国家对于城市和地区的划分的有：省直辖市和省直辖县级行政区（区级别的）划分的，淘宝这边根据这个差异保存在不同字段里面比如：广东广州：广州属于一个直辖市是放在的receiver_city的字段里面；而河南济源：济源属于省直辖县级行政区划分，是区级别的，放在了receiver_district里面<br/>建议：程序依赖于城市字段做物流等判断的操作，最好加一个判断逻辑：如果返回值里面只有receiver_district参数，该参数作为城市     */
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
        买家邮件地址     */
    BuyerEmail  *string `json:"buyer_email,omitempty" `

    /*
        卖家支付宝账号     */
    SellerAlipayNo  *string `json:"seller_alipay_no,omitempty" `

    /*
        卖家手机     */
    SellerMobile  *string `json:"seller_mobile,omitempty" `

    /*
        卖家电话     */
    SellerPhone  *string `json:"seller_phone,omitempty" `

    /*
        卖家姓名     */
    SellerName  *string `json:"seller_name,omitempty" `

    /*
        卖家邮件地址     */
    SellerEmail  *string `json:"seller_email,omitempty" `

    /*
        交易中剩余的确认收货金额（这个金额会随着子订单确认收货而不断减少，交易成功后会变为零）。精确到2位小数;单位:元。如:200.07，表示:200 元7分     */
    AvailableConfirmFee  *string `json:"available_confirm_fee,omitempty" `

    /*
        是否包含邮费。与available_confirm_fee同时使用。可选值:true(包含),false(不包含)     */
    HasPostFee  *bool `json:"has_post_fee,omitempty" `

    /*
        超时到期时间。格式:yyyy-MM-dd HH:mm:ss。业务规则：前提条件：只有在买家已付款，卖家已发货的情况下才有效如果申请了退款，那么超时会落在子订单上；比如说3笔ABC，A申请了，那么返回的是BC的列表, 主定单不存在如果没有申请过退款，那么超时会挂在主定单上；比如ABC，返回主定单，ABC的超时和主定单相同     */
    TimeoutActionTime  *util.LocalTime `json:"timeout_action_time,omitempty" `

    /*
        交易快照地址     */
    SnapshotUrl  *string `json:"snapshot_url,omitempty" `

    /*
        交易备注。     */
    TradeMemo  *string `json:"trade_memo,omitempty" `

    /*
        交易促销详细信息     */
    Promotion  *string `json:"promotion,omitempty" `

    /*
        是否3D交易     */
    Is3D  *bool `json:"is_3D,omitempty" `

    /*
        是否保障速递，如果为true，则为保障速递订单，使用线下联系发货接口发货，如果未false，则该订单非保障速递，根据卖家设置的订单流转规则可使用物流宝或者常规物流发货。     */
    IsLgtype  *bool `json:"is_lgtype,omitempty" `

    /*
        表示是否是品牌特卖（常规特卖，不包括特卖惠和特实惠）订单，如果是返回true，如果不是返回false。当此字段与is_force_wlb均为true时，订单强制物流宝发货。     */
    IsBrandSale  *bool `json:"is_brand_sale,omitempty" `

    /*
        订单是否强制使用物流宝发货。当此字段与is_brand_sale均为true时，订单强制物流宝发货。此字段为false时，该订单根据流转规则设置可以使用物流宝或者常规方式发货     */
    IsForceWlb  *bool `json:"is_force_wlb,omitempty" `

    /*
        买家备注旗帜（与淘宝网上订单的买家备注旗帜对应，只有买家才能查看该字段）红、黄、绿、蓝、紫 分别对应 1、2、3、4、5     */
    BuyerFlag  *int64 `json:"buyer_flag,omitempty" `

    /*
        卖家备注旗帜（与淘宝网上订单的卖家备注旗帜对应，只有卖家才能查看该字段）红、黄、绿、蓝、紫 分别对应 1、2、3、4、5     */
    SellerFlag  *int64 `json:"seller_flag,omitempty" `

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
        电子凭证的垂直信息     */
    EticketExt  *string `json:"eticket_ext,omitempty" `

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
        订单将在此时间前发出，主要用于预售订单     */
    SendTime  *string `json:"send_time,omitempty" `

    /*
        买家可以通过此字段查询是否当前交易可以评论，can_rate=true可以评价，false则不能评价。     */
    CanRate  *bool `json:"can_rate,omitempty" `

    /*
        卖家是否可以对订单进行评价     */
    SellerCanRate  *bool `json:"seller_can_rate,omitempty" `

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
        物流到货时效，单位天     */
    ArriveInterval  *int64 `json:"arrive_interval,omitempty" `

    /*
        物流到货时效截单时间，格式 HH:mm     */
    ArriveCutTime  *string `json:"arrive_cut_time,omitempty" `

    /*
        物流发货时效，单位小时     */
    ConsignInterval  *int64 `json:"consign_interval,omitempty" `

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
        线下自提门店编码     */
    ShopCode  *string `json:"shop_code,omitempty" `

    /*
        拼音名     */
    HkEnName  *string `json:"hk_en_name,omitempty" `

    /*
        航班号     */
    HkFlightNo  *string `json:"hk_flight_no,omitempty" `

    /*
        中文名     */
    HkChinaName  *string `json:"hk_china_name,omitempty" `

    /*
        证件号码     */
    HkCardCode  *string `json:"hk_card_code,omitempty" `

    /*
        证件类型001代表港澳通行证类型、002代表入台证003代表护照     */
    HkCardType  *string `json:"hk_card_type,omitempty" `

    /*
        航班飞行时间     */
    HkFlightDate  *string `json:"hk_flight_date,omitempty" `

    /*
        性别M: 男性，F: 女性     */
    HkGender  *string `json:"hk_gender,omitempty" `

    /*
        出生日期     */
    HkBirthday  *string `json:"hk_birthday,omitempty" `

    /*
        提货地点     */
    HkPickup  *string `json:"hk_pickup,omitempty" `

    /*
        提货地点id     */
    HkPickupId  *string `json:"hk_pickup_id,omitempty" `

    /*
        商家的预计发货时间     */
    EstConTime  *string `json:"est_con_time,omitempty" `

    /*
        订单列表     */
    Orders  *[]TaobaoTradeFullinfoGetOrder `json:"orders,omitempty" `

    /*
        优惠详情     */
    PromotionDetails  *[]TaobaoTradeFullinfoGetPromotionDetail `json:"promotion_details,omitempty" `

    /*
        物流标签     */
    ServiceTags  *[]TaobaoTradeFullinfoGetLogisticsTag `json:"service_tags,omitempty" `

    /*
        服务子订单列表     */
    ServiceOrders  *[]TaobaoTradeFullinfoGetServiceOrder `json:"service_orders,omitempty" `

    /*
        交易内部来源。WAP(手机);HITAO(嗨淘);TOP(TOP平台);TAOBAO(普通淘宝);JHS(聚划算)一笔订单可能同时有以上多个标记，则以逗号分隔     */
    TradeFrom  *string `json:"trade_from,omitempty" `

    /*
        交易外部来源：ownshop(商家官网)     */
    TradeSource  *string `json:"trade_source,omitempty" `

    /*
        在返回的trade对象上专门增加一个字段zero_purchase来区分，这个为true的就是0元购机预授权交易     */
    ZeroPurchase  *bool `json:"zero_purchase,omitempty" `

    /*
        付款时使用的支付宝积分的额度,单位分，比如返回1，则为1分钱     */
    AlipayPoint  *int64 `json:"alipay_point,omitempty" `

    /*
        天猫点券卡实付款金额,单位分     */
    PccAf  *int64 `json:"pcc_af,omitempty" `

    /*
        交易扩展表信息     */
    TradeExt  *TaobaoTradeFullinfoGetTradeExt `json:"trade_ext,omitempty" `

    /*
        天猫国际官网直供主订单关税税费     */
    OrderTaxFee  *string `json:"order_tax_fee,omitempty" `

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
        抢单状态<br/>0,未处理待分发;1,抢单中;2,已抢单;3,已发货;-1,超时;-2,处理异常;-3,匹配失败;-4,取消抢单;-5,退款取消;-9,逻辑删除     */
    O2oSnatchStatus  *string `json:"o2o_snatch_status,omitempty" `

    /*
        是否屏蔽发货     */
    IsShShip  *bool `json:"is_sh_ship,omitempty" `

    /*
        天猫电子凭证家装     */
    EticketServiceAddr  *string `json:"eticket_service_addr,omitempty" `

    /*
        电子凭证扫码购-扫码支付订单type     */
    EtType  *string `json:"et_type,omitempty" `

    /*
        垂直市场     */
    Market  *string `json:"market,omitempty" `

    /*
        门店预约自提订单标     */
    Obs  *string `json:"obs,omitempty" `

    /*
        扫码购关联门店     */
    EtShopId  *int64 `json:"et_shop_id,omitempty" `

    /*
        满返红包的金额；如果没有满返红包，则值为 0.00     */
    PaidCouponFee  *string `json:"paid_coupon_fee,omitempty" `

    /*
        线下门店自提     */
    ShopPick  *string `json:"shop_pick,omitempty" `

    /*
        处方药审核状态     */
    RxAuditStatus  *string `json:"rx_audit_status,omitempty" `

    /*
        yyyyMMdd     */
    EsDate  *string `json:"es_date,omitempty" `

    /*
        hh:mm-hh:mm     */
    EsRange  *string `json:"es_range,omitempty" `

    /*
        yyyyMMdd     */
    OsDate  *string `json:"os_date,omitempty" `

    /*
        hh:mm-hh:mm     */
    OsRange  *string `json:"os_range,omitempty" `

    /*
        100     */
    CouponFee  *int64 `json:"coupon_fee,omitempty" `

    /*
        邮关订单     */
    PostGateDeclare  *bool `json:"post_gate_declare,omitempty" `

    /*
        跨境订单     */
    CrossBondedDeclare  *bool `json:"cross_bonded_declare,omitempty" `

    /*
        主订单扩展属性;payCurrency:HKD(港币)、TWD(台币)     */
    TradeAttr  *string `json:"trade_attr,omitempty" `

    /*
        TOP拦截标识，0不拦截，1拦截     */
    TopHold  *int64 `json:"top_hold,omitempty" `

    /*
        星盘标识字段     */
    OmniAttr  *string `json:"omni_attr,omitempty" `

    /*
        星盘业务字段     */
    OmniParam  *string `json:"omni_param,omitempty" `

    /*
        assembly     */
    Assembly  *string `json:"assembly,omitempty" `

    /*
        聚划算一起买字段     */
    ForbidConsign  *int64 `json:"forbid_consign,omitempty" `

    /*
        采购订单标识     */
    Identity  *string `json:"identity,omitempty" `

    /*
        天猫拼团拦截标示     */
    TeamBuyHold  *int64 `json:"team_buy_hold,omitempty" `

    /*
        分享购拦截     */
    ShareGroupHold  *int64 `json:"share_group_hold,omitempty" `

    /*
        天猫国际拦截     */
    OfpHold  *int64 `json:"ofp_hold,omitempty" `

    /*
        组装O2O多阶段尾款订单的明细数据 总阶段数，当前阶数，阶段金额（单位：分），支付状态，例如 3_1_100_paid ; 3_2_2000_nopaid     */
    O2oStepTradeDetail  *string `json:"o2o_step_trade_detail,omitempty" `

    /*
        特权定金订单的尾款订单ID     */
    O2oStepOrderId  *string `json:"o2o_step_order_id,omitempty" `

    /*
        分阶段订单的特权定金订单ID     */
    O2oEtOrderId  *string `json:"o2o_et_order_id,omitempty" `

    /*
        分阶段订单的特权定金抵扣金额，单位：分     */
    O2oVoucherPrice  *string `json:"o2o_voucher_price,omitempty" `

    /*
        天猫国际计税优惠金额     */
    OrderTaxPromotionFee  *string `json:"order_tax_promotion_fee,omitempty" `

    /*
        聚划算火拼标记     */
    DelayCreateDelivery  *int64 `json:"delay_create_delivery,omitempty" `

    /*
        tidStr     */
    TidStr  *string `json:"tid_str,omitempty" `

    /*
        top定义订单类型     */
    Toptype  *int64 `json:"toptype,omitempty" `

    /*
        包含的交易服务类型     */
    ServiceType  *string `json:"service_type,omitempty" `

    /*
        o2oServiceMobile     */
    O2oServiceMobile  *string `json:"o2o_service_mobile,omitempty" `

    /*
        o2oServiceName     */
    O2oServiceName  *string `json:"o2o_service_name,omitempty" `

    /*
        o2oServiceState     */
    O2oServiceState  *string `json:"o2o_service_state,omitempty" `

    /*
        o2oServiceCity     */
    O2oServiceCity  *string `json:"o2o_service_city,omitempty" `

    /*
        o2oServiceDistrict     */
    O2oServiceDistrict  *string `json:"o2o_service_district,omitempty" `

    /*
        o2oServiceTown     */
    O2oServiceTown  *string `json:"o2o_service_town,omitempty" `

    /*
        o2oServiceAddress     */
    O2oServiceAddress  *string `json:"o2o_service_address,omitempty" `

    /*
        o2oStepTradeDetailNew     */
    O2oStepTradeDetailNew  *string `json:"o2o_step_trade_detail_new,omitempty" `

    /*
        o2oXiaopiao     */
    O2oXiaopiao  *string `json:"o2o_xiaopiao,omitempty" `

    /*
        o2oContract     */
    O2oContract  *string `json:"o2o_contract,omitempty" `

    /*
        rechargeFee     */
    RechargeFee  *string `json:"recharge_fee,omitempty" `

    /*
        retailStoreCode     */
    RetailStoreCode  *string `json:"retail_store_code,omitempty" `

    /*
        retailOutOrderId     */
    RetailOutOrderId  *string `json:"retail_out_order_id,omitempty" `

    /*
        platformSubsidyFee     */
    PlatformSubsidyFee  *string `json:"platform_subsidy_fee,omitempty" `

    /*
        nrOffline     */
    NrOffline  *string `json:"nr_offline,omitempty" `

    /*
        网厅订单垂直表信息     */
    WttParam  *string `json:"wtt_param,omitempty" `

    /*
        子单物流发货信息     */
    LogisticsInfos  *[]TaobaoTradeFullinfoGetLogisticsInfo `json:"logistics_infos,omitempty" `

    /*
        sellerNick     */
    SellerNick  *string `json:"seller_nick,omitempty" `

    /*
        buyerNick     */
    BuyerNick  *string `json:"buyer_nick,omitempty" `

    /*
        nrStoreOrderId     */
    NrStoreOrderId  *string `json:"nr_store_order_id,omitempty" `

    /*
        门店 ID     */
    NrShopId  *string `json:"nr_shop_id,omitempty" `

    /*
        门店名称     */
    NrShopName  *string `json:"nr_shop_name,omitempty" `

    /*
        导购员ID     */
    NrShopGuideId  *string `json:"nr_shop_guide_id,omitempty" `

    /*
        导购员名称     */
    NrShopGuideName  *string `json:"nr_shop_guide_name,omitempty" `

    /*
        sortInfo     */
    SortInfo  *TaobaoTradeFullinfoGetSortInfo `json:"sort_info,omitempty" `

    /*
        1已排序 2不排序     */
    Sorted  *int64 `json:"sorted,omitempty" `

    /*
        一小时达不处理订单     */
    NrNoHandle  *string `json:"nr_no_handle,omitempty" `

    /*
        为tmall.daogoubao.cloudstore时表示云店链路     */
    BizCode  *string `json:"biz_code,omitempty" `

    /*
        为1，且bizCode不为tmall.daogoubao.cloudstore时，为旗舰店订单     */
    CloudStore  *string `json:"cloud_store,omitempty" `

    /*
        暂不公开     */
    IsGift  *bool `json:"is_gift,omitempty" `

    /*
        暂不公开     */
    DoneeNick  *string `json:"donee_nick,omitempty" `

    /*
        暂不公开     */
    DoneeOpenUid  *string `json:"donee_open_uid,omitempty" `

    /*
        苏宁自提门店code     */
    SuningShopCode  *string `json:"suning_shop_code,omitempty" `

    /*
        苏宁自提门店是否有效     */
    SuningShopValid  *int64 `json:"suning_shop_valid,omitempty" `

    /*
        允许的appkey，逗号分隔     */
    AllowAppkeys  *string `json:"allow_appkeys,omitempty" `

    /*
        是否预售     */
    NewPresell  *bool `json:"new_presell,omitempty" `

    /*
        是否优享     */
    YouXiang  *bool `json:"you_xiang,omitempty" `

    /*
        天猫未来店线下店铺 ID     */
    RetailStoreId  *string `json:"retail_store_id,omitempty" `

    /*
        区分istore订单和普通订单     */
    IsIstore  *bool `json:"is_istore,omitempty" `

    /*
        区分istore订单来源     */
    Ua  *string `json:"ua,omitempty" `

    /*
        linkedmall透传参数     */
    LinkedmallExtInfo  *string `json:"linkedmall_ext_info,omitempty" `

    /*
        支付渠道：0 用户主动支付 1 系统代扣 2 保险赔付     */
    PayChannel  *string `json:"pay_channel,omitempty" `

    /*
        新零售全渠道订单：订单类型，自提订单：pickUp，电商发货：tmall，门店发货（配送、骑手）：storeSend     */
    RtOmniSendType  *string `json:"rt_omni_send_type,omitempty" `

    /*
        新零售全渠道订单：发货门店ID     */
    RtOmniStoreId  *string `json:"rt_omni_store_id,omitempty" `

    /*
        新零售全渠道订单：商家自有发货门店编码     */
    RtOmniOuterStoreId  *string `json:"rt_omni_outer_store_id,omitempty" `

    /*
        同城预约配送开始时间     */
    TcpsStart  *string `json:"tcps_start,omitempty" `

    /*
        同城业务类型,com.tmall.dsd:定时送,storeDsd-fn-3-1:淘速达3公里蜂鸟配送     */
    TcpsCode  *string `json:"tcps_code,omitempty" `

    /*
        同城预约配送结束时间     */
    TcpsEnd  *string `json:"tcps_end,omitempty" `

    /*
        主订单商家代缴税费     */
    MTariffFee  *string `json:"m_tariff_fee,omitempty" `

    /*
        时效服务身份，如tmallPromise代表天猫时效承诺     */
    TimingPromise  *string `json:"timing_promise,omitempty" `

    /*
        时效服务字段，服务字段，会有多个服务值，以英文半角逗号","切割     */
    PromiseService  *string `json:"promise_service,omitempty" `

    /*
        物流截单时间，分钟     */
    CutoffMinutes  *string `json:"cutoff_minutes,omitempty" `

    /*
        物流时效，相对时间，单位是天     */
    EsTime  *string `json:"es_time,omitempty" `

    /*
        最晚发货时间，日期     */
    DeliveryTime  *string `json:"delivery_time,omitempty" `

    /*
        最晚揽收时间，日期     */
    CollectTime  *string `json:"collect_time,omitempty" `

    /*
        最晚派送时间，日期     */
    DispatchTime  *string `json:"dispatch_time,omitempty" `

    /*
        最晚签收时间，日期     */
    SignTime  *string `json:"sign_time,omitempty" `

    /*
        外部会员id     */
    OuterPartnerMemberId  *string `json:"outer_partner_member_id,omitempty" `

    /*
        叶子分类     */
    RootCat  *string `json:"root_cat,omitempty" `

    /*
        1-gifting订单     */
    Gifting  *string `json:"gifting,omitempty" `

    /*
        1-coffee gifting订单     */
    GiftingTakeout  *string `json:"gifting_takeout,omitempty" `

    /*
        预约安装时间     */
    OiDate  *string `json:"oi_date,omitempty" `

    /*
        预约安装时间段     */
    OiRange  *string `json:"oi_range,omitempty" `

    /*
        暂不安装     */
    HoldInstall  *string `json:"hold_install,omitempty" `

    /*
        订单来源     */
    AppName  *string `json:"app_name,omitempty" `

    /*
        同城站类型     */
    EasyHomeCityType  *string `json:"easy_home_city_type,omitempty" `

    /*
        同城站关联订单号     */
    NrDepositOrderId  *string `json:"nr_deposit_order_id,omitempty" `

    /*
        摊位id     */
    NrStoreCode  *string `json:"nr_store_code,omitempty" `

    /*
        使用淘金币的数量，以分为单位，和订单标propoint中间那一段一样，没有返回null     */
    Propoint  *string `json:"propoint,omitempty" `

    /*
        是否周期送订单     */
    ZqsOrderTag  *string `json:"zqs_order_tag,omitempty" `

    /*
        天鲜配冰柜id     */
    TxpFreezerId  *string `json:"txp_freezer_id,omitempty" `

    /*
        天鲜配自提方式     */
    TxpReceiveMethod  *string `json:"txp_receive_method,omitempty" `

    /*
        购物金信息输出     */
    ExpandcardInfo  *TaobaoTradeFullinfoGetExpandCardInfo `json:"expandcard_info,omitempty" `

    /*
        透出的额外信息     */
    ExtendInfo  *string `json:"extend_info,omitempty" `

    /*
        收货地址有变更，返回"1"     */
    Lm  *string `json:"lm,omitempty" `

    /*
        同城购订单来源     */
    BrandLightShopSource  *string `json:"brand_light_shop_source,omitempty" `

    /*
        同城购渠道店id     */
    BrandLightShopStoreId  *string `json:"brand_light_shop_store_id,omitempty" `

    /*
        标识完美履约订单     */
    IsWmly  *string `json:"is_wmly,omitempty" `

    /*
        全渠道包裹信息     */
    OmniPackage  *string `json:"omni_package,omitempty" `

    /*
        新康众扩展数据     */
    NczExtAttr  *string `json:"ncz_ext_attr,omitempty" `

    /*
        苹果发票详情     */
    InvoiceDetailPay  *string `json:"invoice_detail_pay,omitempty" `

    /*
        苹果发票详情     */
    InvoiceDetailMidRefund  *string `json:"invoice_detail_mid_refund,omitempty" `

    /*
        苹果发票详情     */
    InvoiceDetailAfterRefund  *string `json:"invoice_detail_after_refund,omitempty" `

    /*
        买卡订单本金     */
    ExpandCardBasicPrice  *string `json:"expand_card_basic_price,omitempty" `

    /*
        买卡订单权益金     */
    ExpandCardExpandPrice  *string `json:"expand_card_expand_price,omitempty" `

    /*
        用卡订单所用的本金     */
    ExpandCardBasicPriceUsed  *string `json:"expand_card_basic_price_used,omitempty" `

    /*
        用卡订单所用的权益金     */
    ExpandCardExpandPriceUsed  *string `json:"expand_card_expand_price_used,omitempty" `

    /*
        配送cp     */
    DeliveryCps  *string `json:"delivery_cps,omitempty" `

    /*
        业务身份     */
    AsdpBizType  *string `json:"asdp_biz_type,omitempty" `

    /*
        是否是Openmall订单     */
    IsOpenmall  *bool `json:"is_openmall,omitempty" `

    /*
        是否是码上收订单     */
    VLogisticsCreate  *bool `json:"v_logistics_create,omitempty" `

    /*
        是否是非物流订单     */
    QRPay  *bool `json:"q_r_pay,omitempty" `

    /*
        关联下单订单     */
    OrderFollowId  *string `json:"order_follow_id,omitempty" `

    /*
        无物流信息返回true，平台属性，业务不要依赖     */
    NoShipping  *bool `json:"no_shipping,omitempty" `

    /*
        送货上门标     */
    AsdpAds  *string `json:"asdp_ads,omitempty" `

    /*
        消费者催发货标识，lg表示消费者做过催发货     */
    ObTag  *string `json:"ob_tag,omitempty" `

    /*
        地址aid字段     */
    Aid  *string `json:"aid,omitempty" `

    /*
        通用的是否预售，默认是false，需要传general_new_presell参数才能真正识别是否是预售订单     */
    GeneralNewPresell  *bool `json:"general_new_presell,omitempty" `

    /*
        是否疫情登记的订单。0=未登记，1=已登记     */
    DrugRegister  *string `json:"drug_register,omitempty" `

    /*
        分阶段支付详情     */
    StepPayDetails  *[]TaobaoTradeFullinfoGetStepPayDetail `json:"step_pay_details,omitempty" `

    /*
        同意退款检查标识字段     */
    AgreeRefundChecks  *[]TaobaoTradeFullinfoGetAgreeRefundCheck `json:"agree_refund_checks,omitempty" `

    /*
        阶段地址详情字段     */
    AddressDetails  *[]TaobaoTradeFullinfoGetAddressDetail `json:"address_details,omitempty" `

    /*
        阶段收货地址标识字段     */
    StageAddressType  *string `json:"stage_address_type,omitempty" `

    /*
        订单分组ID     */
    OgId  *string `json:"og_id,omitempty" `

    /*
        承诺/最晚送达时间     */
    PromiseSignTime  *string `json:"promise_sign_time,omitempty" `

    /*
        订单是否属于b2b代销     */
    B2bDaixiao  *int64 `json:"b2b_daixiao,omitempty" `

    /*
        天猫商家使用，订单使用的红包信息     */
    TmallCouponFee  *int64 `json:"tmall_coupon_fee,omitempty" `

    /*
        交付计划     */
    DeliveryPlan  *[]TaobaoTradeFullinfoGetDeliveryPlan `json:"delivery_plan,omitempty" `

    /*
        全渠道订单相关字段     */
    OmnichannelParam  *string `json:"omnichannel_param,omitempty" `

    /*
        入参fields字段必须包含receiver_name、receiver_address、created、receiver_mobile、receiver_phone 5个字段，否则无法生成oaid。     */
    Oaid  *string `json:"oaid,omitempty" `

    /*
        CRM系统中特有的ouid     */
    Ouid  *string `json:"ouid,omitempty" `

    /*
        是否是周期购订单     */
    IsCycleBuy  *bool `json:"is_cycle_buy,omitempty" `

    /*
        淘鲜达生鲜半日达     */
    ScenarioGroup  *string `json:"scenario_group,omitempty" `

    /*
        拼团玩法垂直标     */
    PlayType  *string `json:"play_type,omitempty" `

    /*
        ascp会自动流转到菜鸟仓发货     */
    IsForceDc  *bool `json:"is_force_dc,omitempty" `

}

func (s *TaobaoTradeFullinfoGetTrade) SetBuyerOpenUid(v string) *TaobaoTradeFullinfoGetTrade {
    s.BuyerOpenUid = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetTitle(v string) *TaobaoTradeFullinfoGetTrade {
    s.Title = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetType(v string) *TaobaoTradeFullinfoGetTrade {
    s.Type = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetCreated(v util.LocalTime) *TaobaoTradeFullinfoGetTrade {
    s.Created = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetSid(v string) *TaobaoTradeFullinfoGetTrade {
    s.Sid = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetTid(v int64) *TaobaoTradeFullinfoGetTrade {
    s.Tid = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetAcookieId(v string) *TaobaoTradeFullinfoGetTrade {
    s.AcookieId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetSellerRate(v bool) *TaobaoTradeFullinfoGetTrade {
    s.SellerRate = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetBuyerRate(v bool) *TaobaoTradeFullinfoGetTrade {
    s.BuyerRate = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetStatus(v string) *TaobaoTradeFullinfoGetTrade {
    s.Status = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetPayment(v string) *TaobaoTradeFullinfoGetTrade {
    s.Payment = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetDiscountFee(v string) *TaobaoTradeFullinfoGetTrade {
    s.DiscountFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetAdjustFee(v string) *TaobaoTradeFullinfoGetTrade {
    s.AdjustFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetPostFee(v string) *TaobaoTradeFullinfoGetTrade {
    s.PostFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetTotalFee(v string) *TaobaoTradeFullinfoGetTrade {
    s.TotalFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetPayTime(v util.LocalTime) *TaobaoTradeFullinfoGetTrade {
    s.PayTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetEndTime(v util.LocalTime) *TaobaoTradeFullinfoGetTrade {
    s.EndTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetModified(v util.LocalTime) *TaobaoTradeFullinfoGetTrade {
    s.Modified = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetConsignTime(v util.LocalTime) *TaobaoTradeFullinfoGetTrade {
    s.ConsignTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetReceivedPayment(v string) *TaobaoTradeFullinfoGetTrade {
    s.ReceivedPayment = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetCommissionFee(v string) *TaobaoTradeFullinfoGetTrade {
    s.CommissionFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetBuyerMemo(v string) *TaobaoTradeFullinfoGetTrade {
    s.BuyerMemo = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetSellerMemo(v string) *TaobaoTradeFullinfoGetTrade {
    s.SellerMemo = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetBuyerArea(v string) *TaobaoTradeFullinfoGetTrade {
    s.BuyerArea = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetAlipayNo(v string) *TaobaoTradeFullinfoGetTrade {
    s.AlipayNo = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetBuyerMessage(v string) *TaobaoTradeFullinfoGetTrade {
    s.BuyerMessage = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetPicPath(v string) *TaobaoTradeFullinfoGetTrade {
    s.PicPath = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetNumIid(v int64) *TaobaoTradeFullinfoGetTrade {
    s.NumIid = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetPrice(v string) *TaobaoTradeFullinfoGetTrade {
    s.Price = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetCodFee(v string) *TaobaoTradeFullinfoGetTrade {
    s.CodFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetCodStatus(v string) *TaobaoTradeFullinfoGetTrade {
    s.CodStatus = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetBuyerCodFee(v string) *TaobaoTradeFullinfoGetTrade {
    s.BuyerCodFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetSellerCodFee(v string) *TaobaoTradeFullinfoGetTrade {
    s.SellerCodFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetExpressAgencyFee(v string) *TaobaoTradeFullinfoGetTrade {
    s.ExpressAgencyFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetShippingType(v string) *TaobaoTradeFullinfoGetTrade {
    s.ShippingType = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetNum(v int64) *TaobaoTradeFullinfoGetTrade {
    s.Num = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetPointFee(v int64) *TaobaoTradeFullinfoGetTrade {
    s.PointFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetRealPointFee(v int64) *TaobaoTradeFullinfoGetTrade {
    s.RealPointFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetBuyerObtainPointFee(v int64) *TaobaoTradeFullinfoGetTrade {
    s.BuyerObtainPointFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetBuyerAlipayNo(v string) *TaobaoTradeFullinfoGetTrade {
    s.BuyerAlipayNo = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetReceiverName(v string) *TaobaoTradeFullinfoGetTrade {
    s.ReceiverName = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetReceiverCountry(v string) *TaobaoTradeFullinfoGetTrade {
    s.ReceiverCountry = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetReceiverState(v string) *TaobaoTradeFullinfoGetTrade {
    s.ReceiverState = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetReceiverCity(v string) *TaobaoTradeFullinfoGetTrade {
    s.ReceiverCity = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetReceiverDistrict(v string) *TaobaoTradeFullinfoGetTrade {
    s.ReceiverDistrict = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetReceiverTown(v string) *TaobaoTradeFullinfoGetTrade {
    s.ReceiverTown = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetReceiverAddress(v string) *TaobaoTradeFullinfoGetTrade {
    s.ReceiverAddress = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetReceiverZip(v string) *TaobaoTradeFullinfoGetTrade {
    s.ReceiverZip = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetReceiverMobile(v string) *TaobaoTradeFullinfoGetTrade {
    s.ReceiverMobile = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetReceiverPhone(v string) *TaobaoTradeFullinfoGetTrade {
    s.ReceiverPhone = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetBuyerEmail(v string) *TaobaoTradeFullinfoGetTrade {
    s.BuyerEmail = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetSellerAlipayNo(v string) *TaobaoTradeFullinfoGetTrade {
    s.SellerAlipayNo = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetSellerMobile(v string) *TaobaoTradeFullinfoGetTrade {
    s.SellerMobile = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetSellerPhone(v string) *TaobaoTradeFullinfoGetTrade {
    s.SellerPhone = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetSellerName(v string) *TaobaoTradeFullinfoGetTrade {
    s.SellerName = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetSellerEmail(v string) *TaobaoTradeFullinfoGetTrade {
    s.SellerEmail = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetAvailableConfirmFee(v string) *TaobaoTradeFullinfoGetTrade {
    s.AvailableConfirmFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetHasPostFee(v bool) *TaobaoTradeFullinfoGetTrade {
    s.HasPostFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetTimeoutActionTime(v util.LocalTime) *TaobaoTradeFullinfoGetTrade {
    s.TimeoutActionTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetSnapshotUrl(v string) *TaobaoTradeFullinfoGetTrade {
    s.SnapshotUrl = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetTradeMemo(v string) *TaobaoTradeFullinfoGetTrade {
    s.TradeMemo = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetPromotion(v string) *TaobaoTradeFullinfoGetTrade {
    s.Promotion = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetIs3D(v bool) *TaobaoTradeFullinfoGetTrade {
    s.Is3D = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetIsLgtype(v bool) *TaobaoTradeFullinfoGetTrade {
    s.IsLgtype = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetIsBrandSale(v bool) *TaobaoTradeFullinfoGetTrade {
    s.IsBrandSale = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetIsForceWlb(v bool) *TaobaoTradeFullinfoGetTrade {
    s.IsForceWlb = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetBuyerFlag(v int64) *TaobaoTradeFullinfoGetTrade {
    s.BuyerFlag = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetSellerFlag(v int64) *TaobaoTradeFullinfoGetTrade {
    s.SellerFlag = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetCreditCardFee(v string) *TaobaoTradeFullinfoGetTrade {
    s.CreditCardFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetStepTradeStatus(v string) *TaobaoTradeFullinfoGetTrade {
    s.StepTradeStatus = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetStepPaidFee(v string) *TaobaoTradeFullinfoGetTrade {
    s.StepPaidFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetEticketExt(v string) *TaobaoTradeFullinfoGetTrade {
    s.EticketExt = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetMarkDesc(v string) *TaobaoTradeFullinfoGetTrade {
    s.MarkDesc = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetHasYfx(v bool) *TaobaoTradeFullinfoGetTrade {
    s.HasYfx = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetYfxFee(v string) *TaobaoTradeFullinfoGetTrade {
    s.YfxFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetYfxId(v string) *TaobaoTradeFullinfoGetTrade {
    s.YfxId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetYfxType(v string) *TaobaoTradeFullinfoGetTrade {
    s.YfxType = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetSendTime(v string) *TaobaoTradeFullinfoGetTrade {
    s.SendTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetCanRate(v bool) *TaobaoTradeFullinfoGetTrade {
    s.CanRate = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetSellerCanRate(v bool) *TaobaoTradeFullinfoGetTrade {
    s.SellerCanRate = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetIsPartConsign(v bool) *TaobaoTradeFullinfoGetTrade {
    s.IsPartConsign = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetIsDaixiao(v bool) *TaobaoTradeFullinfoGetTrade {
    s.IsDaixiao = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetIsWt(v bool) *TaobaoTradeFullinfoGetTrade {
    s.IsWt = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetArriveInterval(v int64) *TaobaoTradeFullinfoGetTrade {
    s.ArriveInterval = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetArriveCutTime(v string) *TaobaoTradeFullinfoGetTrade {
    s.ArriveCutTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetConsignInterval(v int64) *TaobaoTradeFullinfoGetTrade {
    s.ConsignInterval = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetO2o(v string) *TaobaoTradeFullinfoGetTrade {
    s.O2o = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetO2oGuideId(v string) *TaobaoTradeFullinfoGetTrade {
    s.O2oGuideId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetO2oGuideName(v string) *TaobaoTradeFullinfoGetTrade {
    s.O2oGuideName = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetO2oShopId(v string) *TaobaoTradeFullinfoGetTrade {
    s.O2oShopId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetO2oShopName(v string) *TaobaoTradeFullinfoGetTrade {
    s.O2oShopName = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetO2oDelivery(v string) *TaobaoTradeFullinfoGetTrade {
    s.O2oDelivery = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetO2oOutTradeId(v string) *TaobaoTradeFullinfoGetTrade {
    s.O2oOutTradeId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetShopCode(v string) *TaobaoTradeFullinfoGetTrade {
    s.ShopCode = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetHkEnName(v string) *TaobaoTradeFullinfoGetTrade {
    s.HkEnName = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetHkFlightNo(v string) *TaobaoTradeFullinfoGetTrade {
    s.HkFlightNo = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetHkChinaName(v string) *TaobaoTradeFullinfoGetTrade {
    s.HkChinaName = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetHkCardCode(v string) *TaobaoTradeFullinfoGetTrade {
    s.HkCardCode = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetHkCardType(v string) *TaobaoTradeFullinfoGetTrade {
    s.HkCardType = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetHkFlightDate(v string) *TaobaoTradeFullinfoGetTrade {
    s.HkFlightDate = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetHkGender(v string) *TaobaoTradeFullinfoGetTrade {
    s.HkGender = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetHkBirthday(v string) *TaobaoTradeFullinfoGetTrade {
    s.HkBirthday = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetHkPickup(v string) *TaobaoTradeFullinfoGetTrade {
    s.HkPickup = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetHkPickupId(v string) *TaobaoTradeFullinfoGetTrade {
    s.HkPickupId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetEstConTime(v string) *TaobaoTradeFullinfoGetTrade {
    s.EstConTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetOrders(v []TaobaoTradeFullinfoGetOrder) *TaobaoTradeFullinfoGetTrade {
    s.Orders = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetPromotionDetails(v []TaobaoTradeFullinfoGetPromotionDetail) *TaobaoTradeFullinfoGetTrade {
    s.PromotionDetails = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetServiceTags(v []TaobaoTradeFullinfoGetLogisticsTag) *TaobaoTradeFullinfoGetTrade {
    s.ServiceTags = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetServiceOrders(v []TaobaoTradeFullinfoGetServiceOrder) *TaobaoTradeFullinfoGetTrade {
    s.ServiceOrders = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetTradeFrom(v string) *TaobaoTradeFullinfoGetTrade {
    s.TradeFrom = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetTradeSource(v string) *TaobaoTradeFullinfoGetTrade {
    s.TradeSource = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetZeroPurchase(v bool) *TaobaoTradeFullinfoGetTrade {
    s.ZeroPurchase = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetAlipayPoint(v int64) *TaobaoTradeFullinfoGetTrade {
    s.AlipayPoint = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetPccAf(v int64) *TaobaoTradeFullinfoGetTrade {
    s.PccAf = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetTradeExt(v TaobaoTradeFullinfoGetTradeExt) *TaobaoTradeFullinfoGetTrade {
    s.TradeExt = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetOrderTaxFee(v string) *TaobaoTradeFullinfoGetTrade {
    s.OrderTaxFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetEtSerTime(v string) *TaobaoTradeFullinfoGetTrade {
    s.EtSerTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetEtShopName(v string) *TaobaoTradeFullinfoGetTrade {
    s.EtShopName = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetEtVerifiedShopName(v string) *TaobaoTradeFullinfoGetTrade {
    s.EtVerifiedShopName = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetEtPlateNumber(v string) *TaobaoTradeFullinfoGetTrade {
    s.EtPlateNumber = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetO2oSnatchStatus(v string) *TaobaoTradeFullinfoGetTrade {
    s.O2oSnatchStatus = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetIsShShip(v bool) *TaobaoTradeFullinfoGetTrade {
    s.IsShShip = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetEticketServiceAddr(v string) *TaobaoTradeFullinfoGetTrade {
    s.EticketServiceAddr = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetEtType(v string) *TaobaoTradeFullinfoGetTrade {
    s.EtType = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetMarket(v string) *TaobaoTradeFullinfoGetTrade {
    s.Market = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetObs(v string) *TaobaoTradeFullinfoGetTrade {
    s.Obs = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetEtShopId(v int64) *TaobaoTradeFullinfoGetTrade {
    s.EtShopId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetPaidCouponFee(v string) *TaobaoTradeFullinfoGetTrade {
    s.PaidCouponFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetShopPick(v string) *TaobaoTradeFullinfoGetTrade {
    s.ShopPick = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetRxAuditStatus(v string) *TaobaoTradeFullinfoGetTrade {
    s.RxAuditStatus = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetEsDate(v string) *TaobaoTradeFullinfoGetTrade {
    s.EsDate = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetEsRange(v string) *TaobaoTradeFullinfoGetTrade {
    s.EsRange = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetOsDate(v string) *TaobaoTradeFullinfoGetTrade {
    s.OsDate = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetOsRange(v string) *TaobaoTradeFullinfoGetTrade {
    s.OsRange = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetCouponFee(v int64) *TaobaoTradeFullinfoGetTrade {
    s.CouponFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetPostGateDeclare(v bool) *TaobaoTradeFullinfoGetTrade {
    s.PostGateDeclare = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetCrossBondedDeclare(v bool) *TaobaoTradeFullinfoGetTrade {
    s.CrossBondedDeclare = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetTradeAttr(v string) *TaobaoTradeFullinfoGetTrade {
    s.TradeAttr = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetTopHold(v int64) *TaobaoTradeFullinfoGetTrade {
    s.TopHold = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetOmniAttr(v string) *TaobaoTradeFullinfoGetTrade {
    s.OmniAttr = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetOmniParam(v string) *TaobaoTradeFullinfoGetTrade {
    s.OmniParam = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetAssembly(v string) *TaobaoTradeFullinfoGetTrade {
    s.Assembly = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetForbidConsign(v int64) *TaobaoTradeFullinfoGetTrade {
    s.ForbidConsign = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetIdentity(v string) *TaobaoTradeFullinfoGetTrade {
    s.Identity = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetTeamBuyHold(v int64) *TaobaoTradeFullinfoGetTrade {
    s.TeamBuyHold = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetShareGroupHold(v int64) *TaobaoTradeFullinfoGetTrade {
    s.ShareGroupHold = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetOfpHold(v int64) *TaobaoTradeFullinfoGetTrade {
    s.OfpHold = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetO2oStepTradeDetail(v string) *TaobaoTradeFullinfoGetTrade {
    s.O2oStepTradeDetail = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetO2oStepOrderId(v string) *TaobaoTradeFullinfoGetTrade {
    s.O2oStepOrderId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetO2oEtOrderId(v string) *TaobaoTradeFullinfoGetTrade {
    s.O2oEtOrderId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetO2oVoucherPrice(v string) *TaobaoTradeFullinfoGetTrade {
    s.O2oVoucherPrice = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetOrderTaxPromotionFee(v string) *TaobaoTradeFullinfoGetTrade {
    s.OrderTaxPromotionFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetDelayCreateDelivery(v int64) *TaobaoTradeFullinfoGetTrade {
    s.DelayCreateDelivery = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetTidStr(v string) *TaobaoTradeFullinfoGetTrade {
    s.TidStr = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetToptype(v int64) *TaobaoTradeFullinfoGetTrade {
    s.Toptype = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetServiceType(v string) *TaobaoTradeFullinfoGetTrade {
    s.ServiceType = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetO2oServiceMobile(v string) *TaobaoTradeFullinfoGetTrade {
    s.O2oServiceMobile = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetO2oServiceName(v string) *TaobaoTradeFullinfoGetTrade {
    s.O2oServiceName = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetO2oServiceState(v string) *TaobaoTradeFullinfoGetTrade {
    s.O2oServiceState = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetO2oServiceCity(v string) *TaobaoTradeFullinfoGetTrade {
    s.O2oServiceCity = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetO2oServiceDistrict(v string) *TaobaoTradeFullinfoGetTrade {
    s.O2oServiceDistrict = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetO2oServiceTown(v string) *TaobaoTradeFullinfoGetTrade {
    s.O2oServiceTown = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetO2oServiceAddress(v string) *TaobaoTradeFullinfoGetTrade {
    s.O2oServiceAddress = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetO2oStepTradeDetailNew(v string) *TaobaoTradeFullinfoGetTrade {
    s.O2oStepTradeDetailNew = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetO2oXiaopiao(v string) *TaobaoTradeFullinfoGetTrade {
    s.O2oXiaopiao = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetO2oContract(v string) *TaobaoTradeFullinfoGetTrade {
    s.O2oContract = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetRechargeFee(v string) *TaobaoTradeFullinfoGetTrade {
    s.RechargeFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetRetailStoreCode(v string) *TaobaoTradeFullinfoGetTrade {
    s.RetailStoreCode = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetRetailOutOrderId(v string) *TaobaoTradeFullinfoGetTrade {
    s.RetailOutOrderId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetPlatformSubsidyFee(v string) *TaobaoTradeFullinfoGetTrade {
    s.PlatformSubsidyFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetNrOffline(v string) *TaobaoTradeFullinfoGetTrade {
    s.NrOffline = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetWttParam(v string) *TaobaoTradeFullinfoGetTrade {
    s.WttParam = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetLogisticsInfos(v []TaobaoTradeFullinfoGetLogisticsInfo) *TaobaoTradeFullinfoGetTrade {
    s.LogisticsInfos = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetSellerNick(v string) *TaobaoTradeFullinfoGetTrade {
    s.SellerNick = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetBuyerNick(v string) *TaobaoTradeFullinfoGetTrade {
    s.BuyerNick = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetNrStoreOrderId(v string) *TaobaoTradeFullinfoGetTrade {
    s.NrStoreOrderId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetNrShopId(v string) *TaobaoTradeFullinfoGetTrade {
    s.NrShopId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetNrShopName(v string) *TaobaoTradeFullinfoGetTrade {
    s.NrShopName = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetNrShopGuideId(v string) *TaobaoTradeFullinfoGetTrade {
    s.NrShopGuideId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetNrShopGuideName(v string) *TaobaoTradeFullinfoGetTrade {
    s.NrShopGuideName = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetSortInfo(v TaobaoTradeFullinfoGetSortInfo) *TaobaoTradeFullinfoGetTrade {
    s.SortInfo = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetSorted(v int64) *TaobaoTradeFullinfoGetTrade {
    s.Sorted = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetNrNoHandle(v string) *TaobaoTradeFullinfoGetTrade {
    s.NrNoHandle = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetBizCode(v string) *TaobaoTradeFullinfoGetTrade {
    s.BizCode = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetCloudStore(v string) *TaobaoTradeFullinfoGetTrade {
    s.CloudStore = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetIsGift(v bool) *TaobaoTradeFullinfoGetTrade {
    s.IsGift = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetDoneeNick(v string) *TaobaoTradeFullinfoGetTrade {
    s.DoneeNick = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetDoneeOpenUid(v string) *TaobaoTradeFullinfoGetTrade {
    s.DoneeOpenUid = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetSuningShopCode(v string) *TaobaoTradeFullinfoGetTrade {
    s.SuningShopCode = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetSuningShopValid(v int64) *TaobaoTradeFullinfoGetTrade {
    s.SuningShopValid = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetAllowAppkeys(v string) *TaobaoTradeFullinfoGetTrade {
    s.AllowAppkeys = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetNewPresell(v bool) *TaobaoTradeFullinfoGetTrade {
    s.NewPresell = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetYouXiang(v bool) *TaobaoTradeFullinfoGetTrade {
    s.YouXiang = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetRetailStoreId(v string) *TaobaoTradeFullinfoGetTrade {
    s.RetailStoreId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetIsIstore(v bool) *TaobaoTradeFullinfoGetTrade {
    s.IsIstore = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetUa(v string) *TaobaoTradeFullinfoGetTrade {
    s.Ua = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetLinkedmallExtInfo(v string) *TaobaoTradeFullinfoGetTrade {
    s.LinkedmallExtInfo = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetPayChannel(v string) *TaobaoTradeFullinfoGetTrade {
    s.PayChannel = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetRtOmniSendType(v string) *TaobaoTradeFullinfoGetTrade {
    s.RtOmniSendType = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetRtOmniStoreId(v string) *TaobaoTradeFullinfoGetTrade {
    s.RtOmniStoreId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetRtOmniOuterStoreId(v string) *TaobaoTradeFullinfoGetTrade {
    s.RtOmniOuterStoreId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetTcpsStart(v string) *TaobaoTradeFullinfoGetTrade {
    s.TcpsStart = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetTcpsCode(v string) *TaobaoTradeFullinfoGetTrade {
    s.TcpsCode = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetTcpsEnd(v string) *TaobaoTradeFullinfoGetTrade {
    s.TcpsEnd = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetMTariffFee(v string) *TaobaoTradeFullinfoGetTrade {
    s.MTariffFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetTimingPromise(v string) *TaobaoTradeFullinfoGetTrade {
    s.TimingPromise = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetPromiseService(v string) *TaobaoTradeFullinfoGetTrade {
    s.PromiseService = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetCutoffMinutes(v string) *TaobaoTradeFullinfoGetTrade {
    s.CutoffMinutes = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetEsTime(v string) *TaobaoTradeFullinfoGetTrade {
    s.EsTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetDeliveryTime(v string) *TaobaoTradeFullinfoGetTrade {
    s.DeliveryTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetCollectTime(v string) *TaobaoTradeFullinfoGetTrade {
    s.CollectTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetDispatchTime(v string) *TaobaoTradeFullinfoGetTrade {
    s.DispatchTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetSignTime(v string) *TaobaoTradeFullinfoGetTrade {
    s.SignTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetOuterPartnerMemberId(v string) *TaobaoTradeFullinfoGetTrade {
    s.OuterPartnerMemberId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetRootCat(v string) *TaobaoTradeFullinfoGetTrade {
    s.RootCat = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetGifting(v string) *TaobaoTradeFullinfoGetTrade {
    s.Gifting = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetGiftingTakeout(v string) *TaobaoTradeFullinfoGetTrade {
    s.GiftingTakeout = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetOiDate(v string) *TaobaoTradeFullinfoGetTrade {
    s.OiDate = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetOiRange(v string) *TaobaoTradeFullinfoGetTrade {
    s.OiRange = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetHoldInstall(v string) *TaobaoTradeFullinfoGetTrade {
    s.HoldInstall = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetAppName(v string) *TaobaoTradeFullinfoGetTrade {
    s.AppName = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetEasyHomeCityType(v string) *TaobaoTradeFullinfoGetTrade {
    s.EasyHomeCityType = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetNrDepositOrderId(v string) *TaobaoTradeFullinfoGetTrade {
    s.NrDepositOrderId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetNrStoreCode(v string) *TaobaoTradeFullinfoGetTrade {
    s.NrStoreCode = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetPropoint(v string) *TaobaoTradeFullinfoGetTrade {
    s.Propoint = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetZqsOrderTag(v string) *TaobaoTradeFullinfoGetTrade {
    s.ZqsOrderTag = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetTxpFreezerId(v string) *TaobaoTradeFullinfoGetTrade {
    s.TxpFreezerId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetTxpReceiveMethod(v string) *TaobaoTradeFullinfoGetTrade {
    s.TxpReceiveMethod = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetExpandcardInfo(v TaobaoTradeFullinfoGetExpandCardInfo) *TaobaoTradeFullinfoGetTrade {
    s.ExpandcardInfo = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetExtendInfo(v string) *TaobaoTradeFullinfoGetTrade {
    s.ExtendInfo = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetLm(v string) *TaobaoTradeFullinfoGetTrade {
    s.Lm = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetBrandLightShopSource(v string) *TaobaoTradeFullinfoGetTrade {
    s.BrandLightShopSource = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetBrandLightShopStoreId(v string) *TaobaoTradeFullinfoGetTrade {
    s.BrandLightShopStoreId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetIsWmly(v string) *TaobaoTradeFullinfoGetTrade {
    s.IsWmly = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetOmniPackage(v string) *TaobaoTradeFullinfoGetTrade {
    s.OmniPackage = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetNczExtAttr(v string) *TaobaoTradeFullinfoGetTrade {
    s.NczExtAttr = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetInvoiceDetailPay(v string) *TaobaoTradeFullinfoGetTrade {
    s.InvoiceDetailPay = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetInvoiceDetailMidRefund(v string) *TaobaoTradeFullinfoGetTrade {
    s.InvoiceDetailMidRefund = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetInvoiceDetailAfterRefund(v string) *TaobaoTradeFullinfoGetTrade {
    s.InvoiceDetailAfterRefund = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetExpandCardBasicPrice(v string) *TaobaoTradeFullinfoGetTrade {
    s.ExpandCardBasicPrice = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetExpandCardExpandPrice(v string) *TaobaoTradeFullinfoGetTrade {
    s.ExpandCardExpandPrice = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetExpandCardBasicPriceUsed(v string) *TaobaoTradeFullinfoGetTrade {
    s.ExpandCardBasicPriceUsed = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetExpandCardExpandPriceUsed(v string) *TaobaoTradeFullinfoGetTrade {
    s.ExpandCardExpandPriceUsed = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetDeliveryCps(v string) *TaobaoTradeFullinfoGetTrade {
    s.DeliveryCps = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetAsdpBizType(v string) *TaobaoTradeFullinfoGetTrade {
    s.AsdpBizType = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetIsOpenmall(v bool) *TaobaoTradeFullinfoGetTrade {
    s.IsOpenmall = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetVLogisticsCreate(v bool) *TaobaoTradeFullinfoGetTrade {
    s.VLogisticsCreate = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetQRPay(v bool) *TaobaoTradeFullinfoGetTrade {
    s.QRPay = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetOrderFollowId(v string) *TaobaoTradeFullinfoGetTrade {
    s.OrderFollowId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetNoShipping(v bool) *TaobaoTradeFullinfoGetTrade {
    s.NoShipping = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetAsdpAds(v string) *TaobaoTradeFullinfoGetTrade {
    s.AsdpAds = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetObTag(v string) *TaobaoTradeFullinfoGetTrade {
    s.ObTag = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetAid(v string) *TaobaoTradeFullinfoGetTrade {
    s.Aid = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetGeneralNewPresell(v bool) *TaobaoTradeFullinfoGetTrade {
    s.GeneralNewPresell = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetDrugRegister(v string) *TaobaoTradeFullinfoGetTrade {
    s.DrugRegister = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetStepPayDetails(v []TaobaoTradeFullinfoGetStepPayDetail) *TaobaoTradeFullinfoGetTrade {
    s.StepPayDetails = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetAgreeRefundChecks(v []TaobaoTradeFullinfoGetAgreeRefundCheck) *TaobaoTradeFullinfoGetTrade {
    s.AgreeRefundChecks = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetAddressDetails(v []TaobaoTradeFullinfoGetAddressDetail) *TaobaoTradeFullinfoGetTrade {
    s.AddressDetails = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetStageAddressType(v string) *TaobaoTradeFullinfoGetTrade {
    s.StageAddressType = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetOgId(v string) *TaobaoTradeFullinfoGetTrade {
    s.OgId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetPromiseSignTime(v string) *TaobaoTradeFullinfoGetTrade {
    s.PromiseSignTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetB2bDaixiao(v int64) *TaobaoTradeFullinfoGetTrade {
    s.B2bDaixiao = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetTmallCouponFee(v int64) *TaobaoTradeFullinfoGetTrade {
    s.TmallCouponFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetDeliveryPlan(v []TaobaoTradeFullinfoGetDeliveryPlan) *TaobaoTradeFullinfoGetTrade {
    s.DeliveryPlan = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetOmnichannelParam(v string) *TaobaoTradeFullinfoGetTrade {
    s.OmnichannelParam = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetOaid(v string) *TaobaoTradeFullinfoGetTrade {
    s.Oaid = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetOuid(v string) *TaobaoTradeFullinfoGetTrade {
    s.Ouid = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetIsCycleBuy(v bool) *TaobaoTradeFullinfoGetTrade {
    s.IsCycleBuy = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetScenarioGroup(v string) *TaobaoTradeFullinfoGetTrade {
    s.ScenarioGroup = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetPlayType(v string) *TaobaoTradeFullinfoGetTrade {
    s.PlayType = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTrade) SetIsForceDc(v bool) *TaobaoTradeFullinfoGetTrade {
    s.IsForceDc = &v
    return s
}
