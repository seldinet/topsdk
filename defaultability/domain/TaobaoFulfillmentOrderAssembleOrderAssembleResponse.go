package domain


type TaobaoFulfillmentOrderAssembleOrderAssembleResponse struct {
    /*
        回传结果List     */
    OrderGroupResponses  *[]TaobaoFulfillmentOrderAssembleOrderGroupResponse `json:"order_group_responses,omitempty" `

}

func (s *TaobaoFulfillmentOrderAssembleOrderAssembleResponse) SetOrderGroupResponses(v []TaobaoFulfillmentOrderAssembleOrderGroupResponse) *TaobaoFulfillmentOrderAssembleOrderAssembleResponse {
    s.OrderGroupResponses = &v
    return s
}
