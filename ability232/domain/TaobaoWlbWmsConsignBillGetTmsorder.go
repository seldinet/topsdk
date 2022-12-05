package domain


type TaobaoWlbWmsConsignBillGetTmsorder struct {
    /*
        包裹里面的商品信息列表     */
    TmsItemList  *[]TaobaoWlbWmsConsignBillGetTmsitemlist `json:"tms_item_list,omitempty" `

    /*
        包材信息     */
    PackageMaterialList  *[]TaobaoWlbWmsConsignBillGetPackagemateriallist `json:"package_material_list,omitempty" `

    /*
        包裹高度，单位：毫米     */
    PackageHeight  *int64 `json:"package_height,omitempty" `

    /*
        包裹宽度，单位：毫米     */
    PackageWidth  *int64 `json:"package_width,omitempty" `

    /*
        包裹长度，单位：毫米     */
    PackageLength  *int64 `json:"package_length,omitempty" `

    /*
        包裹重量，单位：克     */
    PackageWeight  *int64 `json:"package_weight,omitempty" `

    /*
        包裹号     */
    PackageCode  *string `json:"package_code,omitempty" `

    /*
        运单编码     */
    TmsOrderCode  *string `json:"tms_order_code,omitempty" `

    /*
        快递公司服务编码     */
    TmsCode  *string `json:"tms_code,omitempty" `

}

func (s *TaobaoWlbWmsConsignBillGetTmsorder) SetTmsItemList(v []TaobaoWlbWmsConsignBillGetTmsitemlist) *TaobaoWlbWmsConsignBillGetTmsorder {
    s.TmsItemList = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetTmsorder) SetPackageMaterialList(v []TaobaoWlbWmsConsignBillGetPackagemateriallist) *TaobaoWlbWmsConsignBillGetTmsorder {
    s.PackageMaterialList = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetTmsorder) SetPackageHeight(v int64) *TaobaoWlbWmsConsignBillGetTmsorder {
    s.PackageHeight = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetTmsorder) SetPackageWidth(v int64) *TaobaoWlbWmsConsignBillGetTmsorder {
    s.PackageWidth = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetTmsorder) SetPackageLength(v int64) *TaobaoWlbWmsConsignBillGetTmsorder {
    s.PackageLength = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetTmsorder) SetPackageWeight(v int64) *TaobaoWlbWmsConsignBillGetTmsorder {
    s.PackageWeight = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetTmsorder) SetPackageCode(v string) *TaobaoWlbWmsConsignBillGetTmsorder {
    s.PackageCode = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetTmsorder) SetTmsOrderCode(v string) *TaobaoWlbWmsConsignBillGetTmsorder {
    s.TmsOrderCode = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetTmsorder) SetTmsCode(v string) *TaobaoWlbWmsConsignBillGetTmsorder {
    s.TmsCode = &v
    return s
}
