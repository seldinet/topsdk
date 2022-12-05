package request


type TmallItemSizemappingTemplateCreateRequest struct {
    /*
        尺码表模板名称     */
    TemplateName  *string `json:"template_name" required:"true" `
    /*
        尺码表模板内容，格式为"尺码值:维度名称:数值,尺码值:维度名称:数值"。其中，数值的单位，长度单位为厘米（cm），体重单位为公斤（kg）。尺码值，维度数据不能包含数字，特殊字符。数值为0-999.9的数字，且最多一位小数。     */
    TemplateContent  *string `json:"template_content" required:"true" `
}

func (s *TmallItemSizemappingTemplateCreateRequest) SetTemplateName(v string) *TmallItemSizemappingTemplateCreateRequest {
    s.TemplateName = &v
    return s
}
func (s *TmallItemSizemappingTemplateCreateRequest) SetTemplateContent(v string) *TmallItemSizemappingTemplateCreateRequest {
    s.TemplateContent = &v
    return s
}

func (req *TmallItemSizemappingTemplateCreateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.TemplateName != nil) {
        paramMap["template_name"] = *req.TemplateName
    }
    if(req.TemplateContent != nil) {
        paramMap["template_content"] = *req.TemplateContent
    }
    return paramMap
}

func (req *TmallItemSizemappingTemplateCreateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}