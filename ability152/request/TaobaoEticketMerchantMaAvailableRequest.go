package request


type TaobaoEticketMerchantMaAvailableRequest struct {
    /*
        业务类型 defalutValue��3001    */
    BizType  *int64 `json:"biz_type,omitempty" required:"false" `
    /*
        需要被核销的码     */
    Code  *string `json:"code" required:"true" `
    /*
        核销份数     */
    ConsumeNum  *int64 `json:"consume_num" required:"true" `
    /*
        业务id（订单号）     */
    OuterId  *string `json:"outer_id" required:"true" `
    /*
        机具编号     */
    PosId  *string `json:"pos_id,omitempty" required:"false" `
    /*
        核销序列号，需要保证唯一     */
    SerialNum  *string `json:"serial_num,omitempty" required:"false" `
    /*
        需要跟发码通知获取到的参数一致     */
    Token  *string `json:"token" required:"true" `
}

func (s *TaobaoEticketMerchantMaAvailableRequest) SetBizType(v int64) *TaobaoEticketMerchantMaAvailableRequest {
    s.BizType = &v
    return s
}
func (s *TaobaoEticketMerchantMaAvailableRequest) SetCode(v string) *TaobaoEticketMerchantMaAvailableRequest {
    s.Code = &v
    return s
}
func (s *TaobaoEticketMerchantMaAvailableRequest) SetConsumeNum(v int64) *TaobaoEticketMerchantMaAvailableRequest {
    s.ConsumeNum = &v
    return s
}
func (s *TaobaoEticketMerchantMaAvailableRequest) SetOuterId(v string) *TaobaoEticketMerchantMaAvailableRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoEticketMerchantMaAvailableRequest) SetPosId(v string) *TaobaoEticketMerchantMaAvailableRequest {
    s.PosId = &v
    return s
}
func (s *TaobaoEticketMerchantMaAvailableRequest) SetSerialNum(v string) *TaobaoEticketMerchantMaAvailableRequest {
    s.SerialNum = &v
    return s
}
func (s *TaobaoEticketMerchantMaAvailableRequest) SetToken(v string) *TaobaoEticketMerchantMaAvailableRequest {
    s.Token = &v
    return s
}

func (req *TaobaoEticketMerchantMaAvailableRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.BizType != nil) {
        paramMap["biz_type"] = *req.BizType
    }
    if(req.Code != nil) {
        paramMap["code"] = *req.Code
    }
    if(req.ConsumeNum != nil) {
        paramMap["consume_num"] = *req.ConsumeNum
    }
    if(req.OuterId != nil) {
        paramMap["outer_id"] = *req.OuterId
    }
    if(req.PosId != nil) {
        paramMap["pos_id"] = *req.PosId
    }
    if(req.SerialNum != nil) {
        paramMap["serial_num"] = *req.SerialNum
    }
    if(req.Token != nil) {
        paramMap["token"] = *req.Token
    }
    return paramMap
}

func (req *TaobaoEticketMerchantMaAvailableRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}