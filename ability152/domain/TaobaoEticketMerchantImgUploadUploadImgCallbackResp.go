package domain


type TaobaoEticketMerchantImgUploadUploadImgCallbackResp struct {
    /*
        扩展属性     */
    AttributeMap  *string `json:"attribute_map,omitempty" `

    /*
        图片在淘宝的文件名     */
    FileName  *string `json:"file_name,omitempty" `

}

func (s *TaobaoEticketMerchantImgUploadUploadImgCallbackResp) SetAttributeMap(v string) *TaobaoEticketMerchantImgUploadUploadImgCallbackResp {
    s.AttributeMap = &v
    return s
}
func (s *TaobaoEticketMerchantImgUploadUploadImgCallbackResp) SetFileName(v string) *TaobaoEticketMerchantImgUploadUploadImgCallbackResp {
    s.FileName = &v
    return s
}
