package domain


type TaobaoItempropsGetItemTaosirDO struct {
    /*
        时间类型：0表示非时间，1表示时间点，2表示时间范围     */
    Type  *int64 `json:"type,omitempty" `

    /*
        数值小数点精度     */
    Precision  *int64 `json:"precision,omitempty" `

    /*
        卖家可选单位List<单位id，单位名>     */
    StdUnitList  *[]TaobaoItempropsGetStdUnit `json:"std_unit_list,omitempty" `

    /*
        表达式元素list     */
    ExprElList  *[]TaobaoItempropsGetExprEl `json:"expr_el_list,omitempty" `

}

func (s *TaobaoItempropsGetItemTaosirDO) SetType(v int64) *TaobaoItempropsGetItemTaosirDO {
    s.Type = &v
    return s
}
func (s *TaobaoItempropsGetItemTaosirDO) SetPrecision(v int64) *TaobaoItempropsGetItemTaosirDO {
    s.Precision = &v
    return s
}
func (s *TaobaoItempropsGetItemTaosirDO) SetStdUnitList(v []TaobaoItempropsGetStdUnit) *TaobaoItempropsGetItemTaosirDO {
    s.StdUnitList = &v
    return s
}
func (s *TaobaoItempropsGetItemTaosirDO) SetExprElList(v []TaobaoItempropsGetExprEl) *TaobaoItempropsGetItemTaosirDO {
    s.ExprElList = &v
    return s
}
