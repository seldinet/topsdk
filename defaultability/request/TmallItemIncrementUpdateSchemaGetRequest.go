package request


type TmallItemIncrementUpdateSchemaGetRequest struct {
    /*
        需要编辑的商品ID     */
    ItemId  *int64 `json:"item_id" required:"true" `
    /*
        如果入参xml_data指定了更新的字段，则只返回指定字段的规则（ISV如果功能性很强，如明确更新Title，请拼装好此字段以提升API整体性能）     */
    XmlData  *string `json:"xml_data,omitempty" required:"false" `
}

func (s *TmallItemIncrementUpdateSchemaGetRequest) SetItemId(v int64) *TmallItemIncrementUpdateSchemaGetRequest {
    s.ItemId = &v
    return s
}
func (s *TmallItemIncrementUpdateSchemaGetRequest) SetXmlData(v string) *TmallItemIncrementUpdateSchemaGetRequest {
    s.XmlData = &v
    return s
}

func (req *TmallItemIncrementUpdateSchemaGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.XmlData != nil) {
        paramMap["xml_data"] = *req.XmlData
    }
    return paramMap
}

func (req *TmallItemIncrementUpdateSchemaGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}