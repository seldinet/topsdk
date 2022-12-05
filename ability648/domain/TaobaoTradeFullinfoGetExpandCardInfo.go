package domain


type TaobaoTradeFullinfoGetExpandCardInfo struct {
    /*
        买卡订单本金     */
    BasicPrice  *string `json:"basic_price,omitempty" `

    /*
        买卡订单权益金     */
    ExpandPrice  *string `json:"expand_price,omitempty" `

    /*
        用卡订单使用的本金     */
    BasicPriceUsed  *string `json:"basic_price_used,omitempty" `

    /*
        用卡订单使用的权益金     */
    ExpandPriceUsed  *string `json:"expand_price_used,omitempty" `

}

func (s *TaobaoTradeFullinfoGetExpandCardInfo) SetBasicPrice(v string) *TaobaoTradeFullinfoGetExpandCardInfo {
    s.BasicPrice = &v
    return s
}
func (s *TaobaoTradeFullinfoGetExpandCardInfo) SetExpandPrice(v string) *TaobaoTradeFullinfoGetExpandCardInfo {
    s.ExpandPrice = &v
    return s
}
func (s *TaobaoTradeFullinfoGetExpandCardInfo) SetBasicPriceUsed(v string) *TaobaoTradeFullinfoGetExpandCardInfo {
    s.BasicPriceUsed = &v
    return s
}
func (s *TaobaoTradeFullinfoGetExpandCardInfo) SetExpandPriceUsed(v string) *TaobaoTradeFullinfoGetExpandCardInfo {
    s.ExpandPriceUsed = &v
    return s
}
