package domain


type TmallItemSizemappingTemplateGetModel struct {
    /*
        尺码表模板内容，格式为"尺码值:维度名称:数值,尺码值:维度名称:数值"。其中，数值的单位，长度单位为厘米（cm），体重单位为公斤（kg）。     */
    TemplateContent  *string `json:"template_content,omitempty" `

    /*
        尺码表模板名称     */
    TemplateName  *string `json:"template_name,omitempty" `

    /*
        尺码表模板ID     */
    TemplateId  *int64 `json:"template_id,omitempty" `

}

func (s *TmallItemSizemappingTemplateGetModel) SetTemplateContent(v string) *TmallItemSizemappingTemplateGetModel {
    s.TemplateContent = &v
    return s
}
func (s *TmallItemSizemappingTemplateGetModel) SetTemplateName(v string) *TmallItemSizemappingTemplateGetModel {
    s.TemplateName = &v
    return s
}
func (s *TmallItemSizemappingTemplateGetModel) SetTemplateId(v int64) *TmallItemSizemappingTemplateGetModel {
    s.TemplateId = &v
    return s
}
