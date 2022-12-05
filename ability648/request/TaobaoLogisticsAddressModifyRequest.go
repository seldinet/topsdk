package request


type TaobaoLogisticsAddressModifyRequest struct {
    /*
        地址库ID     */
    ContactId  *int64 `json:"contact_id" required:"true" `
    /*
        联系人姓名
<font color='red'>长度不可超过20个字节</font>     */
    ContactName  *string `json:"contact_name" required:"true" `
    /*
        所在省     */
    Province  *string `json:"province" required:"true" `
    /*
        所在市     */
    City  *string `json:"city" required:"true" `
    /*
        区、县
<br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font>     */
    Country  *string `json:"country,omitempty" required:"false" `
    /*
        详细街道地址，不需要重复填写省/市/区     */
    Addr  *string `json:"addr" required:"true" `
    /*
        地区邮政编码
<br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font>     */
    ZipCode  *string `json:"zip_code,omitempty" required:"false" `
    /*
        电话号码,手机与电话必需有一个     */
    Phone  *string `json:"phone,omitempty" required:"false" `
    /*
        手机号码，手机与电话必需有一个 <br><font color='red'>手机号码不能超过20位</font>     */
    MobilePhone  *string `json:"mobile_phone,omitempty" required:"false" `
    /*
        公司名称,
<br><font color='red'>公司名称长能不能超过20字节</font>     */
    SellerCompany  *string `json:"seller_company,omitempty" required:"false" `
    /*
        备注,<br><font color='red'>备注不能超过256字节</font>     */
    Memo  *string `json:"memo,omitempty" required:"false" `
    /*
        默认发货地址。<br>
<font color='red'>选择此项(true)，将当前地址设为默认发货地址，撤消原默认发货地址</font>     */
    SendDef  *bool `json:"send_def,omitempty" required:"false" `
    /*
        默认取货地址。<br>
<font color='red'>选择此项(true)，将当前地址设为默认取货地址，撤消原默认取货地址</font>     */
    GetDef  *bool `json:"get_def,omitempty" required:"false" `
    /*
        默认退货地址。<br>
<font color='red'>选择此项(true)，将当前地址设为默认退货地址，撤消原默认退货地址</font>     */
    CancelDef  *bool `json:"cancel_def,omitempty" required:"false" `
}

func (s *TaobaoLogisticsAddressModifyRequest) SetContactId(v int64) *TaobaoLogisticsAddressModifyRequest {
    s.ContactId = &v
    return s
}
func (s *TaobaoLogisticsAddressModifyRequest) SetContactName(v string) *TaobaoLogisticsAddressModifyRequest {
    s.ContactName = &v
    return s
}
func (s *TaobaoLogisticsAddressModifyRequest) SetProvince(v string) *TaobaoLogisticsAddressModifyRequest {
    s.Province = &v
    return s
}
func (s *TaobaoLogisticsAddressModifyRequest) SetCity(v string) *TaobaoLogisticsAddressModifyRequest {
    s.City = &v
    return s
}
func (s *TaobaoLogisticsAddressModifyRequest) SetCountry(v string) *TaobaoLogisticsAddressModifyRequest {
    s.Country = &v
    return s
}
func (s *TaobaoLogisticsAddressModifyRequest) SetAddr(v string) *TaobaoLogisticsAddressModifyRequest {
    s.Addr = &v
    return s
}
func (s *TaobaoLogisticsAddressModifyRequest) SetZipCode(v string) *TaobaoLogisticsAddressModifyRequest {
    s.ZipCode = &v
    return s
}
func (s *TaobaoLogisticsAddressModifyRequest) SetPhone(v string) *TaobaoLogisticsAddressModifyRequest {
    s.Phone = &v
    return s
}
func (s *TaobaoLogisticsAddressModifyRequest) SetMobilePhone(v string) *TaobaoLogisticsAddressModifyRequest {
    s.MobilePhone = &v
    return s
}
func (s *TaobaoLogisticsAddressModifyRequest) SetSellerCompany(v string) *TaobaoLogisticsAddressModifyRequest {
    s.SellerCompany = &v
    return s
}
func (s *TaobaoLogisticsAddressModifyRequest) SetMemo(v string) *TaobaoLogisticsAddressModifyRequest {
    s.Memo = &v
    return s
}
func (s *TaobaoLogisticsAddressModifyRequest) SetSendDef(v bool) *TaobaoLogisticsAddressModifyRequest {
    s.SendDef = &v
    return s
}
func (s *TaobaoLogisticsAddressModifyRequest) SetGetDef(v bool) *TaobaoLogisticsAddressModifyRequest {
    s.GetDef = &v
    return s
}
func (s *TaobaoLogisticsAddressModifyRequest) SetCancelDef(v bool) *TaobaoLogisticsAddressModifyRequest {
    s.CancelDef = &v
    return s
}

func (req *TaobaoLogisticsAddressModifyRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ContactId != nil) {
        paramMap["contact_id"] = *req.ContactId
    }
    if(req.ContactName != nil) {
        paramMap["contact_name"] = *req.ContactName
    }
    if(req.Province != nil) {
        paramMap["province"] = *req.Province
    }
    if(req.City != nil) {
        paramMap["city"] = *req.City
    }
    if(req.Country != nil) {
        paramMap["country"] = *req.Country
    }
    if(req.Addr != nil) {
        paramMap["addr"] = *req.Addr
    }
    if(req.ZipCode != nil) {
        paramMap["zip_code"] = *req.ZipCode
    }
    if(req.Phone != nil) {
        paramMap["phone"] = *req.Phone
    }
    if(req.MobilePhone != nil) {
        paramMap["mobile_phone"] = *req.MobilePhone
    }
    if(req.SellerCompany != nil) {
        paramMap["seller_company"] = *req.SellerCompany
    }
    if(req.Memo != nil) {
        paramMap["memo"] = *req.Memo
    }
    if(req.SendDef != nil) {
        paramMap["send_def"] = *req.SendDef
    }
    if(req.GetDef != nil) {
        paramMap["get_def"] = *req.GetDef
    }
    if(req.CancelDef != nil) {
        paramMap["cancel_def"] = *req.CancelDef
    }
    return paramMap
}

func (req *TaobaoLogisticsAddressModifyRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}