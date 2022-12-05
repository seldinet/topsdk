package domain

import (
        "topsdk/util"
    )

type TaobaoFenxiaoRefundGetTopDpRefundDetailDO struct {
    /*
        分销子订单号，如果是by子单发起退款，就会在退款主单上记录分销子订单号     */
    SubOrderId  *int64 `json:"sub_order_id,omitempty" `

    /*
        是否退货,如果是已发货退货退款/售后退货退款，就是true     */
    IsReturnGoods  *bool `json:"is_return_goods,omitempty" `

    /*
        退款创建时间     */
    RefundCreateTime  *util.LocalTime `json:"refund_create_time,omitempty" `

    /*
        退款状态1：分销商已经申请退款，等待供应商确认2：供应商已经同意退货，等待分销商退货3：分销商已经退货，等待供应商确认收货4：退款关闭5：退款成功  6：供应商拒绝退款12：供应商同意退款，待系统打款  9：没有申请退款 10：供应商拒绝确认收货,待分销商重新修改     */
    RefundStatus  *int64 `json:"refund_status,omitempty" `

    /*
        退款的金额(元)     */
    RefundFee  *string `json:"refund_fee,omitempty" `

    /*
        支付给供应商的金额(元)，分销订单子单实付金额-退款金额     */
    PaySupFee  *string `json:"pay_sup_fee,omitempty" `

    /*
        退款原因     */
    RefundReason  *string `json:"refund_reason,omitempty" `

    /*
        退款说明     */
    RefundDesc  *string `json:"refund_desc,omitempty" `

    /*
        分销商nick     */
    DistributorNick  *string `json:"distributor_nick,omitempty" `

    /*
        供应商nick     */
    SupplierNick  *string `json:"supplier_nick,omitempty" `

    /*
        退款修改时间。格式:yyyy-MM-dd HH:mm:ss     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

    /*
        分销主订单号     */
    PurchaseOrderId  *int64 `json:"purchase_order_id,omitempty" `

    /*
        退款流程类型：4：未发货退款；1：已发货仅退款；2：已发货退货退款；3：售后仅退款；5：物流拒收；6：售后退货退款     */
    RefundFlowType  *int64 `json:"refund_flow_type,omitempty" `

    /*
        超时时间     */
    Timeout  *util.LocalTime `json:"timeout,omitempty" `

    /*
        超时类型：
1：供应商同意退款/同意退货超时；
2：供应商确认收货超时     */
    ToType  *int64 `json:"to_type,omitempty" `

    /*
        前台消费者订单对应的退款详情     */
    BuyerRefund  *TaobaoFenxiaoRefundGetBuyerRefund `json:"buyer_refund,omitempty" `

    /*
        分销退款单号     */
    RefundId  *int64 `json:"refund_id,omitempty" `

    /*
        退货的物流信息     */
    ReturnLogistics  *[]TaobaoFenxiaoRefundGetRefundLogistics `json:"return_logistics,omitempty" `

    /*
        退款明细项，记录退款涉及的订单     */
    RefundItems  *[]TaobaoFenxiaoRefundGetRefundItem `json:"refund_items,omitempty" `

}

func (s *TaobaoFenxiaoRefundGetTopDpRefundDetailDO) SetSubOrderId(v int64) *TaobaoFenxiaoRefundGetTopDpRefundDetailDO {
    s.SubOrderId = &v
    return s
}
func (s *TaobaoFenxiaoRefundGetTopDpRefundDetailDO) SetIsReturnGoods(v bool) *TaobaoFenxiaoRefundGetTopDpRefundDetailDO {
    s.IsReturnGoods = &v
    return s
}
func (s *TaobaoFenxiaoRefundGetTopDpRefundDetailDO) SetRefundCreateTime(v util.LocalTime) *TaobaoFenxiaoRefundGetTopDpRefundDetailDO {
    s.RefundCreateTime = &v
    return s
}
func (s *TaobaoFenxiaoRefundGetTopDpRefundDetailDO) SetRefundStatus(v int64) *TaobaoFenxiaoRefundGetTopDpRefundDetailDO {
    s.RefundStatus = &v
    return s
}
func (s *TaobaoFenxiaoRefundGetTopDpRefundDetailDO) SetRefundFee(v string) *TaobaoFenxiaoRefundGetTopDpRefundDetailDO {
    s.RefundFee = &v
    return s
}
func (s *TaobaoFenxiaoRefundGetTopDpRefundDetailDO) SetPaySupFee(v string) *TaobaoFenxiaoRefundGetTopDpRefundDetailDO {
    s.PaySupFee = &v
    return s
}
func (s *TaobaoFenxiaoRefundGetTopDpRefundDetailDO) SetRefundReason(v string) *TaobaoFenxiaoRefundGetTopDpRefundDetailDO {
    s.RefundReason = &v
    return s
}
func (s *TaobaoFenxiaoRefundGetTopDpRefundDetailDO) SetRefundDesc(v string) *TaobaoFenxiaoRefundGetTopDpRefundDetailDO {
    s.RefundDesc = &v
    return s
}
func (s *TaobaoFenxiaoRefundGetTopDpRefundDetailDO) SetDistributorNick(v string) *TaobaoFenxiaoRefundGetTopDpRefundDetailDO {
    s.DistributorNick = &v
    return s
}
func (s *TaobaoFenxiaoRefundGetTopDpRefundDetailDO) SetSupplierNick(v string) *TaobaoFenxiaoRefundGetTopDpRefundDetailDO {
    s.SupplierNick = &v
    return s
}
func (s *TaobaoFenxiaoRefundGetTopDpRefundDetailDO) SetModified(v util.LocalTime) *TaobaoFenxiaoRefundGetTopDpRefundDetailDO {
    s.Modified = &v
    return s
}
func (s *TaobaoFenxiaoRefundGetTopDpRefundDetailDO) SetPurchaseOrderId(v int64) *TaobaoFenxiaoRefundGetTopDpRefundDetailDO {
    s.PurchaseOrderId = &v
    return s
}
func (s *TaobaoFenxiaoRefundGetTopDpRefundDetailDO) SetRefundFlowType(v int64) *TaobaoFenxiaoRefundGetTopDpRefundDetailDO {
    s.RefundFlowType = &v
    return s
}
func (s *TaobaoFenxiaoRefundGetTopDpRefundDetailDO) SetTimeout(v util.LocalTime) *TaobaoFenxiaoRefundGetTopDpRefundDetailDO {
    s.Timeout = &v
    return s
}
func (s *TaobaoFenxiaoRefundGetTopDpRefundDetailDO) SetToType(v int64) *TaobaoFenxiaoRefundGetTopDpRefundDetailDO {
    s.ToType = &v
    return s
}
func (s *TaobaoFenxiaoRefundGetTopDpRefundDetailDO) SetBuyerRefund(v TaobaoFenxiaoRefundGetBuyerRefund) *TaobaoFenxiaoRefundGetTopDpRefundDetailDO {
    s.BuyerRefund = &v
    return s
}
func (s *TaobaoFenxiaoRefundGetTopDpRefundDetailDO) SetRefundId(v int64) *TaobaoFenxiaoRefundGetTopDpRefundDetailDO {
    s.RefundId = &v
    return s
}
func (s *TaobaoFenxiaoRefundGetTopDpRefundDetailDO) SetReturnLogistics(v []TaobaoFenxiaoRefundGetRefundLogistics) *TaobaoFenxiaoRefundGetTopDpRefundDetailDO {
    s.ReturnLogistics = &v
    return s
}
func (s *TaobaoFenxiaoRefundGetTopDpRefundDetailDO) SetRefundItems(v []TaobaoFenxiaoRefundGetRefundItem) *TaobaoFenxiaoRefundGetTopDpRefundDetailDO {
    s.RefundItems = &v
    return s
}
