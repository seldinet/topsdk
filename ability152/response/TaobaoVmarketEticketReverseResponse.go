package response

import (
)

type TaobaoVmarketEticketReverseResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        0:失败，1:成功
    */
    RetCode  int64 `json:"ret_code,omitempty" `
    /*
        整个订单的剩余可核销数量
    */
    LeftNum  int64 `json:"left_num,omitempty" `
    /*
        宝贝标题
    */
    ItemTitle  string `json:"item_title,omitempty" `
}
