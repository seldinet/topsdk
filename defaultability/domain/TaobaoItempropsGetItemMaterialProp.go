package domain


type TaobaoItempropsGetItemMaterialProp struct {
    /*
        材质值列表     */
    Materials  *[]TaobaoItempropsGetItemMateriaValueDO `json:"materials,omitempty" `

}

func (s *TaobaoItempropsGetItemMaterialProp) SetMaterials(v []TaobaoItempropsGetItemMateriaValueDO) *TaobaoItempropsGetItemMaterialProp {
    s.Materials = &v
    return s
}
