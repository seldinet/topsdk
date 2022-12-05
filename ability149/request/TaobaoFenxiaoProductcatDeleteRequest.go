package request


type TaobaoFenxiaoProductcatDeleteRequest struct {
    /*
        产品线ID     */
    ProductLineId  *int64 `json:"product_line_id" required:"true" `
}

func (s *TaobaoFenxiaoProductcatDeleteRequest) SetProductLineId(v int64) *TaobaoFenxiaoProductcatDeleteRequest {
    s.ProductLineId = &v
    return s
}

func (req *TaobaoFenxiaoProductcatDeleteRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ProductLineId != nil) {
        paramMap["product_line_id"] = *req.ProductLineId
    }
    return paramMap
}

func (req *TaobaoFenxiaoProductcatDeleteRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}