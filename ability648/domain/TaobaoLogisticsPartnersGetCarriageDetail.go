package domain


type TaobaoLogisticsPartnersGetCarriageDetail struct {
    /*
        续费（单位：元）     */
    AddFee  *int64 `json:"add_fee,omitempty" `

    /*
        续重（单位：千克）     */
    AddWeight  *int64 `json:"add_weight,omitempty" `

    /*
        破损赔付     */
    DamagePayment  *string `json:"damage_payment,omitempty" `

    /*
        物流公司揽收时间段     */
    GotTime  *string `json:"got_time,omitempty" `

    /*
        首费（单位：元）     */
    InitialFee  *int64 `json:"initial_fee,omitempty" `

    /*
        首重（单位：千克）     */
    InitialWeight  *int64 `json:"initial_weight,omitempty" `

    /*
        丢单赔付     */
    LostPayment  *string `json:"lost_payment,omitempty" `

    /*
        快件送达所需的时间(单位：天)     */
    WayDay  *string `json:"way_day,omitempty" `

}

func (s *TaobaoLogisticsPartnersGetCarriageDetail) SetAddFee(v int64) *TaobaoLogisticsPartnersGetCarriageDetail {
    s.AddFee = &v
    return s
}
func (s *TaobaoLogisticsPartnersGetCarriageDetail) SetAddWeight(v int64) *TaobaoLogisticsPartnersGetCarriageDetail {
    s.AddWeight = &v
    return s
}
func (s *TaobaoLogisticsPartnersGetCarriageDetail) SetDamagePayment(v string) *TaobaoLogisticsPartnersGetCarriageDetail {
    s.DamagePayment = &v
    return s
}
func (s *TaobaoLogisticsPartnersGetCarriageDetail) SetGotTime(v string) *TaobaoLogisticsPartnersGetCarriageDetail {
    s.GotTime = &v
    return s
}
func (s *TaobaoLogisticsPartnersGetCarriageDetail) SetInitialFee(v int64) *TaobaoLogisticsPartnersGetCarriageDetail {
    s.InitialFee = &v
    return s
}
func (s *TaobaoLogisticsPartnersGetCarriageDetail) SetInitialWeight(v int64) *TaobaoLogisticsPartnersGetCarriageDetail {
    s.InitialWeight = &v
    return s
}
func (s *TaobaoLogisticsPartnersGetCarriageDetail) SetLostPayment(v string) *TaobaoLogisticsPartnersGetCarriageDetail {
    s.LostPayment = &v
    return s
}
func (s *TaobaoLogisticsPartnersGetCarriageDetail) SetWayDay(v string) *TaobaoLogisticsPartnersGetCarriageDetail {
    s.WayDay = &v
    return s
}
