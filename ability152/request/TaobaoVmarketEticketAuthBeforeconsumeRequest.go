package request


type TaobaoVmarketEticketAuthBeforeconsumeRequest struct {
    /*
        核销的码，只支持单个码，多个码核销需要多次调用     */
    VerifyCode  *string `json:"verify_code" required:"true" `
    /*
        核销方的ID，如果是普通码商必须传入机具ID,如果是私有码商家（即原有的信任商家）可默认传入私有码商ID     */
    Operatorid  *string `json:"operatorid" required:"true" `
    /*
        网点ID,网点授权核销时，必须传入；其他核销方式可不传     */
    Storeid  *string `json:"storeid,omitempty" required:"false" `
}

func (s *TaobaoVmarketEticketAuthBeforeconsumeRequest) SetVerifyCode(v string) *TaobaoVmarketEticketAuthBeforeconsumeRequest {
    s.VerifyCode = &v
    return s
}
func (s *TaobaoVmarketEticketAuthBeforeconsumeRequest) SetOperatorid(v string) *TaobaoVmarketEticketAuthBeforeconsumeRequest {
    s.Operatorid = &v
    return s
}
func (s *TaobaoVmarketEticketAuthBeforeconsumeRequest) SetStoreid(v string) *TaobaoVmarketEticketAuthBeforeconsumeRequest {
    s.Storeid = &v
    return s
}

func (req *TaobaoVmarketEticketAuthBeforeconsumeRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.VerifyCode != nil) {
        paramMap["verify_code"] = *req.VerifyCode
    }
    if(req.Operatorid != nil) {
        paramMap["operatorid"] = *req.Operatorid
    }
    if(req.Storeid != nil) {
        paramMap["storeid"] = *req.Storeid
    }
    return paramMap
}

func (req *TaobaoVmarketEticketAuthBeforeconsumeRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}