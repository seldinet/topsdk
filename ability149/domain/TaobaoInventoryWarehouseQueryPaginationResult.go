package domain


type TaobaoInventoryWarehouseQueryPaginationResult struct {
    /*
        仓库信息数组     */
    Data  *string `json:"data,omitempty" `

    /*
        错误码     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        错误信息     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

    /*
        页码     */
    PageNo  *int64 `json:"page_no,omitempty" `

    /*
        页大小     */
    PageSize  *int64 `json:"page_size,omitempty" `

    /*
        是否成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        总条数     */
    TotalCount  *int64 `json:"total_count,omitempty" `

}

func (s *TaobaoInventoryWarehouseQueryPaginationResult) SetData(v string) *TaobaoInventoryWarehouseQueryPaginationResult {
    s.Data = &v
    return s
}
func (s *TaobaoInventoryWarehouseQueryPaginationResult) SetErrorCode(v string) *TaobaoInventoryWarehouseQueryPaginationResult {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoInventoryWarehouseQueryPaginationResult) SetErrorMsg(v string) *TaobaoInventoryWarehouseQueryPaginationResult {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoInventoryWarehouseQueryPaginationResult) SetPageNo(v int64) *TaobaoInventoryWarehouseQueryPaginationResult {
    s.PageNo = &v
    return s
}
func (s *TaobaoInventoryWarehouseQueryPaginationResult) SetPageSize(v int64) *TaobaoInventoryWarehouseQueryPaginationResult {
    s.PageSize = &v
    return s
}
func (s *TaobaoInventoryWarehouseQueryPaginationResult) SetSuccess(v bool) *TaobaoInventoryWarehouseQueryPaginationResult {
    s.Success = &v
    return s
}
func (s *TaobaoInventoryWarehouseQueryPaginationResult) SetTotalCount(v int64) *TaobaoInventoryWarehouseQueryPaginationResult {
    s.TotalCount = &v
    return s
}
