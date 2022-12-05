package response

import (
)

type TaobaoVmarketEticketConsumeResponse struct {

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
        整个订单剩余的可核销数量
    */
    LeftNum  int64 `json:"left_num,omitempty" `
    /*
        宝贝标题
    */
    ItemTitle  string `json:"item_title,omitempty" `
    /*
        返回码消费后，需要发送的短信的模版
    */
    SmsTpl  string `json:"sms_tpl,omitempty" `
    /*
        服务内容，用在凭证验证成功后pos机打印小票给消费者
    */
    PrintTpl  string `json:"print_tpl,omitempty" `
    /*
        核销流水号,可以通过该流水号来撤销对应的核销操作
    */
    ConsumeSecialNum  string `json:"consume_secial_num,omitempty" `
    /*
        该核销码在核销后剩余的可核销份数，如果传了new_code来重新生成码，那么这些可核销份数会累积到新的码上
    */
    CodeLeftNum  int64 `json:"code_left_num,omitempty" `
}
