package domain


type TaobaoWlbWmsCainiaoBillQueryCainiaoBillQueryOrderinfolist struct {
    /*
        订单信息     */
    OrderInfo  *TaobaoWlbWmsCainiaoBillQueryCainiaoBillQueryOrderinfo `json:"order_info,omitempty" `

}

func (s *TaobaoWlbWmsCainiaoBillQueryCainiaoBillQueryOrderinfolist) SetOrderInfo(v TaobaoWlbWmsCainiaoBillQueryCainiaoBillQueryOrderinfo) *TaobaoWlbWmsCainiaoBillQueryCainiaoBillQueryOrderinfolist {
    s.OrderInfo = &v
    return s
}
