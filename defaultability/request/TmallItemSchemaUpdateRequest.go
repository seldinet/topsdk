package request


type TmallItemSchemaUpdateRequest struct {
    /*
        需要编辑的商品ID     */
    ItemId  *int64 `json:"item_id" required:"true" `
    /*
        商品发布的目标类目，必须是叶子类目。如果没有切换类目需求不需要填写     */
    CategoryId  *int64 `json:"category_id,omitempty" required:"false" `
    /*
        商品发布的目标product_id。如果没有切换类目或者切换产品的需求，参数不用填写     */
    ProductId  *int64 `json:"product_id,omitempty" required:"false" `
    /*
        根据tmall.item.update.schema.get生成的商品编辑规则入参数据     */
    XmlData  *string `json:"xml_data" required:"true" `
}

func (s *TmallItemSchemaUpdateRequest) SetItemId(v int64) *TmallItemSchemaUpdateRequest {
    s.ItemId = &v
    return s
}
func (s *TmallItemSchemaUpdateRequest) SetCategoryId(v int64) *TmallItemSchemaUpdateRequest {
    s.CategoryId = &v
    return s
}
func (s *TmallItemSchemaUpdateRequest) SetProductId(v int64) *TmallItemSchemaUpdateRequest {
    s.ProductId = &v
    return s
}
func (s *TmallItemSchemaUpdateRequest) SetXmlData(v string) *TmallItemSchemaUpdateRequest {
    s.XmlData = &v
    return s
}

func (req *TmallItemSchemaUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
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

func (req *TmallItemSchemaUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}