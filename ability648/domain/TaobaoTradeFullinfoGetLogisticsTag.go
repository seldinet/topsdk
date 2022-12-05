package domain


type TaobaoTradeFullinfoGetLogisticsTag struct {
    /*
        主订单或子订单的订单号     */
    OrderId  *string `json:"order_id,omitempty" `

    /*
        服务标签     */
    LogisticServiceTagList  *[]TaobaoTradeFullinfoGetLogisticServiceTag `json:"logistic_service_tag_list,omitempty" `

}

func (s *TaobaoTradeFullinfoGetLogisticsTag) SetOrderId(v string) *TaobaoTradeFullinfoGetLogisticsTag {
    s.OrderId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsTag) SetLogisticServiceTagList(v []TaobaoTradeFullinfoGetLogisticServiceTag) *TaobaoTradeFullinfoGetLogisticsTag {
    s.LogisticServiceTagList = &v
    return s
}
