package response

import (
)

type TaobaoWlbWmsStockInOrderNotifyResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        错误详细
    */
    WlErrorMsg  string `json:"wl_error_msg,omitempty" `
    /*
        错误编码
    */
    WlErrorCode  string `json:"wl_error_code,omitempty" `
    /*
        是否成功
    */
    WlSuccess  bool `json:"wl_success,omitempty" `
    /*
        系统自动生成
    */
    OrderCode  string `json:"order_code,omitempty" `
}
