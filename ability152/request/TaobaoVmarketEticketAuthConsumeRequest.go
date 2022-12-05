package request


type TaobaoVmarketEticketAuthConsumeRequest struct {
    /*
        核销的码，只支持单个码，多个码核销需要多次调用     */
    VerifyCode  *string `json:"verify_code" required:"true" `
    /*
        核销方的ID，如果是普通码商必须传入机具ID,如果是私有码商家（即原有的信任商家）可默认传入私有码商ID     */
    Operatorid  *string `json:"operatorid" required:"true" `
    /*
        核销份数     */
    ConsumeNum  *int64 `json:"consume_num" required:"true" `
    /*
        自定义核销流水号，需要小于等于100个字符(a-zA-Z0-9_)     */
    SerialNum  *string `json:"serial_num" required:"true" `
    /*
        网点ID,网点授权核销时，必须传入；其他核销方式可不传     */
    Storeid  *string `json:"storeid,omitempty" required:"false" `
}

func (s *TaobaoVmarketEticketAuthConsumeRequest) SetVerifyCode(v string) *TaobaoVmarketEticketAuthConsumeRequest {
    s.VerifyCode = &v
    return s
}
func (s *TaobaoVmarketEticketAuthConsumeRequest) SetOperatorid(v string) *TaobaoVmarketEticketAuthConsumeRequest {
    s.Operatorid = &v
    return s
}
func (s *TaobaoVmarketEticketAuthConsumeRequest) SetConsumeNum(v int64) *TaobaoVmarketEticketAuthConsumeRequest {
    s.ConsumeNum = &v
    return s
}
func (s *TaobaoVmarketEticketAuthConsumeRequest) SetSerialNum(v string) *TaobaoVmarketEticketAuthConsumeRequest {
    s.SerialNum = &v
    return s
}
func (s *TaobaoVmarketEticketAuthConsumeRequest) SetStoreid(v string) *TaobaoVmarketEticketAuthConsumeRequest {
    s.Storeid = &v
    return s
}

func (req *TaobaoVmarketEticketAuthConsumeRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.VerifyCode != nil) {
        paramMap["verify_code"] = *req.VerifyCode
    }
    if(req.Operatorid != nil) {
        paramMap["operatorid"] = *req.Operatorid
    }
    if(req.ConsumeNum != nil) {
        paramMap["consume_num"] = *req.ConsumeNum
    }
    if(req.SerialNum != nil) {
        paramMap["serial_num"] = *req.SerialNum
    }
    if(req.Storeid != nil) {
        paramMap["storeid"] = *req.Storeid
    }
    return paramMap
}

func (req *TaobaoVmarketEticketAuthConsumeRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}