package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoProductUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回product数据结构中的：product_id,modified
    */
    Product  domain.TaobaoProductUpdateProduct `json:"product,omitempty" `
}
