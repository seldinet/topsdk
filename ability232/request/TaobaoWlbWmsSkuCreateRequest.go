package request


type TaobaoWlbWmsSkuCreateRequest struct {
    /*
        是否启用序列号管理     */
    IsSnMgt  *bool `json:"is_sn_mgt,omitempty" required:"false" `
    /*
        商品保质期天数     */
    Lifecycle  *int64 `json:"lifecycle,omitempty" required:"false" `
    /*
        是否启用保质期管理     */
    IsShelflife  *bool `json:"is_shelflife,omitempty" required:"false" `
    /*
        批准文号     */
    ApprovalNumber  *string `json:"approval_number,omitempty" required:"false" `
    /*
        产地     */
    OriginAddress  *int64 `json:"origin_address,omitempty" required:"false" `
    /*
        毛重，单位克     */
    GrossWeight  *int64 `json:"gross_weight,omitempty" required:"false" `
    /*
        尺码     */
    Size  *string `json:"size,omitempty" required:"false" `
    /*
        颜色     */
    Color  *string `json:"color,omitempty" required:"false" `
    /*
        规格     */
    Specification  *string `json:"specification,omitempty" required:"false" `
    /*
        品牌名称     */
    BrandName  *string `json:"brand_name,omitempty" required:"false" `
    /*
        品牌编码     */
    Brand  *string `json:"brand,omitempty" required:"false" `
    /*
        商品类别名称     */
    CategoryName  *string `json:"category_name,omitempty" required:"false" `
    /*
        商品类别编码（外部系统类别）     */
    Category  *string `json:"category,omitempty" required:"false" `
    /*
        商品类别NORMAL：普通商品、COMBINE：组合商品、DISTRIBUTION：分销商品、HAOCAI耗材、FUSHUPIN附属品、BAOCAI 包材、XUNI虚拟商品、QITA其他)     */
    Type  *string `json:"type" required:"true" `
    /*
        商品标题     */
    Title  *string `json:"title,omitempty" required:"false" `
    /*
        商品名称     */
    Name  *string `json:"name" required:"true" `
    /*
        条形码，多条码请用”；”分隔；     */
    BarCode  *string `json:"bar_code,omitempty" required:"false" `
    /*
        商家商品编码     */
    ItemCode  *string `json:"item_code" required:"true" `
    /*
        箱规     */
    Pcs  *int64 `json:"pcs,omitempty" required:"false" `
    /*
        体积，单位立方厘米     */
    Volume  *int64 `json:"volume,omitempty" required:"false" `
    /*
        高度，单位毫米     */
    Height  *int64 `json:"height,omitempty" required:"false" `
    /*
        拓展属性     */
    ExtendFields  *string `json:"extend_fields,omitempty" required:"false" `
    /*
        启用标识     */
    UseYn  *bool `json:"use_yn" required:"true" `
    /*
        是否启用批次管理     */
    IsBatchMgt  *bool `json:"is_batch_mgt,omitempty" required:"false" `
    /*
        成本价，单位分     */
    CostPrice  *int64 `json:"cost_price,omitempty" required:"false" `
    /*
        零售价，单位分     */
    ItemPrice  *int64 `json:"item_price,omitempty" required:"false" `
    /*
        吊牌价，单位分     */
    TagPrice  *int64 `json:"tag_price,omitempty" required:"false" `
    /*
        是否危险品     */
    IsDanger  *bool `json:"is_danger,omitempty" required:"false" `
    /*
        是否易碎品     */
    IsHygroscopic  *bool `json:"is_hygroscopic,omitempty" required:"false" `
    /*
        宽度，单位毫米     */
    Width  *int64 `json:"width,omitempty" required:"false" `
    /*
        长度，单位毫米     */
    Length  *int64 `json:"length,omitempty" required:"false" `
    /*
        净重，单位克     */
    NetWeight  *int64 `json:"net_weight,omitempty" required:"false" `
    /*
        仓库编码     */
    StoreCode  *string `json:"store_code,omitempty" required:"false" `
    /*
        保质期预警天数     */
    AdventLifecycle  *int64 `json:"advent_lifecycle,omitempty" required:"false" `
    /*
        保质期禁收天数     */
    RejectLifecycle  *int64 `json:"reject_lifecycle,omitempty" required:"false" `
    /*
        保质期禁售天数     */
    LockupLifecycle  *int64 `json:"lockup_lifecycle,omitempty" required:"false" `
    /*
        商家商品ID     */
    ItemId  *string `json:"item_id,omitempty" required:"false" `
    /*
        是否区域销售     */
    IsAreaSale  *bool `json:"is_area_sale,omitempty" required:"false" `
}

func (s *TaobaoWlbWmsSkuCreateRequest) SetIsSnMgt(v bool) *TaobaoWlbWmsSkuCreateRequest {
    s.IsSnMgt = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetLifecycle(v int64) *TaobaoWlbWmsSkuCreateRequest {
    s.Lifecycle = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetIsShelflife(v bool) *TaobaoWlbWmsSkuCreateRequest {
    s.IsShelflife = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetApprovalNumber(v string) *TaobaoWlbWmsSkuCreateRequest {
    s.ApprovalNumber = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetOriginAddress(v int64) *TaobaoWlbWmsSkuCreateRequest {
    s.OriginAddress = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetGrossWeight(v int64) *TaobaoWlbWmsSkuCreateRequest {
    s.GrossWeight = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetSize(v string) *TaobaoWlbWmsSkuCreateRequest {
    s.Size = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetColor(v string) *TaobaoWlbWmsSkuCreateRequest {
    s.Color = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetSpecification(v string) *TaobaoWlbWmsSkuCreateRequest {
    s.Specification = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetBrandName(v string) *TaobaoWlbWmsSkuCreateRequest {
    s.BrandName = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetBrand(v string) *TaobaoWlbWmsSkuCreateRequest {
    s.Brand = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetCategoryName(v string) *TaobaoWlbWmsSkuCreateRequest {
    s.CategoryName = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetCategory(v string) *TaobaoWlbWmsSkuCreateRequest {
    s.Category = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetType(v string) *TaobaoWlbWmsSkuCreateRequest {
    s.Type = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetTitle(v string) *TaobaoWlbWmsSkuCreateRequest {
    s.Title = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetName(v string) *TaobaoWlbWmsSkuCreateRequest {
    s.Name = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetBarCode(v string) *TaobaoWlbWmsSkuCreateRequest {
    s.BarCode = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetItemCode(v string) *TaobaoWlbWmsSkuCreateRequest {
    s.ItemCode = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetPcs(v int64) *TaobaoWlbWmsSkuCreateRequest {
    s.Pcs = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetVolume(v int64) *TaobaoWlbWmsSkuCreateRequest {
    s.Volume = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetHeight(v int64) *TaobaoWlbWmsSkuCreateRequest {
    s.Height = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetExtendFields(v string) *TaobaoWlbWmsSkuCreateRequest {
    s.ExtendFields = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetUseYn(v bool) *TaobaoWlbWmsSkuCreateRequest {
    s.UseYn = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetIsBatchMgt(v bool) *TaobaoWlbWmsSkuCreateRequest {
    s.IsBatchMgt = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetCostPrice(v int64) *TaobaoWlbWmsSkuCreateRequest {
    s.CostPrice = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetItemPrice(v int64) *TaobaoWlbWmsSkuCreateRequest {
    s.ItemPrice = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetTagPrice(v int64) *TaobaoWlbWmsSkuCreateRequest {
    s.TagPrice = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetIsDanger(v bool) *TaobaoWlbWmsSkuCreateRequest {
    s.IsDanger = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetIsHygroscopic(v bool) *TaobaoWlbWmsSkuCreateRequest {
    s.IsHygroscopic = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetWidth(v int64) *TaobaoWlbWmsSkuCreateRequest {
    s.Width = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetLength(v int64) *TaobaoWlbWmsSkuCreateRequest {
    s.Length = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetNetWeight(v int64) *TaobaoWlbWmsSkuCreateRequest {
    s.NetWeight = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetStoreCode(v string) *TaobaoWlbWmsSkuCreateRequest {
    s.StoreCode = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetAdventLifecycle(v int64) *TaobaoWlbWmsSkuCreateRequest {
    s.AdventLifecycle = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetRejectLifecycle(v int64) *TaobaoWlbWmsSkuCreateRequest {
    s.RejectLifecycle = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetLockupLifecycle(v int64) *TaobaoWlbWmsSkuCreateRequest {
    s.LockupLifecycle = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetItemId(v string) *TaobaoWlbWmsSkuCreateRequest {
    s.ItemId = &v
    return s
}
func (s *TaobaoWlbWmsSkuCreateRequest) SetIsAreaSale(v bool) *TaobaoWlbWmsSkuCreateRequest {
    s.IsAreaSale = &v
    return s
}

func (req *TaobaoWlbWmsSkuCreateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.IsSnMgt != nil) {
        paramMap["is_sn_mgt"] = *req.IsSnMgt
    }
    if(req.Lifecycle != nil) {
        paramMap["lifecycle"] = *req.Lifecycle
    }
    if(req.IsShelflife != nil) {
        paramMap["is_shelflife"] = *req.IsShelflife
    }
    if(req.ApprovalNumber != nil) {
        paramMap["approval_number"] = *req.ApprovalNumber
    }
    if(req.OriginAddress != nil) {
        paramMap["origin_address"] = *req.OriginAddress
    }
    if(req.GrossWeight != nil) {
        paramMap["gross_weight"] = *req.GrossWeight
    }
    if(req.Size != nil) {
        paramMap["size"] = *req.Size
    }
    if(req.Color != nil) {
        paramMap["color"] = *req.Color
    }
    if(req.Specification != nil) {
        paramMap["specification"] = *req.Specification
    }
    if(req.BrandName != nil) {
        paramMap["brand_name"] = *req.BrandName
    }
    if(req.Brand != nil) {
        paramMap["brand"] = *req.Brand
    }
    if(req.CategoryName != nil) {
        paramMap["category_name"] = *req.CategoryName
    }
    if(req.Category != nil) {
        paramMap["category"] = *req.Category
    }
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    if(req.Title != nil) {
        paramMap["title"] = *req.Title
    }
    if(req.Name != nil) {
        paramMap["name"] = *req.Name
    }
    if(req.BarCode != nil) {
        paramMap["bar_code"] = *req.BarCode
    }
    if(req.ItemCode != nil) {
        paramMap["item_code"] = *req.ItemCode
    }
    if(req.Pcs != nil) {
        paramMap["pcs"] = *req.Pcs
    }
    if(req.Volume != nil) {
        paramMap["volume"] = *req.Volume
    }
    if(req.Height != nil) {
        paramMap["height"] = *req.Height
    }
    if(req.ExtendFields != nil) {
        paramMap["extend_fields"] = *req.ExtendFields
    }
    if(req.UseYn != nil) {
        paramMap["use_yn"] = *req.UseYn
    }
    if(req.IsBatchMgt != nil) {
        paramMap["is_batch_mgt"] = *req.IsBatchMgt
    }
    if(req.CostPrice != nil) {
        paramMap["cost_price"] = *req.CostPrice
    }
    if(req.ItemPrice != nil) {
        paramMap["item_price"] = *req.ItemPrice
    }
    if(req.TagPrice != nil) {
        paramMap["tag_price"] = *req.TagPrice
    }
    if(req.IsDanger != nil) {
        paramMap["is_danger"] = *req.IsDanger
    }
    if(req.IsHygroscopic != nil) {
        paramMap["is_hygroscopic"] = *req.IsHygroscopic
    }
    if(req.Width != nil) {
        paramMap["width"] = *req.Width
    }
    if(req.Length != nil) {
        paramMap["length"] = *req.Length
    }
    if(req.NetWeight != nil) {
        paramMap["net_weight"] = *req.NetWeight
    }
    if(req.StoreCode != nil) {
        paramMap["store_code"] = *req.StoreCode
    }
    if(req.AdventLifecycle != nil) {
        paramMap["advent_lifecycle"] = *req.AdventLifecycle
    }
    if(req.RejectLifecycle != nil) {
        paramMap["reject_lifecycle"] = *req.RejectLifecycle
    }
    if(req.LockupLifecycle != nil) {
        paramMap["lockup_lifecycle"] = *req.LockupLifecycle
    }
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.IsAreaSale != nil) {
        paramMap["is_area_sale"] = *req.IsAreaSale
    }
    return paramMap
}

func (req *TaobaoWlbWmsSkuCreateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}