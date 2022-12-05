package request


type TmallItemSimpleschemaUpdateRequest struct {
    /*
        商品id     */
    ItemId  *int64 `json:"item_id" required:"true" `
    /*
        编辑商品时提交的xml信息     */
    SchemaXmlFields  *string `json:"schema_xml_fields" required:"true" `
}

func (s *TmallItemSimpleschemaUpdateRequest) SetItemId(v int64) *TmallItemSimpleschemaUpdateRequest {
    s.ItemId = &v
    return s
}
func (s *TmallItemSimpleschemaUpdateRequest) SetSchemaXmlFields(v string) *TmallItemSimpleschemaUpdateRequest {
    s.SchemaXmlFields = &v
    return s
}

func (req *TmallItemSimpleschemaUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.SchemaXmlFields != nil) {
        paramMap["schema_xml_fields"] = *req.SchemaXmlFields
    }
    return paramMap
}

func (req *TmallItemSimpleschemaUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}