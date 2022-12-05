package domain


type TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossOrderitemlist struct {
    /*
        损益详情     */
    OrderItem  *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossOrderitem `json:"order_item,omitempty" `

}

func (s *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossOrderitemlist) SetOrderItem(v TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossOrderitem) *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossOrderitemlist {
    s.OrderItem = &v
    return s
}
