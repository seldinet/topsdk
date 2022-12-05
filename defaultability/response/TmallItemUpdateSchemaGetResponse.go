package response

import (
)

type TmallItemUpdateSchemaGetResponse struct {

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
    UpdateItemResult  string `json:"update_item_result,omitempty" `
}
