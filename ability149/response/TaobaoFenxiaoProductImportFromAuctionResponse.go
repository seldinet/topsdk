package response

import (
    "topsdk/util"
)

type TaobaoFenxiaoProductImportFromAuctionResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        生成的产品id
    */
    Pid  int64 `json:"pid,omitempty" `
    /*
        操作时间
    */
    OptTime  util.LocalTime `json:"opt_time,omitempty" `
}
