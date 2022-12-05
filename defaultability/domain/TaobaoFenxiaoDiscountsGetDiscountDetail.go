package domain

import (
        "topsdk/util"
    )

type TaobaoFenxiaoDiscountsGetDiscountDetail struct {
    /*
        折扣详情ID     */
    DetailId  *int64 `json:"detail_id,omitempty" `

    /*
        折扣类型:GRADE（按会员等级优惠）、DISTRIBUTOR（按分销商优惠）     */
    TargetType  *string `json:"target_type,omitempty" `

    /*
        会员等级的id或者分销商id     */
    TargetId  *int64 `json:"target_id,omitempty" `

    /*
        等级名称或者分销商名称     */
    TargetName  *string `json:"target_name,omitempty" `

    /*
        优惠方式:PERCENT（按折扣优惠）、PRICE（按减价优惠）     */
    DiscountType  *string `json:"discount_type,omitempty" `

    /*
        优惠比率或者优惠价格 10%或10     */
    DiscountValue  *int64 `json:"discount_value,omitempty" `

    /*
        创建时间     */
    Created  *util.LocalTime `json:"created,omitempty" `

    /*
        修改时间     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

}

func (s *TaobaoFenxiaoDiscountsGetDiscountDetail) SetDetailId(v int64) *TaobaoFenxiaoDiscountsGetDiscountDetail {
    s.DetailId = &v
    return s
}
func (s *TaobaoFenxiaoDiscountsGetDiscountDetail) SetTargetType(v string) *TaobaoFenxiaoDiscountsGetDiscountDetail {
    s.TargetType = &v
    return s
}
func (s *TaobaoFenxiaoDiscountsGetDiscountDetail) SetTargetId(v int64) *TaobaoFenxiaoDiscountsGetDiscountDetail {
    s.TargetId = &v
    return s
}
func (s *TaobaoFenxiaoDiscountsGetDiscountDetail) SetTargetName(v string) *TaobaoFenxiaoDiscountsGetDiscountDetail {
    s.TargetName = &v
    return s
}
func (s *TaobaoFenxiaoDiscountsGetDiscountDetail) SetDiscountType(v string) *TaobaoFenxiaoDiscountsGetDiscountDetail {
    s.DiscountType = &v
    return s
}
func (s *TaobaoFenxiaoDiscountsGetDiscountDetail) SetDiscountValue(v int64) *TaobaoFenxiaoDiscountsGetDiscountDetail {
    s.DiscountValue = &v
    return s
}
func (s *TaobaoFenxiaoDiscountsGetDiscountDetail) SetCreated(v util.LocalTime) *TaobaoFenxiaoDiscountsGetDiscountDetail {
    s.Created = &v
    return s
}
func (s *TaobaoFenxiaoDiscountsGetDiscountDetail) SetModified(v util.LocalTime) *TaobaoFenxiaoDiscountsGetDiscountDetail {
    s.Modified = &v
    return s
}
