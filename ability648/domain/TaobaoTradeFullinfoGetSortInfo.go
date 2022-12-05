package domain


type TaobaoTradeFullinfoGetSortInfo struct {
    /*
        daySortVal     */
    DaySortVal  *int64 `json:"day_sort_val,omitempty" `

    /*
        hourSortVal     */
    HourSortVal  *int64 `json:"hour_sort_val,omitempty" `

}

func (s *TaobaoTradeFullinfoGetSortInfo) SetDaySortVal(v int64) *TaobaoTradeFullinfoGetSortInfo {
    s.DaySortVal = &v
    return s
}
func (s *TaobaoTradeFullinfoGetSortInfo) SetHourSortVal(v int64) *TaobaoTradeFullinfoGetSortInfo {
    s.HourSortVal = &v
    return s
}
