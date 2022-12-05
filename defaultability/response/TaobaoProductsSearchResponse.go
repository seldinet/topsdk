package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoProductsSearchResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        结果总数
    */
    TotalResults  int64 `json:"total_results,omitempty" `
    /*
        返回具体信息为入参fields请求的字段信息
    */
    Products  []domain.TaobaoProductsSearchProduct `json:"products,omitempty" `
}
