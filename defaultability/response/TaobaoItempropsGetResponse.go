package response

import (
    "topsdk/defaultability/domain"
    "topsdk/util"
)

type TaobaoItempropsGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        最近修改时间(只有取全量或增量的时候会返回该字段)。格式:yyyy-MM-dd HH:mm:ss
    */
    LastModified  util.LocalTime `json:"last_modified,omitempty" `
    /*
        商品属性
    */
    ItemProps  []domain.TaobaoItempropsGetItemProp `json:"item_props,omitempty" `
}
