package response

import (
)

type AlibabaItemPublishMarketGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        商家可发布的市场列表,多个以逗号(,)分隔
    */
    Markets  string `json:"markets,omitempty" `
}
