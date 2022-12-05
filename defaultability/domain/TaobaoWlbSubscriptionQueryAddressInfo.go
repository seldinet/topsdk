package domain


type TaobaoWlbSubscriptionQueryAddressInfo struct {
    /*
        省     */
    Province  *string `json:"province,omitempty" `

    /*
        市     */
    City  *string `json:"city,omitempty" `

    /*
        区     */
    Borough  *string `json:"borough,omitempty" `

    /*
        详细地址     */
    Address  *string `json:"address,omitempty" `

    /*
        邮编     */
    Zip  *string `json:"zip,omitempty" `

}

func (s *TaobaoWlbSubscriptionQueryAddressInfo) SetProvince(v string) *TaobaoWlbSubscriptionQueryAddressInfo {
    s.Province = &v
    return s
}
func (s *TaobaoWlbSubscriptionQueryAddressInfo) SetCity(v string) *TaobaoWlbSubscriptionQueryAddressInfo {
    s.City = &v
    return s
}
func (s *TaobaoWlbSubscriptionQueryAddressInfo) SetBorough(v string) *TaobaoWlbSubscriptionQueryAddressInfo {
    s.Borough = &v
    return s
}
func (s *TaobaoWlbSubscriptionQueryAddressInfo) SetAddress(v string) *TaobaoWlbSubscriptionQueryAddressInfo {
    s.Address = &v
    return s
}
func (s *TaobaoWlbSubscriptionQueryAddressInfo) SetZip(v string) *TaobaoWlbSubscriptionQueryAddressInfo {
    s.Zip = &v
    return s
}
