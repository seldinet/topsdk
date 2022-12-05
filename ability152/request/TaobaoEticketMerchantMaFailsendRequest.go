package request


type TaobaoEticketMerchantMaFailsendRequest struct {
    /*
        业务类型 defalutValue��3001    */
    BizType  *int64 `json:"biz_type,omitempty" required:"false" `
    /*
        业务id（订单号）     */
    OuterId  *string `json:"outer_id" required:"true" `
    /*
        错误原因码     */
    SubErrCode  *string `json:"sub_err_code" required:"true" `
    /*
        错误码描述     */
    SubErrMsg  *string `json:"sub_err_msg" required:"true" `
    /*
        需要与发码通知获取的值一致     */
    Token  *string `json:"token" required:"true" `
}

func (s *TaobaoEticketMerchantMaFailsendRequest) SetBizType(v int64) *TaobaoEticketMerchantMaFailsendRequest {
    s.BizType = &v
    return s
}
func (s *TaobaoEticketMerchantMaFailsendRequest) SetOuterId(v string) *TaobaoEticketMerchantMaFailsendRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoEticketMerchantMaFailsendRequest) SetSubErrCode(v string) *TaobaoEticketMerchantMaFailsendRequest {
    s.SubErrCode = &v
    return s
}
func (s *TaobaoEticketMerchantMaFailsendRequest) SetSubErrMsg(v string) *TaobaoEticketMerchantMaFailsendRequest {
    s.SubErrMsg = &v
    return s
}
func (s *TaobaoEticketMerchantMaFailsendRequest) SetToken(v string) *TaobaoEticketMerchantMaFailsendRequest {
    s.Token = &v
    return s
}

func (req *TaobaoEticketMerchantMaFailsendRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.BizType != nil) {
        paramMap["biz_type"] = *req.BizType
    }
    if(req.OuterId != nil) {
        paramMap["outer_id"] = *req.OuterId
    }
    if(req.SubErrCode != nil) {
        paramMap["sub_err_code"] = *req.SubErrCode
    }
    if(req.SubErrMsg != nil) {
        paramMap["sub_err_msg"] = *req.SubErrMsg
    }
    if(req.Token != nil) {
        paramMap["token"] = *req.Token
    }
    return paramMap
}

func (req *TaobaoEticketMerchantMaFailsendRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}