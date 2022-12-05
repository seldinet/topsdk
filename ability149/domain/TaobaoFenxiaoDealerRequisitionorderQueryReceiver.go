package domain


type TaobaoFenxiaoDealerRequisitionorderQueryReceiver struct {
    /*
        邮政编码     */
    Zip  *string `json:"zip,omitempty" `

    /*
        移动电话     */
    MobilePhone  *string `json:"mobile_phone,omitempty" `

    /*
        固定电话     */
    Phone  *string `json:"phone,omitempty" `

    /*
        收货人的详细地址信息     */
    Address  *string `json:"address,omitempty" `

    /*
        收货人全名     */
    Name  *string `json:"name,omitempty" `

    /*
        收货人所在省份     */
    State  *string `json:"state,omitempty" `

    /*
        区/县     */
    District  *string `json:"district,omitempty" `

    /*
        证件号     */
    CardId  *string `json:"card_id,omitempty" `

    /*
        收货人的城市     */
    City  *string `json:"city,omitempty" `

}

func (s *TaobaoFenxiaoDealerRequisitionorderQueryReceiver) SetZip(v string) *TaobaoFenxiaoDealerRequisitionorderQueryReceiver {
    s.Zip = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryReceiver) SetMobilePhone(v string) *TaobaoFenxiaoDealerRequisitionorderQueryReceiver {
    s.MobilePhone = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryReceiver) SetPhone(v string) *TaobaoFenxiaoDealerRequisitionorderQueryReceiver {
    s.Phone = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryReceiver) SetAddress(v string) *TaobaoFenxiaoDealerRequisitionorderQueryReceiver {
    s.Address = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryReceiver) SetName(v string) *TaobaoFenxiaoDealerRequisitionorderQueryReceiver {
    s.Name = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryReceiver) SetState(v string) *TaobaoFenxiaoDealerRequisitionorderQueryReceiver {
    s.State = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryReceiver) SetDistrict(v string) *TaobaoFenxiaoDealerRequisitionorderQueryReceiver {
    s.District = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryReceiver) SetCardId(v string) *TaobaoFenxiaoDealerRequisitionorderQueryReceiver {
    s.CardId = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryReceiver) SetCity(v string) *TaobaoFenxiaoDealerRequisitionorderQueryReceiver {
    s.City = &v
    return s
}
