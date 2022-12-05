package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoPictureGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        图片信息列表
    */
    Pictures  []domain.TaobaoPictureGetPicture `json:"pictures,omitempty" `
}
