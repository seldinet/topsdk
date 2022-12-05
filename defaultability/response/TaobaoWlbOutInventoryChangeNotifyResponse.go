package response

import (
    "topsdk/util"
)

type TaobaoWlbOutInventoryChangeNotifyResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        库存变化通知成功时间
    */
    GmtModified  util.LocalTime `json:"gmt_modified,omitempty" `
}
