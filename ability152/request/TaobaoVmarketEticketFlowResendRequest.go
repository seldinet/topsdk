package request


type TaobaoVmarketEticketFlowResendRequest struct {
    /*
        业务单号     */
    OuterId  *string `json:"outer_id" required:"true" `
    /*
        业务类型值，可联系淘宝业务运营取得具体值     */
    BizType  *int64 `json:"biz_type" required:"true" `
}

func (s *TaobaoVmarketEticketFlowResendRequest) SetOuterId(v string) *TaobaoVmarketEticketFlowResendRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoVmarketEticketFlowResendRequest) SetBizType(v int64) *TaobaoVmarketEticketFlowResendRequest {
    s.BizType = &v
    return s
}

func (req *TaobaoVmarketEticketFlowResendRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OuterId != nil) {
        paramMap["outer_id"] = *req.OuterId
    }
    if(req.BizType != nil) {
        paramMap["biz_type"] = *req.BizType
    }
    return paramMap
}

func (req *TaobaoVmarketEticketFlowResendRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}