package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoProductsGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回具体信息为入参fields请求的字段信息
    */
    Products  []domain.TaobaoProductsGetProduct `json:"products,omitempty" `
}
