package request


type TaobaoScitemAddRequest struct {
    /*
        商品名称     */
    ItemName  *string `json:"item_name" required:"true" `
    /*
        商家编码     */
    OuterCode  *string `json:"outer_code" required:"true" `
    /*
        0.普通供应链商品 1.供应链组合主商品 defalutValue��0    */
    ItemType  *int64 `json:"item_type,omitempty" required:"false" `
    /*
        商品属性格式是  p1:v1,p2:v2,p3:v3     */
    Properties  *string `json:"properties,omitempty" required:"false" `
    /*
        条形码     */
    BarCode  *string `json:"bar_code,omitempty" required:"false" `
    /*
        仓储商编码     */
    WmsCode  *string `json:"wms_code,omitempty" required:"false" `
    /*
        是否易碎 0：不是  1：是     */
    IsFriable  *int64 `json:"is_friable,omitempty" required:"false" `
    /*
        是否危险 0：不是  1：是     */
    IsDangerous  *int64 `json:"is_dangerous,omitempty" required:"false" `
    /*
        是否是贵重品 0:不是 1：是     */
    IsCostly  *int64 `json:"is_costly,omitempty" required:"false" `
    /*
        是否保质期：0:不是 1：是     */
    IsWarranty  *int64 `json:"is_warranty,omitempty" required:"false" `
    /*
        重量 单位：g     */
    Weight  *int64 `json:"weight,omitempty" required:"false" `
    /*
        长度 单位：mm     */
    Length  *int64 `json:"length,omitempty" required:"false" `
    /*
        宽 单位：mm     */
    Width  *int64 `json:"width,omitempty" required:"false" `
    /*
        高 单位：mm     */
    Height  *int64 `json:"height,omitempty" required:"false" `
    /*
        体积：立方厘米     */
    Volume  *int64 `json:"volume,omitempty" required:"false" `
    /*
        价格 单位：分     */
    Price  *int64 `json:"price,omitempty" required:"false" `
    /*
        remark     */
    Remark  *string `json:"remark,omitempty" required:"false" `
    /*
        0:液体，1：粉体，2：固体     */
    MatterStatus  *int64 `json:"matter_status,omitempty" required:"false" `
    /*
        品牌id     */
    BrandId  *int64 `json:"brand_id,omitempty" required:"false" `
    /*
        brand_Name     */
    BrandName  *string `json:"brand_name,omitempty" required:"false" `
    /*
        spuId或是cspuid     */
    SpuId  *int64 `json:"spu_id,omitempty" required:"false" `
    /*
        1表示区域销售，0或是空是非区域销售 defalutValue��0    */
    IsAreaSale  *int64 `json:"is_area_sale,omitempty" required:"false" `
}

func (s *TaobaoScitemAddRequest) SetItemName(v string) *TaobaoScitemAddRequest {
    s.ItemName = &v
    return s
}
func (s *TaobaoScitemAddRequest) SetOuterCode(v string) *TaobaoScitemAddRequest {
    s.OuterCode = &v
    return s
}
func (s *TaobaoScitemAddRequest) SetItemType(v int64) *TaobaoScitemAddRequest {
    s.ItemType = &v
    return s
}
func (s *TaobaoScitemAddRequest) SetProperties(v string) *TaobaoScitemAddRequest {
    s.Properties = &v
    return s
}
func (s *TaobaoScitemAddRequest) SetBarCode(v string) *TaobaoScitemAddRequest {
    s.BarCode = &v
    return s
}
func (s *TaobaoScitemAddRequest) SetWmsCode(v string) *TaobaoScitemAddRequest {
    s.WmsCode = &v
    return s
}
func (s *TaobaoScitemAddRequest) SetIsFriable(v int64) *TaobaoScitemAddRequest {
    s.IsFriable = &v
    return s
}
func (s *TaobaoScitemAddRequest) SetIsDangerous(v int64) *TaobaoScitemAddRequest {
    s.IsDangerous = &v
    return s
}
func (s *TaobaoScitemAddRequest) SetIsCostly(v int64) *TaobaoScitemAddRequest {
    s.IsCostly = &v
    return s
}
func (s *TaobaoScitemAddRequest) SetIsWarranty(v int64) *TaobaoScitemAddRequest {
    s.IsWarranty = &v
    return s
}
func (s *TaobaoScitemAddRequest) SetWeight(v int64) *TaobaoScitemAddRequest {
    s.Weight = &v
    return s
}
func (s *TaobaoScitemAddRequest) SetLength(v int64) *TaobaoScitemAddRequest {
    s.Length = &v
    return s
}
func (s *TaobaoScitemAddRequest) SetWidth(v int64) *TaobaoScitemAddRequest {
    s.Width = &v
    return s
}
func (s *TaobaoScitemAddRequest) SetHeight(v int64) *TaobaoScitemAddRequest {
    s.Height = &v
    return s
}
func (s *TaobaoScitemAddRequest) SetVolume(v int64) *TaobaoScitemAddRequest {
    s.Volume = &v
    return s
}
func (s *TaobaoScitemAddRequest) SetPrice(v int64) *TaobaoScitemAddRequest {
    s.Price = &v
    return s
}
func (s *TaobaoScitemAddRequest) SetRemark(v string) *TaobaoScitemAddRequest {
    s.Remark = &v
    return s
}
func (s *TaobaoScitemAddRequest) SetMatterStatus(v int64) *TaobaoScitemAddRequest {
    s.MatterStatus = &v
    return s
}
func (s *TaobaoScitemAddRequest) SetBrandId(v int64) *TaobaoScitemAddRequest {
    s.BrandId = &v
    return s
}
func (s *TaobaoScitemAddRequest) SetBrandName(v string) *TaobaoScitemAddRequest {
    s.BrandName = &v
    return s
}
func (s *TaobaoScitemAddRequest) SetSpuId(v int64) *TaobaoScitemAddRequest {
    s.SpuId = &v
    return s
}
func (s *TaobaoScitemAddRequest) SetIsAreaSale(v int64) *TaobaoScitemAddRequest {
    s.IsAreaSale = &v
    return s
}

func (req *TaobaoScitemAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemName != nil) {
        paramMap["item_name"] = *req.ItemName
    }
    if(req.OuterCode != nil) {
        paramMap["outer_code"] = *req.OuterCode
    }
    if(req.ItemType != nil) {
        paramMap["item_type"] = *req.ItemType
    }
    if(req.Properties != nil) {
        paramMap["properties"] = *req.Properties
    }
    if(req.BarCode != nil) {
        paramMap["bar_code"] = *req.BarCode
    }
    if(req.WmsCode != nil) {
        paramMap["wms_code"] = *req.WmsCode
    }
    if(req.IsFriable != nil) {
        paramMap["is_friable"] = *req.IsFriable
    }
    if(req.IsDangerous != nil) {
        paramMap["is_dangerous"] = *req.IsDangerous
    }
    if(req.IsCostly != nil) {
        paramMap["is_costly"] = *req.IsCostly
    }
    if(req.IsWarranty != nil) {
        paramMap["is_warranty"] = *req.IsWarranty
    }
    if(req.Weight != nil) {
        paramMap["weight"] = *req.Weight
    }
    if(req.Length != nil) {
        paramMap["length"] = *req.Length
    }
    if(req.Width != nil) {
        paramMap["width"] = *req.Width
    }
    if(req.Height != nil) {
        paramMap["height"] = *req.Height
    }
    if(req.Volume != nil) {
        paramMap["volume"] = *req.Volume
    }
    if(req.Price != nil) {
        paramMap["price"] = *req.Price
    }
    if(req.Remark != nil) {
        paramMap["remark"] = *req.Remark
    }
    if(req.MatterStatus != nil) {
        paramMap["matter_status"] = *req.MatterStatus
    }
    if(req.BrandId != nil) {
        paramMap["brand_id"] = *req.BrandId
    }
    if(req.BrandName != nil) {
        paramMap["brand_name"] = *req.BrandName
    }
    if(req.SpuId != nil) {
        paramMap["spu_id"] = *req.SpuId
    }
    if(req.IsAreaSale != nil) {
        paramMap["is_area_sale"] = *req.IsAreaSale
    }
    return paramMap
}

func (req *TaobaoScitemAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}