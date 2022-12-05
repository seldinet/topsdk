package response

import (
)

type AlibabaItemOperateUpshelfResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        商品上架是否成功
    */
    Result  string `json:"result,omitempty" `
}
