package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoWlbItemMapGetResponse struct {

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
    IsSuccess  bool `json:"is_success,omitempty" `
    /*
        外部商品实体
    */
    OutEntityItemList  []domain.TaobaoWlbItemMapGetOutEntityItem `json:"out_entity_item_list,omitempty" `
}
