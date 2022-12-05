package domain


type TaobaoFenxiaoDealerRequisitionorderQueryFeature struct {
    /*
        属性键     */
    AttrKey  *string `json:"attr_key,omitempty" `

    /*
        属性值     */
    AttrValue  *string `json:"attr_value,omitempty" `

}

func (s *TaobaoFenxiaoDealerRequisitionorderQueryFeature) SetAttrKey(v string) *TaobaoFenxiaoDealerRequisitionorderQueryFeature {
    s.AttrKey = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryFeature) SetAttrValue(v string) *TaobaoFenxiaoDealerRequisitionorderQueryFeature {
    s.AttrValue = &v
    return s
}
