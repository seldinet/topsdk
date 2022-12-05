package request


type TaobaoFenxiaoProductGradepriceGetRequest struct {
    /*
        产品id     */
    ProductId  *int64 `json:"product_id" required:"true" `
    /*
        skuId     */
    SkuId  *int64 `json:"sku_id,omitempty" required:"false" `
    /*
        经、代销模式（1：代销、2：经销）     */
    TradeType  *int64 `json:"trade_type,omitempty" required:"false" `
}

func (s *TaobaoFenxiaoProductGradepriceGetRequest) SetProductId(v int64) *TaobaoFenxiaoProductGradepriceGetRequest {
    s.ProductId = &v
    return s
}
func (s *TaobaoFenxiaoProductGradepriceGetRequest) SetSkuId(v int64) *TaobaoFenxiaoProductGradepriceGetRequest {
    s.SkuId = &v
    return s
}
func (s *TaobaoFenxiaoProductGradepriceGetRequest) SetTradeType(v int64) *TaobaoFenxiaoProductGradepriceGetRequest {
    s.TradeType = &v
    return s
}

func (req *TaobaoFenxiaoProductGradepriceGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ProductId != nil) {
        paramMap["product_id"] = *req.ProductId
    }
    if(req.SkuId != nil) {
        paramMap["sku_id"] = *req.SkuId
    }
    if(req.TradeType != nil) {
        paramMap["trade_type"] = *req.TradeType
    }
    return paramMap
}

func (req *TaobaoFenxiaoProductGradepriceGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}