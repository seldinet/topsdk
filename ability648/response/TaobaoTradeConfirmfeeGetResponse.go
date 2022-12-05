package response

import (
    "topsdk/ability648/domain"
)

type TaobaoTradeConfirmfeeGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        获取到的交易确认收货费用
    */
    TradeConfirmFee  domain.TaobaoTradeConfirmfeeGetTradeConfirmFee `json:"trade_confirm_fee,omitempty" `
}
