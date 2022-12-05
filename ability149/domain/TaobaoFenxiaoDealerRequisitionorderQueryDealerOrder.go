package domain

import (
        "topsdk/util"
    )

type TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder struct {
    /*
        物流方式：SELF_PICKUP（自提）、LOGISTICS（物流发货)     */
    LogisticsType  *string `json:"logistics_type,omitempty" `

    /*
        总采购数量     */
    QuantityCount  *int64 `json:"quantity_count,omitempty" `

    /*
        收货人信息     */
    Receiver  *TaobaoFenxiaoDealerRequisitionorderQueryReceiver `json:"receiver,omitempty" `

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
    DealerOrderDetails  *[]TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail `json:"dealer_order_details,omitempty" `

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
        属性信息列表，key-value形式。     */
    Features  *[]TaobaoFenxiaoDealerRequisitionorderQueryFeature `json:"features,omitempty" `

    /*
        distMemo     */
    DistMemo  *string `json:"dist_memo,omitempty" `

}

func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder) SetLogisticsType(v string) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder {
    s.LogisticsType = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder) SetQuantityCount(v int64) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder {
    s.QuantityCount = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder) SetReceiver(v TaobaoFenxiaoDealerRequisitionorderQueryReceiver) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder {
    s.Receiver = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder) SetLogisticsFee(v string) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder {
    s.LogisticsFee = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder) SetRebateFee(v string) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder {
    s.RebateFee = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder) SetModifiedTime(v util.LocalTime) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder {
    s.ModifiedTime = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder) SetAppliedTime(v util.LocalTime) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder {
    s.AppliedTime = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder) SetAuditTimeApplier(v util.LocalTime) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder {
    s.AuditTimeApplier = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder) SetTotalPrice(v string) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder {
    s.TotalPrice = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder) SetOrderStatus(v string) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder {
    s.OrderStatus = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder) SetCloseReason(v string) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder {
    s.CloseReason = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder) SetDeliveredQuantityCount(v int64) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder {
    s.DeliveredQuantityCount = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder) SetPayType(v string) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder {
    s.PayType = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder) SetSupplierNick(v string) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder {
    s.SupplierNick = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder) SetDealerOrderId(v int64) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder {
    s.DealerOrderId = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder) SetRefuseReasonSupplier(v string) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder {
    s.RefuseReasonSupplier = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder) SetRefuseReasonApplier(v string) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder {
    s.RefuseReasonApplier = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder) SetDealerOrderDetails(v []TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder {
    s.DealerOrderDetails = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder) SetApplierNick(v string) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder {
    s.ApplierNick = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder) SetAuditTimeSupplier(v util.LocalTime) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder {
    s.AuditTimeSupplier = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder) SetAlipayNo(v string) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder {
    s.AlipayNo = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder) SetPayTime(v util.LocalTime) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder {
    s.PayTime = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder) SetSupplierMemo(v string) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder {
    s.SupplierMemo = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder) SetSupplierMemoFlag(v int64) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder {
    s.SupplierMemoFlag = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder) SetFeatures(v []TaobaoFenxiaoDealerRequisitionorderQueryFeature) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder {
    s.Features = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder) SetDistMemo(v string) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrder {
    s.DistMemo = &v
    return s
}
