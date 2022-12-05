package response

import (
    "topsdk/ability648/domain"
)

type TaobaoLogisticsAddressSearchResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        一组地址库数据，
    */
    Addresses  []domain.TaobaoLogisticsAddressSearchAddressResult `json:"addresses,omitempty" `
}
