package domain


type TaobaoWlbWmsConsignBillGetPackagemateriallist struct {
    /*
        包裹包材信息     */
    PackageMaterial  *TaobaoWlbWmsConsignBillGetPackagematerial `json:"package_material,omitempty" `

}

func (s *TaobaoWlbWmsConsignBillGetPackagemateriallist) SetPackageMaterial(v TaobaoWlbWmsConsignBillGetPackagematerial) *TaobaoWlbWmsConsignBillGetPackagemateriallist {
    s.PackageMaterial = &v
    return s
}
