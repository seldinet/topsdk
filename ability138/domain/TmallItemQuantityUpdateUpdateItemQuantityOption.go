package domain


type TmallItemQuantityUpdateUpdateItemQuantityOption struct {
    /*
        增量更新时幂等外部编码，只能包含十六进制字符(0-9,a-f,A-F)     */
    OuterBizKey  *string `json:"outer_biz_key,omitempty" `

    /*
        库存更新方式: 1是全量更新 2是增量更新；默认是1。     */
    Type  *int64 `json:"type,omitempty" `

}

func (s *TmallItemQuantityUpdateUpdateItemQuantityOption) SetOuterBizKey(v string) *TmallItemQuantityUpdateUpdateItemQuantityOption {
    s.OuterBizKey = &v
    return s
}
func (s *TmallItemQuantityUpdateUpdateItemQuantityOption) SetType(v int64) *TmallItemQuantityUpdateUpdateItemQuantityOption {
    s.Type = &v
    return s
}
