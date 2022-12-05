package domain

import (
        "topsdk/util"
    )

type TaobaoWlbOrderstatusGetWlbProcessStatus struct {
    /*
        物流宝订单编码     */
    OrderCode  *string `json:"order_code,omitempty" `

    /*
        操作人     */
    Operator  *string `json:"operator,omitempty" `

    /*
        操作时间     */
    OperateTime  *util.LocalTime `json:"operate_time,omitempty" `

    /*
        仓库合作公司编码     */
    StoreTpCode  *string `json:"store_tp_code,omitempty" `

    /*
        仓库编码     */
    StoreCode  *string `json:"store_code,omitempty" `

    /*
        tms合作公司编码     */
    TmsTpCode  *string `json:"tms_tp_code,omitempty" `

    /*
        tms合作公司订单编码     */
    TmsOrderCode  *string `json:"tms_order_code,omitempty" `

    /*
        状态内容     */
    Content  *string `json:"content,omitempty" `

    /*
        备注     */
    Remark  *string `json:"remark,omitempty" `

    /*
        订单操作状态：WMS_ACCEPT;WMS_PRINT;WMS_PICK;WMS_CHECK;WMS_PACKAGE;WMS_CONSIGN;WMS_CANCEL;WMS_UNKNOWN;WMS_CONFIRMEDTMS_ACCEPT;TMS_STATION_IN;TMS_STATION_OUT;TMS_SIGN;TMS_REJECT;TMS_CANCEL;TMS_UNKNOW;SYS_UNKNOWN     */
    StatusCode  *string `json:"status_code,omitempty" `

}

func (s *TaobaoWlbOrderstatusGetWlbProcessStatus) SetOrderCode(v string) *TaobaoWlbOrderstatusGetWlbProcessStatus {
    s.OrderCode = &v
    return s
}
func (s *TaobaoWlbOrderstatusGetWlbProcessStatus) SetOperator(v string) *TaobaoWlbOrderstatusGetWlbProcessStatus {
    s.Operator = &v
    return s
}
func (s *TaobaoWlbOrderstatusGetWlbProcessStatus) SetOperateTime(v util.LocalTime) *TaobaoWlbOrderstatusGetWlbProcessStatus {
    s.OperateTime = &v
    return s
}
func (s *TaobaoWlbOrderstatusGetWlbProcessStatus) SetStoreTpCode(v string) *TaobaoWlbOrderstatusGetWlbProcessStatus {
    s.StoreTpCode = &v
    return s
}
func (s *TaobaoWlbOrderstatusGetWlbProcessStatus) SetStoreCode(v string) *TaobaoWlbOrderstatusGetWlbProcessStatus {
    s.StoreCode = &v
    return s
}
func (s *TaobaoWlbOrderstatusGetWlbProcessStatus) SetTmsTpCode(v string) *TaobaoWlbOrderstatusGetWlbProcessStatus {
    s.TmsTpCode = &v
    return s
}
func (s *TaobaoWlbOrderstatusGetWlbProcessStatus) SetTmsOrderCode(v string) *TaobaoWlbOrderstatusGetWlbProcessStatus {
    s.TmsOrderCode = &v
    return s
}
func (s *TaobaoWlbOrderstatusGetWlbProcessStatus) SetContent(v string) *TaobaoWlbOrderstatusGetWlbProcessStatus {
    s.Content = &v
    return s
}
func (s *TaobaoWlbOrderstatusGetWlbProcessStatus) SetRemark(v string) *TaobaoWlbOrderstatusGetWlbProcessStatus {
    s.Remark = &v
    return s
}
func (s *TaobaoWlbOrderstatusGetWlbProcessStatus) SetStatusCode(v string) *TaobaoWlbOrderstatusGetWlbProcessStatus {
    s.StatusCode = &v
    return s
}
