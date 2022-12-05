package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoItemAnchorGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回的宝贝描述模板结果数目
    */
    TotalResults  int64 `json:"total_results,omitempty" `
    /*
        宝贝描述规范化可使用打标模块的锚点信息
    */
    AnchorModules  []domain.TaobaoItemAnchorGetIdsModule `json:"anchor_modules,omitempty" `
}
