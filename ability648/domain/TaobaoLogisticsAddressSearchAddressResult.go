package domain

import (
        "topsdk/util"
    )

type TaobaoLogisticsAddressSearchAddressResult struct {
    /*
        地址库ID     */
    ContactId  *int64 `json:"contact_id,omitempty" `

    /*
        联系人姓名     */
    ContactName  *string `json:"contact_name,omitempty" `

    /*
        省     */
    Province  *string `json:"province,omitempty" `

    /*
        市     */
    City  *string `json:"city,omitempty" `

    /*
        区、县     */
    Country  *string `json:"country,omitempty" `

    /*
        详细街道地址，不需要重复填写省/市/区     */
    Addr  *string `json:"addr,omitempty" `

    /*
        地区邮政编码     */
    ZipCode  *string `json:"zip_code,omitempty" `

    /*
        电话号码,手机与电话必需有一个     */
    Phone  *string `json:"phone,omitempty" `

    /*
        手机号码，手机与电话必需有一个 
手机号码不能超过20位     */
    MobilePhone  *string `json:"mobile_phone,omitempty" `

    /*
        公司名称,     */
    SellerCompany  *string `json:"seller_company,omitempty" `

    /*
        备注     */
    Memo  *string `json:"memo,omitempty" `

    /*
        区域ID     */
    AreaId  *int64 `json:"area_id,omitempty" `

    /*
        是否默认发货地址     */
    SendDef  *bool `json:"send_def,omitempty" `

    /*
        是否默认取货地址     */
    GetDef  *bool `json:"get_def,omitempty" `

    /*
        是否默认退货地址     */
    CancelDef  *bool `json:"cancel_def,omitempty" `

    /*
        修改日期时间     */
    ModifyDate  *util.LocalTime `json:"modify_date,omitempty" `

}

func (s *TaobaoLogisticsAddressSearchAddressResult) SetContactId(v int64) *TaobaoLogisticsAddressSearchAddressResult {
    s.ContactId = &v
    return s
}
func (s *TaobaoLogisticsAddressSearchAddressResult) SetContactName(v string) *TaobaoLogisticsAddressSearchAddressResult {
    s.ContactName = &v
    return s
}
func (s *TaobaoLogisticsAddressSearchAddressResult) SetProvince(v string) *TaobaoLogisticsAddressSearchAddressResult {
    s.Province = &v
    return s
}
func (s *TaobaoLogisticsAddressSearchAddressResult) SetCity(v string) *TaobaoLogisticsAddressSearchAddressResult {
    s.City = &v
    return s
}
func (s *TaobaoLogisticsAddressSearchAddressResult) SetCountry(v string) *TaobaoLogisticsAddressSearchAddressResult {
    s.Country = &v
    return s
}
func (s *TaobaoLogisticsAddressSearchAddressResult) SetAddr(v string) *TaobaoLogisticsAddressSearchAddressResult {
    s.Addr = &v
    return s
}
func (s *TaobaoLogisticsAddressSearchAddressResult) SetZipCode(v string) *TaobaoLogisticsAddressSearchAddressResult {
    s.ZipCode = &v
    return s
}
func (s *TaobaoLogisticsAddressSearchAddressResult) SetPhone(v string) *TaobaoLogisticsAddressSearchAddressResult {
    s.Phone = &v
    return s
}
func (s *TaobaoLogisticsAddressSearchAddressResult) SetMobilePhone(v string) *TaobaoLogisticsAddressSearchAddressResult {
    s.MobilePhone = &v
    return s
}
func (s *TaobaoLogisticsAddressSearchAddressResult) SetSellerCompany(v string) *TaobaoLogisticsAddressSearchAddressResult {
    s.SellerCompany = &v
    return s
}
func (s *TaobaoLogisticsAddressSearchAddressResult) SetMemo(v string) *TaobaoLogisticsAddressSearchAddressResult {
    s.Memo = &v
    return s
}
func (s *TaobaoLogisticsAddressSearchAddressResult) SetAreaId(v int64) *TaobaoLogisticsAddressSearchAddressResult {
    s.AreaId = &v
    return s
}
func (s *TaobaoLogisticsAddressSearchAddressResult) SetSendDef(v bool) *TaobaoLogisticsAddressSearchAddressResult {
    s.SendDef = &v
    return s
}
func (s *TaobaoLogisticsAddressSearchAddressResult) SetGetDef(v bool) *TaobaoLogisticsAddressSearchAddressResult {
    s.GetDef = &v
    return s
}
func (s *TaobaoLogisticsAddressSearchAddressResult) SetCancelDef(v bool) *TaobaoLogisticsAddressSearchAddressResult {
    s.CancelDef = &v
    return s
}
func (s *TaobaoLogisticsAddressSearchAddressResult) SetModifyDate(v util.LocalTime) *TaobaoLogisticsAddressSearchAddressResult {
    s.ModifyDate = &v
    return s
}
