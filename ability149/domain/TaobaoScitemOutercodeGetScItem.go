package domain


type TaobaoScitemOutercodeGetScItem struct {
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
        1表示区域销售，0或是空是非区域销售     */
    IsAreaSale  *int64 `json:"is_area_sale,omitempty" `

    /*
        1.普通供应链商品 2.供应链组合主商品     */
    ItemType  *int64 `json:"item_type,omitempty" `

    /*
        后端商品options字段     */
    Options  *int64 `json:"options,omitempty" `

}

func (s *TaobaoScitemOutercodeGetScItem) SetWeight(v int64) *TaobaoScitemOutercodeGetScItem {
    s.Weight = &v
    return s
}
func (s *TaobaoScitemOutercodeGetScItem) SetBarCode(v string) *TaobaoScitemOutercodeGetScItem {
    s.BarCode = &v
    return s
}
func (s *TaobaoScitemOutercodeGetScItem) SetRemark(v string) *TaobaoScitemOutercodeGetScItem {
    s.Remark = &v
    return s
}
func (s *TaobaoScitemOutercodeGetScItem) SetWidth(v int64) *TaobaoScitemOutercodeGetScItem {
    s.Width = &v
    return s
}
func (s *TaobaoScitemOutercodeGetScItem) SetBrandId(v int64) *TaobaoScitemOutercodeGetScItem {
    s.BrandId = &v
    return s
}
func (s *TaobaoScitemOutercodeGetScItem) SetProperties(v string) *TaobaoScitemOutercodeGetScItem {
    s.Properties = &v
    return s
}
func (s *TaobaoScitemOutercodeGetScItem) SetBrandName(v string) *TaobaoScitemOutercodeGetScItem {
    s.BrandName = &v
    return s
}
func (s *TaobaoScitemOutercodeGetScItem) SetWmsCode(v string) *TaobaoScitemOutercodeGetScItem {
    s.WmsCode = &v
    return s
}
func (s *TaobaoScitemOutercodeGetScItem) SetIsFriable(v bool) *TaobaoScitemOutercodeGetScItem {
    s.IsFriable = &v
    return s
}
func (s *TaobaoScitemOutercodeGetScItem) SetIsWarranty(v bool) *TaobaoScitemOutercodeGetScItem {
    s.IsWarranty = &v
    return s
}
func (s *TaobaoScitemOutercodeGetScItem) SetPrice(v int64) *TaobaoScitemOutercodeGetScItem {
    s.Price = &v
    return s
}
func (s *TaobaoScitemOutercodeGetScItem) SetHeight(v int64) *TaobaoScitemOutercodeGetScItem {
    s.Height = &v
    return s
}
func (s *TaobaoScitemOutercodeGetScItem) SetIsDangerous(v bool) *TaobaoScitemOutercodeGetScItem {
    s.IsDangerous = &v
    return s
}
func (s *TaobaoScitemOutercodeGetScItem) SetItemName(v string) *TaobaoScitemOutercodeGetScItem {
    s.ItemName = &v
    return s
}
func (s *TaobaoScitemOutercodeGetScItem) SetItemId(v int64) *TaobaoScitemOutercodeGetScItem {
    s.ItemId = &v
    return s
}
func (s *TaobaoScitemOutercodeGetScItem) SetIsCostly(v bool) *TaobaoScitemOutercodeGetScItem {
    s.IsCostly = &v
    return s
}
func (s *TaobaoScitemOutercodeGetScItem) SetOuterCode(v string) *TaobaoScitemOutercodeGetScItem {
    s.OuterCode = &v
    return s
}
func (s *TaobaoScitemOutercodeGetScItem) SetVolume(v int64) *TaobaoScitemOutercodeGetScItem {
    s.Volume = &v
    return s
}
func (s *TaobaoScitemOutercodeGetScItem) SetLength(v int64) *TaobaoScitemOutercodeGetScItem {
    s.Length = &v
    return s
}
func (s *TaobaoScitemOutercodeGetScItem) SetMatterStatus(v string) *TaobaoScitemOutercodeGetScItem {
    s.MatterStatus = &v
    return s
}
func (s *TaobaoScitemOutercodeGetScItem) SetIsAreaSale(v int64) *TaobaoScitemOutercodeGetScItem {
    s.IsAreaSale = &v
    return s
}
func (s *TaobaoScitemOutercodeGetScItem) SetItemType(v int64) *TaobaoScitemOutercodeGetScItem {
    s.ItemType = &v
    return s
}
func (s *TaobaoScitemOutercodeGetScItem) SetOptions(v int64) *TaobaoScitemOutercodeGetScItem {
    s.Options = &v
    return s
}
