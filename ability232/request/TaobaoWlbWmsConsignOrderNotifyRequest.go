package request

import (
        "topsdk/ability232/domain"
        "topsdk/util"
    )

type TaobaoWlbWmsConsignOrderNotifyRequest struct {
    /*
        拓展属性     */
    ExtendFields  *string `json:"extend_fields,omitempty" required:"false" `
    /*
        订单商品信息列表     */
    OrderItemList  *[]domain.TaobaoWlbWmsConsignOrderNotifyOrderitemlistwlbwmsconsignordernotify `json:"order_item_list,omitempty" required:"false" `
    /*
        发票信息列表     */
    InvoiceInfoList  *[]domain.TaobaoWlbWmsConsignOrderNotifyInvoicelistwlbwmsconsignordernotify `json:"invoice_info_list,omitempty" required:"false" `
    /*
        收件人信息     */
    ReceiverInfo  *domain.TaobaoWlbWmsConsignOrderNotifyReceiverwlbwmsconsignordernotify `json:"receiver_info,omitempty" required:"false" `
    /*
        废弃，车牌号     */
    CarNo  *string `json:"car_no,omitempty" required:"false" `
    /*
        发货方信息     */
    SenderInfo  *domain.TaobaoWlbWmsConsignOrderNotifySenderwlbwmsconsignordernotify `json:"sender_info,omitempty" required:"false" `
    /*
        废弃，取件人身份证     */
    PickerId  *string `json:"picker_id,omitempty" required:"false" `
    /*
        废弃，承运商名称     */
    CarriersName  *string `json:"carriers_name,omitempty" required:"false" `
    /*
        废弃，取件人电话     */
    PickerCall  *string `json:"picker_call,omitempty" required:"false" `
    /*
        废弃，取件人姓名     */
    PickerName  *string `json:"picker_name,omitempty" required:"false" `
    /*
        废弃，出库方式（自提，快递，销毁）     */
    TransportMode  *string `json:"transport_mode,omitempty" required:"false" `
    /*
        配送要求     */
    DeliverRequirements  *domain.TaobaoWlbWmsConsignOrderNotifyDeliverrequirementswlbwmsconsignordernotify `json:"deliver_requirements,omitempty" required:"false" `
    /*
        快递公司名称     */
    TmsServiceName  *string `json:"tms_service_name,omitempty" required:"false" `
    /*
        快递公司编码，此字段为空时，由菜鸟选择快递配送     */
    TmsServiceCode  *string `json:"tms_service_code,omitempty" required:"false" `
    /*
        快递费用，单位分     */
    Postfee  *int64 `json:"postfee,omitempty" required:"false" `
    /*
        订单已付金额，消费者已经支付的金额，单位分     */
    GotAmount  *int64 `json:"got_amount,omitempty" required:"false" `
    /*
        订单应收金额，消费者还需要付的金额，单位分     */
    ArAmount  *int64 `json:"ar_amount,omitempty" required:"false" `
    /*
        订单优惠金额，整单优惠金额，单位分     */
    DiscountAmount  *int64 `json:"discount_amount,omitempty" required:"false" `
    /*
        订单总金额,=总商品金额-订单优惠金额+快递费用，单位分     */
    OrderAmount  *int64 `json:"order_amount,omitempty" required:"false" `
    /*
        废弃，支付宝交易号     */
    AlipayNo  *string `json:"alipay_no,omitempty" required:"false" `
    /*
        下单时间，订单在交易平台创建时间     */
    OrderShopCreateTime  *util.LocalTime `json:"order_shop_create_time,omitempty" required:"false" `
    /*
        订单审核时间,ERP创建支付时间     */
    OrderExaminationTime  *util.LocalTime `json:"order_examination_time,omitempty" required:"false" `
    /*
        订单支付时间     */
    OrderPayTime  *util.LocalTime `json:"order_pay_time,omitempty" required:"false" `
    /*
        订单创建时间,ERP创建订单时间     */
    OrderCreateTime  *util.LocalTime `json:"order_create_time,omitempty" required:"false" `
    /*
        废弃，订单优先级0 -10，根据订单作业优先级设置，数字越大，优先级越高     */
    OrderPriority  *int64 `json:"order_priority,omitempty" required:"false" `
    /*
        废弃，交易内部来源，文本透传 WAP(手机);HITAO(嗨淘);TOP(TOP平台);TAOBAO(普通淘宝);JHS(聚划算) 一笔订单可能同时有以上多个标记，则以逗号分隔     */
    OrderSubSource  *string `json:"order_sub_source,omitempty" required:"false" `
    /*
        订单来源（213 天猫，201 淘宝，214 京东，202 1688 阿里中文站 ，203 苏宁在线，204 亚马逊中国，205 当当，208 1号店，207 唯品会，209 国美在线，210 拍拍，206 易贝ebay，211 聚美优品，212 乐蜂网，215 邮乐，216 凡客，217 优购，218 银泰，219 易讯，221 聚尚网，222 蘑菇街，223 POS门店，301 其他）     */
    OrderSource  *int64 `json:"order_source,omitempty" required:"false" `
    /*
        订单标识 (1: cod –货到付款，4:invoiceinfo-需要发票)     */
    OrderFlag  *string `json:"order_flag,omitempty" required:"false" `
    /*
        单据类型 201 一般交易出库单 202 B2B交易出库单 502 换货出库单 503 补发出库单     */
    OrderType  *int64 `json:"order_type" required:"true" `
    /*
        ERP订单号     */
    OrderCode  *string `json:"order_code" required:"true" `
    /*
        仓库编码，此字段为空时，由菜鸟路由仓库发货     */
    StoreCode  *string `json:"store_code,omitempty" required:"false" `
    /*
        备注     */
    Remark  *string `json:"remark,omitempty" required:"false" `
    /*
        前物流订单号，订单类型为502 换货出库单 503 补发出库单时，需求传入此内容     */
    PrevOrderCode  *string `json:"prev_order_code,omitempty" required:"false" `
    /*
        COD服务费，单位分     */
    ServiceFee  *int64 `json:"service_fee,omitempty" required:"false" `
}

func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetExtendFields(v string) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.ExtendFields = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetOrderItemList(v []domain.TaobaoWlbWmsConsignOrderNotifyOrderitemlistwlbwmsconsignordernotify) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.OrderItemList = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetInvoiceInfoList(v []domain.TaobaoWlbWmsConsignOrderNotifyInvoicelistwlbwmsconsignordernotify) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.InvoiceInfoList = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetReceiverInfo(v domain.TaobaoWlbWmsConsignOrderNotifyReceiverwlbwmsconsignordernotify) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.ReceiverInfo = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetCarNo(v string) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.CarNo = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetSenderInfo(v domain.TaobaoWlbWmsConsignOrderNotifySenderwlbwmsconsignordernotify) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.SenderInfo = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetPickerId(v string) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.PickerId = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetCarriersName(v string) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.CarriersName = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetPickerCall(v string) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.PickerCall = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetPickerName(v string) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.PickerName = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetTransportMode(v string) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.TransportMode = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetDeliverRequirements(v domain.TaobaoWlbWmsConsignOrderNotifyDeliverrequirementswlbwmsconsignordernotify) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.DeliverRequirements = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetTmsServiceName(v string) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.TmsServiceName = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetTmsServiceCode(v string) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.TmsServiceCode = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetPostfee(v int64) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.Postfee = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetGotAmount(v int64) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.GotAmount = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetArAmount(v int64) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.ArAmount = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetDiscountAmount(v int64) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.DiscountAmount = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetOrderAmount(v int64) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.OrderAmount = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetAlipayNo(v string) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.AlipayNo = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetOrderShopCreateTime(v util.LocalTime) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.OrderShopCreateTime = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetOrderExaminationTime(v util.LocalTime) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.OrderExaminationTime = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetOrderPayTime(v util.LocalTime) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.OrderPayTime = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetOrderCreateTime(v util.LocalTime) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.OrderCreateTime = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetOrderPriority(v int64) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.OrderPriority = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetOrderSubSource(v string) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.OrderSubSource = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetOrderSource(v int64) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.OrderSource = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetOrderFlag(v string) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.OrderFlag = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetOrderType(v int64) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.OrderType = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetOrderCode(v string) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.OrderCode = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetStoreCode(v string) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.StoreCode = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetRemark(v string) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.Remark = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetPrevOrderCode(v string) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.PrevOrderCode = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyRequest) SetServiceFee(v int64) *TaobaoWlbWmsConsignOrderNotifyRequest {
    s.ServiceFee = &v
    return s
}

func (req *TaobaoWlbWmsConsignOrderNotifyRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ExtendFields != nil) {
        paramMap["extend_fields"] = *req.ExtendFields
    }
    if(req.OrderItemList != nil) {
        paramMap["order_item_list"] = util.ConvertStructList(*req.OrderItemList)
    }
    if(req.InvoiceInfoList != nil) {
        paramMap["invoice_info_list"] = util.ConvertStructList(*req.InvoiceInfoList)
    }
    if(req.ReceiverInfo != nil) {
        paramMap["receiver_info"] = util.ConvertStruct(*req.ReceiverInfo)
    }
    if(req.CarNo != nil) {
        paramMap["car_no"] = *req.CarNo
    }
    if(req.SenderInfo != nil) {
        paramMap["sender_info"] = util.ConvertStruct(*req.SenderInfo)
    }
    if(req.PickerId != nil) {
        paramMap["picker_id"] = *req.PickerId
    }
    if(req.CarriersName != nil) {
        paramMap["carriers_name"] = *req.CarriersName
    }
    if(req.PickerCall != nil) {
        paramMap["picker_call"] = *req.PickerCall
    }
    if(req.PickerName != nil) {
        paramMap["picker_name"] = *req.PickerName
    }
    if(req.TransportMode != nil) {
        paramMap["transport_mode"] = *req.TransportMode
    }
    if(req.DeliverRequirements != nil) {
        paramMap["deliver_requirements"] = util.ConvertStruct(*req.DeliverRequirements)
    }
    if(req.TmsServiceName != nil) {
        paramMap["tms_service_name"] = *req.TmsServiceName
    }
    if(req.TmsServiceCode != nil) {
        paramMap["tms_service_code"] = *req.TmsServiceCode
    }
    if(req.Postfee != nil) {
        paramMap["postfee"] = *req.Postfee
    }
    if(req.GotAmount != nil) {
        paramMap["got_amount"] = *req.GotAmount
    }
    if(req.ArAmount != nil) {
        paramMap["ar_amount"] = *req.ArAmount
    }
    if(req.DiscountAmount != nil) {
        paramMap["discount_amount"] = *req.DiscountAmount
    }
    if(req.OrderAmount != nil) {
        paramMap["order_amount"] = *req.OrderAmount
    }
    if(req.AlipayNo != nil) {
        paramMap["alipay_no"] = *req.AlipayNo
    }
    if(req.OrderShopCreateTime != nil) {
        paramMap["order_shop_create_time"] = *req.OrderShopCreateTime
    }
    if(req.OrderExaminationTime != nil) {
        paramMap["order_examination_time"] = *req.OrderExaminationTime
    }
    if(req.OrderPayTime != nil) {
        paramMap["order_pay_time"] = *req.OrderPayTime
    }
    if(req.OrderCreateTime != nil) {
        paramMap["order_create_time"] = *req.OrderCreateTime
    }
    if(req.OrderPriority != nil) {
        paramMap["order_priority"] = *req.OrderPriority
    }
    if(req.OrderSubSource != nil) {
        paramMap["order_sub_source"] = *req.OrderSubSource
    }
    if(req.OrderSource != nil) {
        paramMap["order_source"] = *req.OrderSource
    }
    if(req.OrderFlag != nil) {
        paramMap["order_flag"] = *req.OrderFlag
    }
    if(req.OrderType != nil) {
        paramMap["order_type"] = *req.OrderType
    }
    if(req.OrderCode != nil) {
        paramMap["order_code"] = *req.OrderCode
    }
    if(req.StoreCode != nil) {
        paramMap["store_code"] = *req.StoreCode
    }
    if(req.Remark != nil) {
        paramMap["remark"] = *req.Remark
    }
    if(req.PrevOrderCode != nil) {
        paramMap["prev_order_code"] = *req.PrevOrderCode
    }
    if(req.ServiceFee != nil) {
        paramMap["service_fee"] = *req.ServiceFee
    }
    return paramMap
}

func (req *TaobaoWlbWmsConsignOrderNotifyRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}