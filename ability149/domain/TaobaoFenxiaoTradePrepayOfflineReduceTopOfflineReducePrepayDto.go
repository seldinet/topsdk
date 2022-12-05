package domain

import (
        "topsdk/util"
    )

type TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto struct {
    /*
        资金流水类型：1.纸质承兑； 2.电子承兑；3.现金；4.优惠返点；5.奖励     */
    FlowType  *int64 `json:"flow_type,omitempty" `

    /*
        线下流水类型 1票据作废 2线下使用     */
    OfflinePrepayDetailType  *int64 `json:"offline_prepay_detail_type,omitempty" `

    /*
        金额，单位分（必须为负数）     */
    TicketMoney  *int64 `json:"ticket_money,omitempty" `

    /*
        收款人账号     */
    ReceiverAccountNum  *string `json:"receiver_account_num,omitempty" `

    /*
        外部系统支付流水Id，用于商家上传流水时去重(外部保证唯一）     */
    OuterPayId  *string `json:"outer_pay_id,omitempty" `

    /*
        销商淘宝nick     */
    DistNick  *string `json:"dist_nick,omitempty" `

    /*
        出票人全称     */
    DrawerFullName  *string `json:"drawer_full_name,omitempty" `

    /*
        付款行行号     */
    PayBankNum  *string `json:"pay_bank_num,omitempty" `

    /*
        出票人账号     */
    DrawerAccountNum  *string `json:"drawer_account_num,omitempty" `

    /*
        支付时间     */
    PayTime  *util.LocalTime `json:"pay_time,omitempty" `

    /*
        出票日期     */
    TicketIssueDate  *util.LocalTime `json:"ticket_issue_date,omitempty" `

    /*
        收款开户银行     */
    ReceiverBankFullName  *string `json:"receiver_bank_full_name,omitempty" `

    /*
        承兑票据号     */
    TicketId  *string `json:"ticket_id,omitempty" `

    /*
        汇票到期日期     */
    AcceptDate  *util.LocalTime `json:"accept_date,omitempty" `

    /*
        收款人全称     */
    ReceiverFullName  *string `json:"receiver_full_name,omitempty" `

    /*
        付款行全称     */
    PayBankFullName  *string `json:"pay_bank_full_name,omitempty" `

}

func (s *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto) SetFlowType(v int64) *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto {
    s.FlowType = &v
    return s
}
func (s *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto) SetOfflinePrepayDetailType(v int64) *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto {
    s.OfflinePrepayDetailType = &v
    return s
}
func (s *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto) SetTicketMoney(v int64) *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto {
    s.TicketMoney = &v
    return s
}
func (s *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto) SetReceiverAccountNum(v string) *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto {
    s.ReceiverAccountNum = &v
    return s
}
func (s *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto) SetOuterPayId(v string) *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto {
    s.OuterPayId = &v
    return s
}
func (s *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto) SetDistNick(v string) *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto {
    s.DistNick = &v
    return s
}
func (s *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto) SetDrawerFullName(v string) *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto {
    s.DrawerFullName = &v
    return s
}
func (s *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto) SetPayBankNum(v string) *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto {
    s.PayBankNum = &v
    return s
}
func (s *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto) SetDrawerAccountNum(v string) *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto {
    s.DrawerAccountNum = &v
    return s
}
func (s *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto) SetPayTime(v util.LocalTime) *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto {
    s.PayTime = &v
    return s
}
func (s *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto) SetTicketIssueDate(v util.LocalTime) *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto {
    s.TicketIssueDate = &v
    return s
}
func (s *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto) SetReceiverBankFullName(v string) *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto {
    s.ReceiverBankFullName = &v
    return s
}
func (s *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto) SetTicketId(v string) *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto {
    s.TicketId = &v
    return s
}
func (s *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto) SetAcceptDate(v util.LocalTime) *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto {
    s.AcceptDate = &v
    return s
}
func (s *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto) SetReceiverFullName(v string) *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto {
    s.ReceiverFullName = &v
    return s
}
func (s *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto) SetPayBankFullName(v string) *TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto {
    s.PayBankFullName = &v
    return s
}
