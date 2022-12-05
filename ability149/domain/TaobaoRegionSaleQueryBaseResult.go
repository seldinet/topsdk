package domain


type TaobaoRegionSaleQueryBaseResult struct {
    /*
        data     */
    Data  *string `json:"data,omitempty" `

    /*
        errorCode     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        errorMsg     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

    /*
        success     */
    Success  *bool `json:"success,omitempty" `

}

func (s *TaobaoRegionSaleQueryBaseResult) SetData(v string) *TaobaoRegionSaleQueryBaseResult {
    s.Data = &v
    return s
}
func (s *TaobaoRegionSaleQueryBaseResult) SetErrorCode(v string) *TaobaoRegionSaleQueryBaseResult {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoRegionSaleQueryBaseResult) SetErrorMsg(v string) *TaobaoRegionSaleQueryBaseResult {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoRegionSaleQueryBaseResult) SetSuccess(v bool) *TaobaoRegionSaleQueryBaseResult {
    s.Success = &v
    return s
}
