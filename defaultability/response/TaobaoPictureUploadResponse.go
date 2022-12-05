package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoPictureUploadResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        当前上传的一张图片信息
    */
    Picture  domain.TaobaoPictureUploadPicture `json:"picture,omitempty" `
}
