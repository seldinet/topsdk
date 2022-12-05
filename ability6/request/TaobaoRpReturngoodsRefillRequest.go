package request


type TaobaoRpReturngoodsRefillRequest struct {
    /*
        退款单编号     */
    RefundId  *int64 `json:"refund_id" required:"true" `
    /*
        退款阶段，可选值：售中：onsale，售后：aftersale     */
    RefundPhase  *string `json:"refund_phase" required:"true" `
    /*
        物流公司运单号     */
    LogisticsWaybillNo  *string `json:"logistics_waybill_no" required:"true" `
    /*
        物流公司编号     */
    LogisticsCompanyCode  *string `json:"logistics_company_code" required:"true" `
}

func (s *TaobaoRpReturngoodsRefillRequest) SetRefundId(v int64) *TaobaoRpReturngoodsRefillRequest {
    s.RefundId = &v
    return s
}
func (s *TaobaoRpReturngoodsRefillRequest) SetRefundPhase(v string) *TaobaoRpReturngoodsRefillRequest {
    s.RefundPhase = &v
    return s
}
func (s *TaobaoRpReturngoodsRefillRequest) SetLogisticsWaybillNo(v string) *TaobaoRpReturngoodsRefillRequest {
    s.LogisticsWaybillNo = &v
    return s
}
func (s *TaobaoRpReturngoodsRefillRequest) SetLogisticsCompanyCode(v string) *TaobaoRpReturngoodsRefillRequest {
    s.LogisticsCompanyCode = &v
    return s
}

func (req *TaobaoRpReturngoodsRefillRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefundId != nil) {
        paramMap["refund_id"] = *req.RefundId
    }
    if(req.RefundPhase != nil) {
        paramMap["refund_phase"] = *req.RefundPhase
    }
    if(req.LogisticsWaybillNo != nil) {
        paramMap["logistics_waybill_no"] = *req.LogisticsWaybillNo
    }
    if(req.LogisticsCompanyCode != nil) {
        paramMap["logistics_company_code"] = *req.LogisticsCompanyCode
    }
    return paramMap
}

func (req *TaobaoRpReturngoodsRefillRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}