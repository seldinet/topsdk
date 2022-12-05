package domain


type TaobaoWlbTmsorderQueryWlbTmsOrder struct {
    /*
        物流订单编号     */
    OrderCode  *string `json:"order_code,omitempty" `

    /*
        物流订单的所有者ID,货主     */
    UserId  *int64 `json:"user_id,omitempty" `

    /*
        运单号     */
    Code  *string `json:"code,omitempty" `

    /*
        物流公司编码     */
    CompanyCode  *string `json:"company_code,omitempty" `

}

func (s *TaobaoWlbTmsorderQueryWlbTmsOrder) SetOrderCode(v string) *TaobaoWlbTmsorderQueryWlbTmsOrder {
    s.OrderCode = &v
    return s
}
func (s *TaobaoWlbTmsorderQueryWlbTmsOrder) SetUserId(v int64) *TaobaoWlbTmsorderQueryWlbTmsOrder {
    s.UserId = &v
    return s
}
func (s *TaobaoWlbTmsorderQueryWlbTmsOrder) SetCode(v string) *TaobaoWlbTmsorderQueryWlbTmsOrder {
    s.Code = &v
    return s
}
func (s *TaobaoWlbTmsorderQueryWlbTmsOrder) SetCompanyCode(v string) *TaobaoWlbTmsorderQueryWlbTmsOrder {
    s.CompanyCode = &v
    return s
}
