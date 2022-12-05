package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoFenxiaoProductsGetResponse struct {

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
        产品对象记录集。返回 FenxiaoProduct 包含的字段信息。
    */
    Products  []domain.TaobaoFenxiaoProductsGetFenxiaoProduct `json:"products,omitempty" `
}
