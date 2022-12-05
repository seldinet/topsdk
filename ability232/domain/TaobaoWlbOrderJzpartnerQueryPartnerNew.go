package domain


type TaobaoWlbOrderJzpartnerQueryPartnerNew struct {
    /*
        是否虚拟物流商     */
    IsVirtualTp  *bool `json:"is_virtual_tp,omitempty" `

    /*
        服务类型     */
    ServiceType  *int64 `json:"service_type,omitempty" `

    /*
        物流商名称     */
    TpName  *string `json:"tp_name,omitempty" `

    /*
        物流商编码     */
    TpCode  *string `json:"tp_code,omitempty" `

}

func (s *TaobaoWlbOrderJzpartnerQueryPartnerNew) SetIsVirtualTp(v bool) *TaobaoWlbOrderJzpartnerQueryPartnerNew {
    s.IsVirtualTp = &v
    return s
}
func (s *TaobaoWlbOrderJzpartnerQueryPartnerNew) SetServiceType(v int64) *TaobaoWlbOrderJzpartnerQueryPartnerNew {
    s.ServiceType = &v
    return s
}
func (s *TaobaoWlbOrderJzpartnerQueryPartnerNew) SetTpName(v string) *TaobaoWlbOrderJzpartnerQueryPartnerNew {
    s.TpName = &v
    return s
}
func (s *TaobaoWlbOrderJzpartnerQueryPartnerNew) SetTpCode(v string) *TaobaoWlbOrderJzpartnerQueryPartnerNew {
    s.TpCode = &v
    return s
}
