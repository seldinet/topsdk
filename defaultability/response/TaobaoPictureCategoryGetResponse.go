package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoPictureCategoryGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        图片分类
    */
    PictureCategories  []domain.TaobaoPictureCategoryGetPictureCategory `json:"picture_categories,omitempty" `
}
