package response

import (
)

type TaobaoVmarketEticketFlowConsumeResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        执行成功
    */
    RetCode  int64 `json:"ret_code,omitempty" `
    /*
        错误提示信息
    */
    ErrorMsg  string `json:"error_msg,omitempty" `
}
