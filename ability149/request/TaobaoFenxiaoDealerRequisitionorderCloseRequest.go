package request


type TaobaoFenxiaoDealerRequisitionorderCloseRequest struct {
    /*
        经销采购单编号     */
    DealerOrderId  *int64 `json:"dealer_order_id" required:"true" `
    /*
        关闭原因：
1：长时间无法联系到分销商，取消交易。
2：分销商错误提交申请，取消交易。
3：缺货，近段时间都无法发货。
4：分销商恶意提交申请单。
5：其他原因。     */
    Reason  *int64 `json:"reason" required:"true" `
    /*
        关闭详细原因，字数5-200字     */
    ReasonDetail  *string `json:"reason_detail" required:"true" `
}

func (s *TaobaoFenxiaoDealerRequisitionorderCloseRequest) SetDealerOrderId(v int64) *TaobaoFenxiaoDealerRequisitionorderCloseRequest {
    s.DealerOrderId = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderCloseRequest) SetReason(v int64) *TaobaoFenxiaoDealerRequisitionorderCloseRequest {
    s.Reason = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderCloseRequest) SetReasonDetail(v string) *TaobaoFenxiaoDealerRequisitionorderCloseRequest {
    s.ReasonDetail = &v
    return s
}

func (req *TaobaoFenxiaoDealerRequisitionorderCloseRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.DealerOrderId != nil) {
        paramMap["dealer_order_id"] = *req.DealerOrderId
    }
    if(req.Reason != nil) {
        paramMap["reason"] = *req.Reason
    }
    if(req.ReasonDetail != nil) {
        paramMap["reason_detail"] = *req.ReasonDetail
    }
    return paramMap
}

func (req *TaobaoFenxiaoDealerRequisitionorderCloseRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}