package response

import (
)

type TaobaoPictureCategoryUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        更新图片分类是否成功
    */
    Done  bool `json:"done,omitempty" `
}
