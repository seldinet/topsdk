package request


type TaobaoSkusQuantityUpdateRequest struct {
    /*
        商品数字ID，必填参数     */
    NumIid  *int64 `json:"num_iid" required:"true" `
    /*
        sku库存批量修改入参，用于指定一批sku和每个sku的库存修改值，特殊可填。格式为skuId:库存修改值;skuId:库存修改值。最多支持20个SKU同时修改。     */
    SkuidQuantities  *string `json:"skuid_quantities,omitempty" required:"false" `
    /*
        特殊可选，skuIdQuantities为空的时候用该字段通过outerId来指定sku和其库存修改值。格式为outerId:库存修改值;outerId:库存修改值。当skuIdQuantities不为空的时候该字段失效。当一个outerId对应多个sku时，所有匹配到的sku都会被修改库存。最多支持20个SKU同时修改。     */
    OuteridQuantities  *string `json:"outerid_quantities,omitempty" required:"false" `
    /*
        库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新。当选择全量更新时，如果库存更新值传入的是负数，会出错并返回错误码；当选择增量更新时，如果库存更新值为负数且绝对值大于当前库存，则sku库存会设置为0. defalutValue��1    */
    Type  *int64 `json:"type,omitempty" required:"false" `
}

func (s *TaobaoSkusQuantityUpdateRequest) SetNumIid(v int64) *TaobaoSkusQuantityUpdateRequest {
    s.NumIid = &v
    return s
}
func (s *TaobaoSkusQuantityUpdateRequest) SetSkuidQuantities(v string) *TaobaoSkusQuantityUpdateRequest {
    s.SkuidQuantities = &v
    return s
}
func (s *TaobaoSkusQuantityUpdateRequest) SetOuteridQuantities(v string) *TaobaoSkusQuantityUpdateRequest {
    s.OuteridQuantities = &v
    return s
}
func (s *TaobaoSkusQuantityUpdateRequest) SetType(v int64) *TaobaoSkusQuantityUpdateRequest {
    s.Type = &v
    return s
}

func (req *TaobaoSkusQuantityUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.NumIid != nil) {
        paramMap["num_iid"] = *req.NumIid
    }
    if(req.SkuidQuantities != nil) {
        paramMap["skuid_quantities"] = *req.SkuidQuantities
    }
    if(req.OuteridQuantities != nil) {
        paramMap["outerid_quantities"] = *req.OuteridQuantities
    }
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    return paramMap
}

func (req *TaobaoSkusQuantityUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}