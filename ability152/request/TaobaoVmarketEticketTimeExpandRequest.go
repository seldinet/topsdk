package request


type TaobaoVmarketEticketTimeExpandRequest struct {
    /*
        订单ID     */
    OrderId  *int64 `json:"order_id" required:"true" `
    /*
        延长天数，延长时间=当前过期时间+延长天数     */
    ExpandDays  *int64 `json:"expand_days" required:"true" `
}

func (s *TaobaoVmarketEticketTimeExpandRequest) SetOrderId(v int64) *TaobaoVmarketEticketTimeExpandRequest {
    s.OrderId = &v
    return s
}
func (s *TaobaoVmarketEticketTimeExpandRequest) SetExpandDays(v int64) *TaobaoVmarketEticketTimeExpandRequest {
    s.ExpandDays = &v
    return s
}

func (req *TaobaoVmarketEticketTimeExpandRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OrderId != nil) {
        paramMap["order_id"] = *req.OrderId
    }
    if(req.ExpandDays != nil) {
        paramMap["expand_days"] = *req.ExpandDays
    }
    return paramMap
}

func (req *TaobaoVmarketEticketTimeExpandRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}