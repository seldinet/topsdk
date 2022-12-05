package domain

import (
        "topsdk/util"
    )

type TaobaoFenxiaoDiscountsGetDiscount struct {
    /*
        折扣ID     */
    DiscountId  *int64 `json:"discount_id,omitempty" `

    /*
        折扣名称     */
    Name  *string `json:"name,omitempty" `

    /*
        折扣详情     */
    Details  *[]TaobaoFenxiaoDiscountsGetDiscountDetail `json:"details,omitempty" `

    /*
        创建时间     */
    Created  *util.LocalTime `json:"created,omitempty" `

    /*
        修改时间     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

}

func (s *TaobaoFenxiaoDiscountsGetDiscount) SetDiscountId(v int64) *TaobaoFenxiaoDiscountsGetDiscount {
    s.DiscountId = &v
    return s
}
func (s *TaobaoFenxiaoDiscountsGetDiscount) SetName(v string) *TaobaoFenxiaoDiscountsGetDiscount {
    s.Name = &v
    return s
}
func (s *TaobaoFenxiaoDiscountsGetDiscount) SetDetails(v []TaobaoFenxiaoDiscountsGetDiscountDetail) *TaobaoFenxiaoDiscountsGetDiscount {
    s.Details = &v
    return s
}
func (s *TaobaoFenxiaoDiscountsGetDiscount) SetCreated(v util.LocalTime) *TaobaoFenxiaoDiscountsGetDiscount {
    s.Created = &v
    return s
}
func (s *TaobaoFenxiaoDiscountsGetDiscount) SetModified(v util.LocalTime) *TaobaoFenxiaoDiscountsGetDiscount {
    s.Modified = &v
    return s
}
