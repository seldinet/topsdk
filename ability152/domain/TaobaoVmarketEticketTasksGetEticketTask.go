package domain


type TaobaoVmarketEticketTasksGetEticketTask struct {
    /*
        订单ID     */
    OrderId  *int64 `json:"order_id,omitempty" `

}

func (s *TaobaoVmarketEticketTasksGetEticketTask) SetOrderId(v int64) *TaobaoVmarketEticketTasksGetEticketTask {
    s.OrderId = &v
    return s
}
