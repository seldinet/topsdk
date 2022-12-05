package request


type TaobaoFenxiaoRefundGetRequest struct {
    /*
        要查询的退款对应的分销子订单号     */
    SubOrderId  *int64 `json:"sub_order_id,omitempty" required:"false" `
    /*
        是否查询下游消费者订单对应退款信息     */
    QuerySellerRefund  *bool `json:"query_seller_refund,omitempty" required:"false" `
    /*
        退款单id（分销子订单号和退款单id，两者至少必填一个，都填的情况下，以退款单id为准）     */
    RefundId  *int64 `json:"refund_id,omitempty" required:"false" `
}

func (s *TaobaoFenxiaoRefundGetRequest) SetSubOrderId(v int64) *TaobaoFenxiaoRefundGetRequest {
    s.SubOrderId = &v
    return s
}
func (s *TaobaoFenxiaoRefundGetRequest) SetQuerySellerRefund(v bool) *TaobaoFenxiaoRefundGetRequest {
    s.QuerySellerRefund = &v
    return s
}
func (s *TaobaoFenxiaoRefundGetRequest) SetRefundId(v int64) *TaobaoFenxiaoRefundGetRequest {
    s.RefundId = &v
    return s
}

func (req *TaobaoFenxiaoRefundGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.SubOrderId != nil) {
        paramMap["sub_order_id"] = *req.SubOrderId
    }
    if(req.QuerySellerRefund != nil) {
        paramMap["query_seller_refund"] = *req.QuerySellerRefund
    }
    if(req.RefundId != nil) {
        paramMap["refund_id"] = *req.RefundId
    }
    return paramMap
}

func (req *TaobaoFenxiaoRefundGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}