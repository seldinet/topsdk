package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoFenxiaoDiscountsGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        记录数
    */
    TotalResults  int64 `json:"total_results,omitempty" `
    /*
        折扣数据结构
    */
    Discounts  []domain.TaobaoFenxiaoDiscountsGetDiscount `json:"discounts,omitempty" `
}
