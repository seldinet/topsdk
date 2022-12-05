package domain


type TaobaoLocationRelationQuerySingleResult struct {
    /*
        地点关系     */
    LocationRelationList  *[]TaobaoLocationRelationQueryLocationRelationDto `json:"location_relation_list,omitempty" `

    /*
        错误码     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        是否成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        错误信息     */
    ErrorMessage  *string `json:"error_message,omitempty" `

}

func (s *TaobaoLocationRelationQuerySingleResult) SetLocationRelationList(v []TaobaoLocationRelationQueryLocationRelationDto) *TaobaoLocationRelationQuerySingleResult {
    s.LocationRelationList = &v
    return s
}
func (s *TaobaoLocationRelationQuerySingleResult) SetErrorCode(v string) *TaobaoLocationRelationQuerySingleResult {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoLocationRelationQuerySingleResult) SetSuccess(v bool) *TaobaoLocationRelationQuerySingleResult {
    s.Success = &v
    return s
}
func (s *TaobaoLocationRelationQuerySingleResult) SetErrorMessage(v string) *TaobaoLocationRelationQuerySingleResult {
    s.ErrorMessage = &v
    return s
}
