package response

import (
    "topsdk/ability648/domain"
)

type TaobaoTradesSoldGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        搜索到的交易信息总数
    */
    TotalResults  int64 `json:"total_results,omitempty" `
    /*
        搜索到的交易信息列表，返回的Trade和Order中包含的具体信息为入参fields请求的字段信息
    */
    Trades  []domain.TaobaoTradesSoldGetTrade `json:"trades,omitempty" `
    /*
        是否存在下一页
    */
    HasNext  bool `json:"has_next,omitempty" `
}
