package request


type TaobaoFenxiaoProductImageDeleteRequest struct {
    /*
        产品ID     */
    ProductId  *int64 `json:"product_id" required:"true" `
    /*
        图片位置     */
    Position  *int64 `json:"position" required:"true" `
    /*
        properties表示sku图片的属性。key:value形式，key是pid，value是vid。如果position是0的话，则properties需要是必传项     */
    Properties  *string `json:"properties,omitempty" required:"false" `
}

func (s *TaobaoFenxiaoProductImageDeleteRequest) SetProductId(v int64) *TaobaoFenxiaoProductImageDeleteRequest {
    s.ProductId = &v
    return s
}
func (s *TaobaoFenxiaoProductImageDeleteRequest) SetPosition(v int64) *TaobaoFenxiaoProductImageDeleteRequest {
    s.Position = &v
    return s
}
func (s *TaobaoFenxiaoProductImageDeleteRequest) SetProperties(v string) *TaobaoFenxiaoProductImageDeleteRequest {
    s.Properties = &v
    return s
}

func (req *TaobaoFenxiaoProductImageDeleteRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ProductId != nil) {
        paramMap["product_id"] = *req.ProductId
    }
    if(req.Position != nil) {
        paramMap["position"] = *req.Position
    }
    if(req.Properties != nil) {
        paramMap["properties"] = *req.Properties
    }
    return paramMap
}

func (req *TaobaoFenxiaoProductImageDeleteRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}