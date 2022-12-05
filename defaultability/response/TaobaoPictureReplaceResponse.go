package response

import (
)

type TaobaoPictureReplaceResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        图片替换是否成功
    */
    Done  bool `json:"done,omitempty" `
}
