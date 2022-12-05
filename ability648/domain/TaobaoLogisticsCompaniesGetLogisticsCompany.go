package domain


type TaobaoLogisticsCompaniesGetLogisticsCompany struct {
    /*
        物流公司标识     */
    Id  *int64 `json:"id,omitempty" `

    /*
        物流公司代码     */
    Code  *string `json:"code,omitempty" `

    /*
        物流公司简称     */
    Name  *string `json:"name,omitempty" `

    /*
        运单号验证正则表达式     */
    RegMailNo  *string `json:"reg_mail_no,omitempty" `

}

func (s *TaobaoLogisticsCompaniesGetLogisticsCompany) SetId(v int64) *TaobaoLogisticsCompaniesGetLogisticsCompany {
    s.Id = &v
    return s
}
func (s *TaobaoLogisticsCompaniesGetLogisticsCompany) SetCode(v string) *TaobaoLogisticsCompaniesGetLogisticsCompany {
    s.Code = &v
    return s
}
func (s *TaobaoLogisticsCompaniesGetLogisticsCompany) SetName(v string) *TaobaoLogisticsCompaniesGetLogisticsCompany {
    s.Name = &v
    return s
}
func (s *TaobaoLogisticsCompaniesGetLogisticsCompany) SetRegMailNo(v string) *TaobaoLogisticsCompaniesGetLogisticsCompany {
    s.RegMailNo = &v
    return s
}
