package domain


type TmallChannelProductsQueryPageResultDto struct {
    /*
        异常信息     */
    ErrorMessage  *string `json:"error_message,omitempty" `

    /*
        产品信息     */
    ProductList  *[]TmallChannelProductsQueryProductTopDto `json:"product_list,omitempty" `

    /*
        总数     */
    TotalCount  *int64 `json:"total_count,omitempty" `

    /*
        是否有下一页     */
    HasNext  *bool `json:"has_next,omitempty" `

    /*
        错误码     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        是否查询成功     */
    Success  *bool `json:"success,omitempty" `

}

func (s *TmallChannelProductsQueryPageResultDto) SetErrorMessage(v string) *TmallChannelProductsQueryPageResultDto {
    s.ErrorMessage = &v
    return s
}
func (s *TmallChannelProductsQueryPageResultDto) SetProductList(v []TmallChannelProductsQueryProductTopDto) *TmallChannelProductsQueryPageResultDto {
    s.ProductList = &v
    return s
}
func (s *TmallChannelProductsQueryPageResultDto) SetTotalCount(v int64) *TmallChannelProductsQueryPageResultDto {
    s.TotalCount = &v
    return s
}
func (s *TmallChannelProductsQueryPageResultDto) SetHasNext(v bool) *TmallChannelProductsQueryPageResultDto {
    s.HasNext = &v
    return s
}
func (s *TmallChannelProductsQueryPageResultDto) SetErrorCode(v string) *TmallChannelProductsQueryPageResultDto {
    s.ErrorCode = &v
    return s
}
func (s *TmallChannelProductsQueryPageResultDto) SetSuccess(v bool) *TmallChannelProductsQueryPageResultDto {
    s.Success = &v
    return s
}
