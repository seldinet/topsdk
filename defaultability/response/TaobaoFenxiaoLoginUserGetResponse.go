package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoFenxiaoLoginUserGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        登录用户信息
    */
    LoginUser  domain.TaobaoFenxiaoLoginUserGetLoginUser `json:"login_user,omitempty" `
}
