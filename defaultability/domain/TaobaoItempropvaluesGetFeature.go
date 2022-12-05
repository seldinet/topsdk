package domain


type TaobaoItempropvaluesGetFeature struct {
    /*
        属性键     */
    AttrKey  *string `json:"attr_key,omitempty" `

    /*
        属性值     */
    AttrValue  *string `json:"attr_value,omitempty" `

}

func (s *TaobaoItempropvaluesGetFeature) SetAttrKey(v string) *TaobaoItempropvaluesGetFeature {
    s.AttrKey = &v
    return s
}
func (s *TaobaoItempropvaluesGetFeature) SetAttrValue(v string) *TaobaoItempropvaluesGetFeature {
    s.AttrValue = &v
    return s
}
