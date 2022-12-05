package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoPictureUserinfoGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        用户使用图片空间的信息
    */
    UserInfo  domain.TaobaoPictureUserinfoGetUserInfo `json:"user_info,omitempty" `
}
