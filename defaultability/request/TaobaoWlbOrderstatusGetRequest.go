package request


type TaobaoWlbOrderstatusGetRequest struct {
    /*
        物流宝订单编码     */
    OrderCode  *string `json:"order_code" required:"true" `
}

func (s *TaobaoWlbOrderstatusGetRequest) SetOrderCode(v string) *TaobaoWlbOrderstatusGetRequest {
    s.OrderCode = &v
    return s
}

func (req *TaobaoWlbOrderstatusGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OrderCode != nil) {
        paramMap["order_code"] = *req.OrderCode
    }
    return paramMap
}

func (req *TaobaoWlbOrderstatusGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}