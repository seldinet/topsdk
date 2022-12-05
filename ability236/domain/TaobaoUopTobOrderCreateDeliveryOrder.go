package domain

import (
        "topsdk/util"
    )

type TaobaoUopTobOrderCreateDeliveryOrder struct {
    /*
        ERP出库单号,ERP系统内本次出库的唯一标示     */
    DeliveryOrderCode  *string `json:"delivery_order_code,omitempty" `

    /*
        交接入库单号,例如唯品会入库单号或者门店收货单号、商家仓入库单号等     */
    RelInBoundOrderCode  *string `json:"rel_in_bound_order_code,omitempty" `

    /*
        发货仓库     */
    WarehouseCode  *string `json:"warehouse_code,omitempty" `

    /*
        单据类型,出库单类型(JYCK=一般交易出库单;HHCK=换货出库单;BFCK=补发出库单;QTCK=其他出库单;TOBCK=TOB出库;BIGTOBCK=大B2B发货)     */
    OrderType  *string `json:"order_type,omitempty" `

    /*
        到货渠道类型，VIP＝1、门店＝2、经销商＝3、大润发=4、猫超=5、零售通=6、AE=7、京东=8、亚马逊=9、经销=10、代理=11、商超=12、其他=99     */
    ArriveChannelType  *string `json:"arrive_channel_type,omitempty" `

    /*
        发货单创建时间     */
    CreateTime  *util.LocalTime `json:"create_time,omitempty" `

    /*
        收货人信息     */
    ReceiverInfo  *TaobaoUopTobOrderCreateReceiverInfo `json:"receiver_info,omitempty" `

    /*
        物流公司编码     */
    LogisticsCode  *string `json:"logistics_code,omitempty" `

    /*
        物流公司名称     */
    LogisticsName  *string `json:"logistics_name,omitempty" `

    /*
        最晚到货时间     */
    LastArriveDate  *util.LocalTime `json:"last_arrive_date,omitempty" `

    /*
        订单信息     */
    OrderLine  *[]TaobaoUopTobOrderCreateOrderLine `json:"order_line,omitempty" `

    /*
        扩展信息     */
    ExtendProps  *string `json:"extend_props,omitempty" `

    /*
        收货时间区间     */
    SignTime  *string `json:"sign_time,omitempty" `

    /*
        是否自提     */
    IsSelfLifting  *string `json:"is_self_lifting,omitempty" `

    /*
        配送方式，1=整车直送、2=拼车直送、3=零担、4=快递、5=自提     */
    TransportMode  *string `json:"transport_mode,omitempty" `

}

func (s *TaobaoUopTobOrderCreateDeliveryOrder) SetDeliveryOrderCode(v string) *TaobaoUopTobOrderCreateDeliveryOrder {
    s.DeliveryOrderCode = &v
    return s
}
func (s *TaobaoUopTobOrderCreateDeliveryOrder) SetRelInBoundOrderCode(v string) *TaobaoUopTobOrderCreateDeliveryOrder {
    s.RelInBoundOrderCode = &v
    return s
}
func (s *TaobaoUopTobOrderCreateDeliveryOrder) SetWarehouseCode(v string) *TaobaoUopTobOrderCreateDeliveryOrder {
    s.WarehouseCode = &v
    return s
}
func (s *TaobaoUopTobOrderCreateDeliveryOrder) SetOrderType(v string) *TaobaoUopTobOrderCreateDeliveryOrder {
    s.OrderType = &v
    return s
}
func (s *TaobaoUopTobOrderCreateDeliveryOrder) SetArriveChannelType(v string) *TaobaoUopTobOrderCreateDeliveryOrder {
    s.ArriveChannelType = &v
    return s
}
func (s *TaobaoUopTobOrderCreateDeliveryOrder) SetCreateTime(v util.LocalTime) *TaobaoUopTobOrderCreateDeliveryOrder {
    s.CreateTime = &v
    return s
}
func (s *TaobaoUopTobOrderCreateDeliveryOrder) SetReceiverInfo(v TaobaoUopTobOrderCreateReceiverInfo) *TaobaoUopTobOrderCreateDeliveryOrder {
    s.ReceiverInfo = &v
    return s
}
func (s *TaobaoUopTobOrderCreateDeliveryOrder) SetLogisticsCode(v string) *TaobaoUopTobOrderCreateDeliveryOrder {
    s.LogisticsCode = &v
    return s
}
func (s *TaobaoUopTobOrderCreateDeliveryOrder) SetLogisticsName(v string) *TaobaoUopTobOrderCreateDeliveryOrder {
    s.LogisticsName = &v
    return s
}
func (s *TaobaoUopTobOrderCreateDeliveryOrder) SetLastArriveDate(v util.LocalTime) *TaobaoUopTobOrderCreateDeliveryOrder {
    s.LastArriveDate = &v
    return s
}
func (s *TaobaoUopTobOrderCreateDeliveryOrder) SetOrderLine(v []TaobaoUopTobOrderCreateOrderLine) *TaobaoUopTobOrderCreateDeliveryOrder {
    s.OrderLine = &v
    return s
}
func (s *TaobaoUopTobOrderCreateDeliveryOrder) SetExtendProps(v string) *TaobaoUopTobOrderCreateDeliveryOrder {
    s.ExtendProps = &v
    return s
}
func (s *TaobaoUopTobOrderCreateDeliveryOrder) SetSignTime(v string) *TaobaoUopTobOrderCreateDeliveryOrder {
    s.SignTime = &v
    return s
}
func (s *TaobaoUopTobOrderCreateDeliveryOrder) SetIsSelfLifting(v string) *TaobaoUopTobOrderCreateDeliveryOrder {
    s.IsSelfLifting = &v
    return s
}
func (s *TaobaoUopTobOrderCreateDeliveryOrder) SetTransportMode(v string) *TaobaoUopTobOrderCreateDeliveryOrder {
    s.TransportMode = &v
    return s
}
