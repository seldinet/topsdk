package domain


type TaobaoWlbWmsConsignBillGetTmsorderlist struct {
    /*
        运单信息列表     */
    TmsOrder  *TaobaoWlbWmsConsignBillGetTmsorder `json:"tms_order,omitempty" `

}

func (s *TaobaoWlbWmsConsignBillGetTmsorderlist) SetTmsOrder(v TaobaoWlbWmsConsignBillGetTmsorder) *TaobaoWlbWmsConsignBillGetTmsorderlist {
    s.TmsOrder = &v
    return s
}
