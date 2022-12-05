package domain

import (
        "topsdk/util"
    )

type TaobaoFenxiaoCooperationGetCooperation struct {
    /*
        合作关系ID     */
    CooperateId  *int64 `json:"cooperate_id,omitempty" `

    /*
        分销商ID     */
    DistributorId  *int64 `json:"distributor_id,omitempty" `

    /*
        分销商nick     */
    DistributorNick  *string `json:"distributor_nick,omitempty" `

    /*
        供应商ID     */
    SupplierId  *int64 `json:"supplier_id,omitempty" `

    /*
        供应商NICK     */
    SupplierNick  *string `json:"supplier_nick,omitempty" `

    /*
        授权产品线     */
    ProductLine  *string `json:"product_line,omitempty" `

    /*
        授权产品线名称，和product_line中的值按序对应     */
    ProductLineName  *[]string `json:"product_line_name,omitempty" `

    /*
        等级ID     */
    GradeId  *int64 `json:"grade_id,omitempty" `

    /*
        分销方式： AGENT(代销) 、DEALER(经销)     */
    TradeType  *string `json:"trade_type,omitempty" `

    /*
        供应商授权的支付方式：ALIPAY(支付宝)、OFFPREPAY(预付款)、OFFTRANSFER(转帐)、OFFSETTLEMENT(后期结算)     */
    AuthPayway  *[]string `json:"auth_payway,omitempty" `

    /*
        合作起始时间     */
    StartDate  *util.LocalTime `json:"start_date,omitempty" `

    /*
        合作终止时间     */
    EndDate  *util.LocalTime `json:"end_date,omitempty" `

    /*
        合作状态： NORMAL(合作中)、 ENDING(终止中) 、END (终止)     */
    Status  *string `json:"status,omitempty" `

}

func (s *TaobaoFenxiaoCooperationGetCooperation) SetCooperateId(v int64) *TaobaoFenxiaoCooperationGetCooperation {
    s.CooperateId = &v
    return s
}
func (s *TaobaoFenxiaoCooperationGetCooperation) SetDistributorId(v int64) *TaobaoFenxiaoCooperationGetCooperation {
    s.DistributorId = &v
    return s
}
func (s *TaobaoFenxiaoCooperationGetCooperation) SetDistributorNick(v string) *TaobaoFenxiaoCooperationGetCooperation {
    s.DistributorNick = &v
    return s
}
func (s *TaobaoFenxiaoCooperationGetCooperation) SetSupplierId(v int64) *TaobaoFenxiaoCooperationGetCooperation {
    s.SupplierId = &v
    return s
}
func (s *TaobaoFenxiaoCooperationGetCooperation) SetSupplierNick(v string) *TaobaoFenxiaoCooperationGetCooperation {
    s.SupplierNick = &v
    return s
}
func (s *TaobaoFenxiaoCooperationGetCooperation) SetProductLine(v string) *TaobaoFenxiaoCooperationGetCooperation {
    s.ProductLine = &v
    return s
}
func (s *TaobaoFenxiaoCooperationGetCooperation) SetProductLineName(v []string) *TaobaoFenxiaoCooperationGetCooperation {
    s.ProductLineName = &v
    return s
}
func (s *TaobaoFenxiaoCooperationGetCooperation) SetGradeId(v int64) *TaobaoFenxiaoCooperationGetCooperation {
    s.GradeId = &v
    return s
}
func (s *TaobaoFenxiaoCooperationGetCooperation) SetTradeType(v string) *TaobaoFenxiaoCooperationGetCooperation {
    s.TradeType = &v
    return s
}
func (s *TaobaoFenxiaoCooperationGetCooperation) SetAuthPayway(v []string) *TaobaoFenxiaoCooperationGetCooperation {
    s.AuthPayway = &v
    return s
}
func (s *TaobaoFenxiaoCooperationGetCooperation) SetStartDate(v util.LocalTime) *TaobaoFenxiaoCooperationGetCooperation {
    s.StartDate = &v
    return s
}
func (s *TaobaoFenxiaoCooperationGetCooperation) SetEndDate(v util.LocalTime) *TaobaoFenxiaoCooperationGetCooperation {
    s.EndDate = &v
    return s
}
func (s *TaobaoFenxiaoCooperationGetCooperation) SetStatus(v string) *TaobaoFenxiaoCooperationGetCooperation {
    s.Status = &v
    return s
}
