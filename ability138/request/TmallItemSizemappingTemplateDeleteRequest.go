package request


type TmallItemSizemappingTemplateDeleteRequest struct {
    /*
        尺码表模板ID     */
    TemplateId  *int64 `json:"template_id" required:"true" `
}

func (s *TmallItemSizemappingTemplateDeleteRequest) SetTemplateId(v int64) *TmallItemSizemappingTemplateDeleteRequest {
    s.TemplateId = &v
    return s
}

func (req *TmallItemSizemappingTemplateDeleteRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.TemplateId != nil) {
        paramMap["template_id"] = *req.TemplateId
    }
    return paramMap
}

func (req *TmallItemSizemappingTemplateDeleteRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}