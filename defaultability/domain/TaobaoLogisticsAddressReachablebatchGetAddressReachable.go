package domain


type TaobaoLogisticsAddressReachablebatchGetAddressReachable struct {
    /*
        发货地，标准区域编码(四级行政)，可以通过taobao.areas.get获取，如浙江省杭州市余杭区闲林街道为 330110011     */
    SourceAreaCode  *string `json:"source_area_code,omitempty" `

    /*
        服务对应的数字编码，如揽派范围对应88     */
    ServiceType  *int64 `json:"service_type,omitempty" `

    /*
        物流公司编码ID，可以从这个接口获取所有物流公司的标准编码taobao.logistics.companies.get，可以传入多个值，以英文逗号分隔，如“1000000952,1000000953”     */
    PartnerId  *string `json:"partner_id,omitempty" `

    /*
        标准区域编码(三级行政区编码或是四级行政区)，可以通过taobao.areas.get获取，如北京市朝阳区为110105     */
    AreaCode  *string `json:"area_code,omitempty" `

    /*
        详细地址     */
    Address  *string `json:"address,omitempty" `

}

func (s *TaobaoLogisticsAddressReachablebatchGetAddressReachable) SetSourceAreaCode(v string) *TaobaoLogisticsAddressReachablebatchGetAddressReachable {
    s.SourceAreaCode = &v
    return s
}
func (s *TaobaoLogisticsAddressReachablebatchGetAddressReachable) SetServiceType(v int64) *TaobaoLogisticsAddressReachablebatchGetAddressReachable {
    s.ServiceType = &v
    return s
}
func (s *TaobaoLogisticsAddressReachablebatchGetAddressReachable) SetPartnerId(v string) *TaobaoLogisticsAddressReachablebatchGetAddressReachable {
    s.PartnerId = &v
    return s
}
func (s *TaobaoLogisticsAddressReachablebatchGetAddressReachable) SetAreaCode(v string) *TaobaoLogisticsAddressReachablebatchGetAddressReachable {
    s.AreaCode = &v
    return s
}
func (s *TaobaoLogisticsAddressReachablebatchGetAddressReachable) SetAddress(v string) *TaobaoLogisticsAddressReachablebatchGetAddressReachable {
    s.Address = &v
    return s
}
