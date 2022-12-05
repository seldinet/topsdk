package domain


type TaobaoFenxiaoDistributorProductQuantityGetResultDTO struct {
    /*
        是否成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        库存数量     */
    Module  *string `json:"module,omitempty" `

    /*
        错误信息     */
    ErrorMessage  *string `json:"error_message,omitempty" `

    /*
        错误码     */
    ErrorCode  *string `json:"error_code,omitempty" `

}

func (s *TaobaoFenxiaoDistributorProductQuantityGetResultDTO) SetSuccess(v bool) *TaobaoFenxiaoDistributorProductQuantityGetResultDTO {
    s.Success = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductQuantityGetResultDTO) SetModule(v string) *TaobaoFenxiaoDistributorProductQuantityGetResultDTO {
    s.Module = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductQuantityGetResultDTO) SetErrorMessage(v string) *TaobaoFenxiaoDistributorProductQuantityGetResultDTO {
    s.ErrorMessage = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductQuantityGetResultDTO) SetErrorCode(v string) *TaobaoFenxiaoDistributorProductQuantityGetResultDTO {
    s.ErrorCode = &v
    return s
}
