package request

import (
        "topsdk/util"
    )

type TaobaoWlbItemAddRequest struct {
    /*
        商品名称     */
    Name  *string `json:"name" required:"true" `
    /*
        商品标题     */
    Title  *string `json:"title,omitempty" required:"false" `
    /*
        商品编码     */
    ItemCode  *string `json:"item_code" required:"true" `
    /*
        商品备注     */
    Remark  *string `json:"remark,omitempty" required:"false" `
    /*
        NORMAL--普通商品
COMBINE--组合商品
DISTRIBUTION--分销 defalutValue��NORMAL    */
    Type  *string `json:"type,omitempty" required:"false" `
    /*
        是否sku     */
    IsSku  *string `json:"is_sku" required:"true" `
    /*
        属性名列表,目前支持的属性：
毛重:GWeight	
净重:Nweight
皮重: Tweight
自定义属性：
paramkey1
paramkey2
paramkey3
paramkey4     */
    ProNameList  *[]string `json:"pro_name_list,omitempty" required:"false" `
    /*
        属性值列表：
10,8     */
    ProValueList  *[]string `json:"pro_value_list,omitempty" required:"false" `
    /*
        是否易碎品     */
    IsFriable  *bool `json:"is_friable,omitempty" required:"false" `
    /*
        是否危险品     */
    IsDangerous  *bool `json:"is_dangerous,omitempty" required:"false" `
    /*
        商品颜色     */
    Color  *string `json:"color,omitempty" required:"false" `
    /*
        商品重量，单位G     */
    Weight  *int64 `json:"weight,omitempty" required:"false" `
    /*
        商品长度，单位毫米     */
    Length  *int64 `json:"length,omitempty" required:"false" `
    /*
        商品宽度，单位毫米     */
    Width  *int64 `json:"width,omitempty" required:"false" `
    /*
        商品高度，单位毫米     */
    Height  *int64 `json:"height,omitempty" required:"false" `
    /*
        商品体积，单位立方厘米     */
    Volume  *int64 `json:"volume,omitempty" required:"false" `
    /*
        货类     */
    GoodsCat  *string `json:"goods_cat,omitempty" required:"false" `
    /*
        计价货类     */
    PricingCat  *string `json:"pricing_cat,omitempty" required:"false" `
    /*
        商品包装材料类型     */
    PackageMaterial  *string `json:"package_material,omitempty" required:"false" `
    /*
        商品价格，单位：分     */
    Price  *int64 `json:"price,omitempty" required:"false" `
    /*
        是否支持批次     */
    SupportBatch  *bool `json:"support_batch,omitempty" required:"false" `
}

func (s *TaobaoWlbItemAddRequest) SetName(v string) *TaobaoWlbItemAddRequest {
    s.Name = &v
    return s
}
func (s *TaobaoWlbItemAddRequest) SetTitle(v string) *TaobaoWlbItemAddRequest {
    s.Title = &v
    return s
}
func (s *TaobaoWlbItemAddRequest) SetItemCode(v string) *TaobaoWlbItemAddRequest {
    s.ItemCode = &v
    return s
}
func (s *TaobaoWlbItemAddRequest) SetRemark(v string) *TaobaoWlbItemAddRequest {
    s.Remark = &v
    return s
}
func (s *TaobaoWlbItemAddRequest) SetType(v string) *TaobaoWlbItemAddRequest {
    s.Type = &v
    return s
}
func (s *TaobaoWlbItemAddRequest) SetIsSku(v string) *TaobaoWlbItemAddRequest {
    s.IsSku = &v
    return s
}
func (s *TaobaoWlbItemAddRequest) SetProNameList(v []string) *TaobaoWlbItemAddRequest {
    s.ProNameList = &v
    return s
}
func (s *TaobaoWlbItemAddRequest) SetProValueList(v []string) *TaobaoWlbItemAddRequest {
    s.ProValueList = &v
    return s
}
func (s *TaobaoWlbItemAddRequest) SetIsFriable(v bool) *TaobaoWlbItemAddRequest {
    s.IsFriable = &v
    return s
}
func (s *TaobaoWlbItemAddRequest) SetIsDangerous(v bool) *TaobaoWlbItemAddRequest {
    s.IsDangerous = &v
    return s
}
func (s *TaobaoWlbItemAddRequest) SetColor(v string) *TaobaoWlbItemAddRequest {
    s.Color = &v
    return s
}
func (s *TaobaoWlbItemAddRequest) SetWeight(v int64) *TaobaoWlbItemAddRequest {
    s.Weight = &v
    return s
}
func (s *TaobaoWlbItemAddRequest) SetLength(v int64) *TaobaoWlbItemAddRequest {
    s.Length = &v
    return s
}
func (s *TaobaoWlbItemAddRequest) SetWidth(v int64) *TaobaoWlbItemAddRequest {
    s.Width = &v
    return s
}
func (s *TaobaoWlbItemAddRequest) SetHeight(v int64) *TaobaoWlbItemAddRequest {
    s.Height = &v
    return s
}
func (s *TaobaoWlbItemAddRequest) SetVolume(v int64) *TaobaoWlbItemAddRequest {
    s.Volume = &v
    return s
}
func (s *TaobaoWlbItemAddRequest) SetGoodsCat(v string) *TaobaoWlbItemAddRequest {
    s.GoodsCat = &v
    return s
}
func (s *TaobaoWlbItemAddRequest) SetPricingCat(v string) *TaobaoWlbItemAddRequest {
    s.PricingCat = &v
    return s
}
func (s *TaobaoWlbItemAddRequest) SetPackageMaterial(v string) *TaobaoWlbItemAddRequest {
    s.PackageMaterial = &v
    return s
}
func (s *TaobaoWlbItemAddRequest) SetPrice(v int64) *TaobaoWlbItemAddRequest {
    s.Price = &v
    return s
}
func (s *TaobaoWlbItemAddRequest) SetSupportBatch(v bool) *TaobaoWlbItemAddRequest {
    s.SupportBatch = &v
    return s
}

func (req *TaobaoWlbItemAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Name != nil) {
        paramMap["name"] = *req.Name
    }
    if(req.Title != nil) {
        paramMap["title"] = *req.Title
    }
    if(req.ItemCode != nil) {
        paramMap["item_code"] = *req.ItemCode
    }
    if(req.Remark != nil) {
        paramMap["remark"] = *req.Remark
    }
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    if(req.IsSku != nil) {
        paramMap["is_sku"] = *req.IsSku
    }
    if(req.ProNameList != nil) {
        paramMap["pro_name_list"] = util.ConvertBasicList(*req.ProNameList)
    }
    if(req.ProValueList != nil) {
        paramMap["pro_value_list"] = util.ConvertBasicList(*req.ProValueList)
    }
    if(req.IsFriable != nil) {
        paramMap["is_friable"] = *req.IsFriable
    }
    if(req.IsDangerous != nil) {
        paramMap["is_dangerous"] = *req.IsDangerous
    }
    if(req.Color != nil) {
        paramMap["color"] = *req.Color
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
    if(req.GoodsCat != nil) {
        paramMap["goods_cat"] = *req.GoodsCat
    }
    if(req.PricingCat != nil) {
        paramMap["pricing_cat"] = *req.PricingCat
    }
    if(req.PackageMaterial != nil) {
        paramMap["package_material"] = *req.PackageMaterial
    }
    if(req.Price != nil) {
        paramMap["price"] = *req.Price
    }
    if(req.SupportBatch != nil) {
        paramMap["support_batch"] = *req.SupportBatch
    }
    return paramMap
}

func (req *TaobaoWlbItemAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}