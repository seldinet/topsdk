package domain


type TaobaoRpRefundsAgreeRefundMappingResult struct {
    /*
        结果信息     */
    Message  *string `json:"message,omitempty" `

    /*
        退款ID     */
    RefundId  *string `json:"refund_id,omitempty" `

    /*
        是否成功     */
    Succ  *bool `json:"succ,omitempty" `

}

func (s *TaobaoRpRefundsAgreeRefundMappingResult) SetMessage(v string) *TaobaoRpRefundsAgreeRefundMappingResult {
    s.Message = &v
    return s
}
func (s *TaobaoRpRefundsAgreeRefundMappingResult) SetRefundId(v string) *TaobaoRpRefundsAgreeRefundMappingResult {
    s.RefundId = &v
    return s
}
func (s *TaobaoRpRefundsAgreeRefundMappingResult) SetSucc(v bool) *TaobaoRpRefundsAgreeRefundMappingResult {
    s.Succ = &v
    return s
}
