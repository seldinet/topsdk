package response

import (
)

type TaobaoWlbItemCombinationGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        组合子商品id列表
    */
    ItemIdList  []int64 `json:"item_id_list,omitempty" `
}
