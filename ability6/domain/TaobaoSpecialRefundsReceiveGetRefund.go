package domain

import (
        "topsdk/util"
    )

type TaobaoSpecialRefundsReceiveGetRefund struct {
    /*
        disputeRequest ：1 是仅退款   3 是退货退款  4 是换货     */
    Attribute  *string `json:"attribute,omitempty" `

    /*
        买家昵称     */
    BuyerNick  *string `json:"buyer_nick,omitempty" `

    /*
        物流公司名称     */
    CompanyName  *string `json:"company_name,omitempty" `

    /*
        退款申请时间。格式:yyyy-MM-dd HH:mm:ss     */
    Created  *util.LocalTime `json:"created,omitempty" `

    /*
        退款说明     */
    Desc  *string `json:"desc,omitempty" `

    /*
        货物状态。可选值BUYER_NOT_RECEIVED (买家未收到货) BUYER_RECEIVED (买家已收到货) BUYER_RETURNED_GOODS (买家已退货)     */
    GoodStatus  *string `json:"good_status,omitempty" `

    /*
        买家是否需要退货。可选值:true(是),false(否)     */
    HasGoodReturn  *bool `json:"has_good_return,omitempty" `

    /*
        更新时间。格式:yyyy-MM-dd HH:mm:ss     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

    /*
        商品购买数量     */
    Num  *int64 `json:"num,omitempty" `

    /*
        子订单号。如果是单笔交易oid会等于tid     */
    Oid  *int64 `json:"oid,omitempty" `

    /*
        退款约束，可选值：cannot_refuse（不允许操作），refund_onweb（需要到网页版操作）     */
    OperationContraint  *string `json:"operation_contraint,omitempty" `

    /*
        退款对应的订单交易状态。可选值TRADE_NO_CREATE_PAY(没有创建支付宝交易) WAIT_BUYER_PAY(等待买家付款) WAIT_SELLER_SEND_GOODS(等待卖家发货,即:买家已付款) WAIT_BUYER_CONFIRM_GOODS(等待买家确认收货,即:卖家已发货) TRADE_BUYER_SIGNED(买家已签收,货到付款专用) TRADE_FINISHED(交易成功) TRADE_CLOSED(交易关闭) TRADE_CLOSED_BY_TAOBAO(交易被淘宝关闭) ALL_WAIT_PAY(包含：WAIT_BUYER_PAY、TRADE_NO_CREATE_PAY) ALL_CLOSED(包含：TRADE_CLOSED、TRADE_CLOSED_BY_TAOBAO) 取自&quot;http://open.taobao.com/dev/index.php/%E4%BA%A4%E6%98%93%E7%8A%B6%E6%80%81&quot;     */
    OrderStatus  *string `json:"order_status,omitempty" `

    /*
        商品外部商家编码     */
    OuterId  *string `json:"outer_id,omitempty" `

    /*
        支付给卖家的金额(交易总金额-退还给买家的金额)。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
    Payment  *string `json:"payment,omitempty" `

    /*
        退款原因     */
    Reason  *string `json:"reason,omitempty" `

    /*
        退还金额(退还给买家的金额)。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
    RefundFee  *string `json:"refund_fee,omitempty" `

    /*
        退款单号     */
    RefundId  *int64 `json:"refund_id,omitempty" `

    /*
        退款阶段，可选值：onsale/aftersale     */
    RefundPhase  *string `json:"refund_phase,omitempty" `

    /*
        退款版本号（时间戳）     */
    RefundVersion  *int64 `json:"refund_version,omitempty" `

    /*
        卖家昵称     */
    SellerNick  *string `json:"seller_nick,omitempty" `

    /*
        退货运单号     */
    Sid  *string `json:"sid,omitempty" `

    /*
        商品SKU信息     */
    Sku  *string `json:"sku,omitempty" `

    /*
        退款状态。可选值WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意) WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货) WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货) SELLER_REFUSE_BUYER(卖家拒绝退款) CLOSED(退款关闭) SUCCESS(退款成功)     */
    Status  *string `json:"status,omitempty" `

    /*
        淘宝交易单号     */
    Tid  *int64 `json:"tid,omitempty" `

    /*
        商品标题     */
    Title  *string `json:"title,omitempty" `

    /*
        交易总金额。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
    TotalFee  *string `json:"total_fee,omitempty" `

    /*
        特殊退款类型：退差返现、价保服务     */
    SpecialRefundType  *string `json:"special_refund_type,omitempty" `

    /*
        买家openId     */
    BuyerOpenUid  *string `json:"buyer_open_uid,omitempty" `

}

func (s *TaobaoSpecialRefundsReceiveGetRefund) SetAttribute(v string) *TaobaoSpecialRefundsReceiveGetRefund {
    s.Attribute = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRefund) SetBuyerNick(v string) *TaobaoSpecialRefundsReceiveGetRefund {
    s.BuyerNick = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRefund) SetCompanyName(v string) *TaobaoSpecialRefundsReceiveGetRefund {
    s.CompanyName = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRefund) SetCreated(v util.LocalTime) *TaobaoSpecialRefundsReceiveGetRefund {
    s.Created = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRefund) SetDesc(v string) *TaobaoSpecialRefundsReceiveGetRefund {
    s.Desc = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRefund) SetGoodStatus(v string) *TaobaoSpecialRefundsReceiveGetRefund {
    s.GoodStatus = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRefund) SetHasGoodReturn(v bool) *TaobaoSpecialRefundsReceiveGetRefund {
    s.HasGoodReturn = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRefund) SetModified(v util.LocalTime) *TaobaoSpecialRefundsReceiveGetRefund {
    s.Modified = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRefund) SetNum(v int64) *TaobaoSpecialRefundsReceiveGetRefund {
    s.Num = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRefund) SetOid(v int64) *TaobaoSpecialRefundsReceiveGetRefund {
    s.Oid = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRefund) SetOperationContraint(v string) *TaobaoSpecialRefundsReceiveGetRefund {
    s.OperationContraint = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRefund) SetOrderStatus(v string) *TaobaoSpecialRefundsReceiveGetRefund {
    s.OrderStatus = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRefund) SetOuterId(v string) *TaobaoSpecialRefundsReceiveGetRefund {
    s.OuterId = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRefund) SetPayment(v string) *TaobaoSpecialRefundsReceiveGetRefund {
    s.Payment = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRefund) SetReason(v string) *TaobaoSpecialRefundsReceiveGetRefund {
    s.Reason = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRefund) SetRefundFee(v string) *TaobaoSpecialRefundsReceiveGetRefund {
    s.RefundFee = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRefund) SetRefundId(v int64) *TaobaoSpecialRefundsReceiveGetRefund {
    s.RefundId = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRefund) SetRefundPhase(v string) *TaobaoSpecialRefundsReceiveGetRefund {
    s.RefundPhase = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRefund) SetRefundVersion(v int64) *TaobaoSpecialRefundsReceiveGetRefund {
    s.RefundVersion = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRefund) SetSellerNick(v string) *TaobaoSpecialRefundsReceiveGetRefund {
    s.SellerNick = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRefund) SetSid(v string) *TaobaoSpecialRefundsReceiveGetRefund {
    s.Sid = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRefund) SetSku(v string) *TaobaoSpecialRefundsReceiveGetRefund {
    s.Sku = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRefund) SetStatus(v string) *TaobaoSpecialRefundsReceiveGetRefund {
    s.Status = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRefund) SetTid(v int64) *TaobaoSpecialRefundsReceiveGetRefund {
    s.Tid = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRefund) SetTitle(v string) *TaobaoSpecialRefundsReceiveGetRefund {
    s.Title = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRefund) SetTotalFee(v string) *TaobaoSpecialRefundsReceiveGetRefund {
    s.TotalFee = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRefund) SetSpecialRefundType(v string) *TaobaoSpecialRefundsReceiveGetRefund {
    s.SpecialRefundType = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetRefund) SetBuyerOpenUid(v string) *TaobaoSpecialRefundsReceiveGetRefund {
    s.BuyerOpenUid = &v
    return s
}
