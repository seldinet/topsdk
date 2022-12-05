package domain

import (
        "topsdk/util"
    )

type TaobaoWlbWmsConsignBillGetConsignsendinfo struct {
    /*
        发票确认信息列表     */
    InvoinceConfirmList  *[]TaobaoWlbWmsConsignBillGetInvoinceconfirmlist `json:"invoince_confirm_list,omitempty" `

    /*
        运单信息     */
    TmsOrderList  *[]TaobaoWlbWmsConsignBillGetTmsorderlist `json:"tms_order_list,omitempty" `

    /*
        订单信息     */
    OrderItemList  *[]TaobaoWlbWmsConsignBillGetOrderitemlist `json:"order_item_list,omitempty" `

    /*
        仓库订单完成时间     */
    ConfirmTime  *util.LocalTime `json:"confirm_time,omitempty" `

    /*
        订单状态 WMS_ACCEPT 接单成功 WMS_REJECT 接单失败 WMS_CONFIRMED 仓库生产完成     */
    Status  *string `json:"status,omitempty" `

    /*
        仓储编码     */
    StoreCode  *string `json:"store_code,omitempty" `

    /*
        订单类型 201 销售出库 502 换货出库 503 补发出库     */
    OrderType  *int64 `json:"order_type,omitempty" `

    /*
        菜鸟订单编码     */
    CnOrderCode  *string `json:"cn_order_code,omitempty" `

    /*
        ERP订单编码     */
    OrderCode  *string `json:"order_code,omitempty" `

}

func (s *TaobaoWlbWmsConsignBillGetConsignsendinfo) SetInvoinceConfirmList(v []TaobaoWlbWmsConsignBillGetInvoinceconfirmlist) *TaobaoWlbWmsConsignBillGetConsignsendinfo {
    s.InvoinceConfirmList = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetConsignsendinfo) SetTmsOrderList(v []TaobaoWlbWmsConsignBillGetTmsorderlist) *TaobaoWlbWmsConsignBillGetConsignsendinfo {
    s.TmsOrderList = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetConsignsendinfo) SetOrderItemList(v []TaobaoWlbWmsConsignBillGetOrderitemlist) *TaobaoWlbWmsConsignBillGetConsignsendinfo {
    s.OrderItemList = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetConsignsendinfo) SetConfirmTime(v util.LocalTime) *TaobaoWlbWmsConsignBillGetConsignsendinfo {
    s.ConfirmTime = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetConsignsendinfo) SetStatus(v string) *TaobaoWlbWmsConsignBillGetConsignsendinfo {
    s.Status = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetConsignsendinfo) SetStoreCode(v string) *TaobaoWlbWmsConsignBillGetConsignsendinfo {
    s.StoreCode = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetConsignsendinfo) SetOrderType(v int64) *TaobaoWlbWmsConsignBillGetConsignsendinfo {
    s.OrderType = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetConsignsendinfo) SetCnOrderCode(v string) *TaobaoWlbWmsConsignBillGetConsignsendinfo {
    s.CnOrderCode = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetConsignsendinfo) SetOrderCode(v string) *TaobaoWlbWmsConsignBillGetConsignsendinfo {
    s.OrderCode = &v
    return s
}
