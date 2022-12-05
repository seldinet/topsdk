package request


type TaobaoWlbOrderCancelRequest struct {
    /*
        物流宝订单编号     */
    WlbOrderCode  *string `json:"wlb_order_code" required:"true" `
}

func (s *TaobaoWlbOrderCancelRequest) SetWlbOrderCode(v string) *TaobaoWlbOrderCancelRequest {
    s.WlbOrderCode = &v
    return s
}

func (req *TaobaoWlbOrderCancelRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.WlbOrderCode != nil) {
        paramMap["wlb_order_code"] = *req.WlbOrderCode
    }
    return paramMap
}

func (req *TaobaoWlbOrderCancelRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}