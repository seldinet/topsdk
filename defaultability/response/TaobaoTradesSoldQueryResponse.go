package response

import (
)

type TaobaoTradesSoldQueryResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        订单ID列表。按照订单创建时间倒序，最多返回最近的100笔订单。
    */
    TidList  []string `json:"tid_list,omitempty" `
}
