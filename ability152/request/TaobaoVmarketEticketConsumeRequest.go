package request


type TaobaoVmarketEticketConsumeRequest struct {
    /*
        进行验码的电子凭证订单的订单ID     */
    OrderId  *int64 `json:"order_id" required:"true" `
    /*
        核销的码，只支持单个码，多个码核销需要多次调用     */
    VerifyCode  *string `json:"verify_code" required:"true" `
    /*
        核销份数     */
    ConsumeNum  *int64 `json:"consume_num" required:"true" `
    /*
        安全验证token,需要和发码通知中的token一致     */
    Token  *string `json:"token" required:"true" `
    /*
        码商ID,是码商的话必须传递,如果是信任卖家不需要传     */
    CodemerchantId  *int64 `json:"codemerchant_id,omitempty" required:"false" `
    /*
        机具ID(此参数信任卖家可不传递，码商必须传递)     */
    Posid  *string `json:"posid,omitempty" required:"false" `
    /*
        手机后四位(没有特殊说明请不要传该参数)     */
    Mobile  *string `json:"mobile,omitempty" required:"false" `
    /*
        核销后需要重新生成的码，如果不需要重新生成码，不要传该参数     */
    NewCode  *string `json:"new_code,omitempty" required:"false" `
    /*
        自定义核销流水号，如果核销调用失败，可以用该核销流水号进行冲正操作，需要小于等于100个字符(a-zA-Z0-9_)；每次核销都是唯一的流水号     */
    SerialNum  *string `json:"serial_num,omitempty" required:"false" `
    /*
        不需要上传二维码图片或者核销后不需重新生成码码商请不要传，需要传入二维码的码商请先调用taobao.vmarket.eticket.qrcode.upload接口，将返回的img_filename文件名称作为参数（如果二维码不变的话，也可将将发码时传入二维码文件名作为参数传入），文件名与参数new_code必须相互对应。     */
    QrImages  *string `json:"qr_images,omitempty" required:"false" `
}

func (s *TaobaoVmarketEticketConsumeRequest) SetOrderId(v int64) *TaobaoVmarketEticketConsumeRequest {
    s.OrderId = &v
    return s
}
func (s *TaobaoVmarketEticketConsumeRequest) SetVerifyCode(v string) *TaobaoVmarketEticketConsumeRequest {
    s.VerifyCode = &v
    return s
}
func (s *TaobaoVmarketEticketConsumeRequest) SetConsumeNum(v int64) *TaobaoVmarketEticketConsumeRequest {
    s.ConsumeNum = &v
    return s
}
func (s *TaobaoVmarketEticketConsumeRequest) SetToken(v string) *TaobaoVmarketEticketConsumeRequest {
    s.Token = &v
    return s
}
func (s *TaobaoVmarketEticketConsumeRequest) SetCodemerchantId(v int64) *TaobaoVmarketEticketConsumeRequest {
    s.CodemerchantId = &v
    return s
}
func (s *TaobaoVmarketEticketConsumeRequest) SetPosid(v string) *TaobaoVmarketEticketConsumeRequest {
    s.Posid = &v
    return s
}
func (s *TaobaoVmarketEticketConsumeRequest) SetMobile(v string) *TaobaoVmarketEticketConsumeRequest {
    s.Mobile = &v
    return s
}
func (s *TaobaoVmarketEticketConsumeRequest) SetNewCode(v string) *TaobaoVmarketEticketConsumeRequest {
    s.NewCode = &v
    return s
}
func (s *TaobaoVmarketEticketConsumeRequest) SetSerialNum(v string) *TaobaoVmarketEticketConsumeRequest {
    s.SerialNum = &v
    return s
}
func (s *TaobaoVmarketEticketConsumeRequest) SetQrImages(v string) *TaobaoVmarketEticketConsumeRequest {
    s.QrImages = &v
    return s
}

func (req *TaobaoVmarketEticketConsumeRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OrderId != nil) {
        paramMap["order_id"] = *req.OrderId
    }
    if(req.VerifyCode != nil) {
        paramMap["verify_code"] = *req.VerifyCode
    }
    if(req.ConsumeNum != nil) {
        paramMap["consume_num"] = *req.ConsumeNum
    }
    if(req.Token != nil) {
        paramMap["token"] = *req.Token
    }
    if(req.CodemerchantId != nil) {
        paramMap["codemerchant_id"] = *req.CodemerchantId
    }
    if(req.Posid != nil) {
        paramMap["posid"] = *req.Posid
    }
    if(req.Mobile != nil) {
        paramMap["mobile"] = *req.Mobile
    }
    if(req.NewCode != nil) {
        paramMap["new_code"] = *req.NewCode
    }
    if(req.SerialNum != nil) {
        paramMap["serial_num"] = *req.SerialNum
    }
    if(req.QrImages != nil) {
        paramMap["qr_images"] = *req.QrImages
    }
    return paramMap
}

func (req *TaobaoVmarketEticketConsumeRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}