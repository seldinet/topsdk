package response

import (
)

type TaobaoWlbItemAddResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        新增的商品
    */
    ItemId  string `json:"item_id,omitempty" `
}
