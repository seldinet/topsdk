package response

import (
    "topsdk/util"
)

type TaobaoVmarketEticketAuthBeforeconsumeResponse struct {

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
        商品标题
    */
    ItemTitle  string `json:"item_title,omitempty" `
    /*
        订单ID
    */
    OrderId  int64 `json:"order_id,omitempty" `
    /*
        淘宝卖家ID
    */
    TaobaoSid  int64 `json:"taobao_sid,omitempty" `
    /*
        淘宝卖家旺旺名称
    */
    SellerNick  string `json:"seller_nick,omitempty" `
    /*
        有效期结束时间
    */
    ValidEnds  util.LocalTime `json:"valid_ends,omitempty" `
    /*
        有效期开始时间
    */
    ValidStart  util.LocalTime `json:"valid_start,omitempty" `
    /*
        当前码剩余可核销数量
    */
    CodeLeftNum  int64 `json:"code_left_num,omitempty" `
}
