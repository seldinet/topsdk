package request


type TaobaoFenxiaoProductToChannelImportRequest struct {
    /*
        要导入的渠道[21 零售PLUS]目前仅支持此渠道     */
    Channel  *int64 `json:"channel" required:"true" `
    /*
        要导入的产品id     */
    ProductId  *int64 `json:"product_id" required:"true" `
}

func (s *TaobaoFenxiaoProductToChannelImportRequest) SetChannel(v int64) *TaobaoFenxiaoProductToChannelImportRequest {
    s.Channel = &v
    return s
}
func (s *TaobaoFenxiaoProductToChannelImportRequest) SetProductId(v int64) *TaobaoFenxiaoProductToChannelImportRequest {
    s.ProductId = &v
    return s
}

func (req *TaobaoFenxiaoProductToChannelImportRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Channel != nil) {
        paramMap["channel"] = *req.Channel
    }
    if(req.ProductId != nil) {
        paramMap["product_id"] = *req.ProductId
    }
    return paramMap
}

func (req *TaobaoFenxiaoProductToChannelImportRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}