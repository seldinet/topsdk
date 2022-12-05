package domain


type TaobaoScitemQueryQueryPagination struct {
    /*
        当前页码数     */
    PageIndex  *int64 `json:"page_index,omitempty" `

    /*
        分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录     */
    PageSize  *int64 `json:"page_size,omitempty" `

}

func (s *TaobaoScitemQueryQueryPagination) SetPageIndex(v int64) *TaobaoScitemQueryQueryPagination {
    s.PageIndex = &v
    return s
}
func (s *TaobaoScitemQueryQueryPagination) SetPageSize(v int64) *TaobaoScitemQueryQueryPagination {
    s.PageSize = &v
    return s
}
