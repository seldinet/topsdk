package request


type TaobaoVmarketEticketBeforeconsumeRequest struct {
    /*
        需要验码的电子凭证订单ID     */
    OrderId  *int64 `json:"order_id" required:"true" `
    /*
        需要验的码     */
    VerifyCode  *string `json:"verify_code" required:"true" `
    /*
        安全验证token，需要和发码通知中的token一致     */
    Token  *string `json:"token" required:"true" `
    /*
        码商ID,是码商的话必须传递,如果是信任卖家不需要传     */
    CodemerchantId  *int64 `json:"codemerchant_id,omitempty" required:"false" `
    /*
        操作员身份ID，如果是码商必须传,如果是信任卖家不需要传     */
    Posid  *string `json:"posid,omitempty" required:"false" `
    /*
        手机号码后四位,没有特殊说明请不要传     */
    Mobile  *string `json:"mobile,omitempty" required:"false" `
}

func (s *TaobaoVmarketEticketBeforeconsumeRequest) SetOrderId(v int64) *TaobaoVmarketEticketBeforeconsumeRequest {
    s.OrderId = &v
    return s
}
func (s *TaobaoVmarketEticketBeforeconsumeRequest) SetVerifyCode(v string) *TaobaoVmarketEticketBeforeconsumeRequest {
    s.VerifyCode = &v
    return s
}
func (s *TaobaoVmarketEticketBeforeconsumeRequest) SetToken(v string) *TaobaoVmarketEticketBeforeconsumeRequest {
    s.Token = &v
    return s
}
func (s *TaobaoVmarketEticketBeforeconsumeRequest) SetCodemerchantId(v int64) *TaobaoVmarketEticketBeforeconsumeRequest {
    s.CodemerchantId = &v
    return s
}
func (s *TaobaoVmarketEticketBeforeconsumeRequest) SetPosid(v string) *TaobaoVmarketEticketBeforeconsumeRequest {
    s.Posid = &v
    return s
}
func (s *TaobaoVmarketEticketBeforeconsumeRequest) SetMobile(v string) *TaobaoVmarketEticketBeforeconsumeRequest {
    s.Mobile = &v
    return s
}

func (req *TaobaoVmarketEticketBeforeconsumeRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OrderId != nil) {
        paramMap["order_id"] = *req.OrderId
    }
    if(req.VerifyCode != nil) {
        paramMap["verify_code"] = *req.VerifyCode
    }
    if(req.Token != nil) {
        paramMap["token"] = *req.Token
    }
    if(req.CodemerchantId != nil) {
        paramMap["codemerchant_id"] = *req.CodemerchantId
    }
    if(req.Posid != nil) {
        paramMap["posid"] = *req.Posid
    }
    if(req.Mobile != nil) {
        paramMap["mobile"] = *req.Mobile
    }
    return paramMap
}

func (req *TaobaoVmarketEticketBeforeconsumeRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}