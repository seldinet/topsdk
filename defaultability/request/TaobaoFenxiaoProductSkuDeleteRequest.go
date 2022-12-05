package request


type TaobaoFenxiaoProductSkuDeleteRequest struct {
    /*
        产品id     */
    ProductId  *int64 `json:"product_id" required:"true" `
    /*
        sku属性     */
    Properties  *string `json:"properties" required:"true" `
}

func (s *TaobaoFenxiaoProductSkuDeleteRequest) SetProductId(v int64) *TaobaoFenxiaoProductSkuDeleteRequest {
    s.ProductId = &v
    return s
}
func (s *TaobaoFenxiaoProductSkuDeleteRequest) SetProperties(v string) *TaobaoFenxiaoProductSkuDeleteRequest {
    s.Properties = &v
    return s
}

func (req *TaobaoFenxiaoProductSkuDeleteRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ProductId != nil) {
        paramMap["product_id"] = *req.ProductId
    }
    if(req.Properties != nil) {
        paramMap["properties"] = *req.Properties
    }
    return paramMap
}

func (req *TaobaoFenxiaoProductSkuDeleteRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}