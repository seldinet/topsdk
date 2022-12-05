package request


type TaobaoWlbOrderConsignRequest struct {
    /*
        物流宝订单编号     */
    WlbOrderCode  *string `json:"wlb_order_code" required:"true" `
}

func (s *TaobaoWlbOrderConsignRequest) SetWlbOrderCode(v string) *TaobaoWlbOrderConsignRequest {
    s.WlbOrderCode = &v
    return s
}

func (req *TaobaoWlbOrderConsignRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.WlbOrderCode != nil) {
        paramMap["wlb_order_code"] = *req.WlbOrderCode
    }
    return paramMap
}

func (req *TaobaoWlbOrderConsignRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}