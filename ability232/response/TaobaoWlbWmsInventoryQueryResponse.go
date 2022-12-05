package response

import (
    "topsdk/ability232/domain"
)

type TaobaoWlbWmsInventoryQueryResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        商品详情列表
    */
    ItemList  []domain.TaobaoWlbWmsInventoryQueryWmsInventoryQueryItemlist `json:"item_list,omitempty" `
    /*
        总数
    */
    TotalCount  int64 `json:"total_count,omitempty" `
    /*
        错误代码
    */
    WlErrorCode  string `json:"wl_error_code,omitempty" `
    /*
        错误信息
    */
    WlErrorMsg  string `json:"wl_error_msg,omitempty" `
    /*
        是否成功
    */
    WlSuccess  bool `json:"wl_success,omitempty" `
}
