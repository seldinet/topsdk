package response

import (
    "topsdk/ability648/domain"
)

type TaobaoAreasGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        地址区域信息列表.返回的Area包含的具体信息为入参fields请求的字段信息 。
    */
    Areas  []domain.TaobaoAreasGetArea `json:"areas,omitempty" `
}
