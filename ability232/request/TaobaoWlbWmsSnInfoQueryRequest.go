package request


type TaobaoWlbWmsSnInfoQueryRequest struct {
    /*
        订单编码     */
    OrderCode  *string `json:"order_code" required:"true" `
    /*
        订单类型（1:仓配订单 10：配送扫码 20：增值扫码 40:ERP单号; 50：交易订单 ）     */
    OrderCodeType  *int64 `json:"order_code_type" required:"true" `
    /*
        页数，默认每页50条     */
    PageIndex  *int64 `json:"page_index" required:"true" `
}

func (s *TaobaoWlbWmsSnInfoQueryRequest) SetOrderCode(v string) *TaobaoWlbWmsSnInfoQueryRequest {
    s.OrderCode = &v
    return s
}
func (s *TaobaoWlbWmsSnInfoQueryRequest) SetOrderCodeType(v int64) *TaobaoWlbWmsSnInfoQueryRequest {
    s.OrderCodeType = &v
    return s
}
func (s *TaobaoWlbWmsSnInfoQueryRequest) SetPageIndex(v int64) *TaobaoWlbWmsSnInfoQueryRequest {
    s.PageIndex = &v
    return s
}

func (req *TaobaoWlbWmsSnInfoQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OrderCode != nil) {
        paramMap["order_code"] = *req.OrderCode
    }
    if(req.OrderCodeType != nil) {
        paramMap["order_code_type"] = *req.OrderCodeType
    }
    if(req.PageIndex != nil) {
        paramMap["page_index"] = *req.PageIndex
    }
    return paramMap
}

func (req *TaobaoWlbWmsSnInfoQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}