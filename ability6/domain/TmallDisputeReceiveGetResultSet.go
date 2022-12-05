package domain


type TmallDisputeReceiveGetResultSet struct {
    /*
        当前页面的纠纷单数量     */
    PageResults  *int64 `json:"page_results,omitempty" `

    /*
        所有符合查询条件的纠纷单数量     */
    TotalResults  *int64 `json:"total_results,omitempty" `

    /*
        是否还有下一页     */
    HasNext  *bool `json:"has_next,omitempty" `

    /*
        错误码     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        错误信息     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

    /*
        results     */
    Results  *[]TmallDisputeReceiveGetDispute `json:"results,omitempty" `

}

func (s *TmallDisputeReceiveGetResultSet) SetPageResults(v int64) *TmallDisputeReceiveGetResultSet {
    s.PageResults = &v
    return s
}
func (s *TmallDisputeReceiveGetResultSet) SetTotalResults(v int64) *TmallDisputeReceiveGetResultSet {
    s.TotalResults = &v
    return s
}
func (s *TmallDisputeReceiveGetResultSet) SetHasNext(v bool) *TmallDisputeReceiveGetResultSet {
    s.HasNext = &v
    return s
}
func (s *TmallDisputeReceiveGetResultSet) SetErrorCode(v string) *TmallDisputeReceiveGetResultSet {
    s.ErrorCode = &v
    return s
}
func (s *TmallDisputeReceiveGetResultSet) SetErrorMsg(v string) *TmallDisputeReceiveGetResultSet {
    s.ErrorMsg = &v
    return s
}
func (s *TmallDisputeReceiveGetResultSet) SetResults(v []TmallDisputeReceiveGetDispute) *TmallDisputeReceiveGetResultSet {
    s.Results = &v
    return s
}
