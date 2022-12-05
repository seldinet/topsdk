package request


type TaobaoVmarketEticketCodesGetRequest struct {
    /*
        订单号     */
    OrderId  *int64 `json:"order_id" required:"true" `
    /*
        码商ID     */
    CodemerchantId  *int64 `json:"codemerchant_id,omitempty" required:"false" `
}

func (s *TaobaoVmarketEticketCodesGetRequest) SetOrderId(v int64) *TaobaoVmarketEticketCodesGetRequest {
    s.OrderId = &v
    return s
}
func (s *TaobaoVmarketEticketCodesGetRequest) SetCodemerchantId(v int64) *TaobaoVmarketEticketCodesGetRequest {
    s.CodemerchantId = &v
    return s
}

func (req *TaobaoVmarketEticketCodesGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OrderId != nil) {
        paramMap["order_id"] = *req.OrderId
    }
    if(req.CodemerchantId != nil) {
        paramMap["codemerchant_id"] = *req.CodemerchantId
    }
    return paramMap
}

func (req *TaobaoVmarketEticketCodesGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}