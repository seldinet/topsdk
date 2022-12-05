package domain


type TaobaoWlbWmsConsignOrderNotifyConsignorderlist struct {
    /*
        发货订单信息     */
    ConsignOrder  *TaobaoWlbWmsConsignOrderNotifyConsignorder `json:"consign_order,omitempty" `

}

func (s *TaobaoWlbWmsConsignOrderNotifyConsignorderlist) SetConsignOrder(v TaobaoWlbWmsConsignOrderNotifyConsignorder) *TaobaoWlbWmsConsignOrderNotifyConsignorderlist {
    s.ConsignOrder = &v
    return s
}
