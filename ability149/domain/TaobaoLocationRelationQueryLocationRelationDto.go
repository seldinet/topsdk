package domain


type TaobaoLocationRelationQueryLocationRelationDto struct {
    /*
        实体类型 2：仓库  6：门店 （target,sorce 二选一填写，都填写报错）     */
    TargetInvStoreType  *int64 `json:"target_inv_store_type,omitempty" `

    /*
        实体code     */
    TargetStoreCode  *string `json:"target_store_code,omitempty" `

    /*
        实体类型 2：仓库  6：门店     */
    SourceInvStoreType  *int64 `json:"source_inv_store_type,omitempty" `

    /*
        实体code     */
    SourceStoreCode  *string `json:"source_store_code,omitempty" `

}

func (s *TaobaoLocationRelationQueryLocationRelationDto) SetTargetInvStoreType(v int64) *TaobaoLocationRelationQueryLocationRelationDto {
    s.TargetInvStoreType = &v
    return s
}
func (s *TaobaoLocationRelationQueryLocationRelationDto) SetTargetStoreCode(v string) *TaobaoLocationRelationQueryLocationRelationDto {
    s.TargetStoreCode = &v
    return s
}
func (s *TaobaoLocationRelationQueryLocationRelationDto) SetSourceInvStoreType(v int64) *TaobaoLocationRelationQueryLocationRelationDto {
    s.SourceInvStoreType = &v
    return s
}
func (s *TaobaoLocationRelationQueryLocationRelationDto) SetSourceStoreCode(v string) *TaobaoLocationRelationQueryLocationRelationDto {
    s.SourceStoreCode = &v
    return s
}
