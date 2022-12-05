package domain


type TaobaoFenxiaoRefundGetRefundItem struct {
    /*
        退款明细ID，针对一笔退款每一个品就映射为一个明细，每一个明细有一个全局唯一的ID     */
    RefundItemId  *int64 `json:"refund_item_id,omitempty" `

    /*
        分销子订单号     */
    SubOrderId  *int64 `json:"sub_order_id,omitempty" `

    /*
        退货数量     */
    RefundQuantity  *int64 `json:"refund_quantity,omitempty" `

}

func (s *TaobaoFenxiaoRefundGetRefundItem) SetRefundItemId(v int64) *TaobaoFenxiaoRefundGetRefundItem {
    s.RefundItemId = &v
    return s
}
func (s *TaobaoFenxiaoRefundGetRefundItem) SetSubOrderId(v int64) *TaobaoFenxiaoRefundGetRefundItem {
    s.SubOrderId = &v
    return s
}
func (s *TaobaoFenxiaoRefundGetRefundItem) SetRefundQuantity(v int64) *TaobaoFenxiaoRefundGetRefundItem {
    s.RefundQuantity = &v
    return s
}
