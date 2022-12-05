package request


type TaobaoVmarketEticketSendRequest struct {
    /*
        订单编号     */
    OrderId  *int64 `json:"order_id" required:"true" `
    /*
        发送成功的验证码及可验证次数的列表，码和可验证次数用英文冒号分隔，多个码之间用英文逗号分隔，所有字符都为英文半角     */
    VerifyCodes  *string `json:"verify_codes" required:"true" `
    /*
        安全验证token，需要和发码通知中的token一致     */
    Token  *string `json:"token" required:"true" `
    /*
        码商ID,是码商的话必须传递,如果是信任卖家,不需要传     */
    CodemerchantId  *int64 `json:"codemerchant_id,omitempty" required:"false" `
    /*
        不需要上传二维码图片的码商请不要传，需要传入二维码的码商请先调用taobao.vmarket.eticket.qrcode.upload接口，将返回的img_filename文件名称作为参数，多个文件名用逗号隔开且与参数verify_codes按从左到有的顺序一一对应。     */
    QrImages  *string `json:"qr_images,omitempty" required:"false" `
}

func (s *TaobaoVmarketEticketSendRequest) SetOrderId(v int64) *TaobaoVmarketEticketSendRequest {
    s.OrderId = &v
    return s
}
func (s *TaobaoVmarketEticketSendRequest) SetVerifyCodes(v string) *TaobaoVmarketEticketSendRequest {
    s.VerifyCodes = &v
    return s
}
func (s *TaobaoVmarketEticketSendRequest) SetToken(v string) *TaobaoVmarketEticketSendRequest {
    s.Token = &v
    return s
}
func (s *TaobaoVmarketEticketSendRequest) SetCodemerchantId(v int64) *TaobaoVmarketEticketSendRequest {
    s.CodemerchantId = &v
    return s
}
func (s *TaobaoVmarketEticketSendRequest) SetQrImages(v string) *TaobaoVmarketEticketSendRequest {
    s.QrImages = &v
    return s
}

func (req *TaobaoVmarketEticketSendRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OrderId != nil) {
        paramMap["order_id"] = *req.OrderId
    }
    if(req.VerifyCodes != nil) {
        paramMap["verify_codes"] = *req.VerifyCodes
    }
    if(req.Token != nil) {
        paramMap["token"] = *req.Token
    }
    if(req.CodemerchantId != nil) {
        paramMap["codemerchant_id"] = *req.CodemerchantId
    }
    if(req.QrImages != nil) {
        paramMap["qr_images"] = *req.QrImages
    }
    return paramMap
}

func (req *TaobaoVmarketEticketSendRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}