package domain


type TaobaoFenxiaoDealerRequisitionorderGetReceiver struct {
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

func (s *TaobaoFenxiaoDealerRequisitionorderGetReceiver) SetZip(v string) *TaobaoFenxiaoDealerRequisitionorderGetReceiver {
    s.Zip = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetReceiver) SetMobilePhone(v string) *TaobaoFenxiaoDealerRequisitionorderGetReceiver {
    s.MobilePhone = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetReceiver) SetPhone(v string) *TaobaoFenxiaoDealerRequisitionorderGetReceiver {
    s.Phone = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetReceiver) SetAddress(v string) *TaobaoFenxiaoDealerRequisitionorderGetReceiver {
    s.Address = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetReceiver) SetName(v string) *TaobaoFenxiaoDealerRequisitionorderGetReceiver {
    s.Name = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetReceiver) SetState(v string) *TaobaoFenxiaoDealerRequisitionorderGetReceiver {
    s.State = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetReceiver) SetDistrict(v string) *TaobaoFenxiaoDealerRequisitionorderGetReceiver {
    s.District = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetReceiver) SetCardId(v string) *TaobaoFenxiaoDealerRequisitionorderGetReceiver {
    s.CardId = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetReceiver) SetCity(v string) *TaobaoFenxiaoDealerRequisitionorderGetReceiver {
    s.City = &v
    return s
}
