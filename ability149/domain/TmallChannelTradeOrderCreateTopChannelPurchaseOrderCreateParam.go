package domain


type TmallChannelTradeOrderCreateTopChannelPurchaseOrderCreateParam struct {
    /*
        是否自动审批 defalutValue:false    */
    AutoAudit  *bool `json:"auto_audit,omitempty" `

    /*
        分销商淘宝数字ID，如为空，down_user_nick必须输入     */
    DownAccountId  *int64 `json:"down_account_id,omitempty" `

    /*
        请求编码     */
    RequestNo  *string `json:"request_no,omitempty" `

    /*
        采购明细     */
    Items  *[]TmallChannelTradeOrderCreateTopChannelSubPurchaseOrderCreateParam `json:"items,omitempty" `

    /*
        分销商渠道角色，默认分销终端商 defalutValue:2    */
    DownRoleType  *int64 `json:"down_role_type,omitempty" `

    /*
        分销商用户类型，默认淘宝用户 defalutValue:1    */
    DownAccountType  *int64 `json:"down_account_type,omitempty" `

    /*
        分销商昵称     */
    DownUserNick  *string `json:"down_user_nick,omitempty" `

    /*
        内部编码     */
    InternalCode  *string `json:"internal_code,omitempty" `

    /*
        渠道编码，11-线下网批     */
    Channel  *int64 `json:"channel,omitempty" `

}

func (s *TmallChannelTradeOrderCreateTopChannelPurchaseOrderCreateParam) SetAutoAudit(v bool) *TmallChannelTradeOrderCreateTopChannelPurchaseOrderCreateParam {
    s.AutoAudit = &v
    return s
}
func (s *TmallChannelTradeOrderCreateTopChannelPurchaseOrderCreateParam) SetDownAccountId(v int64) *TmallChannelTradeOrderCreateTopChannelPurchaseOrderCreateParam {
    s.DownAccountId = &v
    return s
}
func (s *TmallChannelTradeOrderCreateTopChannelPurchaseOrderCreateParam) SetRequestNo(v string) *TmallChannelTradeOrderCreateTopChannelPurchaseOrderCreateParam {
    s.RequestNo = &v
    return s
}
func (s *TmallChannelTradeOrderCreateTopChannelPurchaseOrderCreateParam) SetItems(v []TmallChannelTradeOrderCreateTopChannelSubPurchaseOrderCreateParam) *TmallChannelTradeOrderCreateTopChannelPurchaseOrderCreateParam {
    s.Items = &v
    return s
}
func (s *TmallChannelTradeOrderCreateTopChannelPurchaseOrderCreateParam) SetDownRoleType(v int64) *TmallChannelTradeOrderCreateTopChannelPurchaseOrderCreateParam {
    s.DownRoleType = &v
    return s
}
func (s *TmallChannelTradeOrderCreateTopChannelPurchaseOrderCreateParam) SetDownAccountType(v int64) *TmallChannelTradeOrderCreateTopChannelPurchaseOrderCreateParam {
    s.DownAccountType = &v
    return s
}
func (s *TmallChannelTradeOrderCreateTopChannelPurchaseOrderCreateParam) SetDownUserNick(v string) *TmallChannelTradeOrderCreateTopChannelPurchaseOrderCreateParam {
    s.DownUserNick = &v
    return s
}
func (s *TmallChannelTradeOrderCreateTopChannelPurchaseOrderCreateParam) SetInternalCode(v string) *TmallChannelTradeOrderCreateTopChannelPurchaseOrderCreateParam {
    s.InternalCode = &v
    return s
}
func (s *TmallChannelTradeOrderCreateTopChannelPurchaseOrderCreateParam) SetChannel(v int64) *TmallChannelTradeOrderCreateTopChannelPurchaseOrderCreateParam {
    s.Channel = &v
    return s
}
