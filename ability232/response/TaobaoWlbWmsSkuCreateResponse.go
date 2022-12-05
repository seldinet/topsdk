package response

import (
)

type TaobaoWlbWmsSkuCreateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        系统自动生成
    */
    ItemId  int64 `json:"item_id,omitempty" `
    /*
        错误信息
    */
    WlErrorMsg  string `json:"wl_error_msg,omitempty" `
    /*
        是否成功
    */
    WlSuccess  bool `json:"wl_success,omitempty" `
    /*
        错误码
    */
    WlErrorCode  string `json:"wl_error_code,omitempty" `
}
