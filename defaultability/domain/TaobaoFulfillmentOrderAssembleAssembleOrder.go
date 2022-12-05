package domain


type TaobaoFulfillmentOrderAssembleAssembleOrder struct {
    /*
        组合id，服务商内部的合单操作id，取消合单会根据group_id进行删除操作。     */
    GroupId  *string `json:"group_id,omitempty" `

    /*
        合单订单列表，一个列表最多200     */
    OrderList  *[]TaobaoFulfillmentOrderAssembleOrderGroup `json:"order_list,omitempty" `

}

func (s *TaobaoFulfillmentOrderAssembleAssembleOrder) SetGroupId(v string) *TaobaoFulfillmentOrderAssembleAssembleOrder {
    s.GroupId = &v
    return s
}
func (s *TaobaoFulfillmentOrderAssembleAssembleOrder) SetOrderList(v []TaobaoFulfillmentOrderAssembleOrderGroup) *TaobaoFulfillmentOrderAssembleAssembleOrder {
    s.OrderList = &v
    return s
}
