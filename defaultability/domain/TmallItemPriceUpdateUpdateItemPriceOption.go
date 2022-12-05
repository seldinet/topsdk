package domain


type TmallItemPriceUpdateUpdateItemPriceOption struct {
    /*
        是否忽略涉嫌炒信警告信息     */
    IgnoreFakeCredit  *bool `json:"ignore_fake_credit,omitempty" `

    /*
        目标币种，非必填，仅支持天猫国际官网同购商家将外币价格修改成人民币价格时使用     */
    CurrencyType  *string `json:"currency_type,omitempty" `

}

func (s *TmallItemPriceUpdateUpdateItemPriceOption) SetIgnoreFakeCredit(v bool) *TmallItemPriceUpdateUpdateItemPriceOption {
    s.IgnoreFakeCredit = &v
    return s
}
func (s *TmallItemPriceUpdateUpdateItemPriceOption) SetCurrencyType(v string) *TmallItemPriceUpdateUpdateItemPriceOption {
    s.CurrencyType = &v
    return s
}
