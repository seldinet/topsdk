package domain


type TaobaoRefundRefusereasonGetReason struct {
    /*
        原因编号     */
    ReasonId  *int64 `json:"reason_id,omitempty" `

    /*
        已经影响商品完好     */
    ReasonText  *string `json:"reason_text,omitempty" `

    /*
        已经影响商品完好     */
    ReasonTips  *string `json:"reason_tips,omitempty" `

}

func (s *TaobaoRefundRefusereasonGetReason) SetReasonId(v int64) *TaobaoRefundRefusereasonGetReason {
    s.ReasonId = &v
    return s
}
func (s *TaobaoRefundRefusereasonGetReason) SetReasonText(v string) *TaobaoRefundRefusereasonGetReason {
    s.ReasonText = &v
    return s
}
func (s *TaobaoRefundRefusereasonGetReason) SetReasonTips(v string) *TaobaoRefundRefusereasonGetReason {
    s.ReasonTips = &v
    return s
}
