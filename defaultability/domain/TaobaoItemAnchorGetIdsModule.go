package domain


type TaobaoItemAnchorGetIdsModule struct {
    /*
        宝贝描述规范化模块名     */
    Name  *string `json:"name,omitempty" `

    /*
        宝贝描述规范化模块id     */
    Id  *int64 `json:"id,omitempty" `

    /*
        0为自动打标；
1为人工打标；     */
    Type  *int64 `json:"type,omitempty" `

}

func (s *TaobaoItemAnchorGetIdsModule) SetName(v string) *TaobaoItemAnchorGetIdsModule {
    s.Name = &v
    return s
}
func (s *TaobaoItemAnchorGetIdsModule) SetId(v int64) *TaobaoItemAnchorGetIdsModule {
    s.Id = &v
    return s
}
func (s *TaobaoItemAnchorGetIdsModule) SetType(v int64) *TaobaoItemAnchorGetIdsModule {
    s.Type = &v
    return s
}
