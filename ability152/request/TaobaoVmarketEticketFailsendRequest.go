package request


type TaobaoVmarketEticketFailsendRequest struct {
    /*
        订单号     */
    OrderId  *int64 `json:"order_id" required:"true" `
    /*
        发码通知时的token     */
    Token  *string `json:"token,omitempty" required:"false" `
    /*
        错误码     */
    ErrorCode  *int64 `json:"error_code" required:"true" `
    /*
        错误信息     */
    ErrorMsg  *string `json:"error_msg,omitempty" required:"false" `
}

func (s *TaobaoVmarketEticketFailsendRequest) SetOrderId(v int64) *TaobaoVmarketEticketFailsendRequest {
    s.OrderId = &v
    return s
}
func (s *TaobaoVmarketEticketFailsendRequest) SetToken(v string) *TaobaoVmarketEticketFailsendRequest {
    s.Token = &v
    return s
}
func (s *TaobaoVmarketEticketFailsendRequest) SetErrorCode(v int64) *TaobaoVmarketEticketFailsendRequest {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoVmarketEticketFailsendRequest) SetErrorMsg(v string) *TaobaoVmarketEticketFailsendRequest {
    s.ErrorMsg = &v
    return s
}

func (req *TaobaoVmarketEticketFailsendRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OrderId != nil) {
        paramMap["order_id"] = *req.OrderId
    }
    if(req.Token != nil) {
        paramMap["token"] = *req.Token
    }
    if(req.ErrorCode != nil) {
        paramMap["error_code"] = *req.ErrorCode
    }
    if(req.ErrorMsg != nil) {
        paramMap["error_msg"] = *req.ErrorMsg
    }
    return paramMap
}

func (req *TaobaoVmarketEticketFailsendRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}