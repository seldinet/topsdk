package request

import (
        "topsdk/util"
    )

type TaobaoEticketMerchantImgUploadRequest struct {
    /*
        二维码图片     */
    ImgBytes  *[]byte `json:"img_bytes" required:"true" `
}

func (s *TaobaoEticketMerchantImgUploadRequest) SetImgBytes(v []byte) *TaobaoEticketMerchantImgUploadRequest {
    s.ImgBytes = &v
    return s
}

func (req *TaobaoEticketMerchantImgUploadRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    return paramMap
}

func (req *TaobaoEticketMerchantImgUploadRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    if req.ImgBytes != nil {
        fileMap["img_bytes"] = *req.ImgBytes
    }
    return fileMap
}