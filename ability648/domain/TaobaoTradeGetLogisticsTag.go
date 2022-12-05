package domain


type TaobaoTradeGetLogisticsTag struct {
    /*
        主订单或子订单的订单号     */
    OrderId  *string `json:"order_id,omitempty" `

    /*
        服务标签     */
    LogisticServiceTagList  *[]TaobaoTradeGetLogisticServiceTag `json:"logistic_service_tag_list,omitempty" `

}

func (s *TaobaoTradeGetLogisticsTag) SetOrderId(v string) *TaobaoTradeGetLogisticsTag {
    s.OrderId = &v
    return s
}
func (s *TaobaoTradeGetLogisticsTag) SetLogisticServiceTagList(v []TaobaoTradeGetLogisticServiceTag) *TaobaoTradeGetLogisticsTag {
    s.LogisticServiceTagList = &v
    return s
}
