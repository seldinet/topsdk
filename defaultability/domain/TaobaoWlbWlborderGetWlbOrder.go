package domain

import (
        "topsdk/util"
    )

type TaobaoWlbWlborderGetWlbOrder struct {
    /*
        出库或者入库，IN表示入库，OUT表示出库     */
    OperateType  *string `json:"operate_type,omitempty" `

    /*
        订单编码     */
    OrderCode  *string `json:"order_code,omitempty" `

    /*
        订单来源:
产生物流订单的原因，比如:

订单来源:1:TAOBAO;2:EXT;3:ERP;4:WMS     */
    OrderSource  *string `json:"order_source,omitempty" `

    /*
        对应创建物流宝订单top接口中的的out_biz_code字段，主要是用来去重用     */
    OrderSourceCode  *string `json:"order_source_code,omitempty" `

    /*
        1:正常订单: NORMAL
2:退货订单: RETURN
3:换货订单: CHANGE     */
    OrderType  *string `json:"order_type,omitempty" `

    /*
        (1)其它:    OTHER
(2)淘宝交易: TAOBAO
(3)301:调拨: ALLOCATION
(4)401:盘点:CHECK
(5)501:销售采购:PRUCHASE     */
    OrderSubType  *string `json:"order_sub_type,omitempty" `

    /*
        订单总价     */
    TotalAmount  *int64 `json:"total_amount,omitempty" `

    /*
        卖家ID     */
    UserId  *int64 `json:"user_id,omitempty" `

    /*
        卖家NICK     */
    UserNick  *string `json:"user_nick,omitempty" `

    /*
        仓库编码     */
    StoreCode  *string `json:"store_code,omitempty" `

    /*
        发货物流公司编码，STO,YTO,EMS等     */
    TmsTpCode  *string `json:"tms_tp_code,omitempty" `

    /*
        物流状态，
订单已创建：0
订单已取消: -1
订单关闭:-3
下发中: 10
已下发: 11
TMS拒签 :-100
接收方拒签：-200
已发货:100
签收成功:200     */
    OrderStatus  *string `json:"order_status,omitempty" `

    /*
        卖家取消,仓库取消标识     */
    OrderStatusReason  *string `json:"order_status_reason,omitempty" `

    /*
        订单备注     */
    Remark  *string `json:"remark,omitempty" `

    /*
        原订单编码     */
    PrevOrderCode  *string `json:"prev_order_code,omitempty" `

    /*
        系统自动生成     */
    IsCompleted  *bool `json:"is_completed,omitempty" `

    /*
        买家nick     */
    BuyerNick  *string `json:"buyer_nick,omitempty" `

    /*
        接收人姓名     */
    ReceiverName  *string `json:"receiver_name,omitempty" `

    /*
        接收人旺旺     */
    ReceiverWangwang  *string `json:"receiver_wangwang,omitempty" `

    /*
        收件人邮箱     */
    ReceiverMail  *string `json:"receiver_mail,omitempty" `

    /*
        收件人邮编     */
    ReceiverZipCode  *string `json:"receiver_zip_code,omitempty" `

    /*
        接收人手机号码     */
    ReceiverMobile  *string `json:"receiver_mobile,omitempty" `

    /*
        接收人固定电话     */
    ReceiverPhone  *string `json:"receiver_phone,omitempty" `

    /*
        发票信息     */
    InvoiceInfo  *string `json:"invoice_info,omitempty" `

    /*
        第1位:COD,2:限时配送,3:预售,4:需要发票,5:已投诉,第6位:合单,第7位:拆单 第8位：EXCHANGE-换货， 第9位:VISIT-上门 ， 第10位: MODIFYTRANSPORT-是否可改配送方式，第11位：是否物流代理确认发货     */
    OrderFlag  *int64 `json:"order_flag,omitempty" `

    /*
        应收金额     */
    ReceivableAmount  *int64 `json:"receivable_amount,omitempty" `

    /*
        收件人省份     */
    ReceiverProvince  *string `json:"receiver_province,omitempty" `

    /*
        收件人城市     */
    ReceiverCity  *string `json:"receiver_city,omitempty" `

    /*
        区或者县     */
    ReceiverArea  *string `json:"receiver_area,omitempty" `

    /*
        收件人具体地址     */
    ReceiverAddress  *string `json:"receiver_address,omitempty" `

    /*
        cod服务费     */
    ServiceFee  *int64 `json:"service_fee,omitempty" `

    /*
        订单取消状态：
1-取消中； 
2-取消失败；
3-取消完成     */
    CancelOrderStatus  *int64 `json:"cancel_order_status,omitempty" `

    /*
        发件人所在省份     */
    SenderProvince  *string `json:"sender_province,omitempty" `

    /*
        发件人城市     */
    SenderCity  *string `json:"sender_city,omitempty" `

    /*
        发件人所在区     */
    SenderArea  *string `json:"sender_area,omitempty" `

    /*
        发件人地址     */
    SenderAddress  *string `json:"sender_address,omitempty" `

    /*
        发件人姓名     */
    SenderName  *string `json:"sender_name,omitempty" `

    /*
        发件人电子邮箱     */
    SenderEmail  *string `json:"sender_email,omitempty" `

    /*
        发件人联系电话     */
    SenderPhone  *string `json:"sender_phone,omitempty" `

    /*
        发件人移动电话     */
    SenderMobile  *string `json:"sender_mobile,omitempty" `

    /*
        发件人邮编     */
    SenderZipCode  *string `json:"sender_zip_code,omitempty" `

    /*
        配送开始时间通常是HH:MM格式     */
    ScheduleDay  *string `json:"schedule_day,omitempty" `

    /*
        配送结束时间通常是HH:MM格式     */
    ScheduleEnd  *string `json:"schedule_end,omitempty" `

    /*
        计划送达开始时间     */
    ExpectStartTime  *util.LocalTime `json:"expect_start_time,omitempty" `

    /*
        计划送达结束时间     */
    ExpectEndTime  *util.LocalTime `json:"expect_end_time,omitempty" `

    /*
        发货速度  ，
101-当日达，
102-次晨达，
103-次日达     */
    ScheduleSpeed  *int64 `json:"schedule_speed,omitempty" `

    /*
        openuid     */
    OpenUid  *string `json:"open_uid,omitempty" `

}

func (s *TaobaoWlbWlborderGetWlbOrder) SetOperateType(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.OperateType = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetOrderCode(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.OrderCode = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetOrderSource(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.OrderSource = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetOrderSourceCode(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.OrderSourceCode = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetOrderType(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.OrderType = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetOrderSubType(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.OrderSubType = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetTotalAmount(v int64) *TaobaoWlbWlborderGetWlbOrder {
    s.TotalAmount = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetUserId(v int64) *TaobaoWlbWlborderGetWlbOrder {
    s.UserId = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetUserNick(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.UserNick = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetStoreCode(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.StoreCode = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetTmsTpCode(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.TmsTpCode = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetOrderStatus(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.OrderStatus = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetOrderStatusReason(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.OrderStatusReason = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetRemark(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.Remark = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetPrevOrderCode(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.PrevOrderCode = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetIsCompleted(v bool) *TaobaoWlbWlborderGetWlbOrder {
    s.IsCompleted = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetBuyerNick(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.BuyerNick = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetReceiverName(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.ReceiverName = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetReceiverWangwang(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.ReceiverWangwang = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetReceiverMail(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.ReceiverMail = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetReceiverZipCode(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.ReceiverZipCode = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetReceiverMobile(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.ReceiverMobile = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetReceiverPhone(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.ReceiverPhone = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetInvoiceInfo(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.InvoiceInfo = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetOrderFlag(v int64) *TaobaoWlbWlborderGetWlbOrder {
    s.OrderFlag = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetReceivableAmount(v int64) *TaobaoWlbWlborderGetWlbOrder {
    s.ReceivableAmount = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetReceiverProvince(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.ReceiverProvince = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetReceiverCity(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.ReceiverCity = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetReceiverArea(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.ReceiverArea = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetReceiverAddress(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.ReceiverAddress = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetServiceFee(v int64) *TaobaoWlbWlborderGetWlbOrder {
    s.ServiceFee = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetCancelOrderStatus(v int64) *TaobaoWlbWlborderGetWlbOrder {
    s.CancelOrderStatus = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetSenderProvince(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.SenderProvince = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetSenderCity(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.SenderCity = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetSenderArea(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.SenderArea = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetSenderAddress(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.SenderAddress = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetSenderName(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.SenderName = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetSenderEmail(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.SenderEmail = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetSenderPhone(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.SenderPhone = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetSenderMobile(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.SenderMobile = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetSenderZipCode(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.SenderZipCode = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetScheduleDay(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.ScheduleDay = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetScheduleEnd(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.ScheduleEnd = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetExpectStartTime(v util.LocalTime) *TaobaoWlbWlborderGetWlbOrder {
    s.ExpectStartTime = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetExpectEndTime(v util.LocalTime) *TaobaoWlbWlborderGetWlbOrder {
    s.ExpectEndTime = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetScheduleSpeed(v int64) *TaobaoWlbWlborderGetWlbOrder {
    s.ScheduleSpeed = &v
    return s
}
func (s *TaobaoWlbWlborderGetWlbOrder) SetOpenUid(v string) *TaobaoWlbWlborderGetWlbOrder {
    s.OpenUid = &v
    return s
}
