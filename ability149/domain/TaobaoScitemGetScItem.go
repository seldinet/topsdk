package domain


type TaobaoScitemGetScItem struct {
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

    /*
        1表示区域销售，0或是空是非区域销售     */
    IsAreaSale  *int64 `json:"is_area_sale,omitempty" `

    /*
        后端商品options字段     */
    Options  *int64 `json:"options,omitempty" `

}

func (s *TaobaoScitemGetScItem) SetWeight(v int64) *TaobaoScitemGetScItem {
    s.Weight = &v
    return s
}
func (s *TaobaoScitemGetScItem) SetBarCode(v string) *TaobaoScitemGetScItem {
    s.BarCode = &v
    return s
}
func (s *TaobaoScitemGetScItem) SetRemark(v string) *TaobaoScitemGetScItem {
    s.Remark = &v
    return s
}
func (s *TaobaoScitemGetScItem) SetWidth(v int64) *TaobaoScitemGetScItem {
    s.Width = &v
    return s
}
func (s *TaobaoScitemGetScItem) SetBrandId(v int64) *TaobaoScitemGetScItem {
    s.BrandId = &v
    return s
}
func (s *TaobaoScitemGetScItem) SetProperties(v string) *TaobaoScitemGetScItem {
    s.Properties = &v
    return s
}
func (s *TaobaoScitemGetScItem) SetBrandName(v string) *TaobaoScitemGetScItem {
    s.BrandName = &v
    return s
}
func (s *TaobaoScitemGetScItem) SetWmsCode(v string) *TaobaoScitemGetScItem {
    s.WmsCode = &v
    return s
}
func (s *TaobaoScitemGetScItem) SetIsFriable(v bool) *TaobaoScitemGetScItem {
    s.IsFriable = &v
    return s
}
func (s *TaobaoScitemGetScItem) SetIsWarranty(v bool) *TaobaoScitemGetScItem {
    s.IsWarranty = &v
    return s
}
func (s *TaobaoScitemGetScItem) SetPrice(v int64) *TaobaoScitemGetScItem {
    s.Price = &v
    return s
}
func (s *TaobaoScitemGetScItem) SetHeight(v int64) *TaobaoScitemGetScItem {
    s.Height = &v
    return s
}
func (s *TaobaoScitemGetScItem) SetIsDangerous(v bool) *TaobaoScitemGetScItem {
    s.IsDangerous = &v
    return s
}
func (s *TaobaoScitemGetScItem) SetItemName(v string) *TaobaoScitemGetScItem {
    s.ItemName = &v
    return s
}
func (s *TaobaoScitemGetScItem) SetItemId(v int64) *TaobaoScitemGetScItem {
    s.ItemId = &v
    return s
}
func (s *TaobaoScitemGetScItem) SetIsCostly(v bool) *TaobaoScitemGetScItem {
    s.IsCostly = &v
    return s
}
func (s *TaobaoScitemGetScItem) SetOuterCode(v string) *TaobaoScitemGetScItem {
    s.OuterCode = &v
    return s
}
func (s *TaobaoScitemGetScItem) SetVolume(v int64) *TaobaoScitemGetScItem {
    s.Volume = &v
    return s
}
func (s *TaobaoScitemGetScItem) SetLength(v int64) *TaobaoScitemGetScItem {
    s.Length = &v
    return s
}
func (s *TaobaoScitemGetScItem) SetMatterStatus(v string) *TaobaoScitemGetScItem {
    s.MatterStatus = &v
    return s
}
func (s *TaobaoScitemGetScItem) SetItemType(v int64) *TaobaoScitemGetScItem {
    s.ItemType = &v
    return s
}
func (s *TaobaoScitemGetScItem) SetIsAreaSale(v int64) *TaobaoScitemGetScItem {
    s.IsAreaSale = &v
    return s
}
func (s *TaobaoScitemGetScItem) SetOptions(v int64) *TaobaoScitemGetScItem {
    s.Options = &v
    return s
}
