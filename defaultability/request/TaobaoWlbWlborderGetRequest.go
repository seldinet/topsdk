package request


type TaobaoWlbWlborderGetRequest struct {
    /*
        物流宝订单编码     */
    WlbOrderCode  *string `json:"wlb_order_code" required:"true" `
}

func (s *TaobaoWlbWlborderGetRequest) SetWlbOrderCode(v string) *TaobaoWlbWlborderGetRequest {
    s.WlbOrderCode = &v
    return s
}

func (req *TaobaoWlbWlborderGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.WlbOrderCode != nil) {
        paramMap["wlb_order_code"] = *req.WlbOrderCode
    }
    return paramMap
}

func (req *TaobaoWlbWlborderGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}