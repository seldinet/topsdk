package request

import (
        "topsdk/util"
    )

type TaobaoWlbOrderCreateRequest struct {
    /*
        该字段暂时保留     */
    Attributes  *string `json:"attributes,omitempty" required:"false" `
    /*
        订单商品列表： {"order_item_list":[{"trade_code":"可选,淘宝交易订单，并且不是赠品，必须要传订单来源编号"," sub_trade_code ":"可选,淘宝子交易号","item_id":"必须,商品Id","item_code":"必须,商家编码","item_name":"可选,物流宝商品名称","item_quantity":"必选,计划数量","item_price":"必选,物品价格,单位为分","owner_user_nick
":"可选,货主nick 代销模式下会存在","flag":"判断是否为赠品0 不是1是","remarks":"可选,备注","batch_remark":"可选，批次描述信息会把这个信息带给WMS，但不会跟物流宝库存相关联"，"inventory_type":"库存类型1 可销售库存 101 类型用来定义残次品 201 冻结类型库存 301 在途库存","picture_url":"图片Url","distributor_user_nick": "分销商NICK",必选"ext_order_item_code":"可选，外部商品的商家编码"]} ======================================== 如果订单中的商品条目数大于50条的时候，我们会校验，不能创建成功，需要你按照50个一批的数量传，需要分批调用该接口，第二次传的时候，需要带上wlb_order_code和is_finished和order_item_list三个字段是必传的，is_finished为true表示传输完毕，为false表示还没完全传完。     */
    OrderItemList  *string `json:"order_item_list" required:"true" `
    /*
        用字符串格式来表示订单标记列表：比如COD^PRESELL^SPLIT^LIMIT 等，中间用“^”来隔开 ---------------------------------------- 订单标记list（所有字母全部大写）： 1: COD –货到付款 2: LIMIT-限时配送 3: PRESELL-预售 5:COMPLAIN-已投诉 7:SPLIT-拆单， 8:EXCHANGE-换货， 9:VISIT-上门 ， 10: MODIFYTRANSPORT-是否可改配送方式，
: 是否可改配送方式  默认可更改
11 CONSIGN 物流宝代理发货,自动更改发货状态
12: SELLER_AFFORD 是否卖家承担运费 默认是，即没 13: SYNC_RETURN_BILL，同时退回发票     */
    OrderFlag  *string `json:"order_flag,omitempty" required:"false" `
    /*
        订单类型:  （1）NORMAL_OUT ：正常出库  （2）NORMAL_IN：正常入库  （3）RETURN_IN：退货入库  （4）EXCHANGE_OUT：换货出库     */
    OrderType  *string `json:"order_type" required:"true" `
    /*
        订单子类型：  （1）OTHER： 其他  （2）TAOBAO_TRADE： 淘宝交易  （3）OTHER_TRADE：其他交易  （4）ALLCOATE： 调拨  （5）PURCHASE:采购     */
    OrderSubType  *string `json:"order_sub_type" required:"true" `
    /*
        {"invoince_info": [{"bill_type":"发票类型，必选", "bill_title":"发票抬头，必选", "bill_amount":"发票金额(单位是分)，必选","bill_content":"发票内容，可选"}]}     */
    InvoinceInfo  *string `json:"invoince_info,omitempty" required:"false" `
    /*
        包裹件数，入库单和出库单中会用到     */
    PackageCount  *int64 `json:"package_count,omitempty" required:"false" `
    /*
        出库单中可能会用到
运输公司名称^^^运输公司联系人^^^运输公司运单号^^^运输公司电话^^^运输公司联系人身份证号

========================================
如果某一个字段的数据为空时，必须传NA     */
    TmsInfo  *string `json:"tms_info,omitempty" required:"false" `
    /*
        支付宝交易号     */
    AlipayNo  *string `json:"alipay_no,omitempty" required:"false" `
    /*
        投递时延要求:  （1）INSTANT_ARRIVED： 当日达  （2）TOMMORROY_MORNING_ARRIVED：次晨达  （3）TOMMORROY_ARRIVED：次日达  （4）工作日：WORK_DAY  （5）节假日：WEEKED_DAY     */
    ScheduleType  *string `json:"schedule_type,omitempty" required:"false" `
    /*
        收货方信息，必须传， 手机和电话必选其一。
收货方信息：
邮编^^^省^^^市^^^区^^^具体地址^^^收件方名称^^^手机^^^电话

如果某一个字段的数据为空时，必须传NA     */
    ReceiverInfo  *string `json:"receiver_info,omitempty" required:"false" `
    /*
        发货方信息，发货方信息必须传， 手机和电话必选其一。 发货方信息：
邮编^^^省^^^市^^^区^^^具体地址^^^收件方名称^^^手机^^^电话
如果某一个字段的数据为空时，必须传NA     */
    SenderInfo  *string `json:"sender_info,omitempty" required:"false" `
    /*
        仓库编码     */
    StoreCode  *string `json:"store_code" required:"true" `
    /*
        外部订单业务ID，该编号在isv中是唯一编号， 用来控制并发，去重用     */
    OutBizCode  *string `json:"out_biz_code" required:"true" `
    /*
        源订单编号     */
    PrevOrderCode  *string `json:"prev_order_code,omitempty" required:"false" `
    /*
        物流公司编码     */
    TmsServiceCode  *string `json:"tms_service_code,omitempty" required:"false" `
    /*
        总金额     */
    TotalAmount  *int64 `json:"total_amount,omitempty" required:"false" `
    /*
        应收金额，cod订单必选     */
    PayableAmount  *int64 `json:"payable_amount,omitempty" required:"false" `
    /*
        备注     */
    Remark  *string `json:"remark,omitempty" required:"false" `
    /*
        买家呢称     */
    BuyerNick  *string `json:"buyer_nick,omitempty" required:"false" `
    /*
        计划开始送达时间  在入库单中可能会使用     */
    ExpectStartTime  *util.LocalTime `json:"expect_start_time,omitempty" required:"false" `
    /*
        期望结束时间，在入库单会使用到     */
    ExpectEndTime  *util.LocalTime `json:"expect_end_time,omitempty" required:"false" `
    /*
        投递时间范围要求,格式'13:20'用分号隔开     */
    ScheduleStart  *string `json:"schedule_start,omitempty" required:"false" `
    /*
        投递时间范围要求,格式'15:20'用分号隔开     */
    ScheduleEnd  *string `json:"schedule_end,omitempty" required:"false" `
    /*
        cod服务费，只有cod订单的时候，才需要这个字段     */
    ServiceFee  *int64 `json:"service_fee,omitempty" required:"false" `
    /*
        物流宝订单编号，该接口约定每次最多只能传50条order_item_list，如果一个物流宝订单超过50条商品的时候，需要批量来调用该接口，第一次调用的时候，wlb_order_code为空，如果第一次创建成功，该接口返回wlb_order_code，其后继续再该订单上添加商品条目，需要带上wlb_order_code，out_biz_code，order_item_list,is_finished四个字段。     */
    OrderCode  *string `json:"order_code,omitempty" required:"false" `
    /*
        该物流宝订单是否已完成，如果完成则设置为true，如果为false，则需要等待继续创建订单商品信息。     */
    IsFinished  *bool `json:"is_finished" required:"true" `
    /*
        运单编号，退货单时可能会使用     */
    TmsOrderCode  *string `json:"tms_order_code,omitempty" required:"false" `
}

func (s *TaobaoWlbOrderCreateRequest) SetAttributes(v string) *TaobaoWlbOrderCreateRequest {
    s.Attributes = &v
    return s
}
func (s *TaobaoWlbOrderCreateRequest) SetOrderItemList(v string) *TaobaoWlbOrderCreateRequest {
    s.OrderItemList = &v
    return s
}
func (s *TaobaoWlbOrderCreateRequest) SetOrderFlag(v string) *TaobaoWlbOrderCreateRequest {
    s.OrderFlag = &v
    return s
}
func (s *TaobaoWlbOrderCreateRequest) SetOrderType(v string) *TaobaoWlbOrderCreateRequest {
    s.OrderType = &v
    return s
}
func (s *TaobaoWlbOrderCreateRequest) SetOrderSubType(v string) *TaobaoWlbOrderCreateRequest {
    s.OrderSubType = &v
    return s
}
func (s *TaobaoWlbOrderCreateRequest) SetInvoinceInfo(v string) *TaobaoWlbOrderCreateRequest {
    s.InvoinceInfo = &v
    return s
}
func (s *TaobaoWlbOrderCreateRequest) SetPackageCount(v int64) *TaobaoWlbOrderCreateRequest {
    s.PackageCount = &v
    return s
}
func (s *TaobaoWlbOrderCreateRequest) SetTmsInfo(v string) *TaobaoWlbOrderCreateRequest {
    s.TmsInfo = &v
    return s
}
func (s *TaobaoWlbOrderCreateRequest) SetAlipayNo(v string) *TaobaoWlbOrderCreateRequest {
    s.AlipayNo = &v
    return s
}
func (s *TaobaoWlbOrderCreateRequest) SetScheduleType(v string) *TaobaoWlbOrderCreateRequest {
    s.ScheduleType = &v
    return s
}
func (s *TaobaoWlbOrderCreateRequest) SetReceiverInfo(v string) *TaobaoWlbOrderCreateRequest {
    s.ReceiverInfo = &v
    return s
}
func (s *TaobaoWlbOrderCreateRequest) SetSenderInfo(v string) *TaobaoWlbOrderCreateRequest {
    s.SenderInfo = &v
    return s
}
func (s *TaobaoWlbOrderCreateRequest) SetStoreCode(v string) *TaobaoWlbOrderCreateRequest {
    s.StoreCode = &v
    return s
}
func (s *TaobaoWlbOrderCreateRequest) SetOutBizCode(v string) *TaobaoWlbOrderCreateRequest {
    s.OutBizCode = &v
    return s
}
func (s *TaobaoWlbOrderCreateRequest) SetPrevOrderCode(v string) *TaobaoWlbOrderCreateRequest {
    s.PrevOrderCode = &v
    return s
}
func (s *TaobaoWlbOrderCreateRequest) SetTmsServiceCode(v string) *TaobaoWlbOrderCreateRequest {
    s.TmsServiceCode = &v
    return s
}
func (s *TaobaoWlbOrderCreateRequest) SetTotalAmount(v int64) *TaobaoWlbOrderCreateRequest {
    s.TotalAmount = &v
    return s
}
func (s *TaobaoWlbOrderCreateRequest) SetPayableAmount(v int64) *TaobaoWlbOrderCreateRequest {
    s.PayableAmount = &v
    return s
}
func (s *TaobaoWlbOrderCreateRequest) SetRemark(v string) *TaobaoWlbOrderCreateRequest {
    s.Remark = &v
    return s
}
func (s *TaobaoWlbOrderCreateRequest) SetBuyerNick(v string) *TaobaoWlbOrderCreateRequest {
    s.BuyerNick = &v
    return s
}
func (s *TaobaoWlbOrderCreateRequest) SetExpectStartTime(v util.LocalTime) *TaobaoWlbOrderCreateRequest {
    s.ExpectStartTime = &v
    return s
}
func (s *TaobaoWlbOrderCreateRequest) SetExpectEndTime(v util.LocalTime) *TaobaoWlbOrderCreateRequest {
    s.ExpectEndTime = &v
    return s
}
func (s *TaobaoWlbOrderCreateRequest) SetScheduleStart(v string) *TaobaoWlbOrderCreateRequest {
    s.ScheduleStart = &v
    return s
}
func (s *TaobaoWlbOrderCreateRequest) SetScheduleEnd(v string) *TaobaoWlbOrderCreateRequest {
    s.ScheduleEnd = &v
    return s
}
func (s *TaobaoWlbOrderCreateRequest) SetServiceFee(v int64) *TaobaoWlbOrderCreateRequest {
    s.ServiceFee = &v
    return s
}
func (s *TaobaoWlbOrderCreateRequest) SetOrderCode(v string) *TaobaoWlbOrderCreateRequest {
    s.OrderCode = &v
    return s
}
func (s *TaobaoWlbOrderCreateRequest) SetIsFinished(v bool) *TaobaoWlbOrderCreateRequest {
    s.IsFinished = &v
    return s
}
func (s *TaobaoWlbOrderCreateRequest) SetTmsOrderCode(v string) *TaobaoWlbOrderCreateRequest {
    s.TmsOrderCode = &v
    return s
}

func (req *TaobaoWlbOrderCreateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Attributes != nil) {
        paramMap["attributes"] = *req.Attributes
    }
    if(req.OrderItemList != nil) {
        paramMap["order_item_list"] = *req.OrderItemList
    }
    if(req.OrderFlag != nil) {
        paramMap["order_flag"] = *req.OrderFlag
    }
    if(req.OrderType != nil) {
        paramMap["order_type"] = *req.OrderType
    }
    if(req.OrderSubType != nil) {
        paramMap["order_sub_type"] = *req.OrderSubType
    }
    if(req.InvoinceInfo != nil) {
        paramMap["invoince_info"] = *req.InvoinceInfo
    }
    if(req.PackageCount != nil) {
        paramMap["package_count"] = *req.PackageCount
    }
    if(req.TmsInfo != nil) {
        paramMap["tms_info"] = *req.TmsInfo
    }
    if(req.AlipayNo != nil) {
        paramMap["alipay_no"] = *req.AlipayNo
    }
    if(req.ScheduleType != nil) {
        paramMap["schedule_type"] = *req.ScheduleType
    }
    if(req.ReceiverInfo != nil) {
        paramMap["receiver_info"] = *req.ReceiverInfo
    }
    if(req.SenderInfo != nil) {
        paramMap["sender_info"] = *req.SenderInfo
    }
    if(req.StoreCode != nil) {
        paramMap["store_code"] = *req.StoreCode
    }
    if(req.OutBizCode != nil) {
        paramMap["out_biz_code"] = *req.OutBizCode
    }
    if(req.PrevOrderCode != nil) {
        paramMap["prev_order_code"] = *req.PrevOrderCode
    }
    if(req.TmsServiceCode != nil) {
        paramMap["tms_service_code"] = *req.TmsServiceCode
    }
    if(req.TotalAmount != nil) {
        paramMap["total_amount"] = *req.TotalAmount
    }
    if(req.PayableAmount != nil) {
        paramMap["payable_amount"] = *req.PayableAmount
    }
    if(req.Remark != nil) {
        paramMap["remark"] = *req.Remark
    }
    if(req.BuyerNick != nil) {
        paramMap["buyer_nick"] = *req.BuyerNick
    }
    if(req.ExpectStartTime != nil) {
        paramMap["expect_start_time"] = *req.ExpectStartTime
    }
    if(req.ExpectEndTime != nil) {
        paramMap["expect_end_time"] = *req.ExpectEndTime
    }
    if(req.ScheduleStart != nil) {
        paramMap["schedule_start"] = *req.ScheduleStart
    }
    if(req.ScheduleEnd != nil) {
        paramMap["schedule_end"] = *req.ScheduleEnd
    }
    if(req.ServiceFee != nil) {
        paramMap["service_fee"] = *req.ServiceFee
    }
    if(req.OrderCode != nil) {
        paramMap["order_code"] = *req.OrderCode
    }
    if(req.IsFinished != nil) {
        paramMap["is_finished"] = *req.IsFinished
    }
    if(req.TmsOrderCode != nil) {
        paramMap["tms_order_code"] = *req.TmsOrderCode
    }
    return paramMap
}

func (req *TaobaoWlbOrderCreateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}