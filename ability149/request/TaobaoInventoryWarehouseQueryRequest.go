package request


type TaobaoInventoryWarehouseQueryRequest struct {
    /*
        页码     */
    PageNo  *int64 `json:"page_no" required:"true" `
    /*
        页大小     */
    PageSize  *int64 `json:"page_size" required:"true" `
}

func (s *TaobaoInventoryWarehouseQueryRequest) SetPageNo(v int64) *TaobaoInventoryWarehouseQueryRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoInventoryWarehouseQueryRequest) SetPageSize(v int64) *TaobaoInventoryWarehouseQueryRequest {
    s.PageSize = &v
    return s
}

func (req *TaobaoInventoryWarehouseQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    return paramMap
}

func (req *TaobaoInventoryWarehouseQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}