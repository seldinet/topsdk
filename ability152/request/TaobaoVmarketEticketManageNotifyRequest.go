package request


type TaobaoVmarketEticketManageNotifyRequest struct {
    /*
        订单编号     */
    OrderId  *int64 `json:"order_id" required:"true" `
    /*
        需要调用的通知方法，目前仅支持是send（发码）或resend（重新发码）     */
    NotifyMethod  *string `json:"notify_method" required:"true" `
    /*
        码商ID，如果是码商，必须传，如果是信任卖家，不需要传     */
    CodemerchantId  *int64 `json:"codemerchant_id,omitempty" required:"false" `
}

func (s *TaobaoVmarketEticketManageNotifyRequest) SetOrderId(v int64) *TaobaoVmarketEticketManageNotifyRequest {
    s.OrderId = &v
    return s
}
func (s *TaobaoVmarketEticketManageNotifyRequest) SetNotifyMethod(v string) *TaobaoVmarketEticketManageNotifyRequest {
    s.NotifyMethod = &v
    return s
}
func (s *TaobaoVmarketEticketManageNotifyRequest) SetCodemerchantId(v int64) *TaobaoVmarketEticketManageNotifyRequest {
    s.CodemerchantId = &v
    return s
}

func (req *TaobaoVmarketEticketManageNotifyRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OrderId != nil) {
        paramMap["order_id"] = *req.OrderId
    }
    if(req.NotifyMethod != nil) {
        paramMap["notify_method"] = *req.NotifyMethod
    }
    if(req.CodemerchantId != nil) {
        paramMap["codemerchant_id"] = *req.CodemerchantId
    }
    return paramMap
}

func (req *TaobaoVmarketEticketManageNotifyRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}