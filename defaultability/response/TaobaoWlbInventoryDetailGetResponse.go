package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoWlbInventoryDetailGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        库存详情对象。其中包括货主ID，仓库编码，库存，库存类型等属性
    */
    InventoryList  []domain.TaobaoWlbInventoryDetailGetWlbInventory `json:"inventory_list,omitempty" `
    /*
        入参的item_id
    */
    ItemId  int64 `json:"item_id,omitempty" `
}
