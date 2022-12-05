package response

import (
)

type TaobaoPictureUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        更新是否成功
    */
    Done  bool `json:"done,omitempty" `
}
