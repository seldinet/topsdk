package request


type TaobaoVmarketEticketFlowConsumeRequest struct {
    /*
        业务单号     */
    OuterId  *string `json:"outer_id" required:"true" `
    /*
        凭证码     */
    Code  *string `json:"code" required:"true" `
    /*
        淘宝业务提供的业务类型值，请联系相关业务运营取得     */
    BizType  *int64 `json:"biz_type" required:"true" `
    /*
        核销操作人     */
    Operator  *string `json:"operator,omitempty" required:"false" `
}

func (s *TaobaoVmarketEticketFlowConsumeRequest) SetOuterId(v string) *TaobaoVmarketEticketFlowConsumeRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoVmarketEticketFlowConsumeRequest) SetCode(v string) *TaobaoVmarketEticketFlowConsumeRequest {
    s.Code = &v
    return s
}
func (s *TaobaoVmarketEticketFlowConsumeRequest) SetBizType(v int64) *TaobaoVmarketEticketFlowConsumeRequest {
    s.BizType = &v
    return s
}
func (s *TaobaoVmarketEticketFlowConsumeRequest) SetOperator(v string) *TaobaoVmarketEticketFlowConsumeRequest {
    s.Operator = &v
    return s
}

func (req *TaobaoVmarketEticketFlowConsumeRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OuterId != nil) {
        paramMap["outer_id"] = *req.OuterId
    }
    if(req.Code != nil) {
        paramMap["code"] = *req.Code
    }
    if(req.BizType != nil) {
        paramMap["biz_type"] = *req.BizType
    }
    if(req.Operator != nil) {
        paramMap["operator"] = *req.Operator
    }
    return paramMap
}

func (req *TaobaoVmarketEticketFlowConsumeRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}