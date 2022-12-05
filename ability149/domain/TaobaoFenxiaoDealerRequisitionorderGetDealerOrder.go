package domain

import (
        "topsdk/util"
    )

type TaobaoFenxiaoDealerRequisitionorderGetDealerOrder struct {
    /*
        物流方式：SELF_PICKUP（自提）、LOGISTICS（物流发货)     */
    LogisticsType  *string `json:"logistics_type,omitempty" `

    /*
        总采购数量     */
    QuantityCount  *int64 `json:"quantity_count,omitempty" `

    /*
        收货人信息     */
    Receiver  *TaobaoFenxiaoDealerRequisitionorderGetReceiver `json:"receiver,omitempty" `

    /*
        物流费用(精确到2位小数;单位:元。如:200.07，表示:200元7分 )     */
    LogisticsFee  *string `json:"logistics_fee,omitempty" `

    /*
        折让费用(精确到2位小数;单位:元。如:200.07，表示:200元7分 )     */
    RebateFee  *string `json:"rebate_fee,omitempty" `

    /*
        修改时间     */
    ModifiedTime  *util.LocalTime `json:"modified_time,omitempty" `

    /*
        申请时间     */
    AppliedTime  *util.LocalTime `json:"applied_time,omitempty" `

    /*
        分销商最后一次确认/申请/拒绝的时间     */
    AuditTimeApplier  *util.LocalTime `json:"audit_time_applier,omitempty" `

    /*
        采购总价(精确到2位小数;单位:元。如:200.07，表示:200元7分 )     */
    TotalPrice  *string `json:"total_price,omitempty" `

    /*
        WAIT_FOR_SUPPLIER_AUDIT1：分销商提交申请，待供应商审核；SUPPLIER_REFUSE：供应商驳回申请，待分销商确认；WAIT_FOR_APPLIER_AUDIT：供应商修改后，待分销商确认；WAIT_FOR_SUPPLIER_AUDIT2：分销商拒绝修改，待供应商再审核；BOTH_AGREE_WAIT_PAY：审核通过下单成功，待分销商付款WAIT_FOR_SUPPLIER_DELIVER：付款成功，待供应商发货；WAIT_FOR_APPLIER_STORAGE：供应商发货，待分销商收货；TRADE_FINISHED：分销商收货，交易成功；TRADE_CLOSED：经销采购单关闭。     */
    OrderStatus  *string `json:"order_status,omitempty" `

    /*
        关闭原因     */
    CloseReason  *string `json:"close_reason,omitempty" `

    /*
        已发货数量     */
    DeliveredQuantityCount  *int64 `json:"delivered_quantity_count,omitempty" `

    /*
        支付方式：ALIPAY_SURETY（支付宝担保交易）TRANSFER（线下转账）PREPAY（预存款）IMMEDIATELY（即时到账）     */
    PayType  *string `json:"pay_type,omitempty" `

    /*
        供应商nick     */
    SupplierNick  *string `json:"supplier_nick,omitempty" `

    /*
        经销采购单编号，API发货使用此字段     */
    DealerOrderId  *int64 `json:"dealer_order_id,omitempty" `

    /*
        供应商驳回申请的原因     */
    RefuseReasonSupplier  *string `json:"refuse_reason_supplier,omitempty" `

    /*
        分销商拒绝供应商修改的原因     */
    RefuseReasonApplier  *string `json:"refuse_reason_applier,omitempty" `

    /*
        产品明细     */
    DealerOrderDetails  *[]TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail `json:"dealer_order_details,omitempty" `

    /*
        分销商nick     */
    ApplierNick  *string `json:"applier_nick,omitempty" `

    /*
        供应商最后一次审核通过/修改/驳回的时间     */
    AuditTimeSupplier  *util.LocalTime `json:"audit_time_supplier,omitempty" `

    /*
        支付宝交易号     */
    AlipayNo  *string `json:"alipay_no,omitempty" `

    /*
        付款时间     */
    PayTime  *util.LocalTime `json:"pay_time,omitempty" `

    /*
        供应商备注。仅供应商可见。     */
    SupplierMemo  *string `json:"supplier_memo,omitempty" `

    /*
        供应商备注旗帜。1:红色 2:黄色 3:绿色 4:蓝色 5:粉红色。仅供应商可见。     */
    SupplierMemoFlag  *int64 `json:"supplier_memo_flag,omitempty" `

    /*
        属性列表，key-value形式。     */
    Features  *[]TaobaoFenxiaoDealerRequisitionorderGetFeature `json:"features,omitempty" `

    /*
        属性键     */
    DistMemo  *string `json:"dist_memo,omitempty" `

}

func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder) SetLogisticsType(v string) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder {
    s.LogisticsType = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder) SetQuantityCount(v int64) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder {
    s.QuantityCount = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder) SetReceiver(v TaobaoFenxiaoDealerRequisitionorderGetReceiver) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder {
    s.Receiver = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder) SetLogisticsFee(v string) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder {
    s.LogisticsFee = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder) SetRebateFee(v string) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder {
    s.RebateFee = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder) SetModifiedTime(v util.LocalTime) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder {
    s.ModifiedTime = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder) SetAppliedTime(v util.LocalTime) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder {
    s.AppliedTime = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder) SetAuditTimeApplier(v util.LocalTime) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder {
    s.AuditTimeApplier = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder) SetTotalPrice(v string) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder {
    s.TotalPrice = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder) SetOrderStatus(v string) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder {
    s.OrderStatus = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder) SetCloseReason(v string) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder {
    s.CloseReason = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder) SetDeliveredQuantityCount(v int64) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder {
    s.DeliveredQuantityCount = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder) SetPayType(v string) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder {
    s.PayType = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder) SetSupplierNick(v string) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder {
    s.SupplierNick = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder) SetDealerOrderId(v int64) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder {
    s.DealerOrderId = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder) SetRefuseReasonSupplier(v string) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder {
    s.RefuseReasonSupplier = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder) SetRefuseReasonApplier(v string) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder {
    s.RefuseReasonApplier = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder) SetDealerOrderDetails(v []TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder {
    s.DealerOrderDetails = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder) SetApplierNick(v string) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder {
    s.ApplierNick = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder) SetAuditTimeSupplier(v util.LocalTime) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder {
    s.AuditTimeSupplier = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder) SetAlipayNo(v string) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder {
    s.AlipayNo = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder) SetPayTime(v util.LocalTime) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder {
    s.PayTime = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder) SetSupplierMemo(v string) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder {
    s.SupplierMemo = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder) SetSupplierMemoFlag(v int64) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder {
    s.SupplierMemoFlag = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder) SetFeatures(v []TaobaoFenxiaoDealerRequisitionorderGetFeature) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder {
    s.Features = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder) SetDistMemo(v string) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrder {
    s.DistMemo = &v
    return s
}
