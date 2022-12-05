package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoWlbItemBatchQueryResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回结果记录的数量
    */
    TotalCount  int64 `json:"total_count,omitempty" `
    /*
        商品库存及批次信息查询结果
    */
    ItemInventoryBatchList  []domain.TaobaoWlbItemBatchQueryWlbItemBatchInventory `json:"item_inventory_batch_list,omitempty" `
}
