package domain


type TaobaoWlbWmsStockInOrderNotifyReceiverinfowlbwmsstockinordernotifywl struct {
    /*
        收件人手机     */
    ReceiverPhone  *string `json:"receiver_phone,omitempty" `

    /*
        收件人手机     */
    ReceiverMobile  *string `json:"receiver_mobile,omitempty" `

    /*
        收件人名称     */
    ReceiverName  *string `json:"receiver_name,omitempty" `

    /*
        收件方地址     */
    ReceiverAddress  *string `json:"receiver_address,omitempty" `

    /*
        收件人区县     */
    ReceiverArea  *string `json:"receiver_area,omitempty" `

    /*
        收件人城市     */
    ReceiverCity  *string `json:"receiver_city,omitempty" `

    /*
        收件人省份     */
    ReceiverProvince  *string `json:"receiver_province,omitempty" `

    /*
        收件人邮编     */
    ReceiverZipCode  *string `json:"receiver_zip_code,omitempty" `

    /*
        收件人镇     */
    ReceiverTown  *string `json:"receiver_town,omitempty" `

}

func (s *TaobaoWlbWmsStockInOrderNotifyReceiverinfowlbwmsstockinordernotifywl) SetReceiverPhone(v string) *TaobaoWlbWmsStockInOrderNotifyReceiverinfowlbwmsstockinordernotifywl {
    s.ReceiverPhone = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyReceiverinfowlbwmsstockinordernotifywl) SetReceiverMobile(v string) *TaobaoWlbWmsStockInOrderNotifyReceiverinfowlbwmsstockinordernotifywl {
    s.ReceiverMobile = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyReceiverinfowlbwmsstockinordernotifywl) SetReceiverName(v string) *TaobaoWlbWmsStockInOrderNotifyReceiverinfowlbwmsstockinordernotifywl {
    s.ReceiverName = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyReceiverinfowlbwmsstockinordernotifywl) SetReceiverAddress(v string) *TaobaoWlbWmsStockInOrderNotifyReceiverinfowlbwmsstockinordernotifywl {
    s.ReceiverAddress = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyReceiverinfowlbwmsstockinordernotifywl) SetReceiverArea(v string) *TaobaoWlbWmsStockInOrderNotifyReceiverinfowlbwmsstockinordernotifywl {
    s.ReceiverArea = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyReceiverinfowlbwmsstockinordernotifywl) SetReceiverCity(v string) *TaobaoWlbWmsStockInOrderNotifyReceiverinfowlbwmsstockinordernotifywl {
    s.ReceiverCity = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyReceiverinfowlbwmsstockinordernotifywl) SetReceiverProvince(v string) *TaobaoWlbWmsStockInOrderNotifyReceiverinfowlbwmsstockinordernotifywl {
    s.ReceiverProvince = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyReceiverinfowlbwmsstockinordernotifywl) SetReceiverZipCode(v string) *TaobaoWlbWmsStockInOrderNotifyReceiverinfowlbwmsstockinordernotifywl {
    s.ReceiverZipCode = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyReceiverinfowlbwmsstockinordernotifywl) SetReceiverTown(v string) *TaobaoWlbWmsStockInOrderNotifyReceiverinfowlbwmsstockinordernotifywl {
    s.ReceiverTown = &v
    return s
}
