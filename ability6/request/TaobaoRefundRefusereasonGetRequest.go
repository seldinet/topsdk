package request

import (
        "topsdk/util"
    )

type TaobaoRefundRefusereasonGetRequest struct {
    /*
        售中或售后     */
    RefundPhase  *string `json:"refund_phase" required:"true" `
    /*
        退款编号     */
    RefundId  *int64 `json:"refund_id" required:"true" `
    /*
        返回参数     */
    Fields  *[]string `json:"fields" required:"true" `
}

func (s *TaobaoRefundRefusereasonGetRequest) SetRefundPhase(v string) *TaobaoRefundRefusereasonGetRequest {
    s.RefundPhase = &v
    return s
}
func (s *TaobaoRefundRefusereasonGetRequest) SetRefundId(v int64) *TaobaoRefundRefusereasonGetRequest {
    s.RefundId = &v
    return s
}
func (s *TaobaoRefundRefusereasonGetRequest) SetFields(v []string) *TaobaoRefundRefusereasonGetRequest {
    s.Fields = &v
    return s
}

func (req *TaobaoRefundRefusereasonGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefundPhase != nil) {
        paramMap["refund_phase"] = *req.RefundPhase
    }
    if(req.RefundId != nil) {
        paramMap["refund_id"] = *req.RefundId
    }
    if(req.Fields != nil) {
        paramMap["fields"] = util.ConvertBasicList(*req.Fields)
    }
    return paramMap
}

func (req *TaobaoRefundRefusereasonGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}