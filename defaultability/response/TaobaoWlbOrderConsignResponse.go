package response

import (
    "topsdk/util"
)

type TaobaoWlbOrderConsignResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        修改时间
    */
    ModifyTime  util.LocalTime `json:"modify_time,omitempty" `
}
