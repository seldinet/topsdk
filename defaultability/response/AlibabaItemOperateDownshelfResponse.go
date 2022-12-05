package response

import (
)

type AlibabaItemOperateDownshelfResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        商品下架是否成功
    */
    Result  string `json:"result,omitempty" `
}
