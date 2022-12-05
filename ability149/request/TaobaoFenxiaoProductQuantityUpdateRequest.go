package request


type TaobaoFenxiaoProductQuantityUpdateRequest struct {
    /*
        产品ID     */
    ProductId  *int64 `json:"product_id" required:"true" `
    /*
        库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新。当选择全量更新时，如果库存更新值传入的是负数，会出错并返回错误码；当选择增量更新时，如果库存更新值为负数且绝对值大于当前库存，则sku库存会设置为0     */
    Type  *int64 `json:"type,omitempty" required:"false" `
    /*
        库存修改值。产品有sku时，与sku属性顺序对应，用,分隔。产品无sku时，只写库存值。
当全量更新库存时，quantity必须为大于等于0的正整数；当增量更新库存时，quantity为整数，可小于等于0。若增量更新时传入的库存为负数，则负数与实际库存之和不能小于0。比如当前实际库存为1，传入增量更新quantity=-1，库存改为0     */
    Quantity  *string `json:"quantity" required:"true" `
    /*
        sku属性值，产品有sku时填写，多个sku用,分隔。为空时默认该产品无sku，则只修改产品的库存。请参照taobao.fenxiao.products.get接口返回的properties设置入参     */
    Properties  *string `json:"properties,omitempty" required:"false" `
}

func (s *TaobaoFenxiaoProductQuantityUpdateRequest) SetProductId(v int64) *TaobaoFenxiaoProductQuantityUpdateRequest {
    s.ProductId = &v
    return s
}
func (s *TaobaoFenxiaoProductQuantityUpdateRequest) SetType(v int64) *TaobaoFenxiaoProductQuantityUpdateRequest {
    s.Type = &v
    return s
}
func (s *TaobaoFenxiaoProductQuantityUpdateRequest) SetQuantity(v string) *TaobaoFenxiaoProductQuantityUpdateRequest {
    s.Quantity = &v
    return s
}
func (s *TaobaoFenxiaoProductQuantityUpdateRequest) SetProperties(v string) *TaobaoFenxiaoProductQuantityUpdateRequest {
    s.Properties = &v
    return s
}

func (req *TaobaoFenxiaoProductQuantityUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ProductId != nil) {
        paramMap["product_id"] = *req.ProductId
    }
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    if(req.Quantity != nil) {
        paramMap["quantity"] = *req.Quantity
    }
    if(req.Properties != nil) {
        paramMap["properties"] = *req.Properties
    }
    return paramMap
}

func (req *TaobaoFenxiaoProductQuantityUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}