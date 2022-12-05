package domain


type TaobaoItemPromotionRuleGetItemPromotionRule struct {
    /*
        规则名称     */
    Name  *string `json:"name,omitempty" `

    /*
        规则描叙信息     */
    Message  *string `json:"message,omitempty" `

    /*
        规则生效开始时间     */
    StartTime  *string `json:"start_time,omitempty" `

    /*
        规则生效结束时间     */
    EndTime  *string `json:"end_time,omitempty" `

    /*
        规则类型，常见有SKU锁定规则,下架锁定规则,库存减少锁定规则,库存禁止修改规则,一口价禁止修改规则     */
    Type  *string `json:"type,omitempty" `

}

func (s *TaobaoItemPromotionRuleGetItemPromotionRule) SetName(v string) *TaobaoItemPromotionRuleGetItemPromotionRule {
    s.Name = &v
    return s
}
func (s *TaobaoItemPromotionRuleGetItemPromotionRule) SetMessage(v string) *TaobaoItemPromotionRuleGetItemPromotionRule {
    s.Message = &v
    return s
}
func (s *TaobaoItemPromotionRuleGetItemPromotionRule) SetStartTime(v string) *TaobaoItemPromotionRuleGetItemPromotionRule {
    s.StartTime = &v
    return s
}
func (s *TaobaoItemPromotionRuleGetItemPromotionRule) SetEndTime(v string) *TaobaoItemPromotionRuleGetItemPromotionRule {
    s.EndTime = &v
    return s
}
func (s *TaobaoItemPromotionRuleGetItemPromotionRule) SetType(v string) *TaobaoItemPromotionRuleGetItemPromotionRule {
    s.Type = &v
    return s
}
