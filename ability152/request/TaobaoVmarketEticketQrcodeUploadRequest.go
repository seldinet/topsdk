package request

import (
        "topsdk/util"
    )

type TaobaoVmarketEticketQrcodeUploadRequest struct {
    /*
        上传的图片byte[]  小于300K，图片尺寸400*400以内     */
    ImgBytes  *[]byte `json:"img_bytes" required:"true" `
    /*
        码商ID     */
    CodeMerchantId  *int64 `json:"code_merchant_id" required:"true" `
}

func (s *TaobaoVmarketEticketQrcodeUploadRequest) SetImgBytes(v []byte) *TaobaoVmarketEticketQrcodeUploadRequest {
    s.ImgBytes = &v
    return s
}
func (s *TaobaoVmarketEticketQrcodeUploadRequest) SetCodeMerchantId(v int64) *TaobaoVmarketEticketQrcodeUploadRequest {
    s.CodeMerchantId = &v
    return s
}

func (req *TaobaoVmarketEticketQrcodeUploadRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.CodeMerchantId != nil) {
        paramMap["code_merchant_id"] = *req.CodeMerchantId
    }
    return paramMap
}

func (req *TaobaoVmarketEticketQrcodeUploadRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    if req.ImgBytes != nil {
        fileMap["img_bytes"] = *req.ImgBytes
    }
    return fileMap
}