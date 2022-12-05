package domain


type TaobaoItempropsGetPropValue struct {
    /*
        类目ID     */
    Cid  *int64 `json:"cid,omitempty" `

    /*
        属性 ID     */
    Pid  *int64 `json:"pid,omitempty" `

    /*
        属性值ID     */
    Vid  *int64 `json:"vid,omitempty" `

    /*
        属性值     */
    Name  *string `json:"name,omitempty" `

    /*
        属性名     */
    PropName  *string `json:"prop_name,omitempty" `

    /*
        状态。可选值:normal(正常),deleted(删除)     */
    Status  *string `json:"status,omitempty" `

    /*
        属性值别名     */
    NameAlias  *string `json:"name_alias,omitempty" `

    /*
        是否为父类目属性     */
    IsParent  *bool `json:"is_parent,omitempty" `

    /*
        排列序号。取值范围:大于零的整数     */
    SortOrder  *int64 `json:"sort_order,omitempty" `

    /*
        属性值feature     */
    Features  *[]TaobaoItempropsGetFeature `json:"features,omitempty" `

}

func (s *TaobaoItempropsGetPropValue) SetCid(v int64) *TaobaoItempropsGetPropValue {
    s.Cid = &v
    return s
}
func (s *TaobaoItempropsGetPropValue) SetPid(v int64) *TaobaoItempropsGetPropValue {
    s.Pid = &v
    return s
}
func (s *TaobaoItempropsGetPropValue) SetVid(v int64) *TaobaoItempropsGetPropValue {
    s.Vid = &v
    return s
}
func (s *TaobaoItempropsGetPropValue) SetName(v string) *TaobaoItempropsGetPropValue {
    s.Name = &v
    return s
}
func (s *TaobaoItempropsGetPropValue) SetPropName(v string) *TaobaoItempropsGetPropValue {
    s.PropName = &v
    return s
}
func (s *TaobaoItempropsGetPropValue) SetStatus(v string) *TaobaoItempropsGetPropValue {
    s.Status = &v
    return s
}
func (s *TaobaoItempropsGetPropValue) SetNameAlias(v string) *TaobaoItempropsGetPropValue {
    s.NameAlias = &v
    return s
}
func (s *TaobaoItempropsGetPropValue) SetIsParent(v bool) *TaobaoItempropsGetPropValue {
    s.IsParent = &v
    return s
}
func (s *TaobaoItempropsGetPropValue) SetSortOrder(v int64) *TaobaoItempropsGetPropValue {
    s.SortOrder = &v
    return s
}
func (s *TaobaoItempropsGetPropValue) SetFeatures(v []TaobaoItempropsGetFeature) *TaobaoItempropsGetPropValue {
    s.Features = &v
    return s
}
