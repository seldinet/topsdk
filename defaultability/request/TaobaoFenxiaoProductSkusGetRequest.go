package request


type TaobaoFenxiaoProductSkusGetRequest struct {
    /*
        产品ID     */
    ProductId  *int64 `json:"product_id" required:"true" `
}

func (s *TaobaoFenxiaoProductSkusGetRequest) SetProductId(v int64) *TaobaoFenxiaoProductSkusGetRequest {
    s.ProductId = &v
    return s
}

func (req *TaobaoFenxiaoProductSkusGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ProductId != nil) {
        paramMap["product_id"] = *req.ProductId
    }
    return paramMap
}

func (req *TaobaoFenxiaoProductSkusGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}