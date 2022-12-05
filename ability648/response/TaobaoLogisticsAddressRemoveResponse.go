package response

import (
    "topsdk/ability648/domain"
)

type TaobaoLogisticsAddressRemoveResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        只返回修改日期modify_date
    */
    AddressResult  domain.TaobaoLogisticsAddressRemoveAddressResult `json:"address_result,omitempty" `
}
