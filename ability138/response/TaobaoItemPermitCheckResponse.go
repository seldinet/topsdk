package response

import (
)

type TaobaoItemPermitCheckResponse struct {

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
    Error  bool `json:"error,omitempty" `
    /*
        错误码
    */
    Errorcode  string `json:"errorcode,omitempty" `
}
