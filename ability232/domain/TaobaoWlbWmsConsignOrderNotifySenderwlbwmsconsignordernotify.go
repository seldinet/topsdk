package domain


type TaobaoWlbWmsConsignOrderNotifySenderwlbwmsconsignordernotify struct {
    /*
        发件方电话     */
    SenderPhone  *string `json:"sender_phone,omitempty" `

    /*
        发件方手机     */
    SenderMobile  *string `json:"sender_mobile,omitempty" `

    /*
        发件方名称     */
    SenderName  *string `json:"sender_name,omitempty" `

    /*
        发件方地址     */
    SenderAddress  *string `json:"sender_address,omitempty" `

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

    /*
        发件方镇     */
    SenderTown  *string `json:"sender_town,omitempty" `

}

func (s *TaobaoWlbWmsConsignOrderNotifySenderwlbwmsconsignordernotify) SetSenderPhone(v string) *TaobaoWlbWmsConsignOrderNotifySenderwlbwmsconsignordernotify {
    s.SenderPhone = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifySenderwlbwmsconsignordernotify) SetSenderMobile(v string) *TaobaoWlbWmsConsignOrderNotifySenderwlbwmsconsignordernotify {
    s.SenderMobile = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifySenderwlbwmsconsignordernotify) SetSenderName(v string) *TaobaoWlbWmsConsignOrderNotifySenderwlbwmsconsignordernotify {
    s.SenderName = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifySenderwlbwmsconsignordernotify) SetSenderAddress(v string) *TaobaoWlbWmsConsignOrderNotifySenderwlbwmsconsignordernotify {
    s.SenderAddress = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifySenderwlbwmsconsignordernotify) SetSenderArea(v string) *TaobaoWlbWmsConsignOrderNotifySenderwlbwmsconsignordernotify {
    s.SenderArea = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifySenderwlbwmsconsignordernotify) SetSenderCity(v string) *TaobaoWlbWmsConsignOrderNotifySenderwlbwmsconsignordernotify {
    s.SenderCity = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifySenderwlbwmsconsignordernotify) SetSenderProvince(v string) *TaobaoWlbWmsConsignOrderNotifySenderwlbwmsconsignordernotify {
    s.SenderProvince = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifySenderwlbwmsconsignordernotify) SetSenderZipCode(v string) *TaobaoWlbWmsConsignOrderNotifySenderwlbwmsconsignordernotify {
    s.SenderZipCode = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifySenderwlbwmsconsignordernotify) SetSenderTown(v string) *TaobaoWlbWmsConsignOrderNotifySenderwlbwmsconsignordernotify {
    s.SenderTown = &v
    return s
}
