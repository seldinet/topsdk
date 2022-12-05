package response

import (
    "topsdk/ability149/domain"
)

type TaobaoFenxiaoDistributorProductsGetResponse struct {

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
        产品对象记录集。返回 FenxiaoProduct 包含的字段信息。
    */
    Products  []domain.TaobaoFenxiaoDistributorProductsGetFenxiaoProduct `json:"products,omitempty" `
}
