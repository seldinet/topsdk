package domain


type TaobaoWlbWmsConsignBillGetConsignsendinfolist struct {
    /*
        物流订单发货信息     */
    ConsignSendInfo  *TaobaoWlbWmsConsignBillGetConsignsendinfo `json:"consign_send_info,omitempty" `

}

func (s *TaobaoWlbWmsConsignBillGetConsignsendinfolist) SetConsignSendInfo(v TaobaoWlbWmsConsignBillGetConsignsendinfo) *TaobaoWlbWmsConsignBillGetConsignsendinfolist {
    s.ConsignSendInfo = &v
    return s
}
