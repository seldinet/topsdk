package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoWlbNotifyMessagePageGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        通道消息
    */
    WlbMessages  []domain.TaobaoWlbNotifyMessagePageGetWlbMessage `json:"wlb_messages,omitempty" `
    /*
        2000-01-01 00:00:00
    */
    TotalCount  int64 `json:"total_count,omitempty" `
}
