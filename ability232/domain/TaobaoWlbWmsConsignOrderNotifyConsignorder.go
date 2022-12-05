package domain


type TaobaoWlbWmsConsignOrderNotifyConsignorder struct {
    /*
        仓库物流订单信息列表     */
    ConsignOrderItemList  *[]TaobaoWlbWmsConsignOrderNotifyConsignorderitemlist `json:"consign_order_item_list,omitempty" `

    /*
        是否成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        错误编码     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        错误信息     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

    /*
        仓库订单编码     */
    StoreOrderCode  *string `json:"store_order_code,omitempty" `

    /*
        仓库编码     */
    StoreCode  *string `json:"store_code,omitempty" `

    /*
        配送编码     */
    TmsCode  *string `json:"tms_code,omitempty" `

}

func (s *TaobaoWlbWmsConsignOrderNotifyConsignorder) SetConsignOrderItemList(v []TaobaoWlbWmsConsignOrderNotifyConsignorderitemlist) *TaobaoWlbWmsConsignOrderNotifyConsignorder {
    s.ConsignOrderItemList = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyConsignorder) SetSuccess(v bool) *TaobaoWlbWmsConsignOrderNotifyConsignorder {
    s.Success = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyConsignorder) SetErrorCode(v string) *TaobaoWlbWmsConsignOrderNotifyConsignorder {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyConsignorder) SetErrorMsg(v string) *TaobaoWlbWmsConsignOrderNotifyConsignorder {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyConsignorder) SetStoreOrderCode(v string) *TaobaoWlbWmsConsignOrderNotifyConsignorder {
    s.StoreOrderCode = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyConsignorder) SetStoreCode(v string) *TaobaoWlbWmsConsignOrderNotifyConsignorder {
    s.StoreCode = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyConsignorder) SetTmsCode(v string) *TaobaoWlbWmsConsignOrderNotifyConsignorder {
    s.TmsCode = &v
    return s
}
