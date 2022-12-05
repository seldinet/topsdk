package response

import (
    "topsdk/ability149/domain"
)

type TaobaoFenxiaoProductGradepriceGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        操作是否成功
    */
    IsSuccess  bool `json:"is_success,omitempty" `
    /*
        等级折扣列表
    */
    GradeDiscounts  []domain.TaobaoFenxiaoProductGradepriceGetGradeDiscount `json:"grade_discounts,omitempty" `
}
