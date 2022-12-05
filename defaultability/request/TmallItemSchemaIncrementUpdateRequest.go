package request


type TmallItemSchemaIncrementUpdateRequest struct {
    /*
        需要编辑的商品ID     */
    ItemId  *int64 `json:"item_id" required:"true" `
    /*
        根据tmall.item.increment.update.schema.get生成的商品增量编辑规则入参数据。需要更新的字段，一定要在入参的XML重点update_fields字段中明确指明     */
    XmlData  *string `json:"xml_data" required:"true" `
}

func (s *TmallItemSchemaIncrementUpdateRequest) SetItemId(v int64) *TmallItemSchemaIncrementUpdateRequest {
    s.ItemId = &v
    return s
}
func (s *TmallItemSchemaIncrementUpdateRequest) SetXmlData(v string) *TmallItemSchemaIncrementUpdateRequest {
    s.XmlData = &v
    return s
}

func (req *TmallItemSchemaIncrementUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.XmlData != nil) {
        paramMap["xml_data"] = *req.XmlData
    }
    return paramMap
}

func (req *TmallItemSchemaIncrementUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}