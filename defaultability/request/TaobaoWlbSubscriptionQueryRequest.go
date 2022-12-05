package request


type TaobaoWlbSubscriptionQueryRequest struct {
    /*
        状态 
AUDITING 1-待审核; 
CANCEL 2-撤销 ;
CHECKED 3-审核通过 ;
FAILED 4-审核未通过 ;
SYNCHRONIZING 5-同步中;
只允许输入上面指定的值，且可以为空，为空时查询所有状态。若输错了，则按AUDITING处理。     */
    Status  *string `json:"status,omitempty" required:"false" `
    /*
        当前页 defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录 defalutValue��20    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
}

func (s *TaobaoWlbSubscriptionQueryRequest) SetStatus(v string) *TaobaoWlbSubscriptionQueryRequest {
    s.Status = &v
    return s
}
func (s *TaobaoWlbSubscriptionQueryRequest) SetPageNo(v int64) *TaobaoWlbSubscriptionQueryRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoWlbSubscriptionQueryRequest) SetPageSize(v int64) *TaobaoWlbSubscriptionQueryRequest {
    s.PageSize = &v
    return s
}

func (req *TaobaoWlbSubscriptionQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Status != nil) {
        paramMap["status"] = *req.Status
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    return paramMap
}

func (req *TaobaoWlbSubscriptionQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}