package response

import (
    "topsdk/util"
)

type TmallItemSimpleschemaUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        编辑商品的itemid
    */
    UpdateItemResult  string `json:"update_item_result,omitempty" `
    /*
        编辑商品操作成功时间
    */
    GmtModified  util.LocalTime `json:"gmt_modified,omitempty" `
    /*
        sku与outerId映射信息
    */
    SkuMapJson  string `json:"sku_map_json,omitempty" `
}
