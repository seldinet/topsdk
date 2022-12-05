package request


type TmallItemSizemappingTemplateGetRequest struct {
    /*
        尺码表模板ID     */
    TemplateId  *int64 `json:"template_id" required:"true" `
}

func (s *TmallItemSizemappingTemplateGetRequest) SetTemplateId(v int64) *TmallItemSizemappingTemplateGetRequest {
    s.TemplateId = &v
    return s
}

func (req *TmallItemSizemappingTemplateGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.TemplateId != nil) {
        paramMap["template_id"] = *req.TemplateId
    }
    return paramMap
}

func (req *TmallItemSizemappingTemplateGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}