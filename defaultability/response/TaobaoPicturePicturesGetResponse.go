package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoPicturePicturesGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        图片空间图片数据对象
    */
    Pictures  []domain.TaobaoPicturePicturesGetPicture `json:"pictures,omitempty" `
}
