package response

import (
    "topsdk/util"
)

type TaobaoFenxiaoProductAddResponse struct {

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
        产品创建时间 时间格式：yyyy-MM-dd HH:mm:ss
    */
    Created  util.LocalTime `json:"created,omitempty" `
}
