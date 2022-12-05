package domain


type TaobaoLogisticsPartnersGetLogisticsPartner struct {
    /*
        物流公司揽收和资费详细信息     */
    Carriage  *TaobaoLogisticsPartnersGetCarriageDetail `json:"carriage,omitempty" `

    /*
        揽收说明信息     */
    CoverRemark  *string `json:"cover_remark,omitempty" `

    /*
        物流公司详细信息     */
    Partner  *TaobaoLogisticsPartnersGetPartnerDetail `json:"partner,omitempty" `

    /*
        不可送达的说明信息     */
    UncoverRemark  *string `json:"uncover_remark,omitempty" `

}

func (s *TaobaoLogisticsPartnersGetLogisticsPartner) SetCarriage(v TaobaoLogisticsPartnersGetCarriageDetail) *TaobaoLogisticsPartnersGetLogisticsPartner {
    s.Carriage = &v
    return s
}
func (s *TaobaoLogisticsPartnersGetLogisticsPartner) SetCoverRemark(v string) *TaobaoLogisticsPartnersGetLogisticsPartner {
    s.CoverRemark = &v
    return s
}
func (s *TaobaoLogisticsPartnersGetLogisticsPartner) SetPartner(v TaobaoLogisticsPartnersGetPartnerDetail) *TaobaoLogisticsPartnersGetLogisticsPartner {
    s.Partner = &v
    return s
}
func (s *TaobaoLogisticsPartnersGetLogisticsPartner) SetUncoverRemark(v string) *TaobaoLogisticsPartnersGetLogisticsPartner {
    s.UncoverRemark = &v
    return s
}
