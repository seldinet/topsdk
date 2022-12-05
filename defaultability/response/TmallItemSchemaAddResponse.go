package response

import (
    "topsdk/util"
)

type TmallItemSchemaAddResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回商品发布结果
    */
    AddItemResult  string `json:"add_item_result,omitempty" `
    /*
        发布商品操作成功时间
    */
    GmtCreate  util.LocalTime `json:"gmt_create,omitempty" `
}
