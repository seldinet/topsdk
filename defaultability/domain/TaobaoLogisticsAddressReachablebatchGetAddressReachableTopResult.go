package domain


type TaobaoLogisticsAddressReachablebatchGetAddressReachableTopResult struct {
    /*
        筛单结果l列表     */
    ReachableResultList  *[]TaobaoLogisticsAddressReachablebatchGetAddressReachableResult `json:"reachable_result_list,omitempty" `

}

func (s *TaobaoLogisticsAddressReachablebatchGetAddressReachableTopResult) SetReachableResultList(v []TaobaoLogisticsAddressReachablebatchGetAddressReachableResult) *TaobaoLogisticsAddressReachablebatchGetAddressReachableTopResult {
    s.ReachableResultList = &v
    return s
}
