package response

import (
)

type TaobaoVmarketEticketFlowResendResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        1成功;0失败
    */
    RetCode  int64 `json:"ret_code,omitempty" `
    /*
        错误提示信息
    */
    ErrorMsg  string `json:"error_msg,omitempty" `
}
