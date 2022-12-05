package response

import (
    "topsdk/util"
)

type TmallItemSchemaIncrementUpdateResponse struct {

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
    UpdateItemResult  string `json:"update_item_result,omitempty" `
    /*
        商品更新操作成功时间
    */
    GmtModified  util.LocalTime `json:"gmt_modified,omitempty" `
}
