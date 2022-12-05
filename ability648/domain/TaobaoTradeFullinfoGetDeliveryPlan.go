package domain

import (
        "topsdk/util"
    )

type TaobaoTradeFullinfoGetDeliveryPlan struct {
    /*
        计划id     */
    PlanId  *int64 `json:"plan_id,omitempty" `

    /*
        交易订单号     */
    OrderId  *int64 `json:"order_id,omitempty" `

    /*
        发货期数     */
    CurrPhase  *int64 `json:"curr_phase,omitempty" `

    /*
        外部物流单号     */
    OutLogisticsId  *string `json:"out_logistics_id,omitempty" `

    /*
        备货期开始时间     */
    PrepareTimeBegin  *util.LocalTime `json:"prepare_time_begin,omitempty" `

    /*
        计划发货时间     */
    ShipTimeBegin  *util.LocalTime `json:"ship_time_begin,omitempty" `

    /*
        实际发货时间     */
    ActualShipTime  *util.LocalTime `json:"actual_ship_time,omitempty" `

    /*
        发货商品数量     */
    GoodsNum  *int64 `json:"goods_num,omitempty" `

    /*
        计划状态     */
    PlanStatus  *int64 `json:"plan_status,omitempty" `

    /*
        计划退款状态     */
    PlanRefundStatus  *int64 `json:"plan_refund_status,omitempty" `

    /*
        收货人姓名     */
    ReceiverName  *string `json:"receiver_name,omitempty" `

    /*
        收货人手机号     */
    ReceiverMobile  *string `json:"receiver_mobile,omitempty" `

    /*
        收货人电话     */
    ReceiverPhone  *string `json:"receiver_phone,omitempty" `

    /*
        收货详细地址     */
    ReceiverAddress  *string `json:"receiver_address,omitempty" `

    /*
        收货地址     */
    ReceiverTown  *string `json:"receiver_town,omitempty" `

    /*
        收货地址     */
    ReceiverDistrict  *string `json:"receiver_district,omitempty" `

    /*
        收货地址     */
    ReceiverCity  *string `json:"receiver_city,omitempty" `

    /*
        收货地址     */
    ReceiverState  *string `json:"receiver_state,omitempty" `

    /*
        收货地址     */
    ReceiverCountry  *string `json:"receiver_country,omitempty" `

    /*
        aid     */
    Oaid  *string `json:"oaid,omitempty" `

    /*
        隐私号     */
    PrivacyNum  *string `json:"privacy_num,omitempty" `

    /*
        隐私号过期时间     */
    PrivacyExpireTime  *util.LocalTime `json:"privacy_expire_time,omitempty" `

}

func (s *TaobaoTradeFullinfoGetDeliveryPlan) SetPlanId(v int64) *TaobaoTradeFullinfoGetDeliveryPlan {
    s.PlanId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetDeliveryPlan) SetOrderId(v int64) *TaobaoTradeFullinfoGetDeliveryPlan {
    s.OrderId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetDeliveryPlan) SetCurrPhase(v int64) *TaobaoTradeFullinfoGetDeliveryPlan {
    s.CurrPhase = &v
    return s
}
func (s *TaobaoTradeFullinfoGetDeliveryPlan) SetOutLogisticsId(v string) *TaobaoTradeFullinfoGetDeliveryPlan {
    s.OutLogisticsId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetDeliveryPlan) SetPrepareTimeBegin(v util.LocalTime) *TaobaoTradeFullinfoGetDeliveryPlan {
    s.PrepareTimeBegin = &v
    return s
}
func (s *TaobaoTradeFullinfoGetDeliveryPlan) SetShipTimeBegin(v util.LocalTime) *TaobaoTradeFullinfoGetDeliveryPlan {
    s.ShipTimeBegin = &v
    return s
}
func (s *TaobaoTradeFullinfoGetDeliveryPlan) SetActualShipTime(v util.LocalTime) *TaobaoTradeFullinfoGetDeliveryPlan {
    s.ActualShipTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetDeliveryPlan) SetGoodsNum(v int64) *TaobaoTradeFullinfoGetDeliveryPlan {
    s.GoodsNum = &v
    return s
}
func (s *TaobaoTradeFullinfoGetDeliveryPlan) SetPlanStatus(v int64) *TaobaoTradeFullinfoGetDeliveryPlan {
    s.PlanStatus = &v
    return s
}
func (s *TaobaoTradeFullinfoGetDeliveryPlan) SetPlanRefundStatus(v int64) *TaobaoTradeFullinfoGetDeliveryPlan {
    s.PlanRefundStatus = &v
    return s
}
func (s *TaobaoTradeFullinfoGetDeliveryPlan) SetReceiverName(v string) *TaobaoTradeFullinfoGetDeliveryPlan {
    s.ReceiverName = &v
    return s
}
func (s *TaobaoTradeFullinfoGetDeliveryPlan) SetReceiverMobile(v string) *TaobaoTradeFullinfoGetDeliveryPlan {
    s.ReceiverMobile = &v
    return s
}
func (s *TaobaoTradeFullinfoGetDeliveryPlan) SetReceiverPhone(v string) *TaobaoTradeFullinfoGetDeliveryPlan {
    s.ReceiverPhone = &v
    return s
}
func (s *TaobaoTradeFullinfoGetDeliveryPlan) SetReceiverAddress(v string) *TaobaoTradeFullinfoGetDeliveryPlan {
    s.ReceiverAddress = &v
    return s
}
func (s *TaobaoTradeFullinfoGetDeliveryPlan) SetReceiverTown(v string) *TaobaoTradeFullinfoGetDeliveryPlan {
    s.ReceiverTown = &v
    return s
}
func (s *TaobaoTradeFullinfoGetDeliveryPlan) SetReceiverDistrict(v string) *TaobaoTradeFullinfoGetDeliveryPlan {
    s.ReceiverDistrict = &v
    return s
}
func (s *TaobaoTradeFullinfoGetDeliveryPlan) SetReceiverCity(v string) *TaobaoTradeFullinfoGetDeliveryPlan {
    s.ReceiverCity = &v
    return s
}
func (s *TaobaoTradeFullinfoGetDeliveryPlan) SetReceiverState(v string) *TaobaoTradeFullinfoGetDeliveryPlan {
    s.ReceiverState = &v
    return s
}
func (s *TaobaoTradeFullinfoGetDeliveryPlan) SetReceiverCountry(v string) *TaobaoTradeFullinfoGetDeliveryPlan {
    s.ReceiverCountry = &v
    return s
}
func (s *TaobaoTradeFullinfoGetDeliveryPlan) SetOaid(v string) *TaobaoTradeFullinfoGetDeliveryPlan {
    s.Oaid = &v
    return s
}
func (s *TaobaoTradeFullinfoGetDeliveryPlan) SetPrivacyNum(v string) *TaobaoTradeFullinfoGetDeliveryPlan {
    s.PrivacyNum = &v
    return s
}
func (s *TaobaoTradeFullinfoGetDeliveryPlan) SetPrivacyExpireTime(v util.LocalTime) *TaobaoTradeFullinfoGetDeliveryPlan {
    s.PrivacyExpireTime = &v
    return s
}
