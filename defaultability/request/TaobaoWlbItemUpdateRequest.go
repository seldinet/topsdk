package request

import (
        "topsdk/util"
    )

type TaobaoWlbItemUpdateRequest struct {
    /*
        要修改的商品id     */
    Id  *int64 `json:"id" required:"true" `
    /*
        要修改的商品名称     */
    Name  *string `json:"name,omitempty" required:"false" `
    /*
        要修改的商品标题     */
    Title  *string `json:"title,omitempty" required:"false" `
    /*
        要修改的商品备注     */
    Remark  *string `json:"remark,omitempty" required:"false" `
    /*
        需要修改的商品属性值的列表，如果属性不存在，则新增属性     */
    UpdatePropertyKeyList  *[]string `json:"update_property_key_list,omitempty" required:"false" `
    /*
        需要修改的属性值的列表     */
    UpdatePropertyValueList  *[]string `json:"update_property_value_list,omitempty" required:"false" `
    /*
        需要删除的商品属性key列表     */
    DeletePropertyKeyList  *[]string `json:"delete_property_key_list,omitempty" required:"false" `
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
        商品长度，单位厘米     */
    Length  *int64 `json:"length,omitempty" required:"false" `
    /*
        商品宽度，单位厘米     */
    Width  *int64 `json:"width,omitempty" required:"false" `
    /*
        商品高度，单位厘米     */
    Height  *int64 `json:"height,omitempty" required:"false" `
    /*
        商品体积，单位立方厘米     */
    Volume  *int64 `json:"volume,omitempty" required:"false" `
    /*
        商品货类     */
    GoodsCat  *string `json:"goods_cat,omitempty" required:"false" `
    /*
        商品计价货类     */
    PricingCat  *string `json:"pricing_cat,omitempty" required:"false" `
    /*
        商品包装材料类型     */
    PackageMaterial  *string `json:"package_material,omitempty" required:"false" `
}

func (s *TaobaoWlbItemUpdateRequest) SetId(v int64) *TaobaoWlbItemUpdateRequest {
    s.Id = &v
    return s
}
func (s *TaobaoWlbItemUpdateRequest) SetName(v string) *TaobaoWlbItemUpdateRequest {
    s.Name = &v
    return s
}
func (s *TaobaoWlbItemUpdateRequest) SetTitle(v string) *TaobaoWlbItemUpdateRequest {
    s.Title = &v
    return s
}
func (s *TaobaoWlbItemUpdateRequest) SetRemark(v string) *TaobaoWlbItemUpdateRequest {
    s.Remark = &v
    return s
}
func (s *TaobaoWlbItemUpdateRequest) SetUpdatePropertyKeyList(v []string) *TaobaoWlbItemUpdateRequest {
    s.UpdatePropertyKeyList = &v
    return s
}
func (s *TaobaoWlbItemUpdateRequest) SetUpdatePropertyValueList(v []string) *TaobaoWlbItemUpdateRequest {
    s.UpdatePropertyValueList = &v
    return s
}
func (s *TaobaoWlbItemUpdateRequest) SetDeletePropertyKeyList(v []string) *TaobaoWlbItemUpdateRequest {
    s.DeletePropertyKeyList = &v
    return s
}
func (s *TaobaoWlbItemUpdateRequest) SetIsFriable(v bool) *TaobaoWlbItemUpdateRequest {
    s.IsFriable = &v
    return s
}
func (s *TaobaoWlbItemUpdateRequest) SetIsDangerous(v bool) *TaobaoWlbItemUpdateRequest {
    s.IsDangerous = &v
    return s
}
func (s *TaobaoWlbItemUpdateRequest) SetColor(v string) *TaobaoWlbItemUpdateRequest {
    s.Color = &v
    return s
}
func (s *TaobaoWlbItemUpdateRequest) SetWeight(v int64) *TaobaoWlbItemUpdateRequest {
    s.Weight = &v
    return s
}
func (s *TaobaoWlbItemUpdateRequest) SetLength(v int64) *TaobaoWlbItemUpdateRequest {
    s.Length = &v
    return s
}
func (s *TaobaoWlbItemUpdateRequest) SetWidth(v int64) *TaobaoWlbItemUpdateRequest {
    s.Width = &v
    return s
}
func (s *TaobaoWlbItemUpdateRequest) SetHeight(v int64) *TaobaoWlbItemUpdateRequest {
    s.Height = &v
    return s
}
func (s *TaobaoWlbItemUpdateRequest) SetVolume(v int64) *TaobaoWlbItemUpdateRequest {
    s.Volume = &v
    return s
}
func (s *TaobaoWlbItemUpdateRequest) SetGoodsCat(v string) *TaobaoWlbItemUpdateRequest {
    s.GoodsCat = &v
    return s
}
func (s *TaobaoWlbItemUpdateRequest) SetPricingCat(v string) *TaobaoWlbItemUpdateRequest {
    s.PricingCat = &v
    return s
}
func (s *TaobaoWlbItemUpdateRequest) SetPackageMaterial(v string) *TaobaoWlbItemUpdateRequest {
    s.PackageMaterial = &v
    return s
}

func (req *TaobaoWlbItemUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Id != nil) {
        paramMap["id"] = *req.Id
    }
    if(req.Name != nil) {
        paramMap["name"] = *req.Name
    }
    if(req.Title != nil) {
        paramMap["title"] = *req.Title
    }
    if(req.Remark != nil) {
        paramMap["remark"] = *req.Remark
    }
    if(req.UpdatePropertyKeyList != nil) {
        paramMap["update_property_key_list"] = util.ConvertBasicList(*req.UpdatePropertyKeyList)
    }
    if(req.UpdatePropertyValueList != nil) {
        paramMap["update_property_value_list"] = util.ConvertBasicList(*req.UpdatePropertyValueList)
    }
    if(req.DeletePropertyKeyList != nil) {
        paramMap["delete_property_key_list"] = util.ConvertBasicList(*req.DeletePropertyKeyList)
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
    return paramMap
}

func (req *TaobaoWlbItemUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}