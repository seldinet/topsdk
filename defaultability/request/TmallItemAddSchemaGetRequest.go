package request


type TmallItemAddSchemaGetRequest struct {
    /*
        商品发布的目标类目，必须是叶子类目     */
    CategoryId  *int64 `json:"category_id" required:"true" `
    /*
        商品发布的目标product_id     */
    ProductId  *int64 `json:"product_id" required:"true" `
    /*
        发布商品类型，一口价填“b”，拍卖填"a" defalutValue��b    */
    Type  *string `json:"type,omitempty" required:"false" `
    /*
        正常接口调用时，请忽略这个参数或者填FALSE。这个参数提供给ISV对接Schema时，如果想先获取了解所有字段和规则，可以将此字段设置为true，product_id也就不需要提供了，设置为0即可     */
    IsvInit  *bool `json:"isv_init,omitempty" required:"false" `
}

func (s *TmallItemAddSchemaGetRequest) SetCategoryId(v int64) *TmallItemAddSchemaGetRequest {
    s.CategoryId = &v
    return s
}
func (s *TmallItemAddSchemaGetRequest) SetProductId(v int64) *TmallItemAddSchemaGetRequest {
    s.ProductId = &v
    return s
}
func (s *TmallItemAddSchemaGetRequest) SetType(v string) *TmallItemAddSchemaGetRequest {
    s.Type = &v
    return s
}
func (s *TmallItemAddSchemaGetRequest) SetIsvInit(v bool) *TmallItemAddSchemaGetRequest {
    s.IsvInit = &v
    return s
}

func (req *TmallItemAddSchemaGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.CategoryId != nil) {
        paramMap["category_id"] = *req.CategoryId
    }
    if(req.ProductId != nil) {
        paramMap["product_id"] = *req.ProductId
    }
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    if(req.IsvInit != nil) {
        paramMap["isv_init"] = *req.IsvInit
    }
    return paramMap
}

func (req *TmallItemAddSchemaGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}