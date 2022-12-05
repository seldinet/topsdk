package domain


type TaobaoWlbWmsConsignOrderNotifyConsignorderitemlist struct {
    /*
        仓库物流订单信息列表     */
    ConsignOrderItem  *TaobaoWlbWmsConsignOrderNotifyConsignorderitem `json:"consign_order_item,omitempty" `

}

func (s *TaobaoWlbWmsConsignOrderNotifyConsignorderitemlist) SetConsignOrderItem(v TaobaoWlbWmsConsignOrderNotifyConsignorderitem) *TaobaoWlbWmsConsignOrderNotifyConsignorderitemlist {
    s.ConsignOrderItem = &v
    return s
}
