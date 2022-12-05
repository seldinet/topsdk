package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoPictureCategoryAddResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        图片分类信息
    */
    PictureCategory  domain.TaobaoPictureCategoryAddPictureCategory `json:"picture_category,omitempty" `
}
