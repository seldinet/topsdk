package domain


type TaobaoDeliveryTemplatesGetTopFee struct {
    /*
        可选值：post:平邮; cod:货到付款; ems:EMS; express:快递公司     */
    ServiceType  *string `json:"service_type,omitempty" `

    /*
        邮费子项涉及的地区,多个地区用逗号连接数量串;可以用taobao.areas.get接口获取.或者参考:http://www.stats.gov.cn/tjbz/xzqhdm/t20080215_402462675.htm 例 (110000,310000,320000,330000).就代表邮费子项涉 及(北京,上海,江苏,浙江)四个地区.填写时要注意对照地区代码值,如果填入错误地区代码,将会出现错误信息.     */
    Destination  *string `json:"destination,omitempty" `

    /*
        首费标准：当valuation(记价方式)为0时输入1-9999范围内的整数     */
    StartStandard  *string `json:"start_standard,omitempty" `

    /*
        首费：输入0.01-999.99（最多包含两位小数） 不能为空也不能为0     */
    StartFee  *string `json:"start_fee,omitempty" `

    /*
        增费标准：当valuation(记价方式)为0时输入1-9999范围内的整数     */
    AddStandard  *string `json:"add_standard,omitempty" `

    /*
        增费：输入0.00-999.99（最多包含两位小数） 不能为空不输入运费时必须输入0     */
    AddFee  *string `json:"add_fee,omitempty" `

}

func (s *TaobaoDeliveryTemplatesGetTopFee) SetServiceType(v string) *TaobaoDeliveryTemplatesGetTopFee {
    s.ServiceType = &v
    return s
}
func (s *TaobaoDeliveryTemplatesGetTopFee) SetDestination(v string) *TaobaoDeliveryTemplatesGetTopFee {
    s.Destination = &v
    return s
}
func (s *TaobaoDeliveryTemplatesGetTopFee) SetStartStandard(v string) *TaobaoDeliveryTemplatesGetTopFee {
    s.StartStandard = &v
    return s
}
func (s *TaobaoDeliveryTemplatesGetTopFee) SetStartFee(v string) *TaobaoDeliveryTemplatesGetTopFee {
    s.StartFee = &v
    return s
}
func (s *TaobaoDeliveryTemplatesGetTopFee) SetAddStandard(v string) *TaobaoDeliveryTemplatesGetTopFee {
    s.AddStandard = &v
    return s
}
func (s *TaobaoDeliveryTemplatesGetTopFee) SetAddFee(v string) *TaobaoDeliveryTemplatesGetTopFee {
    s.AddFee = &v
    return s
}
