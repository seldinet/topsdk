package request


type TaobaoVmarketEticketStoreGetRequest struct {
    /*
        订单ID     */
    OrderId  *int64 `json:"order_id" required:"true" `
}

func (s *TaobaoVmarketEticketStoreGetRequest) SetOrderId(v int64) *TaobaoVmarketEticketStoreGetRequest {
    s.OrderId = &v
    return s
}

func (req *TaobaoVmarketEticketStoreGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OrderId != nil) {
        paramMap["order_id"] = *req.OrderId
    }
    return paramMap
}

func (req *TaobaoVmarketEticketStoreGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}