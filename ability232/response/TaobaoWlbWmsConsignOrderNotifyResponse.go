package response

import (
    "topsdk/ability232/domain"
)

type TaobaoWlbWmsConsignOrderNotifyResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        是否成功
    */
    WlSuccess  bool `json:"wl_success,omitempty" `
    /*
        错误编码
    */
    WlErrorCode  string `json:"wl_error_code,omitempty" `
    /*
        错误详细
    */
    WlErrorMsg  string `json:"wl_error_msg,omitempty" `
    /*
        订单编码
    */
    OrderCode  string `json:"order_code,omitempty" `
    /*
        系统自动生成
    */
    ConsignOrderList  []domain.TaobaoWlbWmsConsignOrderNotifyConsignorderlist `json:"consign_order_list,omitempty" `
}
