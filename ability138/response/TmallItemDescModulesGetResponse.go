package response

import (
    "topsdk/ability138/domain"
)

type TmallItemDescModulesGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回描述模块信息
    */
    ModularDescInfo  domain.TmallItemDescModulesGetModularDescInfo `json:"modular_desc_info,omitempty" `
}
