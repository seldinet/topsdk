package response

import (
)

type TaobaoVmarketEticketAuthConsumeResponse struct {

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
}
