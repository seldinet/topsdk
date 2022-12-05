package response

import (
)

type TaobaoLocationRelationEditResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        错误信息
    */
    ErrorMsg  string `json:"error_msg,omitempty" `
    /*
        是否成功
    */
    IsSuccess  bool `json:"is_success,omitempty" `
    /*
        错误码
    */
    Errorcode  string `json:"errorcode,omitempty" `
}
