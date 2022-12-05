package domain

import (
        "topsdk/util"
    )

type TaobaoFenxiaoRefundQueryBuyerRefund struct {
    /*
        分销子订单号     */
    SubOrderId  *int64 `json:"sub_order_id,omitempty" `

    /*
        消费者订单对应的退款单号     */
    RefundId  *int64 `json:"refund_id,omitempty" `

    /*
        消费者订单退款涉及的消费者正向子订单号     */
    BizOrderId  *int64 `json:"biz_order_id,omitempty" `

    /*
        消费者订单退款创建时间     */
    RefundCreateTime  *util.LocalTime `json:"refund_create_time,omitempty" `

    /*
        消费者订单退款状态 1、消费者已经申请退款，等待分销商确认 2、分销商已经同意退货，等待消费者退货  3、消费者已经退货，等待分销商确认收货 4、退款关闭   5、退款成功 6、分销商拒绝退款,待消费者重新修改  7、等待消费者确认重新邮寄的货物       */
    RefundStatus  *int64 `json:"refund_status,omitempty" `

    /*
        货物的状态：买家已收到货买家已退货买家未收到货     */
    GoodsStatusDesc  *string `json:"goods_status_desc,omitempty" `

    /*
        买家是否退货     */
    NeedReturnGoods  *bool `json:"need_return_goods,omitempty" `

    /*
        退还给消费者的金额(分)     */
    ReturnFee  *int64 `json:"return_fee,omitempty" `

    /*
        确认收货后会打款给分销商的金额(分),分摊到子单的实付金额-退款给消费者的金额     */
    ToSellerFee  *int64 `json:"to_seller_fee,omitempty" `

    /*
        消费者退款原因     */
    RefundReason  *string `json:"refund_reason,omitempty" `

    /*
        消费者退款说明     */
    RefundDesc  *string `json:"refund_desc,omitempty" `

    /*
        下游买家nick     */
    BuyerNick  *string `json:"buyer_nick,omitempty" `

    /*
        消费者退款修改时间。格式:yyyy-MM-dd HH:mm:ss     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

    /*
        消费者买家id加密key     */
    OpenBuyerId  *string `json:"open_buyer_id,omitempty" `

    /*
        消费者退货数量     */
    ReturnGoodsQuantity  *int64 `json:"return_goods_quantity,omitempty" `

}

func (s *TaobaoFenxiaoRefundQueryBuyerRefund) SetSubOrderId(v int64) *TaobaoFenxiaoRefundQueryBuyerRefund {
    s.SubOrderId = &v
    return s
}
func (s *TaobaoFenxiaoRefundQueryBuyerRefund) SetRefundId(v int64) *TaobaoFenxiaoRefundQueryBuyerRefund {
    s.RefundId = &v
    return s
}
func (s *TaobaoFenxiaoRefundQueryBuyerRefund) SetBizOrderId(v int64) *TaobaoFenxiaoRefundQueryBuyerRefund {
    s.BizOrderId = &v
    return s
}
func (s *TaobaoFenxiaoRefundQueryBuyerRefund) SetRefundCreateTime(v util.LocalTime) *TaobaoFenxiaoRefundQueryBuyerRefund {
    s.RefundCreateTime = &v
    return s
}
func (s *TaobaoFenxiaoRefundQueryBuyerRefund) SetRefundStatus(v int64) *TaobaoFenxiaoRefundQueryBuyerRefund {
    s.RefundStatus = &v
    return s
}
func (s *TaobaoFenxiaoRefundQueryBuyerRefund) SetGoodsStatusDesc(v string) *TaobaoFenxiaoRefundQueryBuyerRefund {
    s.GoodsStatusDesc = &v
    return s
}
func (s *TaobaoFenxiaoRefundQueryBuyerRefund) SetNeedReturnGoods(v bool) *TaobaoFenxiaoRefundQueryBuyerRefund {
    s.NeedReturnGoods = &v
    return s
}
func (s *TaobaoFenxiaoRefundQueryBuyerRefund) SetReturnFee(v int64) *TaobaoFenxiaoRefundQueryBuyerRefund {
    s.ReturnFee = &v
    return s
}
func (s *TaobaoFenxiaoRefundQueryBuyerRefund) SetToSellerFee(v int64) *TaobaoFenxiaoRefundQueryBuyerRefund {
    s.ToSellerFee = &v
    return s
}
func (s *TaobaoFenxiaoRefundQueryBuyerRefund) SetRefundReason(v string) *TaobaoFenxiaoRefundQueryBuyerRefund {
    s.RefundReason = &v
    return s
}
func (s *TaobaoFenxiaoRefundQueryBuyerRefund) SetRefundDesc(v string) *TaobaoFenxiaoRefundQueryBuyerRefund {
    s.RefundDesc = &v
    return s
}
func (s *TaobaoFenxiaoRefundQueryBuyerRefund) SetBuyerNick(v string) *TaobaoFenxiaoRefundQueryBuyerRefund {
    s.BuyerNick = &v
    return s
}
func (s *TaobaoFenxiaoRefundQueryBuyerRefund) SetModified(v util.LocalTime) *TaobaoFenxiaoRefundQueryBuyerRefund {
    s.Modified = &v
    return s
}
func (s *TaobaoFenxiaoRefundQueryBuyerRefund) SetOpenBuyerId(v string) *TaobaoFenxiaoRefundQueryBuyerRefund {
    s.OpenBuyerId = &v
    return s
}
func (s *TaobaoFenxiaoRefundQueryBuyerRefund) SetReturnGoodsQuantity(v int64) *TaobaoFenxiaoRefundQueryBuyerRefund {
    s.ReturnGoodsQuantity = &v
    return s
}
