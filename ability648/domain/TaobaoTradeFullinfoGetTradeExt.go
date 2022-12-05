package domain


type TaobaoTradeFullinfoGetTradeExt struct {
    /*
        enable前扩展标识位     */
    BeforeEnableFlag  *int64 `json:"before_enable_flag,omitempty" `

    /*
        关闭订单前扩展标识位     */
    BeforeCloseFlag  *int64 `json:"before_close_flag,omitempty" `

    /*
        付款前扩展标识位     */
    BeforePayFlag  *int64 `json:"before_pay_flag,omitempty" `

    /*
        发货前扩展标识位     */
    BeforeShipFlag  *int64 `json:"before_ship_flag,omitempty" `

    /*
        确认收货前扩展标识位     */
    BeforeConfirmFlag  *int64 `json:"before_confirm_flag,omitempty" `

    /*
        评价前扩展标识位     */
    BeforeRateFlag  *int64 `json:"before_rate_flag,omitempty" `

    /*
        退款前扩展标识位     */
    BeforeRefundFlag  *int64 `json:"before_refund_flag,omitempty" `

    /*
        修改前扩展标识位     */
    BeforeModifyFlag  *int64 `json:"before_modify_flag,omitempty" `

    /*
        第三方状态，第三方自由定义     */
    ThirdPartyStatus  *int64 `json:"third_party_status,omitempty" `

    /*
        第三方个性化数据     */
    ExtraData  *string `json:"extra_data,omitempty" `

    /*
        attributes标记     */
    ExtAttributes  *string `json:"ext_attributes,omitempty" `

}

func (s *TaobaoTradeFullinfoGetTradeExt) SetBeforeEnableFlag(v int64) *TaobaoTradeFullinfoGetTradeExt {
    s.BeforeEnableFlag = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTradeExt) SetBeforeCloseFlag(v int64) *TaobaoTradeFullinfoGetTradeExt {
    s.BeforeCloseFlag = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTradeExt) SetBeforePayFlag(v int64) *TaobaoTradeFullinfoGetTradeExt {
    s.BeforePayFlag = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTradeExt) SetBeforeShipFlag(v int64) *TaobaoTradeFullinfoGetTradeExt {
    s.BeforeShipFlag = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTradeExt) SetBeforeConfirmFlag(v int64) *TaobaoTradeFullinfoGetTradeExt {
    s.BeforeConfirmFlag = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTradeExt) SetBeforeRateFlag(v int64) *TaobaoTradeFullinfoGetTradeExt {
    s.BeforeRateFlag = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTradeExt) SetBeforeRefundFlag(v int64) *TaobaoTradeFullinfoGetTradeExt {
    s.BeforeRefundFlag = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTradeExt) SetBeforeModifyFlag(v int64) *TaobaoTradeFullinfoGetTradeExt {
    s.BeforeModifyFlag = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTradeExt) SetThirdPartyStatus(v int64) *TaobaoTradeFullinfoGetTradeExt {
    s.ThirdPartyStatus = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTradeExt) SetExtraData(v string) *TaobaoTradeFullinfoGetTradeExt {
    s.ExtraData = &v
    return s
}
func (s *TaobaoTradeFullinfoGetTradeExt) SetExtAttributes(v string) *TaobaoTradeFullinfoGetTradeExt {
    s.ExtAttributes = &v
    return s
}
