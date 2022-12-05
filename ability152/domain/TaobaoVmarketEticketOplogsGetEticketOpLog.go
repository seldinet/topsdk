package domain

import (
        "topsdk/util"
    )

type TaobaoVmarketEticketOplogsGetEticketOpLog struct {
    /*
        操作流水号     */
    ConsumeSerialNum  *string `json:"consume_serial_num,omitempty" `

    /*
        操作金额     */
    Amount  *string `json:"amount,omitempty" `

    /*
        操作数量     */
    Num  *int64 `json:"num,omitempty" `

    /*
        操作类型 1:核销 2:冲正     */
    OpType  *int64 `json:"op_type,omitempty" `

    /*
        操作时间     */
    OpTime  *util.LocalTime `json:"op_time,omitempty" `

    /*
        操作员身份ID     */
    PosId  *string `json:"pos_id,omitempty" `

    /*
        订单ID     */
    OrderId  *int64 `json:"order_id,omitempty" `

    /*
        手机号码后四位     */
    Mobile  *string `json:"mobile,omitempty" `

}

func (s *TaobaoVmarketEticketOplogsGetEticketOpLog) SetConsumeSerialNum(v string) *TaobaoVmarketEticketOplogsGetEticketOpLog {
    s.ConsumeSerialNum = &v
    return s
}
func (s *TaobaoVmarketEticketOplogsGetEticketOpLog) SetAmount(v string) *TaobaoVmarketEticketOplogsGetEticketOpLog {
    s.Amount = &v
    return s
}
func (s *TaobaoVmarketEticketOplogsGetEticketOpLog) SetNum(v int64) *TaobaoVmarketEticketOplogsGetEticketOpLog {
    s.Num = &v
    return s
}
func (s *TaobaoVmarketEticketOplogsGetEticketOpLog) SetOpType(v int64) *TaobaoVmarketEticketOplogsGetEticketOpLog {
    s.OpType = &v
    return s
}
func (s *TaobaoVmarketEticketOplogsGetEticketOpLog) SetOpTime(v util.LocalTime) *TaobaoVmarketEticketOplogsGetEticketOpLog {
    s.OpTime = &v
    return s
}
func (s *TaobaoVmarketEticketOplogsGetEticketOpLog) SetPosId(v string) *TaobaoVmarketEticketOplogsGetEticketOpLog {
    s.PosId = &v
    return s
}
func (s *TaobaoVmarketEticketOplogsGetEticketOpLog) SetOrderId(v int64) *TaobaoVmarketEticketOplogsGetEticketOpLog {
    s.OrderId = &v
    return s
}
func (s *TaobaoVmarketEticketOplogsGetEticketOpLog) SetMobile(v string) *TaobaoVmarketEticketOplogsGetEticketOpLog {
    s.Mobile = &v
    return s
}
