package domain


type TaobaoWlbWmsConsignBillGetPackagematerial struct {
    /*
        包材的数量     */
    MaterialQuantity  *string `json:"material_quantity,omitempty" `

    /*
        淘宝指定的包材型号     */
    MaterialType  *string `json:"material_type,omitempty" `

}

func (s *TaobaoWlbWmsConsignBillGetPackagematerial) SetMaterialQuantity(v string) *TaobaoWlbWmsConsignBillGetPackagematerial {
    s.MaterialQuantity = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetPackagematerial) SetMaterialType(v string) *TaobaoWlbWmsConsignBillGetPackagematerial {
    s.MaterialType = &v
    return s
}
