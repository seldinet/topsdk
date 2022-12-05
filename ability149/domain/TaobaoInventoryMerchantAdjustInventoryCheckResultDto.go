package domain


type TaobaoInventoryMerchantAdjustInventoryCheckResultDto struct {
    /*
        错误信息     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

    /*
        错误码     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        每个子调整单据是否成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        每个货品的调整子单据号，作为业务调整依据，处理时会幂等     */
    SubOrderId  *string `json:"sub_order_id,omitempty" `

}

func (s *TaobaoInventoryMerchantAdjustInventoryCheckResultDto) SetErrorMsg(v string) *TaobaoInventoryMerchantAdjustInventoryCheckResultDto {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoInventoryMerchantAdjustInventoryCheckResultDto) SetErrorCode(v string) *TaobaoInventoryMerchantAdjustInventoryCheckResultDto {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoInventoryMerchantAdjustInventoryCheckResultDto) SetSuccess(v bool) *TaobaoInventoryMerchantAdjustInventoryCheckResultDto {
    s.Success = &v
    return s
}
func (s *TaobaoInventoryMerchantAdjustInventoryCheckResultDto) SetSubOrderId(v string) *TaobaoInventoryMerchantAdjustInventoryCheckResultDto {
    s.SubOrderId = &v
    return s
}
