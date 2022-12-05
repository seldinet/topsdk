package request


type TaobaoEticketMerchantMaReverseRequest struct {
    /*
        业务类型 defalutValue��3001    */
    BizType  *int64 `json:"biz_type,omitempty" required:"false" `
    /*
        码值     */
    Code  *string `json:"code" required:"true" `
    /*
        业务id（订单号）     */
    OuterId  *string `json:"outer_id" required:"true" `
    /*
        机具编号，如果核销时有则必传     */
    PosId  *string `json:"pos_id,omitempty" required:"false" `
    /*
        冲正份数，需要与核销份数一致     */
    ReverseNum  *int64 `json:"reverse_num" required:"true" `
    /*
        需要冲正的核销序列号     */
    SerialNum  *string `json:"serial_num" required:"true" `
    /*
        需要跟发码通知获取到的参数一致     */
    Token  *string `json:"token" required:"true" `
}

func (s *TaobaoEticketMerchantMaReverseRequest) SetBizType(v int64) *TaobaoEticketMerchantMaReverseRequest {
    s.BizType = &v
    return s
}
func (s *TaobaoEticketMerchantMaReverseRequest) SetCode(v string) *TaobaoEticketMerchantMaReverseRequest {
    s.Code = &v
    return s
}
func (s *TaobaoEticketMerchantMaReverseRequest) SetOuterId(v string) *TaobaoEticketMerchantMaReverseRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoEticketMerchantMaReverseRequest) SetPosId(v string) *TaobaoEticketMerchantMaReverseRequest {
    s.PosId = &v
    return s
}
func (s *TaobaoEticketMerchantMaReverseRequest) SetReverseNum(v int64) *TaobaoEticketMerchantMaReverseRequest {
    s.ReverseNum = &v
    return s
}
func (s *TaobaoEticketMerchantMaReverseRequest) SetSerialNum(v string) *TaobaoEticketMerchantMaReverseRequest {
    s.SerialNum = &v
    return s
}
func (s *TaobaoEticketMerchantMaReverseRequest) SetToken(v string) *TaobaoEticketMerchantMaReverseRequest {
    s.Token = &v
    return s
}

func (req *TaobaoEticketMerchantMaReverseRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.BizType != nil) {
        paramMap["biz_type"] = *req.BizType
    }
    if(req.Code != nil) {
        paramMap["code"] = *req.Code
    }
    if(req.OuterId != nil) {
        paramMap["outer_id"] = *req.OuterId
    }
    if(req.PosId != nil) {
        paramMap["pos_id"] = *req.PosId
    }
    if(req.ReverseNum != nil) {
        paramMap["reverse_num"] = *req.ReverseNum
    }
    if(req.SerialNum != nil) {
        paramMap["serial_num"] = *req.SerialNum
    }
    if(req.Token != nil) {
        paramMap["token"] = *req.Token
    }
    return paramMap
}

func (req *TaobaoEticketMerchantMaReverseRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}