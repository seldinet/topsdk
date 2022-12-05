package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoItemAddResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        商品结构,仅有numIid和created返回
    */
    Item  domain.TaobaoItemAddItem `json:"item,omitempty" `
}
