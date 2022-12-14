package domain

import (
        "topsdk/util"
    )

type TaobaoWlbNotifyMessagePageGetWlbMessage struct {
    /*
        创建时间     */
    GmtCreate  *util.LocalTime `json:"gmt_create,omitempty" `

    /*
        消息描述     */
    MsgDescription  *string `json:"msg_description,omitempty" `

    /*
        通知内容：msg_code为STOCK_IN_NOT_CONSISTENT时,msg_content为:orderCode:code;orderItemId:111;itemId:123;planQuantity:10;relQuantity:6; 
msg_code为CANCEL_ORDER_SUCCESS及其它时,msg_content为：orderCode:code;
msg_code为CANCEL_ORDER_SUCCESS及其它时,msg_content为：orderCode:code; msg_code为INVENTORY_CHECK时，msg_content为orderCode:code;     */
    MsgContent  *string `json:"msg_content,omitempty" `

    /*
        通知消息编码： STOCK_IN_NOT_CONSISTENT---入库单不一致 CANCEL_ORDER_SUCCESS---取消订单成功 CANCEL_ORDER_FAILURE---取消订单失败 INVENTORY_CHECK---盘点   INVENTORY_CHECK---盘点消息  ORDER_REJECT--wms拒单  ORDER_CONFIRMED--订单处理成功 WMS_FAILED--wms处理失败     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        消息状态： 不需要确认：NO_NEED_CONFIRM 已确认：CONFIRMED 待确认：TO_BE_CONFIRM     */
    Status  *string `json:"status,omitempty" `

    /*
        用户ID     */
    UserId  *int64 `json:"user_id,omitempty" `

    /*
        消息通道ID     */
    Id  *int64 `json:"id,omitempty" `

}

func (s *TaobaoWlbNotifyMessagePageGetWlbMessage) SetGmtCreate(v util.LocalTime) *TaobaoWlbNotifyMessagePageGetWlbMessage {
    s.GmtCreate = &v
    return s
}
func (s *TaobaoWlbNotifyMessagePageGetWlbMessage) SetMsgDescription(v string) *TaobaoWlbNotifyMessagePageGetWlbMessage {
    s.MsgDescription = &v
    return s
}
func (s *TaobaoWlbNotifyMessagePageGetWlbMessage) SetMsgContent(v string) *TaobaoWlbNotifyMessagePageGetWlbMessage {
    s.MsgContent = &v
    return s
}
func (s *TaobaoWlbNotifyMessagePageGetWlbMessage) SetMsgCode(v string) *TaobaoWlbNotifyMessagePageGetWlbMessage {
    s.MsgCode = &v
    return s
}
func (s *TaobaoWlbNotifyMessagePageGetWlbMessage) SetStatus(v string) *TaobaoWlbNotifyMessagePageGetWlbMessage {
    s.Status = &v
    return s
}
func (s *TaobaoWlbNotifyMessagePageGetWlbMessage) SetUserId(v int64) *TaobaoWlbNotifyMessagePageGetWlbMessage {
    s.UserId = &v
    return s
}
func (s *TaobaoWlbNotifyMessagePageGetWlbMessage) SetId(v int64) *TaobaoWlbNotifyMessagePageGetWlbMessage {
    s.Id = &v
    return s
}
