package domain

import (
        "topsdk/util"
    )

type TaobaoFenxiaoRequisitionsGetRequisition struct {
    /*
        好评率     */
    DistAppraise  *int64 `json:"dist_appraise,omitempty" `

    /*
        分销申请加盟时，给供应商的留言     */
    DistMessage  *string `json:"dist_message,omitempty" `

    /*
        申请时间     */
    GmtCreate  *util.LocalTime `json:"gmt_create,omitempty" `

    /*
        店铺星级     */
    DistLevel  *int64 `json:"dist_level,omitempty" `

    /*
        分销商nick     */
    DistributorNick  *string `json:"distributor_nick,omitempty" `

    /*
        开店时间     */
    DistOpenDate  *util.LocalTime `json:"dist_open_date,omitempty" `

    /*
        申请状态（1-申请中、2-成功、3-被退回、4-已撤消、5-过期）     */
    Status  *int64 `json:"status,omitempty" `

    /*
        主营类目     */
    DistCategory  *int64 `json:"dist_category,omitempty" `

    /*
        分销商id     */
    DistributorId  *int64 `json:"distributor_id,omitempty" `

    /*
        店铺地址     */
    DistShopAddress  *string `json:"dist_shop_address,omitempty" `

    /*
        合作申请ID     */
    RequisitionId  *int64 `json:"requisition_id,omitempty" `

    /*
        是否消保(0-不是、1-是)     */
    DistIsXiaobao  *int64 `json:"dist_is_xiaobao,omitempty" `

    /*
        主营类目名称     */
    DistCategoryName  *string `json:"dist_category_name,omitempty" `

}

func (s *TaobaoFenxiaoRequisitionsGetRequisition) SetDistAppraise(v int64) *TaobaoFenxiaoRequisitionsGetRequisition {
    s.DistAppraise = &v
    return s
}
func (s *TaobaoFenxiaoRequisitionsGetRequisition) SetDistMessage(v string) *TaobaoFenxiaoRequisitionsGetRequisition {
    s.DistMessage = &v
    return s
}
func (s *TaobaoFenxiaoRequisitionsGetRequisition) SetGmtCreate(v util.LocalTime) *TaobaoFenxiaoRequisitionsGetRequisition {
    s.GmtCreate = &v
    return s
}
func (s *TaobaoFenxiaoRequisitionsGetRequisition) SetDistLevel(v int64) *TaobaoFenxiaoRequisitionsGetRequisition {
    s.DistLevel = &v
    return s
}
func (s *TaobaoFenxiaoRequisitionsGetRequisition) SetDistributorNick(v string) *TaobaoFenxiaoRequisitionsGetRequisition {
    s.DistributorNick = &v
    return s
}
func (s *TaobaoFenxiaoRequisitionsGetRequisition) SetDistOpenDate(v util.LocalTime) *TaobaoFenxiaoRequisitionsGetRequisition {
    s.DistOpenDate = &v
    return s
}
func (s *TaobaoFenxiaoRequisitionsGetRequisition) SetStatus(v int64) *TaobaoFenxiaoRequisitionsGetRequisition {
    s.Status = &v
    return s
}
func (s *TaobaoFenxiaoRequisitionsGetRequisition) SetDistCategory(v int64) *TaobaoFenxiaoRequisitionsGetRequisition {
    s.DistCategory = &v
    return s
}
func (s *TaobaoFenxiaoRequisitionsGetRequisition) SetDistributorId(v int64) *TaobaoFenxiaoRequisitionsGetRequisition {
    s.DistributorId = &v
    return s
}
func (s *TaobaoFenxiaoRequisitionsGetRequisition) SetDistShopAddress(v string) *TaobaoFenxiaoRequisitionsGetRequisition {
    s.DistShopAddress = &v
    return s
}
func (s *TaobaoFenxiaoRequisitionsGetRequisition) SetRequisitionId(v int64) *TaobaoFenxiaoRequisitionsGetRequisition {
    s.RequisitionId = &v
    return s
}
func (s *TaobaoFenxiaoRequisitionsGetRequisition) SetDistIsXiaobao(v int64) *TaobaoFenxiaoRequisitionsGetRequisition {
    s.DistIsXiaobao = &v
    return s
}
func (s *TaobaoFenxiaoRequisitionsGetRequisition) SetDistCategoryName(v string) *TaobaoFenxiaoRequisitionsGetRequisition {
    s.DistCategoryName = &v
    return s
}
