package domain


type TaobaoItempropvaluesGetPropValue struct {
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
    Features  *[]TaobaoItempropvaluesGetFeature `json:"features,omitempty" `

}

func (s *TaobaoItempropvaluesGetPropValue) SetCid(v int64) *TaobaoItempropvaluesGetPropValue {
    s.Cid = &v
    return s
}
func (s *TaobaoItempropvaluesGetPropValue) SetPid(v int64) *TaobaoItempropvaluesGetPropValue {
    s.Pid = &v
    return s
}
func (s *TaobaoItempropvaluesGetPropValue) SetVid(v int64) *TaobaoItempropvaluesGetPropValue {
    s.Vid = &v
    return s
}
func (s *TaobaoItempropvaluesGetPropValue) SetName(v string) *TaobaoItempropvaluesGetPropValue {
    s.Name = &v
    return s
}
func (s *TaobaoItempropvaluesGetPropValue) SetPropName(v string) *TaobaoItempropvaluesGetPropValue {
    s.PropName = &v
    return s
}
func (s *TaobaoItempropvaluesGetPropValue) SetStatus(v string) *TaobaoItempropvaluesGetPropValue {
    s.Status = &v
    return s
}
func (s *TaobaoItempropvaluesGetPropValue) SetNameAlias(v string) *TaobaoItempropvaluesGetPropValue {
    s.NameAlias = &v
    return s
}
func (s *TaobaoItempropvaluesGetPropValue) SetIsParent(v bool) *TaobaoItempropvaluesGetPropValue {
    s.IsParent = &v
    return s
}
func (s *TaobaoItempropvaluesGetPropValue) SetSortOrder(v int64) *TaobaoItempropvaluesGetPropValue {
    s.SortOrder = &v
    return s
}
func (s *TaobaoItempropvaluesGetPropValue) SetFeatures(v []TaobaoItempropvaluesGetFeature) *TaobaoItempropvaluesGetPropValue {
    s.Features = &v
    return s
}
