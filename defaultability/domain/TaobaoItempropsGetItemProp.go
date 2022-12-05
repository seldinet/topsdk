package domain


type TaobaoItempropsGetItemProp struct {
    /*
        属性 ID 例：品牌的PID=20000     */
    Pid  *int64 `json:"pid,omitempty" `

    /*
        上级属性ID     */
    ParentPid  *int64 `json:"parent_pid,omitempty" `

    /*
        上级属性值ID     */
    ParentVid  *int64 `json:"parent_vid,omitempty" `

    /*
        属性名     */
    Name  *string `json:"name,omitempty" `

    /*
        是否关键属性。可选值:true(是),false(否)     */
    IsKeyProp  *bool `json:"is_key_prop,omitempty" `

    /*
        是否销售属性。可选值:true(是),false(否)     */
    IsSaleProp  *bool `json:"is_sale_prop,omitempty" `

    /*
        是否颜色属性。可选值:true(是),false(否)     */
    IsColorProp  *bool `json:"is_color_prop,omitempty" `

    /*
        是否是可枚举属性。可选值:true(是),false(否)     */
    IsEnumProp  *bool `json:"is_enum_prop,omitempty" `

    /*
        在is_enum_prop是true的前提下，是否是卖家可以自行输入的属性（注：如果is_enum_prop返回false，该参数统一返回false）。可选值:true(是),false(否)。<b>对于品牌和型号属性（包括子属性）：如果用户是C卖家，则可自定义属性；如果是B卖家，则不可自定义属性，而必须要授权的属性。</b>     */
    IsInputProp  *bool `json:"is_input_prop,omitempty" `

    /*
        是否商品属性。可选值:true(是),false(否)     */
    IsItemProp  *bool `json:"is_item_prop,omitempty" `

    /*
        发布产品或商品时是否为必选属性。可选值:true(是),false(否)     */
    Must  *bool `json:"must,omitempty" `

    /*
        发布产品或商品时是否可以多选。可选值:true(是),false(否)     */
    Multi  *bool `json:"multi,omitempty" `

    /*
        状态。可选值:normal(正常),deleted(删除)     */
    Status  *string `json:"status,omitempty" `

    /*
        子属性的模板（卖家自行输入属性时需要用到）     */
    ChildTemplate  *string `json:"child_template,omitempty" `

    /*
        排列序号。取值范围:大于零的整排列序号。取值范围:大于零的整数     */
    SortOrder  *int64 `json:"sort_order,omitempty" `

    /*
        是否允许别名。可选值：true（是），false（否）     */
    IsAllowAlias  *bool `json:"is_allow_alias,omitempty" `

    /*
        属性值     */
    PropValues  *[]TaobaoItempropsGetPropValue `json:"prop_values,omitempty" `

    /*
        属性的feature列表     */
    Features  *[]TaobaoItempropsGetFeature `json:"features,omitempty" `

    /*
        是否度量衡属性项     */
    IsTaosir  *bool `json:"is_taosir,omitempty" `

    /*
        度量衡相关信息     */
    TaosirDo  *TaobaoItempropsGetItemTaosirDO `json:"taosir_do,omitempty" `

    /*
        是否是材质 属性项     */
    IsMaterial  *bool `json:"is_material,omitempty" `

    /*
        材质属性信息     */
    MaterialDo  *TaobaoItempropsGetItemMaterialProp `json:"material_do,omitempty" `

}

func (s *TaobaoItempropsGetItemProp) SetPid(v int64) *TaobaoItempropsGetItemProp {
    s.Pid = &v
    return s
}
func (s *TaobaoItempropsGetItemProp) SetParentPid(v int64) *TaobaoItempropsGetItemProp {
    s.ParentPid = &v
    return s
}
func (s *TaobaoItempropsGetItemProp) SetParentVid(v int64) *TaobaoItempropsGetItemProp {
    s.ParentVid = &v
    return s
}
func (s *TaobaoItempropsGetItemProp) SetName(v string) *TaobaoItempropsGetItemProp {
    s.Name = &v
    return s
}
func (s *TaobaoItempropsGetItemProp) SetIsKeyProp(v bool) *TaobaoItempropsGetItemProp {
    s.IsKeyProp = &v
    return s
}
func (s *TaobaoItempropsGetItemProp) SetIsSaleProp(v bool) *TaobaoItempropsGetItemProp {
    s.IsSaleProp = &v
    return s
}
func (s *TaobaoItempropsGetItemProp) SetIsColorProp(v bool) *TaobaoItempropsGetItemProp {
    s.IsColorProp = &v
    return s
}
func (s *TaobaoItempropsGetItemProp) SetIsEnumProp(v bool) *TaobaoItempropsGetItemProp {
    s.IsEnumProp = &v
    return s
}
func (s *TaobaoItempropsGetItemProp) SetIsInputProp(v bool) *TaobaoItempropsGetItemProp {
    s.IsInputProp = &v
    return s
}
func (s *TaobaoItempropsGetItemProp) SetIsItemProp(v bool) *TaobaoItempropsGetItemProp {
    s.IsItemProp = &v
    return s
}
func (s *TaobaoItempropsGetItemProp) SetMust(v bool) *TaobaoItempropsGetItemProp {
    s.Must = &v
    return s
}
func (s *TaobaoItempropsGetItemProp) SetMulti(v bool) *TaobaoItempropsGetItemProp {
    s.Multi = &v
    return s
}
func (s *TaobaoItempropsGetItemProp) SetStatus(v string) *TaobaoItempropsGetItemProp {
    s.Status = &v
    return s
}
func (s *TaobaoItempropsGetItemProp) SetChildTemplate(v string) *TaobaoItempropsGetItemProp {
    s.ChildTemplate = &v
    return s
}
func (s *TaobaoItempropsGetItemProp) SetSortOrder(v int64) *TaobaoItempropsGetItemProp {
    s.SortOrder = &v
    return s
}
func (s *TaobaoItempropsGetItemProp) SetIsAllowAlias(v bool) *TaobaoItempropsGetItemProp {
    s.IsAllowAlias = &v
    return s
}
func (s *TaobaoItempropsGetItemProp) SetPropValues(v []TaobaoItempropsGetPropValue) *TaobaoItempropsGetItemProp {
    s.PropValues = &v
    return s
}
func (s *TaobaoItempropsGetItemProp) SetFeatures(v []TaobaoItempropsGetFeature) *TaobaoItempropsGetItemProp {
    s.Features = &v
    return s
}
func (s *TaobaoItempropsGetItemProp) SetIsTaosir(v bool) *TaobaoItempropsGetItemProp {
    s.IsTaosir = &v
    return s
}
func (s *TaobaoItempropsGetItemProp) SetTaosirDo(v TaobaoItempropsGetItemTaosirDO) *TaobaoItempropsGetItemProp {
    s.TaosirDo = &v
    return s
}
func (s *TaobaoItempropsGetItemProp) SetIsMaterial(v bool) *TaobaoItempropsGetItemProp {
    s.IsMaterial = &v
    return s
}
func (s *TaobaoItempropsGetItemProp) SetMaterialDo(v TaobaoItempropsGetItemMaterialProp) *TaobaoItempropsGetItemProp {
    s.MaterialDo = &v
    return s
}
