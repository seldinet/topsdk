package response

import (
    "topsdk/util"
)

type TaobaoFenxiaoProductUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        产品ID
    */
    Pid  int64 `json:"pid,omitempty" `
    /*
        更新时间，时间格式：yyyy-MM-dd HH:mm:ss
    */
    Modified  util.LocalTime `json:"modified,omitempty" `
}
