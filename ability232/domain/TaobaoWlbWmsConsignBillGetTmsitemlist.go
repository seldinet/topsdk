package domain


type TaobaoWlbWmsConsignBillGetTmsitemlist struct {
    /*
        包裹里面商品     */
    TmsItem  *TaobaoWlbWmsConsignBillGetTmsitem `json:"tms_item,omitempty" `

}

func (s *TaobaoWlbWmsConsignBillGetTmsitemlist) SetTmsItem(v TaobaoWlbWmsConsignBillGetTmsitem) *TaobaoWlbWmsConsignBillGetTmsitemlist {
    s.TmsItem = &v
    return s
}
