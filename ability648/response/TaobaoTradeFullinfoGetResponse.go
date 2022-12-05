package response

import (
    "topsdk/ability648/domain"
)

type TaobaoTradeFullinfoGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        交易主订单信息
    */
    Trade  domain.TaobaoTradeFullinfoGetTrade `json:"trade,omitempty" `
}
