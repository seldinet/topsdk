package response

import (
    "topsdk/ability6/domain"
)

type TaobaoRefundRefusereasonGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        卖家拒绝原因对象
    */
    Reasons  []domain.TaobaoRefundRefusereasonGetReason `json:"reasons,omitempty" `
    /*
        原因个数
    */
    TotalResults  int64 `json:"total_results,omitempty" `
    /*
        是否存在下一页
    */
    HasNext  bool `json:"has_next,omitempty" `
}
