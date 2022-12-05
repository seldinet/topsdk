package response

import (
)

type AlibabaItemPublishSubmitResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        商品创建时间
    */
    CreateTime  string `json:"create_time,omitempty" `
    /*
        商品ID
    */
    ItemId  int64 `json:"item_id,omitempty" `
    /*
        商品所属市场
    */
    Market  string `json:"market,omitempty" `
}
