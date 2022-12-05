package domain


type TaobaoWlbWmsSnInfoQuerySninfolist struct {
    /*
        SN信息     */
    SnInfo  *TaobaoWlbWmsSnInfoQuerySninfo `json:"sn_info,omitempty" `

}

func (s *TaobaoWlbWmsSnInfoQuerySninfolist) SetSnInfo(v TaobaoWlbWmsSnInfoQuerySninfo) *TaobaoWlbWmsSnInfoQuerySninfolist {
    s.SnInfo = &v
    return s
}
