package request


type TaobaoVmarketEticketResendRequest struct {
    /*
        订单编号     */
    OrderId  *int64 `json:"order_id" required:"true" `
    /*
        重新发送的验证码及可验证次数的列表，多个码之间用英文逗号分割，需要包含此订单所有可用的码（如果订单总的有10个码，可用的是5个，那么这里设置的是5个可用的码）     */
    VerifyCodes  *string `json:"verify_codes" required:"true" `
    /*
        安全验证token,回传淘宝发通知时发过来的token串     */
    Token  *string `json:"token" required:"true" `
    /*
        码商ID，如果是码商，必须传，如果是信任卖家，不需要传     */
    CodemerchantId  *int64 `json:"codemerchant_id,omitempty" required:"false" `
    /*
        不需要上传二维码图片的码商请不要传，需要传入二维码的码商请先调用taobao.vmarket.eticket.qrcode.upload接口，将返回的img_filename文件名称作为参数（如果二维码不变的话，也可将将发码时传入二维码文件名作为参数传入），多个文件名用逗号隔开且与参数verify_codes按从左到有的顺序一一对应。     */
    QrImages  *string `json:"qr_images,omitempty" required:"false" `
}

func (s *TaobaoVmarketEticketResendRequest) SetOrderId(v int64) *TaobaoVmarketEticketResendRequest {
    s.OrderId = &v
    return s
}
func (s *TaobaoVmarketEticketResendRequest) SetVerifyCodes(v string) *TaobaoVmarketEticketResendRequest {
    s.VerifyCodes = &v
    return s
}
func (s *TaobaoVmarketEticketResendRequest) SetToken(v string) *TaobaoVmarketEticketResendRequest {
    s.Token = &v
    return s
}
func (s *TaobaoVmarketEticketResendRequest) SetCodemerchantId(v int64) *TaobaoVmarketEticketResendRequest {
    s.CodemerchantId = &v
    return s
}
func (s *TaobaoVmarketEticketResendRequest) SetQrImages(v string) *TaobaoVmarketEticketResendRequest {
    s.QrImages = &v
    return s
}

func (req *TaobaoVmarketEticketResendRequest) ToMap() map[string]interface{} {
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

func (req *TaobaoVmarketEticketResendRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}