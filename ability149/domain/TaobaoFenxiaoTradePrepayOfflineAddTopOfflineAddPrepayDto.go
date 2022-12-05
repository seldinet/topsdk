package domain

import (
        "topsdk/util"
    )

type TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto struct {
    /*
        资金流水类型：1.纸质承兑； 2.电子承兑；3.现金；4.优惠返点；5.奖励     */
    FlowType  *int64 `json:"flow_type,omitempty" `

    /*
        金额，单位分（必须为正数）     */
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
        承兑票据号（承兑必填）     */
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

func (s *TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto) SetFlowType(v int64) *TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto {
    s.FlowType = &v
    return s
}
func (s *TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto) SetTicketMoney(v int64) *TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto {
    s.TicketMoney = &v
    return s
}
func (s *TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto) SetReceiverAccountNum(v string) *TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto {
    s.ReceiverAccountNum = &v
    return s
}
func (s *TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto) SetOuterPayId(v string) *TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto {
    s.OuterPayId = &v
    return s
}
func (s *TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto) SetDistNick(v string) *TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto {
    s.DistNick = &v
    return s
}
func (s *TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto) SetDrawerFullName(v string) *TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto {
    s.DrawerFullName = &v
    return s
}
func (s *TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto) SetPayBankNum(v string) *TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto {
    s.PayBankNum = &v
    return s
}
func (s *TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto) SetDrawerAccountNum(v string) *TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto {
    s.DrawerAccountNum = &v
    return s
}
func (s *TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto) SetPayTime(v util.LocalTime) *TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto {
    s.PayTime = &v
    return s
}
func (s *TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto) SetTicketIssueDate(v util.LocalTime) *TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto {
    s.TicketIssueDate = &v
    return s
}
func (s *TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto) SetReceiverBankFullName(v string) *TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto {
    s.ReceiverBankFullName = &v
    return s
}
func (s *TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto) SetTicketId(v string) *TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto {
    s.TicketId = &v
    return s
}
func (s *TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto) SetAcceptDate(v util.LocalTime) *TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto {
    s.AcceptDate = &v
    return s
}
func (s *TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto) SetReceiverFullName(v string) *TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto {
    s.ReceiverFullName = &v
    return s
}
func (s *TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto) SetPayBankFullName(v string) *TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto {
    s.PayBankFullName = &v
    return s
}
