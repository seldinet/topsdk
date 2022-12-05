package response

import (
    "topsdk/ability648/domain"
)

type TaobaoTradeAmountGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        主订单的财务信息详情
    */
    TradeAmount  domain.TaobaoTradeAmountGetTradeAmount `json:"trade_amount,omitempty" `
}
