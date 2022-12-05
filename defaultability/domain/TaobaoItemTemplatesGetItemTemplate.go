package domain


type TaobaoItemTemplatesGetItemTemplate struct {
    /*
        宝贝详情模板的名称     */
    TemplateName  *string `json:"template_name,omitempty" `

    /*
        宝贝模板的id     */
    TemplateId  *int64 `json:"template_id,omitempty" `

    /*
        用于区分宝贝模板属于内店和外店     */
    ShopType  *int64 `json:"shop_type,omitempty" `

}

func (s *TaobaoItemTemplatesGetItemTemplate) SetTemplateName(v string) *TaobaoItemTemplatesGetItemTemplate {
    s.TemplateName = &v
    return s
}
func (s *TaobaoItemTemplatesGetItemTemplate) SetTemplateId(v int64) *TaobaoItemTemplatesGetItemTemplate {
    s.TemplateId = &v
    return s
}
func (s *TaobaoItemTemplatesGetItemTemplate) SetShopType(v int64) *TaobaoItemTemplatesGetItemTemplate {
    s.ShopType = &v
    return s
}
