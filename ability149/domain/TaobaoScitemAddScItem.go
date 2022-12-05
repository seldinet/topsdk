package domain


type TaobaoScitemAddScItem struct {
    /*
        重量.单位：克     */
    Weight  *int64 `json:"weight,omitempty" `

    /*
        条形码     */
    BarCode  *string `json:"bar_code,omitempty" `

    /*
        备注     */
    Remark  *string `json:"remark,omitempty" `

    /*
        宽 单位：mm     */
    Width  *int64 `json:"width,omitempty" `

    /*
        品牌id     */
    BrandId  *int64 `json:"brand_id,omitempty" `

    /*
        商品属性格式是  p1:v1,p2:v2,p3:v3     */
    Properties  *string `json:"properties,omitempty" `

    /*
        品牌名称     */
    BrandName  *string `json:"brand_name,omitempty" `

    /*
        仓储商编码，可以支持多个，格式wmsname:code     */
    WmsCode  *string `json:"wms_code,omitempty" `

    /*
        是否易碎 false ：不是  true：是     */
    IsFriable  *bool `json:"is_friable,omitempty" `

    /*
        是否保质期：false:不是 true：是     */
    IsWarranty  *bool `json:"is_warranty,omitempty" `

    /*
        价格：分（吊牌价）     */
    Price  *int64 `json:"price,omitempty" `

    /*
        高 单位：mm     */
    Height  *int64 `json:"height,omitempty" `

    /*
        是否危险 false：不是  true：是     */
    IsDangerous  *bool `json:"is_dangerous,omitempty" `

    /*
        商品名称     */
    ItemName  *string `json:"item_name,omitempty" `

    /*
        商品id     */
    ItemId  *int64 `json:"item_id,omitempty" `

    /*
        贵重品:false:不是 true：是     */
    IsCostly  *bool `json:"is_costly,omitempty" `

    /*
        商家编码     */
    OuterCode  *string `json:"outer_code,omitempty" `

    /*
        体积：立方厘米     */
    Volume  *int64 `json:"volume,omitempty" `

    /*
        长度 单位：mm     */
    Length  *int64 `json:"length,omitempty" `

    /*
        LIQUID:液体，1：粉体，SOLID：固体     */
    MatterStatus  *string `json:"matter_status,omitempty" `

    /*
        1.普通供应链商品 2.供应链组合主商品     */
    ItemType  *int64 `json:"item_type,omitempty" `

}

func (s *TaobaoScitemAddScItem) SetWeight(v int64) *TaobaoScitemAddScItem {
    s.Weight = &v
    return s
}
func (s *TaobaoScitemAddScItem) SetBarCode(v string) *TaobaoScitemAddScItem {
    s.BarCode = &v
    return s
}
func (s *TaobaoScitemAddScItem) SetRemark(v string) *TaobaoScitemAddScItem {
    s.Remark = &v
    return s
}
func (s *TaobaoScitemAddScItem) SetWidth(v int64) *TaobaoScitemAddScItem {
    s.Width = &v
    return s
}
func (s *TaobaoScitemAddScItem) SetBrandId(v int64) *TaobaoScitemAddScItem {
    s.BrandId = &v
    return s
}
func (s *TaobaoScitemAddScItem) SetProperties(v string) *TaobaoScitemAddScItem {
    s.Properties = &v
    return s
}
func (s *TaobaoScitemAddScItem) SetBrandName(v string) *TaobaoScitemAddScItem {
    s.BrandName = &v
    return s
}
func (s *TaobaoScitemAddScItem) SetWmsCode(v string) *TaobaoScitemAddScItem {
    s.WmsCode = &v
    return s
}
func (s *TaobaoScitemAddScItem) SetIsFriable(v bool) *TaobaoScitemAddScItem {
    s.IsFriable = &v
    return s
}
func (s *TaobaoScitemAddScItem) SetIsWarranty(v bool) *TaobaoScitemAddScItem {
    s.IsWarranty = &v
    return s
}
func (s *TaobaoScitemAddScItem) SetPrice(v int64) *TaobaoScitemAddScItem {
    s.Price = &v
    return s
}
func (s *TaobaoScitemAddScItem) SetHeight(v int64) *TaobaoScitemAddScItem {
    s.Height = &v
    return s
}
func (s *TaobaoScitemAddScItem) SetIsDangerous(v bool) *TaobaoScitemAddScItem {
    s.IsDangerous = &v
    return s
}
func (s *TaobaoScitemAddScItem) SetItemName(v string) *TaobaoScitemAddScItem {
    s.ItemName = &v
    return s
}
func (s *TaobaoScitemAddScItem) SetItemId(v int64) *TaobaoScitemAddScItem {
    s.ItemId = &v
    return s
}
func (s *TaobaoScitemAddScItem) SetIsCostly(v bool) *TaobaoScitemAddScItem {
    s.IsCostly = &v
    return s
}
func (s *TaobaoScitemAddScItem) SetOuterCode(v string) *TaobaoScitemAddScItem {
    s.OuterCode = &v
    return s
}
func (s *TaobaoScitemAddScItem) SetVolume(v int64) *TaobaoScitemAddScItem {
    s.Volume = &v
    return s
}
func (s *TaobaoScitemAddScItem) SetLength(v int64) *TaobaoScitemAddScItem {
    s.Length = &v
    return s
}
func (s *TaobaoScitemAddScItem) SetMatterStatus(v string) *TaobaoScitemAddScItem {
    s.MatterStatus = &v
    return s
}
func (s *TaobaoScitemAddScItem) SetItemType(v int64) *TaobaoScitemAddScItem {
    s.ItemType = &v
    return s
}
