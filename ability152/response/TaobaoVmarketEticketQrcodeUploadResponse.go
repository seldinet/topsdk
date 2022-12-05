package response

import (
)

type TaobaoVmarketEticketQrcodeUploadResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        1:成功  其它为失败
    */
    RetCode  int64 `json:"ret_code,omitempty" `
    /*
        图片文件名称
    */
    ImgFilename  string `json:"img_filename,omitempty" `
}
