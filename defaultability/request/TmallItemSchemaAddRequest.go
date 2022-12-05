package request


type TmallItemSchemaAddRequest struct {
    /*
        商品发布的目标类目，必须是叶子类目     */
    CategoryId  *int64 `json:"category_id" required:"true" `
    /*
        发布商品的productId，如果tmall.product.match.schema.get获取到得字段为空，这个参数传入0，否则需要通过tmall.product.schema.match查询到得可用productId     */
    ProductId  *int64 `json:"product_id" required:"true" `
    /*
        根据tmall.item.add.schema.get生成的商品发布规则入参数据     */
    XmlData  *string `json:"xml_data" required:"true" `
}

func (s *TmallItemSchemaAddRequest) SetCategoryId(v int64) *TmallItemSchemaAddRequest {
    s.CategoryId = &v
    return s
}
func (s *TmallItemSchemaAddRequest) SetProductId(v int64) *TmallItemSchemaAddRequest {
    s.ProductId = &v
    return s
}
func (s *TmallItemSchemaAddRequest) SetXmlData(v string) *TmallItemSchemaAddRequest {
    s.XmlData = &v
    return s
}

func (req *TmallItemSchemaAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.CategoryId != nil) {
        paramMap["category_id"] = *req.CategoryId
    }
    if(req.ProductId != nil) {
        paramMap["product_id"] = *req.ProductId
    }
    if(req.XmlData != nil) {
        paramMap["xml_data"] = *req.XmlData
    }
    return paramMap
}

func (req *TmallItemSchemaAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}