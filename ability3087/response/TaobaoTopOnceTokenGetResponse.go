package response

import (
)

type TaobaoTopOnceTokenGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        响应编码
    */
    ResultCode  string `json:"result_code,omitempty" `
    /*
        token
    */
    Token  string `json:"token,omitempty" `
    /*
        失败详情
    */
    ResultMsg  string `json:"result_msg,omitempty" `
}
