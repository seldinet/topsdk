package request


type TmallItemUpdateSchemaGetRequest struct {
    /*
        需要编辑的商品ID     */
    ItemId  *int64 `json:"item_id" required:"true" `
    /*
        商品发布的目标类目，必须是叶子类目。如果没有切换类目需求，不需要填写。     */
    CategoryId  *int64 `json:"category_id,omitempty" required:"false" `
    /*
        商品发布的目标product_id。如果没有切换产品的需求，参数可以不填写。     */
    ProductId  *int64 `json:"product_id,omitempty" required:"false" `
}

func (s *TmallItemUpdateSchemaGetRequest) SetItemId(v int64) *TmallItemUpdateSchemaGetRequest {
    s.ItemId = &v
    return s
}
func (s *TmallItemUpdateSchemaGetRequest) SetCategoryId(v int64) *TmallItemUpdateSchemaGetRequest {
    s.CategoryId = &v
    return s
}
func (s *TmallItemUpdateSchemaGetRequest) SetProductId(v int64) *TmallItemUpdateSchemaGetRequest {
    s.ProductId = &v
    return s
}

func (req *TmallItemUpdateSchemaGetRequest) ToMap() map[string]interface{} {
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
    return paramMap
}

func (req *TmallItemUpdateSchemaGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}