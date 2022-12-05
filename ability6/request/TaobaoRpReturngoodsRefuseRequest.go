package request

import (
        "topsdk/util"
    )

type TaobaoRpReturngoodsRefuseRequest struct {
    /*
        退款编号     */
    RefundId  *int64 `json:"refund_id" required:"true" `
    /*
        退款服务状态，售后或者售中     */
    RefundPhase  *string `json:"refund_phase" required:"true" `
    /*
        退款版本号     */
    RefundVersion  *int64 `json:"refund_version" required:"true" `
    /*
        拒绝退货凭证图片，必须图片格式，大小不能超过5M     */
    RefuseProof  *[]byte `json:"refuse_proof" required:"true" `
    /*
        拒绝原因编号，会提供拒绝原因列表供选择     */
    RefuseReasonId  *int64 `json:"refuse_reason_id,omitempty" required:"false" `
}

func (s *TaobaoRpReturngoodsRefuseRequest) SetRefundId(v int64) *TaobaoRpReturngoodsRefuseRequest {
    s.RefundId = &v
    return s
}
func (s *TaobaoRpReturngoodsRefuseRequest) SetRefundPhase(v string) *TaobaoRpReturngoodsRefuseRequest {
    s.RefundPhase = &v
    return s
}
func (s *TaobaoRpReturngoodsRefuseRequest) SetRefundVersion(v int64) *TaobaoRpReturngoodsRefuseRequest {
    s.RefundVersion = &v
    return s
}
func (s *TaobaoRpReturngoodsRefuseRequest) SetRefuseProof(v []byte) *TaobaoRpReturngoodsRefuseRequest {
    s.RefuseProof = &v
    return s
}
func (s *TaobaoRpReturngoodsRefuseRequest) SetRefuseReasonId(v int64) *TaobaoRpReturngoodsRefuseRequest {
    s.RefuseReasonId = &v
    return s
}

func (req *TaobaoRpReturngoodsRefuseRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefundId != nil) {
        paramMap["refund_id"] = *req.RefundId
    }
    if(req.RefundPhase != nil) {
        paramMap["refund_phase"] = *req.RefundPhase
    }
    if(req.RefundVersion != nil) {
        paramMap["refund_version"] = *req.RefundVersion
    }
    if(req.RefuseReasonId != nil) {
        paramMap["refuse_reason_id"] = *req.RefuseReasonId
    }
    return paramMap
}

func (req *TaobaoRpReturngoodsRefuseRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    if req.RefuseProof != nil {
        fileMap["refuse_proof"] = *req.RefuseProof
    }
    return fileMap
}