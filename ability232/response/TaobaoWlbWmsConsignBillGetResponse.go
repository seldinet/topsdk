package response

import (
    "topsdk/ability232/domain"
)

type TaobaoWlbWmsConsignBillGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        商品信息列表
    */
    ConsignSendInfoList  []domain.TaobaoWlbWmsConsignBillGetConsignsendinfolist `json:"consign_send_info_list,omitempty" `
}
