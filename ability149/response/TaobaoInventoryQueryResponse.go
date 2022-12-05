package response

import (
    "topsdk/ability149/domain"
)

type TaobaoInventoryQueryResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        商品总体库存信息
    */
    ItemInventorys  []domain.TaobaoInventoryQueryInventorySum `json:"item_inventorys,omitempty" `
    /*
        提示信息，提示不存在的后端商品
    */
    TipInfos  []domain.TaobaoInventoryQueryTipInfo `json:"tip_infos,omitempty" `
}
