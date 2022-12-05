package request


type TaobaoFenxiaoDistributorProductQuantityGetRequest struct {
    /*
        产品ID     */
    ProductId  *int64 `json:"product_id,omitempty" required:"false" `
    /*
        SKU ID     */
    SkuId  *int64 `json:"sku_id,omitempty" required:"false" `
}

func (s *TaobaoFenxiaoDistributorProductQuantityGetRequest) SetProductId(v int64) *TaobaoFenxiaoDistributorProductQuantityGetRequest {
    s.ProductId = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductQuantityGetRequest) SetSkuId(v int64) *TaobaoFenxiaoDistributorProductQuantityGetRequest {
    s.SkuId = &v
    return s
}

func (req *TaobaoFenxiaoDistributorProductQuantityGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ProductId != nil) {
        paramMap["product_id"] = *req.ProductId
    }
    if(req.SkuId != nil) {
        paramMap["sku_id"] = *req.SkuId
    }
    return paramMap
}

func (req *TaobaoFenxiaoDistributorProductQuantityGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}