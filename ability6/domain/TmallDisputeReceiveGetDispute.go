package domain

import (
        "topsdk/util"
    )

type TmallDisputeReceiveGetDispute struct {
    /*
        卖家收货地址     */
    Address  *string `json:"address,omitempty" `

    /*
        支付宝单号     */
    AlipayNo  *string `json:"alipay_no,omitempty" `

    /*
        纠纷单上的各项属性     */
    Attribute  *string `json:"attribute,omitempty" `

    /*
        正向交易单号     */
    BizOrderId  *string `json:"biz_order_id,omitempty" `

    /*
        买家收货地址（换货）     */
    BuyerAddress  *string `json:"buyer_address,omitempty" `

    /*
        买家发货物流公司（换货）     */
    BuyerLogisticName  *string `json:"buyer_logistic_name,omitempty" `

    /*
        买家发货物流单号（换货）     */
    BuyerLogisticNo  *string `json:"buyer_logistic_no,omitempty" `

    /*
        买家收货人姓名（换货）     */
    BuyerName  *string `json:"buyer_name,omitempty" `

    /*
        买家昵称     */
    BuyerNick  *string `json:"buyer_nick,omitempty" `

    /*
        买家联系方式（换货）     */
    BuyerPhone  *string `json:"buyer_phone,omitempty" `

    /*
        卖家发货物流单号     */
    CompanyName  *string `json:"company_name,omitempty" `

    /*
        纠纷单创建时间     */
    Created  *util.LocalTime `json:"created,omitempty" `

    /*
        纠纷说明     */
    Desc  *string `json:"desc,omitempty" `

    /*
        纠纷类型，常见的有：仅退款、退货退款、换货     */
    DisputeRequest  *string `json:"dispute_request,omitempty" `

    /*
        买家退货时间     */
    GoodReturnTime  *util.LocalTime `json:"good_return_time,omitempty" `

    /*
        货物状态     */
    GoodStatus  *string `json:"good_status,omitempty" `

    /*
        买家是否需要退货。可选值:true(是),false(否)     */
    HasGoodReturn  *bool `json:"has_good_return,omitempty" `

    /*
        纠纷单修改时间     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

    /*
        买家购买数量     */
    Num  *int64 `json:"num,omitempty" `

    /*
        子订单号。如果是单笔交易oid会等于tid     */
    Oid  *int64 `json:"oid,omitempty" `

    /*
        退款对应的订单交易状态。可选值TRADE_NO_CREATE_PAY(没有创建支付宝交易) WAIT_BUYER_PAY(等待买家付款) WAIT_SELLER_SEND_GOODS(等待卖家发货,即:买家已付款) WAIT_BUYER_CONFIRM_GOODS(等待买家确认收货,即:卖家已发货) TRADE_BUYER_SIGNED(买家已签收,货到付款专用) TRADE_FINISHED(交易成功) TRADE_CLOSED(交易关闭) TRADE_CLOSED_BY_TAOBAO(交易被淘宝关闭) ALL_WAIT_PAY(包含：WAIT_BUYER_PAY、TRADE_NO_CREATE_PAY) ALL_CLOSED(包含：TRADE_CLOSED、TRADE_CLOSED_BY_TAOBAO) 取自&quot;http://open.taobao.com/dev/index.php/%E4%BA%A4%E6%98%93%E7%8A%B6%E6%80%81&quot;     */
    OrderStatus  *string `json:"order_status,omitempty" `

    /*
        退款原因     */
    Reason  *string `json:"reason,omitempty" `

    /*
        退还金额(退还给买家的金额)。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
    RefundFee  *string `json:"refund_fee,omitempty" `

    /*
        纠纷单号id     */
    RefundId  *string `json:"refund_id,omitempty" `

    /*
        退款阶段，可选值：onsale/aftersale     */
    RefundPhase  *string `json:"refund_phase,omitempty" `

    /*
        卖家发货物流公司（换货）     */
    SellerLogisticName  *string `json:"seller_logistic_name,omitempty" `

    /*
        卖家发货物流单号（换货）     */
    SellerLogisticNo  *string `json:"seller_logistic_no,omitempty" `

    /*
        逆向纠纷状态。其中，仅退款/退货退款的状态为：(1, "买家已经申请退款，等待卖家同意"),(2, "卖家已经同意退款，等待买家退货"),(3, "买家已经退货，等待卖家确认收货"),(4, "退款关闭"),(5, "退款成功"),(6, "卖家拒绝退款”),(7, "等待买家确认重新邮寄的货物"),(8, "等待卖家确认退货地址")；换货的状态为：(1, "换货待处理"),(2, "待买家退货"),(3, "买家已退货，待收货"),(4, "换货关闭"),(5, "换货成功"),(6, "待买家修改"),(12, "待发出换货商品"),(13, "待买家收货"),(14, "请退款")     */
    Status  *string `json:"status,omitempty" `

    /*
        超时时间     */
    TimeOut  *string `json:"time_out,omitempty" `

    /*
        商品名称     */
    Title  *string `json:"title,omitempty" `

    /*
        买家openId     */
    BuyerOpenUid  *string `json:"buyer_open_uid,omitempty" `

}

func (s *TmallDisputeReceiveGetDispute) SetAddress(v string) *TmallDisputeReceiveGetDispute {
    s.Address = &v
    return s
}
func (s *TmallDisputeReceiveGetDispute) SetAlipayNo(v string) *TmallDisputeReceiveGetDispute {
    s.AlipayNo = &v
    return s
}
func (s *TmallDisputeReceiveGetDispute) SetAttribute(v string) *TmallDisputeReceiveGetDispute {
    s.Attribute = &v
    return s
}
func (s *TmallDisputeReceiveGetDispute) SetBizOrderId(v string) *TmallDisputeReceiveGetDispute {
    s.BizOrderId = &v
    return s
}
func (s *TmallDisputeReceiveGetDispute) SetBuyerAddress(v string) *TmallDisputeReceiveGetDispute {
    s.BuyerAddress = &v
    return s
}
func (s *TmallDisputeReceiveGetDispute) SetBuyerLogisticName(v string) *TmallDisputeReceiveGetDispute {
    s.BuyerLogisticName = &v
    return s
}
func (s *TmallDisputeReceiveGetDispute) SetBuyerLogisticNo(v string) *TmallDisputeReceiveGetDispute {
    s.BuyerLogisticNo = &v
    return s
}
func (s *TmallDisputeReceiveGetDispute) SetBuyerName(v string) *TmallDisputeReceiveGetDispute {
    s.BuyerName = &v
    return s
}
func (s *TmallDisputeReceiveGetDispute) SetBuyerNick(v string) *TmallDisputeReceiveGetDispute {
    s.BuyerNick = &v
    return s
}
func (s *TmallDisputeReceiveGetDispute) SetBuyerPhone(v string) *TmallDisputeReceiveGetDispute {
    s.BuyerPhone = &v
    return s
}
func (s *TmallDisputeReceiveGetDispute) SetCompanyName(v string) *TmallDisputeReceiveGetDispute {
    s.CompanyName = &v
    return s
}
func (s *TmallDisputeReceiveGetDispute) SetCreated(v util.LocalTime) *TmallDisputeReceiveGetDispute {
    s.Created = &v
    return s
}
func (s *TmallDisputeReceiveGetDispute) SetDesc(v string) *TmallDisputeReceiveGetDispute {
    s.Desc = &v
    return s
}
func (s *TmallDisputeReceiveGetDispute) SetDisputeRequest(v string) *TmallDisputeReceiveGetDispute {
    s.DisputeRequest = &v
    return s
}
func (s *TmallDisputeReceiveGetDispute) SetGoodReturnTime(v util.LocalTime) *TmallDisputeReceiveGetDispute {
    s.GoodReturnTime = &v
    return s
}
func (s *TmallDisputeReceiveGetDispute) SetGoodStatus(v string) *TmallDisputeReceiveGetDispute {
    s.GoodStatus = &v
    return s
}
func (s *TmallDisputeReceiveGetDispute) SetHasGoodReturn(v bool) *TmallDisputeReceiveGetDispute {
    s.HasGoodReturn = &v
    return s
}
func (s *TmallDisputeReceiveGetDispute) SetModified(v util.LocalTime) *TmallDisputeReceiveGetDispute {
    s.Modified = &v
    return s
}
func (s *TmallDisputeReceiveGetDispute) SetNum(v int64) *TmallDisputeReceiveGetDispute {
    s.Num = &v
    return s
}
func (s *TmallDisputeReceiveGetDispute) SetOid(v int64) *TmallDisputeReceiveGetDispute {
    s.Oid = &v
    return s
}
func (s *TmallDisputeReceiveGetDispute) SetOrderStatus(v string) *TmallDisputeReceiveGetDispute {
    s.OrderStatus = &v
    return s
}
func (s *TmallDisputeReceiveGetDispute) SetReason(v string) *TmallDisputeReceiveGetDispute {
    s.Reason = &v
    return s
}
func (s *TmallDisputeReceiveGetDispute) SetRefundFee(v string) *TmallDisputeReceiveGetDispute {
    s.RefundFee = &v
    return s
}
func (s *TmallDisputeReceiveGetDispute) SetRefundId(v string) *TmallDisputeReceiveGetDispute {
    s.RefundId = &v
    return s
}
func (s *TmallDisputeReceiveGetDispute) SetRefundPhase(v string) *TmallDisputeReceiveGetDispute {
    s.RefundPhase = &v
    return s
}
func (s *TmallDisputeReceiveGetDispute) SetSellerLogisticName(v string) *TmallDisputeReceiveGetDispute {
    s.SellerLogisticName = &v
    return s
}
func (s *TmallDisputeReceiveGetDispute) SetSellerLogisticNo(v string) *TmallDisputeReceiveGetDispute {
    s.SellerLogisticNo = &v
    return s
}
func (s *TmallDisputeReceiveGetDispute) SetStatus(v string) *TmallDisputeReceiveGetDispute {
    s.Status = &v
    return s
}
func (s *TmallDisputeReceiveGetDispute) SetTimeOut(v string) *TmallDisputeReceiveGetDispute {
    s.TimeOut = &v
    return s
}
func (s *TmallDisputeReceiveGetDispute) SetTitle(v string) *TmallDisputeReceiveGetDispute {
    s.Title = &v
    return s
}
func (s *TmallDisputeReceiveGetDispute) SetBuyerOpenUid(v string) *TmallDisputeReceiveGetDispute {
    s.BuyerOpenUid = &v
    return s
}
