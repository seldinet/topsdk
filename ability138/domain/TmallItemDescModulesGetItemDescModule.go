package domain


type TmallItemDescModulesGetItemDescModule struct {
    /*
        一个List<String>的Json串，里面是模块引导提示，不超过3条（isv提交时可忽略不传）     */
    Intros  *string `json:"intros,omitempty" `

    /*
        淘宝图片空间的地址链接，示例模板，最多不超过三个（isv提交时可忽略不传）     */
    TplUrls  *string `json:"tpl_urls,omitempty" `

    /*
        模块id ,(注意:用户自定义模块不用填此项。)     */
    ModuleId  *int64 `json:"module_id,omitempty" `

    /*
        是否必填 （isv提交时可忽略不传）     */
    Required  *bool `json:"required,omitempty" `

    /*
        模块名称     */
    ModuleName  *string `json:"module_name,omitempty" `

    /*
        cat_mod:表示是运营设置的类目维度模块，usr_mod表示的是商家自定义模块。     */
    Type  *string `json:"type,omitempty" `

    /*
        描述具体内容     */
    Content  *string `json:"content,omitempty" `

}

func (s *TmallItemDescModulesGetItemDescModule) SetIntros(v string) *TmallItemDescModulesGetItemDescModule {
    s.Intros = &v
    return s
}
func (s *TmallItemDescModulesGetItemDescModule) SetTplUrls(v string) *TmallItemDescModulesGetItemDescModule {
    s.TplUrls = &v
    return s
}
func (s *TmallItemDescModulesGetItemDescModule) SetModuleId(v int64) *TmallItemDescModulesGetItemDescModule {
    s.ModuleId = &v
    return s
}
func (s *TmallItemDescModulesGetItemDescModule) SetRequired(v bool) *TmallItemDescModulesGetItemDescModule {
    s.Required = &v
    return s
}
func (s *TmallItemDescModulesGetItemDescModule) SetModuleName(v string) *TmallItemDescModulesGetItemDescModule {
    s.ModuleName = &v
    return s
}
func (s *TmallItemDescModulesGetItemDescModule) SetType(v string) *TmallItemDescModulesGetItemDescModule {
    s.Type = &v
    return s
}
func (s *TmallItemDescModulesGetItemDescModule) SetContent(v string) *TmallItemDescModulesGetItemDescModule {
    s.Content = &v
    return s
}
