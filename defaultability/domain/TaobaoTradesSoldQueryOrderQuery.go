package domain

import (
        "topsdk/util"
    )

type TaobaoTradesSoldQueryOrderQuery struct {
    /*
        收件人电话号码     */
    ReceiverPhone  *string `json:"receiver_phone,omitempty" `

    /*
        收件人的手机号     */
    ReceiverMobile  *string `json:"receiver_mobile,omitempty" `

    /*
        收件人的姓名     */
    ReceiverName  *string `json:"receiver_name,omitempty" `

    /*
        查询三个月内交易创建时间开始。格式:yyyy-MM-dd HH:mm:ss     */
    EndCreated  *util.LocalTime `json:"end_created,omitempty" `

    /*
        查询交易创建时间结束。格式:yyyy-MM-dd HH:mm:ss     */
    StartCreated  *util.LocalTime `json:"start_created,omitempty" `

}

func (s *TaobaoTradesSoldQueryOrderQuery) SetReceiverPhone(v string) *TaobaoTradesSoldQueryOrderQuery {
    s.ReceiverPhone = &v
    return s
}
func (s *TaobaoTradesSoldQueryOrderQuery) SetReceiverMobile(v string) *TaobaoTradesSoldQueryOrderQuery {
    s.ReceiverMobile = &v
    return s
}
func (s *TaobaoTradesSoldQueryOrderQuery) SetReceiverName(v string) *TaobaoTradesSoldQueryOrderQuery {
    s.ReceiverName = &v
    return s
}
func (s *TaobaoTradesSoldQueryOrderQuery) SetEndCreated(v util.LocalTime) *TaobaoTradesSoldQueryOrderQuery {
    s.EndCreated = &v
    return s
}
func (s *TaobaoTradesSoldQueryOrderQuery) SetStartCreated(v util.LocalTime) *TaobaoTradesSoldQueryOrderQuery {
    s.StartCreated = &v
    return s
}
