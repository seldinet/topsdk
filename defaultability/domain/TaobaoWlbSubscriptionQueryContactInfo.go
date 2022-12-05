package domain


type TaobaoWlbSubscriptionQueryContactInfo struct {
    /*
        仓库联系人姓名     */
    Name  *string `json:"name,omitempty" `

    /*
        联系电话     */
    Phone  *string `json:"phone,omitempty" `

}

func (s *TaobaoWlbSubscriptionQueryContactInfo) SetName(v string) *TaobaoWlbSubscriptionQueryContactInfo {
    s.Name = &v
    return s
}
func (s *TaobaoWlbSubscriptionQueryContactInfo) SetPhone(v string) *TaobaoWlbSubscriptionQueryContactInfo {
    s.Phone = &v
    return s
}
