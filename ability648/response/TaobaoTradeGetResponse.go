package response

import (
    "topsdk/ability648/domain"
)

type TaobaoTradeGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        搜索到的交易信息列表，返回的Trade和Order中包含的具体信息为入参fields请求的字段信息
    */
    Trade  domain.TaobaoTradeGetTrade `json:"trade,omitempty" `
}
