package domain


type TaobaoTradeFullinfoGetStepPayDetail struct {
    /*
        分阶段支付单Id     */
    Id  *int64 `json:"id,omitempty" `

    /*
        支付状态     */
    PayStatus  *int64 `json:"pay_status,omitempty" `

    /*
        支付顺序     */
    StepNo  *int64 `json:"step_no,omitempty" `

    /*
        支付宝交易号     */
    StepChannelNo  *string `json:"step_channel_no,omitempty" `

    /*
        支付方式     */
    StepInstrumentCode  *string `json:"step_instrument_code,omitempty" `

    /*
        支付金额     */
    StepActualPayFee  *string `json:"step_actual_pay_fee,omitempty" `

}

func (s *TaobaoTradeFullinfoGetStepPayDetail) SetId(v int64) *TaobaoTradeFullinfoGetStepPayDetail {
    s.Id = &v
    return s
}
func (s *TaobaoTradeFullinfoGetStepPayDetail) SetPayStatus(v int64) *TaobaoTradeFullinfoGetStepPayDetail {
    s.PayStatus = &v
    return s
}
func (s *TaobaoTradeFullinfoGetStepPayDetail) SetStepNo(v int64) *TaobaoTradeFullinfoGetStepPayDetail {
    s.StepNo = &v
    return s
}
func (s *TaobaoTradeFullinfoGetStepPayDetail) SetStepChannelNo(v string) *TaobaoTradeFullinfoGetStepPayDetail {
    s.StepChannelNo = &v
    return s
}
func (s *TaobaoTradeFullinfoGetStepPayDetail) SetStepInstrumentCode(v string) *TaobaoTradeFullinfoGetStepPayDetail {
    s.StepInstrumentCode = &v
    return s
}
func (s *TaobaoTradeFullinfoGetStepPayDetail) SetStepActualPayFee(v string) *TaobaoTradeFullinfoGetStepPayDetail {
    s.StepActualPayFee = &v
    return s
}
