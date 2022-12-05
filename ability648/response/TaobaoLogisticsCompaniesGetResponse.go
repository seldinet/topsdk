package response

import (
    "topsdk/ability648/domain"
)

type TaobaoLogisticsCompaniesGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        物流公司信息。返回的LogisticCompany包含的具体信息为入参fields请求的字段信息。
    */
    LogisticsCompanies  []domain.TaobaoLogisticsCompaniesGetLogisticsCompany `json:"logistics_companies,omitempty" `
}
