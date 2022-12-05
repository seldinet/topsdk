package response

import (
    "topsdk/util"
)

type TaobaoWlbOrderCancelResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        修改时间，只有在取消成功的情况下，才可以做
    */
    ModifyTime  util.LocalTime `json:"modify_time,omitempty" `
    /*
        错误编码列表
    */
    ErrorCodeList  string `json:"error_code_list,omitempty" `
}
