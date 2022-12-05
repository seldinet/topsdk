package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoFenxiaoProductcatsGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        查询结果记录数
    */
    TotalResults  int64 `json:"total_results,omitempty" `
    /*
        产品线列表。返回 ProductCat 包含的字段信息。
    */
    Productcats  []domain.TaobaoFenxiaoProductcatsGetProductCat `json:"productcats,omitempty" `
}
