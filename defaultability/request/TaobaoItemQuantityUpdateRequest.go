package request


type TaobaoItemQuantityUpdateRequest struct {
    /*
        商品数字ID，必填参数     */
    NumIid  *int64 `json:"num_iid" required:"true" `
    /*
        要操作的SKU的数字ID，可选。如果不填默认修改宝贝的库存，如果填上则修改该SKU的库存     */
    SkuId  *int64 `json:"sku_id,omitempty" required:"false" `
    /*
        SKU的商家编码，可选参数。如果不填则默认修改宝贝的库存，如果填了则按照商家编码搜索出对应的SKU并修改库存。当sku_id和本字段都填写时以sku_id为准搜索对应SKU     */
    OuterId  *string `json:"outer_id,omitempty" required:"false" `
    /*
        库存修改值，必选。当全量更新库存时，quantity必须为大于等于0的正整数；当增量更新库存时，quantity为整数，可小于等于0。若增量更新时传入的库存为负数，则负数与实际库存之和不能小于0。比如当前实际库存为1，传入增量更新quantity=-1，库存改为0     */
    Quantity  *int64 `json:"quantity" required:"true" `
    /*
        库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新 defalutValue��1    */
    Type  *int64 `json:"type,omitempty" required:"false" `
}

func (s *TaobaoItemQuantityUpdateRequest) SetNumIid(v int64) *TaobaoItemQuantityUpdateRequest {
    s.NumIid = &v
    return s
}
func (s *TaobaoItemQuantityUpdateRequest) SetSkuId(v int64) *TaobaoItemQuantityUpdateRequest {
    s.SkuId = &v
    return s
}
func (s *TaobaoItemQuantityUpdateRequest) SetOuterId(v string) *TaobaoItemQuantityUpdateRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoItemQuantityUpdateRequest) SetQuantity(v int64) *TaobaoItemQuantityUpdateRequest {
    s.Quantity = &v
    return s
}
func (s *TaobaoItemQuantityUpdateRequest) SetType(v int64) *TaobaoItemQuantityUpdateRequest {
    s.Type = &v
    return s
}

func (req *TaobaoItemQuantityUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.NumIid != nil) {
        paramMap["num_iid"] = *req.NumIid
    }
    if(req.SkuId != nil) {
        paramMap["sku_id"] = *req.SkuId
    }
    if(req.OuterId != nil) {
        paramMap["outer_id"] = *req.OuterId
    }
    if(req.Quantity != nil) {
        paramMap["quantity"] = *req.Quantity
    }
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    return paramMap
}

func (req *TaobaoItemQuantityUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}