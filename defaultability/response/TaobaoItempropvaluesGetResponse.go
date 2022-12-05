package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoItempropvaluesGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        最近修改时间。格式:yyyy-MM-dd HH:mm:ss
    */
    LastModified  string `json:"last_modified,omitempty" `
    /*
        属性值
    */
    PropValues  []domain.TaobaoItempropvaluesGetPropValue `json:"prop_values,omitempty" `
}
