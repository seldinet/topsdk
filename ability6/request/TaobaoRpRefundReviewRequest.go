package request


type TaobaoRpRefundReviewRequest struct {
    /*
        审核人姓名     */
    Operator  *string `json:"operator" required:"true" `
    /*
        退款单编号     */
    RefundId  *int64 `json:"refund_id" required:"true" `
    /*
        退款阶段，可选值：售中：onsale，售后：aftersale     */
    RefundPhase  *string `json:"refund_phase" required:"true" `
    /*
        退款最后更新时间，以时间戳的方式表示     */
    RefundVersion  *int64 `json:"refund_version" required:"true" `
    /*
        审核留言     */
    Message  *string `json:"message" required:"true" `
    /*
        审核是否可用于批量退款，可选值：true（审核通过），false（审核不通过或反审核）     */
    Result  *bool `json:"result" required:"true" `
}

func (s *TaobaoRpRefundReviewRequest) SetOperator(v string) *TaobaoRpRefundReviewRequest {
    s.Operator = &v
    return s
}
func (s *TaobaoRpRefundReviewRequest) SetRefundId(v int64) *TaobaoRpRefundReviewRequest {
    s.RefundId = &v
    return s
}
func (s *TaobaoRpRefundReviewRequest) SetRefundPhase(v string) *TaobaoRpRefundReviewRequest {
    s.RefundPhase = &v
    return s
}
func (s *TaobaoRpRefundReviewRequest) SetRefundVersion(v int64) *TaobaoRpRefundReviewRequest {
    s.RefundVersion = &v
    return s
}
func (s *TaobaoRpRefundReviewRequest) SetMessage(v string) *TaobaoRpRefundReviewRequest {
    s.Message = &v
    return s
}
func (s *TaobaoRpRefundReviewRequest) SetResult(v bool) *TaobaoRpRefundReviewRequest {
    s.Result = &v
    return s
}

func (req *TaobaoRpRefundReviewRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Operator != nil) {
        paramMap["operator"] = *req.Operator
    }
    if(req.RefundId != nil) {
        paramMap["refund_id"] = *req.RefundId
    }
    if(req.RefundPhase != nil) {
        paramMap["refund_phase"] = *req.RefundPhase
    }
    if(req.RefundVersion != nil) {
        paramMap["refund_version"] = *req.RefundVersion
    }
    if(req.Message != nil) {
        paramMap["message"] = *req.Message
    }
    if(req.Result != nil) {
        paramMap["result"] = *req.Result
    }
    return paramMap
}

func (req *TaobaoRpRefundReviewRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}