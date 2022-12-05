package request


type TaobaoFenxiaoDealerRequisitionorderAgreeRequest struct {
    /*
        采购申请/经销采购单编号     */
    DealerOrderId  *int64 `json:"dealer_order_id" required:"true" `
}

func (s *TaobaoFenxiaoDealerRequisitionorderAgreeRequest) SetDealerOrderId(v int64) *TaobaoFenxiaoDealerRequisitionorderAgreeRequest {
    s.DealerOrderId = &v
    return s
}

func (req *TaobaoFenxiaoDealerRequisitionorderAgreeRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.DealerOrderId != nil) {
        paramMap["dealer_order_id"] = *req.DealerOrderId
    }
    return paramMap
}

func (req *TaobaoFenxiaoDealerRequisitionorderAgreeRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}