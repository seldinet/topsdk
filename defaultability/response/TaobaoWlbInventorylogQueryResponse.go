package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoWlbInventorylogQueryResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        库存变更记录
    */
    InventoryLogList  []domain.TaobaoWlbInventorylogQueryWlbItemInventoryLog `json:"inventory_log_list,omitempty" `
    /*
        记录数
    */
    TotalCount  int64 `json:"total_count,omitempty" `
}
