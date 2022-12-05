package response

import (
    "topsdk/util"
)

type TaobaoFenxiaoProductSkuUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        操作结果
    */
    Result  bool `json:"result,omitempty" `
    /*
        操作时间
    */
    Created  util.LocalTime `json:"created,omitempty" `
}
