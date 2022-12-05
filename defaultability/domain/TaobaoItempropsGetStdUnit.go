package domain


type TaobaoItempropsGetStdUnit struct {
    /*
        属性键     */
    AttrKey  *string `json:"attr_key,omitempty" `

    /*
        属性值     */
    AttrValue  *string `json:"attr_value,omitempty" `

}

func (s *TaobaoItempropsGetStdUnit) SetAttrKey(v string) *TaobaoItempropsGetStdUnit {
    s.AttrKey = &v
    return s
}
func (s *TaobaoItempropsGetStdUnit) SetAttrValue(v string) *TaobaoItempropsGetStdUnit {
    s.AttrValue = &v
    return s
}
