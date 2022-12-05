package domain


type TaobaoWlbWmsItemCombinationGetResult struct {
    /*
        是否成功     */
    WlSuccess  *bool `json:"wl_success,omitempty" `

    /*
        错误编码     */
    WlErrorCode  *string `json:"wl_error_code,omitempty" `

    /*
        错误信息     */
    WlErrorMsg  *string `json:"wl_error_msg,omitempty" `

    /*
        子货品列表     */
    SubItemList  *[]TaobaoWlbWmsItemCombinationGetSubItemList `json:"sub_item_list,omitempty" `

}

func (s *TaobaoWlbWmsItemCombinationGetResult) SetWlSuccess(v bool) *TaobaoWlbWmsItemCombinationGetResult {
    s.WlSuccess = &v
    return s
}
func (s *TaobaoWlbWmsItemCombinationGetResult) SetWlErrorCode(v string) *TaobaoWlbWmsItemCombinationGetResult {
    s.WlErrorCode = &v
    return s
}
func (s *TaobaoWlbWmsItemCombinationGetResult) SetWlErrorMsg(v string) *TaobaoWlbWmsItemCombinationGetResult {
    s.WlErrorMsg = &v
    return s
}
func (s *TaobaoWlbWmsItemCombinationGetResult) SetSubItemList(v []TaobaoWlbWmsItemCombinationGetSubItemList) *TaobaoWlbWmsItemCombinationGetResult {
    s.SubItemList = &v
    return s
}
