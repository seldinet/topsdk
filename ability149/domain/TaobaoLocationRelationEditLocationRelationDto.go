package domain


type TaobaoLocationRelationEditLocationRelationDto struct {
    /*
        状态  0 正常  -1 删除     */
    Status  *int64 `json:"status,omitempty" `

    /*
        实体类型 2：仓库 6：门店     */
    TargetInvStoreType  *int64 `json:"target_inv_store_type,omitempty" `

    /*
        实体code     */
    TargetStoreCode  *string `json:"target_store_code,omitempty" `

    /*
        实体类型 2：仓库 6：门店     */
    SourceInvStoreType  *int64 `json:"source_inv_store_type,omitempty" `

    /*
        实体code     */
    SourceStoreCode  *string `json:"source_store_code,omitempty" `

}

func (s *TaobaoLocationRelationEditLocationRelationDto) SetStatus(v int64) *TaobaoLocationRelationEditLocationRelationDto {
    s.Status = &v
    return s
}
func (s *TaobaoLocationRelationEditLocationRelationDto) SetTargetInvStoreType(v int64) *TaobaoLocationRelationEditLocationRelationDto {
    s.TargetInvStoreType = &v
    return s
}
func (s *TaobaoLocationRelationEditLocationRelationDto) SetTargetStoreCode(v string) *TaobaoLocationRelationEditLocationRelationDto {
    s.TargetStoreCode = &v
    return s
}
func (s *TaobaoLocationRelationEditLocationRelationDto) SetSourceInvStoreType(v int64) *TaobaoLocationRelationEditLocationRelationDto {
    s.SourceInvStoreType = &v
    return s
}
func (s *TaobaoLocationRelationEditLocationRelationDto) SetSourceStoreCode(v string) *TaobaoLocationRelationEditLocationRelationDto {
    s.SourceStoreCode = &v
    return s
}
