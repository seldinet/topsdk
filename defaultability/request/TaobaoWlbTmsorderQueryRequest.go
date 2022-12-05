package request


type TaobaoWlbTmsorderQueryRequest struct {
    /*
        物流订单编号     */
    OrderCode  *string `json:"order_code" required:"true" `
    /*
        当前页 defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录 defalutValue��20    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
}

func (s *TaobaoWlbTmsorderQueryRequest) SetOrderCode(v string) *TaobaoWlbTmsorderQueryRequest {
    s.OrderCode = &v
    return s
}
func (s *TaobaoWlbTmsorderQueryRequest) SetPageNo(v int64) *TaobaoWlbTmsorderQueryRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoWlbTmsorderQueryRequest) SetPageSize(v int64) *TaobaoWlbTmsorderQueryRequest {
    s.PageSize = &v
    return s
}

func (req *TaobaoWlbTmsorderQueryRequest) ToMap() map[string]interface{} {
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

func (req *TaobaoWlbTmsorderQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}