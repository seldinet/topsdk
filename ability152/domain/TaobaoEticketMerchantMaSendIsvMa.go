package domain


type TaobaoEticketMerchantMaSendIsvMa struct {
    /*
        串码码值，请注意需要回传符合规范的凭证码     */
    Code  *string `json:"code,omitempty" `

    /*
        码的可核销份数     */
    Num  *int64 `json:"num,omitempty" `

    /*
        二维码图片文件名。已经申请了上传二维码的码商必填，其它码商无需关心。这个值是taobao.eticket.merchant.img.upload调用后的file_name     */
    QrImage  *string `json:"qr_image,omitempty" `

}

func (s *TaobaoEticketMerchantMaSendIsvMa) SetCode(v string) *TaobaoEticketMerchantMaSendIsvMa {
    s.Code = &v
    return s
}
func (s *TaobaoEticketMerchantMaSendIsvMa) SetNum(v int64) *TaobaoEticketMerchantMaSendIsvMa {
    s.Num = &v
    return s
}
func (s *TaobaoEticketMerchantMaSendIsvMa) SetQrImage(v string) *TaobaoEticketMerchantMaSendIsvMa {
    s.QrImage = &v
    return s
}
