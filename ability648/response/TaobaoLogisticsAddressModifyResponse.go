package response

import (
    "topsdk/ability648/domain"
)

type TaobaoLogisticsAddressModifyResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        只返回修改时间modify_date
    */
    AddressResult  domain.TaobaoLogisticsAddressModifyAddressResult `json:"address_result,omitempty" `
}
