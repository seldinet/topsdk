package domain


type TaobaoTradeGetLogisticServiceTag struct {
    /*
        物流服务下的标签属性,多个标签之间有";"分隔     */
    ServiceTag  *string `json:"service_tag,omitempty" `

    /*
        消费者选快递请直接判断service_tag是否包含companyCode。而不要判断service_type     */
    ServiceType  *string `json:"service_type,omitempty" `

}

func (s *TaobaoTradeGetLogisticServiceTag) SetServiceTag(v string) *TaobaoTradeGetLogisticServiceTag {
    s.ServiceTag = &v
    return s
}
func (s *TaobaoTradeGetLogisticServiceTag) SetServiceType(v string) *TaobaoTradeGetLogisticServiceTag {
    s.ServiceType = &v
    return s
}
