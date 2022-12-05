package request


type TaobaoVmarketEticketReverseRequest struct {
    /*
        进行验码的电子凭证订单的订单ID     */
    OrderId  *int64 `json:"order_id" required:"true" `
    /*
        冲正的码，只支持单个码     */
    ReverseCode  *string `json:"reverse_code" required:"true" `
    /*
        冲正份数（必须是和被冲正的核销记录的份数一致）     */
    ReverseNum  *int64 `json:"reverse_num" required:"true" `
    /*
        需要冲正的核销记录对应核销流水号（对应的核销操作时候传递的自定义流水号）     */
    ConsumeSecialNum  *string `json:"consume_secial_num" required:"true" `
    /*
        所有冲正后需要重新生成的码和对应的次数。码和次数之间用英文冒号分隔，多个码之间用英文逗号分隔。如果冲正后不需要重新生成码，留空     */
    VerifyCodes  *string `json:"verify_codes,omitempty" required:"false" `
    /*
        安全验证token，需要和该订单发码通知中的token一致     */
    Token  *string `json:"token" required:"true" `
    /*
        码商ID，是码商的话必须传递，如果是信任卖家不要传     */
    CodemerchantId  *int64 `json:"codemerchant_id,omitempty" required:"false" `
    /*
        机具id，如果是码商必须传，如果是信任卖家不要传     */
    Posid  *string `json:"posid,omitempty" required:"false" `
    /*
        不需要上传二维码图片或者冲正后不需要变更码的请不要传，需要传入二维码的码商请先调用taobao.vmarket.eticket.qrcode.upload接口，将返回的img_filename文件名称作为参数，多个文件名用逗号隔开且与参数verify_codes按从左到有的顺序一一对应。     */
    QrImages  *string `json:"qr_images,omitempty" required:"false" `
}

func (s *TaobaoVmarketEticketReverseRequest) SetOrderId(v int64) *TaobaoVmarketEticketReverseRequest {
    s.OrderId = &v
    return s
}
func (s *TaobaoVmarketEticketReverseRequest) SetReverseCode(v string) *TaobaoVmarketEticketReverseRequest {
    s.ReverseCode = &v
    return s
}
func (s *TaobaoVmarketEticketReverseRequest) SetReverseNum(v int64) *TaobaoVmarketEticketReverseRequest {
    s.ReverseNum = &v
    return s
}
func (s *TaobaoVmarketEticketReverseRequest) SetConsumeSecialNum(v string) *TaobaoVmarketEticketReverseRequest {
    s.ConsumeSecialNum = &v
    return s
}
func (s *TaobaoVmarketEticketReverseRequest) SetVerifyCodes(v string) *TaobaoVmarketEticketReverseRequest {
    s.VerifyCodes = &v
    return s
}
func (s *TaobaoVmarketEticketReverseRequest) SetToken(v string) *TaobaoVmarketEticketReverseRequest {
    s.Token = &v
    return s
}
func (s *TaobaoVmarketEticketReverseRequest) SetCodemerchantId(v int64) *TaobaoVmarketEticketReverseRequest {
    s.CodemerchantId = &v
    return s
}
func (s *TaobaoVmarketEticketReverseRequest) SetPosid(v string) *TaobaoVmarketEticketReverseRequest {
    s.Posid = &v
    return s
}
func (s *TaobaoVmarketEticketReverseRequest) SetQrImages(v string) *TaobaoVmarketEticketReverseRequest {
    s.QrImages = &v
    return s
}

func (req *TaobaoVmarketEticketReverseRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OrderId != nil) {
        paramMap["order_id"] = *req.OrderId
    }
    if(req.ReverseCode != nil) {
        paramMap["reverse_code"] = *req.ReverseCode
    }
    if(req.ReverseNum != nil) {
        paramMap["reverse_num"] = *req.ReverseNum
    }
    if(req.ConsumeSecialNum != nil) {
        paramMap["consume_secial_num"] = *req.ConsumeSecialNum
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
    if(req.Posid != nil) {
        paramMap["posid"] = *req.Posid
    }
    if(req.QrImages != nil) {
        paramMap["qr_images"] = *req.QrImages
    }
    return paramMap
}

func (req *TaobaoVmarketEticketReverseRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}