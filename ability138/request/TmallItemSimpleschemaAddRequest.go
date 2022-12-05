package request


type TmallItemSimpleschemaAddRequest struct {
    /*
        根据tmall.item.add.simpleschema.get生成的商品发布规则入参数据     */
    SchemaXmlFields  *string `json:"schema_xml_fields" required:"true" `
}

func (s *TmallItemSimpleschemaAddRequest) SetSchemaXmlFields(v string) *TmallItemSimpleschemaAddRequest {
    s.SchemaXmlFields = &v
    return s
}

func (req *TmallItemSimpleschemaAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.SchemaXmlFields != nil) {
        paramMap["schema_xml_fields"] = *req.SchemaXmlFields
    }
    return paramMap
}

func (req *TmallItemSimpleschemaAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}