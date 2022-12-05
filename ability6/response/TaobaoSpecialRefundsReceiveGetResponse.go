package response

import (
    "topsdk/ability6/domain"
)

type TaobaoSpecialRefundsReceiveGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        是否存在下一页
    */
    HasNext  bool `json:"has_next,omitempty" `
    /*
        搜索到的退款信息列表
    */
    Refunds  []domain.TaobaoSpecialRefundsReceiveGetRefund `json:"refunds,omitempty" `
    /*
        搜索到的退款信息总数
    */
    TotalResults  int64 `json:"total_results,omitempty" `
}
