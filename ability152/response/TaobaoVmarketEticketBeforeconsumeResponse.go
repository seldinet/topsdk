package response

import (
    "topsdk/util"
)

type TaobaoVmarketEticketBeforeconsumeResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        1:可以进行核销码操作
    */
    RetCode  int64 `json:"ret_code,omitempty" `
    /*
        有效期结束时间
    */
    ValidEnds  util.LocalTime `json:"valid_ends,omitempty" `
    /*
        有效期开始时间
    */
    ValidStart  util.LocalTime `json:"valid_start,omitempty" `
    /*
        当前订单剩余可核销数量
    */
    LeftNum  int64 `json:"left_num,omitempty" `
    /*
        扩展字段，暂时预留为0，没有任何意义
    */
    LeftAmount  string `json:"left_amount,omitempty" `
    /*
        商品标题
    */
    ItemTitle  string `json:"item_title,omitempty" `
    /*
        订单ID
    */
    OrderId  int64 `json:"order_id,omitempty" `
    /*
        扩展字段，暂时预留为0，没有任何意义
    */
    ItemType  int64 `json:"item_type,omitempty" `
    /*
        当前码剩余可核销数量
    */
    CodeLeftNum  int64 `json:"code_left_num,omitempty" `
}
