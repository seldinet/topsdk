package domain


type TaobaoItempropsGetFeature struct {
    /*
        属性值     */
    AttrValue  *string `json:"attr_value,omitempty" `

    /*
        属性键     */
    AttrKey  *string `json:"attr_key,omitempty" `

}

func (s *TaobaoItempropsGetFeature) SetAttrValue(v string) *TaobaoItempropsGetFeature {
    s.AttrValue = &v
    return s
}
func (s *TaobaoItempropsGetFeature) SetAttrKey(v string) *TaobaoItempropsGetFeature {
    s.AttrKey = &v
    return s
}
