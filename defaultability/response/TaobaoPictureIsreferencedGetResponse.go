package response

import (
)

type TaobaoPictureIsreferencedGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        图片是否被引用
    */
    IsReferenced  bool `json:"is_referenced,omitempty" `
}
