package domain


type TaobaoLogisticsAddressReachablebatchGetAddressReachableResult struct {
    /*
        是否成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        服务对应的数字编码，如揽派范围对应88     */
    ServiceType  *string `json:"service_type,omitempty" `

    /*
        服务是否可达， 0 - 不可达 1 - 可达 2 - 不确定 3 - 未配置     */
    Reachable  *string `json:"reachable,omitempty" `

    /*
        物流公司编码ID     */
    PartnerId  *int64 `json:"partner_id,omitempty" `

    /*
        物流公司代号     */
    PartnerCode  *string `json:"partner_code,omitempty" `

    /*
        错误信息     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

    /*
        错误码     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        区域编码     */
    DivisionId  *int64 `json:"division_id,omitempty" `

}

func (s *TaobaoLogisticsAddressReachablebatchGetAddressReachableResult) SetSuccess(v bool) *TaobaoLogisticsAddressReachablebatchGetAddressReachableResult {
    s.Success = &v
    return s
}
func (s *TaobaoLogisticsAddressReachablebatchGetAddressReachableResult) SetServiceType(v string) *TaobaoLogisticsAddressReachablebatchGetAddressReachableResult {
    s.ServiceType = &v
    return s
}
func (s *TaobaoLogisticsAddressReachablebatchGetAddressReachableResult) SetReachable(v string) *TaobaoLogisticsAddressReachablebatchGetAddressReachableResult {
    s.Reachable = &v
    return s
}
func (s *TaobaoLogisticsAddressReachablebatchGetAddressReachableResult) SetPartnerId(v int64) *TaobaoLogisticsAddressReachablebatchGetAddressReachableResult {
    s.PartnerId = &v
    return s
}
func (s *TaobaoLogisticsAddressReachablebatchGetAddressReachableResult) SetPartnerCode(v string) *TaobaoLogisticsAddressReachablebatchGetAddressReachableResult {
    s.PartnerCode = &v
    return s
}
func (s *TaobaoLogisticsAddressReachablebatchGetAddressReachableResult) SetErrorMsg(v string) *TaobaoLogisticsAddressReachablebatchGetAddressReachableResult {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoLogisticsAddressReachablebatchGetAddressReachableResult) SetErrorCode(v string) *TaobaoLogisticsAddressReachablebatchGetAddressReachableResult {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoLogisticsAddressReachablebatchGetAddressReachableResult) SetDivisionId(v int64) *TaobaoLogisticsAddressReachablebatchGetAddressReachableResult {
    s.DivisionId = &v
    return s
}
