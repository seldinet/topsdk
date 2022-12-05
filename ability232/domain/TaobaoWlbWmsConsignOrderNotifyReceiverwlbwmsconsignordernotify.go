package domain


type TaobaoWlbWmsConsignOrderNotifyReceiverwlbwmsconsignordernotify struct {
    /*
        收件人电话     */
    ReceiverPhone  *string `json:"receiver_phone,omitempty" `

    /*
        收件人手机     */
    ReceiverMobile  *string `json:"receiver_mobile,omitempty" `

    /*
        收件人Nick     */
    ReceiverNick  *string `json:"receiver_nick,omitempty" `

    /*
        收件人名称     */
    ReceiverName  *string `json:"receiver_name,omitempty" `

    /*
        收件方地址     */
    ReceiverAddress  *string `json:"receiver_address,omitempty" `

    /*
        收件方街道     */
    ReceiverTown  *string `json:"receiver_town,omitempty" `

    /*
        收件方区县     */
    ReceiverArea  *string `json:"receiver_area,omitempty" `

    /*
        收件方城市     */
    ReceiverCity  *string `json:"receiver_city,omitempty" `

    /*
        收件方省份     */
    ReceiverProvince  *string `json:"receiver_province,omitempty" `

    /*
        收件方邮编     */
    ReceiverZipCode  *string `json:"receiver_zip_code,omitempty" `

}

func (s *TaobaoWlbWmsConsignOrderNotifyReceiverwlbwmsconsignordernotify) SetReceiverPhone(v string) *TaobaoWlbWmsConsignOrderNotifyReceiverwlbwmsconsignordernotify {
    s.ReceiverPhone = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyReceiverwlbwmsconsignordernotify) SetReceiverMobile(v string) *TaobaoWlbWmsConsignOrderNotifyReceiverwlbwmsconsignordernotify {
    s.ReceiverMobile = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyReceiverwlbwmsconsignordernotify) SetReceiverNick(v string) *TaobaoWlbWmsConsignOrderNotifyReceiverwlbwmsconsignordernotify {
    s.ReceiverNick = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyReceiverwlbwmsconsignordernotify) SetReceiverName(v string) *TaobaoWlbWmsConsignOrderNotifyReceiverwlbwmsconsignordernotify {
    s.ReceiverName = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyReceiverwlbwmsconsignordernotify) SetReceiverAddress(v string) *TaobaoWlbWmsConsignOrderNotifyReceiverwlbwmsconsignordernotify {
    s.ReceiverAddress = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyReceiverwlbwmsconsignordernotify) SetReceiverTown(v string) *TaobaoWlbWmsConsignOrderNotifyReceiverwlbwmsconsignordernotify {
    s.ReceiverTown = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyReceiverwlbwmsconsignordernotify) SetReceiverArea(v string) *TaobaoWlbWmsConsignOrderNotifyReceiverwlbwmsconsignordernotify {
    s.ReceiverArea = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyReceiverwlbwmsconsignordernotify) SetReceiverCity(v string) *TaobaoWlbWmsConsignOrderNotifyReceiverwlbwmsconsignordernotify {
    s.ReceiverCity = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyReceiverwlbwmsconsignordernotify) SetReceiverProvince(v string) *TaobaoWlbWmsConsignOrderNotifyReceiverwlbwmsconsignordernotify {
    s.ReceiverProvince = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyReceiverwlbwmsconsignordernotify) SetReceiverZipCode(v string) *TaobaoWlbWmsConsignOrderNotifyReceiverwlbwmsconsignordernotify {
    s.ReceiverZipCode = &v
    return s
}
