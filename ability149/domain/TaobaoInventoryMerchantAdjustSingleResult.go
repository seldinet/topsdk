package domain


type TaobaoInventoryMerchantAdjustSingleResult struct {
    /*
        错误信息     */
    ErrorMessage  *string `json:"error_message,omitempty" `

    /*
        data     */
    AdjustResults  *[]TaobaoInventoryMerchantAdjustInventoryCheckResultDto `json:"adjust_results,omitempty" `

    /*
        错误码     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        如果是失败，可能是部分失败。如果是成功，则全部成功     */
    Success  *bool `json:"success,omitempty" `

}

func (s *TaobaoInventoryMerchantAdjustSingleResult) SetErrorMessage(v string) *TaobaoInventoryMerchantAdjustSingleResult {
    s.ErrorMessage = &v
    return s
}
func (s *TaobaoInventoryMerchantAdjustSingleResult) SetAdjustResults(v []TaobaoInventoryMerchantAdjustInventoryCheckResultDto) *TaobaoInventoryMerchantAdjustSingleResult {
    s.AdjustResults = &v
    return s
}
func (s *TaobaoInventoryMerchantAdjustSingleResult) SetErrorCode(v string) *TaobaoInventoryMerchantAdjustSingleResult {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoInventoryMerchantAdjustSingleResult) SetSuccess(v bool) *TaobaoInventoryMerchantAdjustSingleResult {
    s.Success = &v
    return s
}
