package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoWlbItemQueryResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        结果总数
    */
    TotalCount  int64 `json:"total_count,omitempty" `
    /*
        商品信息列表
    */
    ItemList  []domain.TaobaoWlbItemQueryWlbItem `json:"item_list,omitempty" `
}
