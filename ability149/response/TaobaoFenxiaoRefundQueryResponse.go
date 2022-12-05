package response

import (
    "topsdk/ability149/domain"
)

type TaobaoFenxiaoRefundQueryResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        按查询条件查到的记录总数
    */
    TotalResults  int64 `json:"total_results,omitempty" `
    /*
        退款列表
    */
    RefundList  []domain.TaobaoFenxiaoRefundQueryRefundDetail `json:"refund_list,omitempty" `
}
