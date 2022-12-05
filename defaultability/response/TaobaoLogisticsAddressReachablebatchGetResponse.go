package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoLogisticsAddressReachablebatchGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        物流是否可达结果列表
    */
    ReachableResults  []domain.TaobaoLogisticsAddressReachablebatchGetAddressReachableTopResult `json:"reachable_results,omitempty" `
}
