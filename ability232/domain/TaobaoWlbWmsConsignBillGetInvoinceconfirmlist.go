package domain


type TaobaoWlbWmsConsignBillGetInvoinceconfirmlist struct {
    /*
        发票确认信息     */
    InvoinceConfirm  *TaobaoWlbWmsConsignBillGetInvoinceconfirm `json:"invoince_confirm,omitempty" `

}

func (s *TaobaoWlbWmsConsignBillGetInvoinceconfirmlist) SetInvoinceConfirm(v TaobaoWlbWmsConsignBillGetInvoinceconfirm) *TaobaoWlbWmsConsignBillGetInvoinceconfirmlist {
    s.InvoinceConfirm = &v
    return s
}
