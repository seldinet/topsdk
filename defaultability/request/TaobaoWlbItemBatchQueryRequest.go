package request


type TaobaoWlbItemBatchQueryRequest struct {
    /*
        需要查询的商品ID列表，以字符串表示，ID间以;隔开     */
    ItemIds  *string `json:"item_ids" required:"true" `
    /*
        仓库编号     */
    StoreCode  *string `json:"store_code,omitempty" required:"false" `
    /*
        分页查询参数，指定查询页数，默认为1 defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        分页查询参数，每页查询数量，默认20，最大值50,大于50时按照50条查询 defalutValue��20    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
}

func (s *TaobaoWlbItemBatchQueryRequest) SetItemIds(v string) *TaobaoWlbItemBatchQueryRequest {
    s.ItemIds = &v
    return s
}
func (s *TaobaoWlbItemBatchQueryRequest) SetStoreCode(v string) *TaobaoWlbItemBatchQueryRequest {
    s.StoreCode = &v
    return s
}
func (s *TaobaoWlbItemBatchQueryRequest) SetPageNo(v int64) *TaobaoWlbItemBatchQueryRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoWlbItemBatchQueryRequest) SetPageSize(v int64) *TaobaoWlbItemBatchQueryRequest {
    s.PageSize = &v
    return s
}

func (req *TaobaoWlbItemBatchQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemIds != nil) {
        paramMap["item_ids"] = *req.ItemIds
    }
    if(req.StoreCode != nil) {
        paramMap["store_code"] = *req.StoreCode
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    return paramMap
}

func (req *TaobaoWlbItemBatchQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}