package response

import (
)

type TmallItemIncrementUpdateSchemaGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回增量更新商品的规则文档
    */
    UpdateItemResult  string `json:"update_item_result,omitempty" `
}
