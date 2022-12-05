package domain


type TaobaoWlbWmsStockInOrderNotifySenderinfowlbwmsstockinordernotifywl struct {
    /*
        发件方电话     */
    SenderPhone  *string `json:"sender_phone,omitempty" `

    /*
        发件方手机     */
    SenderMobile  *string `json:"sender_mobile,omitempty" `

    /*
        发件方名称，销退场景下填写买家名称； 调拨场景下填写调拨出仓库名称；     */
    SenderName  *string `json:"sender_name,omitempty" `

    /*
        发件方编码，销退场景下填写买家nick，旺旺号等； 调拨场景下填写调拨出仓库编码；     */
    SenderCode  *string `json:"sender_code,omitempty" `

    /*
        发件方地址     */
    SenderAddress  *string `json:"sender_address,omitempty" `

    /*
        发件方镇     */
    SenderTown  *string `json:"sender_town,omitempty" `

    /*
        发件方区县     */
    SenderArea  *string `json:"sender_area,omitempty" `

    /*
        发件方城市     */
    SenderCity  *string `json:"sender_city,omitempty" `

    /*
        发件方省份     */
    SenderProvince  *string `json:"sender_province,omitempty" `

    /*
        发件方邮编     */
    SenderZipCode  *string `json:"sender_zip_code,omitempty" `

}

func (s *TaobaoWlbWmsStockInOrderNotifySenderinfowlbwmsstockinordernotifywl) SetSenderPhone(v string) *TaobaoWlbWmsStockInOrderNotifySenderinfowlbwmsstockinordernotifywl {
    s.SenderPhone = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifySenderinfowlbwmsstockinordernotifywl) SetSenderMobile(v string) *TaobaoWlbWmsStockInOrderNotifySenderinfowlbwmsstockinordernotifywl {
    s.SenderMobile = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifySenderinfowlbwmsstockinordernotifywl) SetSenderName(v string) *TaobaoWlbWmsStockInOrderNotifySenderinfowlbwmsstockinordernotifywl {
    s.SenderName = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifySenderinfowlbwmsstockinordernotifywl) SetSenderCode(v string) *TaobaoWlbWmsStockInOrderNotifySenderinfowlbwmsstockinordernotifywl {
    s.SenderCode = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifySenderinfowlbwmsstockinordernotifywl) SetSenderAddress(v string) *TaobaoWlbWmsStockInOrderNotifySenderinfowlbwmsstockinordernotifywl {
    s.SenderAddress = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifySenderinfowlbwmsstockinordernotifywl) SetSenderTown(v string) *TaobaoWlbWmsStockInOrderNotifySenderinfowlbwmsstockinordernotifywl {
    s.SenderTown = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifySenderinfowlbwmsstockinordernotifywl) SetSenderArea(v string) *TaobaoWlbWmsStockInOrderNotifySenderinfowlbwmsstockinordernotifywl {
    s.SenderArea = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifySenderinfowlbwmsstockinordernotifywl) SetSenderCity(v string) *TaobaoWlbWmsStockInOrderNotifySenderinfowlbwmsstockinordernotifywl {
    s.SenderCity = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifySenderinfowlbwmsstockinordernotifywl) SetSenderProvince(v string) *TaobaoWlbWmsStockInOrderNotifySenderinfowlbwmsstockinordernotifywl {
    s.SenderProvince = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifySenderinfowlbwmsstockinordernotifywl) SetSenderZipCode(v string) *TaobaoWlbWmsStockInOrderNotifySenderinfowlbwmsstockinordernotifywl {
    s.SenderZipCode = &v
    return s
}
