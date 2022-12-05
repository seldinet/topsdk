package response

import (
    "topsdk/ability149/domain"
)

type TaobaoScitemMapQueryResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        后端商品映射列表
    */
    ScItemMaps  []domain.TaobaoScitemMapQueryScItemMap `json:"sc_item_maps,omitempty" `
}
