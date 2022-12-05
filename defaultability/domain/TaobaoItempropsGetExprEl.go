package domain


type TaobaoItempropsGetExprEl struct {
    /*
        0 - 类型为label元素，只用于展示，不用于组装value_name；
1 - 类型为label元素，用于展示，用于组装value_name；
2 - 类型为输入狂元素，主要用于卖家输入数据. 卖家填写完后需要重新设置该元素的文本数据；     */
    Type  *int64 `json:"type,omitempty" `

    /*
        显示文本     */
    Text  *string `json:"text,omitempty" `

    /*
        是否只用于显示的文本元素     */
    IsShowLabel  *bool `json:"is_show_label,omitempty" `

    /*
        是否文本元素，用于显示也用于组装value_name，比如：操作符     */
    IsLabel  *bool `json:"is_label,omitempty" `

    /*
        是否输入框     */
    IsInput  *bool `json:"is_input,omitempty" `

}

func (s *TaobaoItempropsGetExprEl) SetType(v int64) *TaobaoItempropsGetExprEl {
    s.Type = &v
    return s
}
func (s *TaobaoItempropsGetExprEl) SetText(v string) *TaobaoItempropsGetExprEl {
    s.Text = &v
    return s
}
func (s *TaobaoItempropsGetExprEl) SetIsShowLabel(v bool) *TaobaoItempropsGetExprEl {
    s.IsShowLabel = &v
    return s
}
func (s *TaobaoItempropsGetExprEl) SetIsLabel(v bool) *TaobaoItempropsGetExprEl {
    s.IsLabel = &v
    return s
}
func (s *TaobaoItempropsGetExprEl) SetIsInput(v bool) *TaobaoItempropsGetExprEl {
    s.IsInput = &v
    return s
}
