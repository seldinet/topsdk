package response

import (
)

type TmallItemAddSchemaGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回发布商品的规则文档
    */
    AddItemResult  string `json:"add_item_result,omitempty" `
}
