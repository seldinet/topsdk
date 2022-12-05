package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoProductAddResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        产品结构
    */
    Product  domain.TaobaoProductAddProduct `json:"product,omitempty" `
}
