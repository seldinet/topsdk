package request


type TaobaoWlbOrderitemPageGetRequest struct {
    /*
        物流宝订单编码     */
    OrderCode  *string `json:"order_code" required:"true" `
    /*
        分页查询参数，指定查询页数，默认为1 defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        分页查询参数，每页查询数量，默认20，最大值50,大于50时按照50条查询 defalutValue��20    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
}

func (s *TaobaoWlbOrderitemPageGetRequest) SetOrderCode(v string) *TaobaoWlbOrderitemPageGetRequest {
    s.OrderCode = &v
    return s
}
func (s *TaobaoWlbOrderitemPageGetRequest) SetPageNo(v int64) *TaobaoWlbOrderitemPageGetRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoWlbOrderitemPageGetRequest) SetPageSize(v int64) *TaobaoWlbOrderitemPageGetRequest {
    s.PageSize = &v
    return s
}

func (req *TaobaoWlbOrderitemPageGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OrderCode != nil) {
        paramMap["order_code"] = *req.OrderCode
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    return paramMap
}

func (req *TaobaoWlbOrderitemPageGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}