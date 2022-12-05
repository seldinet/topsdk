package domain


type TaobaoFulfillmentOrderAssembleOrderGroupResponse struct {
    /*
        入参中的groupId     */
    GroupId  *string `json:"group_id,omitempty" `

    /*
        回传结果     */
    Result  *bool `json:"result,omitempty" `

    /*
        错误码     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        错误信息描述     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

}

func (s *TaobaoFulfillmentOrderAssembleOrderGroupResponse) SetGroupId(v string) *TaobaoFulfillmentOrderAssembleOrderGroupResponse {
    s.GroupId = &v
    return s
}
func (s *TaobaoFulfillmentOrderAssembleOrderGroupResponse) SetResult(v bool) *TaobaoFulfillmentOrderAssembleOrderGroupResponse {
    s.Result = &v
    return s
}
func (s *TaobaoFulfillmentOrderAssembleOrderGroupResponse) SetErrorCode(v string) *TaobaoFulfillmentOrderAssembleOrderGroupResponse {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoFulfillmentOrderAssembleOrderGroupResponse) SetErrorMsg(v string) *TaobaoFulfillmentOrderAssembleOrderGroupResponse {
    s.ErrorMsg = &v
    return s
}
