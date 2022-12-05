package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoItemPromotionRuleGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        商品规则信息
    */
    Rules  []domain.TaobaoItemPromotionRuleGetItemPromotionRule `json:"rules,omitempty" `
    /*
        商品是否命中更新规则
    */
    Effec  bool `json:"effec,omitempty" `
}
