package response

import (
    "topsdk/ability149/domain"
)

type TmallChannelProductsGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        产品对象记录集
    */
    Products  []domain.TmallChannelProductsGetTopProductDO `json:"products,omitempty" `
    /*
        查询结果记录数
    */
    TotalResults  int64 `json:"total_results,omitempty" `
}
