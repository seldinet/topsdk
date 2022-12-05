package domain


type TaobaoWlbWmsSnInfoQueryResult struct {
    /*
        错误码     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        是否成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        错误信息     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

    /*
        总条数     */
    TotalCount  *int64 `json:"total_count,omitempty" `

    /*
        SN信息列表     */
    SnInfoList  *[]TaobaoWlbWmsSnInfoQuerySninfolist `json:"sn_info_list,omitempty" `

}

func (s *TaobaoWlbWmsSnInfoQueryResult) SetErrorCode(v string) *TaobaoWlbWmsSnInfoQueryResult {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoWlbWmsSnInfoQueryResult) SetSuccess(v bool) *TaobaoWlbWmsSnInfoQueryResult {
    s.Success = &v
    return s
}
func (s *TaobaoWlbWmsSnInfoQueryResult) SetErrorMsg(v string) *TaobaoWlbWmsSnInfoQueryResult {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoWlbWmsSnInfoQueryResult) SetTotalCount(v int64) *TaobaoWlbWmsSnInfoQueryResult {
    s.TotalCount = &v
    return s
}
func (s *TaobaoWlbWmsSnInfoQueryResult) SetSnInfoList(v []TaobaoWlbWmsSnInfoQuerySninfolist) *TaobaoWlbWmsSnInfoQueryResult {
    s.SnInfoList = &v
    return s
}
