package request


type TaobaoScitemMapAddRequest struct {
    /*
        前台ic商品id     */
    ItemId  *int64 `json:"item_id" required:"true" `
    /*
        前台ic商品skuid     */
    SkuId  *int64 `json:"sku_id,omitempty" required:"false" `
    /*
        sc_item_id和outer_code 其中一个不能为空     */
    ScItemId  *int64 `json:"sc_item_id,omitempty" required:"false" `
    /*
        sc_item_id和outer_code 其中一个不能为空     */
    OuterCode  *string `json:"outer_code,omitempty" required:"false" `
    /*
        默认值为false
true:进行高级校验,前端商品或SKU的商家编码必须与后端商品的商家编码一致，否则会拒绝关联
false:不进行高级校验 defalutValue��false    */
    NeedCheck  *bool `json:"need_check,omitempty" required:"false" `
}

func (s *TaobaoScitemMapAddRequest) SetItemId(v int64) *TaobaoScitemMapAddRequest {
    s.ItemId = &v
    return s
}
func (s *TaobaoScitemMapAddRequest) SetSkuId(v int64) *TaobaoScitemMapAddRequest {
    s.SkuId = &v
    return s
}
func (s *TaobaoScitemMapAddRequest) SetScItemId(v int64) *TaobaoScitemMapAddRequest {
    s.ScItemId = &v
    return s
}
func (s *TaobaoScitemMapAddRequest) SetOuterCode(v string) *TaobaoScitemMapAddRequest {
    s.OuterCode = &v
    return s
}
func (s *TaobaoScitemMapAddRequest) SetNeedCheck(v bool) *TaobaoScitemMapAddRequest {
    s.NeedCheck = &v
    return s
}

func (req *TaobaoScitemMapAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.SkuId != nil) {
        paramMap["sku_id"] = *req.SkuId
    }
    if(req.ScItemId != nil) {
        paramMap["sc_item_id"] = *req.ScItemId
    }
    if(req.OuterCode != nil) {
        paramMap["outer_code"] = *req.OuterCode
    }
    if(req.NeedCheck != nil) {
        paramMap["need_check"] = *req.NeedCheck
    }
    return paramMap
}

func (req *TaobaoScitemMapAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}