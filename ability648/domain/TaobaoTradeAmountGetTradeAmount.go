package domain

import (
        "topsdk/util"
    )

type TaobaoTradeAmountGetTradeAmount struct {
    /*
        交易主订单编号     */
    Tid  *int64 `json:"tid,omitempty" `

    /*
        支付宝交易号，如：2009112081173831     */
    AlipayNo  *string `json:"alipay_no,omitempty" `

    /*
        交易创建时间     */
    Created  *util.LocalTime `json:"created,omitempty" `

    /*
        付款时间。格式:yyyy-MM-dd HH:mm:ss     */
    PayTime  *util.LocalTime `json:"pay_time,omitempty" `

    /*
        交易成功时间(更新交易状态为成功的同时更新)/确认收货时间。格式:yyyy-MM-dd HH:mm:ss     */
    EndTime  *util.LocalTime `json:"end_time,omitempty" `

    /*
        主订单的商品金额（各子订单中商品price * num的和，不包括任何优惠信息）。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
    TotalFee  *string `json:"total_fee,omitempty" `

    /*
        邮费。精确到2位小数;单位:元。如:200.07，表示:200元7分。目前只提供整笔交易的邮费，暂不提供各子订单的邮费     */
    PostFee  *string `json:"post_fee,omitempty" `

    /*
        货到付款服务费。精确到2位小数;单位:元。如:12.07，表示:12元7分     */
    CodFee  *string `json:"cod_fee,omitempty" `

    /*
        买家货到付款服务费。精确到2位小数;单位:元。如:12.07，表示:12元7分     */
    BuyerCodFee  *string `json:"buyer_cod_fee,omitempty" `

    /*
        卖家货到付款服务费。精确到2位小数;单位:元。如:12.07，表示:12元7分     */
    SellerCodFee  *string `json:"seller_cod_fee,omitempty" `

    /*
        快递代收款。精确到2位小数;单位:元。如:212.07，表示:212元7分     */
    ExpressAgencyFee  *string `json:"express_agency_fee,omitempty" `

    /*
        主订单实付金额。精确到2位小数;单位:元。如:200.07，表示:200元7分，计算公式如下：如果主订单只有一笔子订单 payment = 子订单的实付金额 + 货到付款服务费(如果是货到付款订单) - 主订单的系统级优惠；如果主订单有多笔子订单 payment = 各子订单的实付金额之和 + 货到付款服务费(如果是货到付款订单) + 邮费 - 主订单的系统级优惠     */
    Payment  *string `json:"payment,omitempty" `

    /*
        交易佣金。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
    CommissionFee  *string `json:"commission_fee,omitempty" `

    /*
        买家获得积分,返点的积分。格式:100;单位:个     */
    BuyerObtainPointFee  *string `json:"buyer_obtain_point_fee,omitempty" `

    /*
        子订单的帐务金额详情列表     */
    OrderAmounts  *[]TaobaoTradeAmountGetOrderAmount `json:"order_amounts,omitempty" `

    /*
        主交易订单的系统级优惠详情     */
    PromotionDetails  *[]TaobaoTradeAmountGetPromotionDetail `json:"promotion_details,omitempty" `

}

func (s *TaobaoTradeAmountGetTradeAmount) SetTid(v int64) *TaobaoTradeAmountGetTradeAmount {
    s.Tid = &v
    return s
}
func (s *TaobaoTradeAmountGetTradeAmount) SetAlipayNo(v string) *TaobaoTradeAmountGetTradeAmount {
    s.AlipayNo = &v
    return s
}
func (s *TaobaoTradeAmountGetTradeAmount) SetCreated(v util.LocalTime) *TaobaoTradeAmountGetTradeAmount {
    s.Created = &v
    return s
}
func (s *TaobaoTradeAmountGetTradeAmount) SetPayTime(v util.LocalTime) *TaobaoTradeAmountGetTradeAmount {
    s.PayTime = &v
    return s
}
func (s *TaobaoTradeAmountGetTradeAmount) SetEndTime(v util.LocalTime) *TaobaoTradeAmountGetTradeAmount {
    s.EndTime = &v
    return s
}
func (s *TaobaoTradeAmountGetTradeAmount) SetTotalFee(v string) *TaobaoTradeAmountGetTradeAmount {
    s.TotalFee = &v
    return s
}
func (s *TaobaoTradeAmountGetTradeAmount) SetPostFee(v string) *TaobaoTradeAmountGetTradeAmount {
    s.PostFee = &v
    return s
}
func (s *TaobaoTradeAmountGetTradeAmount) SetCodFee(v string) *TaobaoTradeAmountGetTradeAmount {
    s.CodFee = &v
    return s
}
func (s *TaobaoTradeAmountGetTradeAmount) SetBuyerCodFee(v string) *TaobaoTradeAmountGetTradeAmount {
    s.BuyerCodFee = &v
    return s
}
func (s *TaobaoTradeAmountGetTradeAmount) SetSellerCodFee(v string) *TaobaoTradeAmountGetTradeAmount {
    s.SellerCodFee = &v
    return s
}
func (s *TaobaoTradeAmountGetTradeAmount) SetExpressAgencyFee(v string) *TaobaoTradeAmountGetTradeAmount {
    s.ExpressAgencyFee = &v
    return s
}
func (s *TaobaoTradeAmountGetTradeAmount) SetPayment(v string) *TaobaoTradeAmountGetTradeAmount {
    s.Payment = &v
    return s
}
func (s *TaobaoTradeAmountGetTradeAmount) SetCommissionFee(v string) *TaobaoTradeAmountGetTradeAmount {
    s.CommissionFee = &v
    return s
}
func (s *TaobaoTradeAmountGetTradeAmount) SetBuyerObtainPointFee(v string) *TaobaoTradeAmountGetTradeAmount {
    s.BuyerObtainPointFee = &v
    return s
}
func (s *TaobaoTradeAmountGetTradeAmount) SetOrderAmounts(v []TaobaoTradeAmountGetOrderAmount) *TaobaoTradeAmountGetTradeAmount {
    s.OrderAmounts = &v
    return s
}
func (s *TaobaoTradeAmountGetTradeAmount) SetPromotionDetails(v []TaobaoTradeAmountGetPromotionDetail) *TaobaoTradeAmountGetTradeAmount {
    s.PromotionDetails = &v
    return s
}
