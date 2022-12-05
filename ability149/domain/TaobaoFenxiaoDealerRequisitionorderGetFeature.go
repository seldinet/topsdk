package domain


type TaobaoFenxiaoDealerRequisitionorderGetFeature struct {
    /*
        属性键     */
    AttrKey  *string `json:"attr_key,omitempty" `

    /*
        属性值     */
    AttrValue  *string `json:"attr_value,omitempty" `

}

func (s *TaobaoFenxiaoDealerRequisitionorderGetFeature) SetAttrKey(v string) *TaobaoFenxiaoDealerRequisitionorderGetFeature {
    s.AttrKey = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetFeature) SetAttrValue(v string) *TaobaoFenxiaoDealerRequisitionorderGetFeature {
    s.AttrValue = &v
    return s
}
