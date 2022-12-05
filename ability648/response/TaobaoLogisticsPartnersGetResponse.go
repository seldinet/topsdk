package response

import (
    "topsdk/ability648/domain"
)

type TaobaoLogisticsPartnersGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        查询揽送范围之内的物流公司信息
    */
    LogisticsPartners  []domain.TaobaoLogisticsPartnersGetLogisticsPartner `json:"logistics_partners,omitempty" `
}
