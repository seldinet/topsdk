package domain

import (
        "topsdk/util"
    )

type TaobaoWlbOrderdetailDateGetWlbOrderDetail struct {
    /*
        出库或者入库，IN表示入库，OUT表示出库     */
    OperateType  *string `json:"operate_type,omitempty" `

    /*
        订单编码     */
    OrderCode  *string `json:"order_code,omitempty" `

    /*
        订单来源:
产生物流订单的原因，比如:

订单来源:1:TAOBAO;2:EXT;3:ERP;4:WMS     */
    OrderSource  *string `json:"order_source,omitempty" `

    /*
        对应创建物流宝订单top接口中的的out_biz_code字段，主要是用来去重用     */
    OrderSourceCode  *string `json:"order_source_code,omitempty" `

    /*
        1:正常订单: NARMAL
2:退货订单: RETURN
3:换货订单: CHANGE     */
    OrderType  *string `json:"order_type,omitempty" `

    /*
        (1)其它:    OTHER
(2)淘宝交易: TAOBAO
(3)301:调拨: ALLOCATION
(4)401:盘点:CHECK
(5)501:销售采购:PRUCHASE     */
    OrderSubType  *string `json:"order_sub_type,omitempty" `

    /*
        卖家ID     */
    UserId  *int64 `json:"user_id,omitempty" `

    /*
        卖家NICK     */
    UserNick  *string `json:"user_nick,omitempty" `

    /*
        仓库编码     */
    StoreCode  *string `json:"store_code,omitempty" `

    /*
        物流状态，
订单已创建：0
订单已取消: -1
订单关闭:-3
下发中: 10
已下发: 11
接收方拒签 :-100
已发货:100
签收成功:200     */
    OrderStatus  *string `json:"order_status,omitempty" `

    /*
        订单备注     */
    Remark  *string `json:"remark,omitempty" `

    /*
        如果是交易单，则显示交易中买家昵称     */
    BuyerNick  *string `json:"buyer_nick,omitempty" `

    /*
        订单创建时间     */
    ModifyTime  *util.LocalTime `json:"modify_time,omitempty" `

    /*
        订单最后一次修改时间     */
    CreateTime  *util.LocalTime `json:"create_time,omitempty" `

    /*
        物流宝订单商品     */
    OrderItemList  *[]TaobaoWlbOrderdetailDateGetWlbOrderItem `json:"order_item_list,omitempty" `

    /*
        买家openId     */
    OpenUid  *string `json:"open_uid,omitempty" `

}

func (s *TaobaoWlbOrderdetailDateGetWlbOrderDetail) SetOperateType(v string) *TaobaoWlbOrderdetailDateGetWlbOrderDetail {
    s.OperateType = &v
    return s
}
func (s *TaobaoWlbOrderdetailDateGetWlbOrderDetail) SetOrderCode(v string) *TaobaoWlbOrderdetailDateGetWlbOrderDetail {
    s.OrderCode = &v
    return s
}
func (s *TaobaoWlbOrderdetailDateGetWlbOrderDetail) SetOrderSource(v string) *TaobaoWlbOrderdetailDateGetWlbOrderDetail {
    s.OrderSource = &v
    return s
}
func (s *TaobaoWlbOrderdetailDateGetWlbOrderDetail) SetOrderSourceCode(v string) *TaobaoWlbOrderdetailDateGetWlbOrderDetail {
    s.OrderSourceCode = &v
    return s
}
func (s *TaobaoWlbOrderdetailDateGetWlbOrderDetail) SetOrderType(v string) *TaobaoWlbOrderdetailDateGetWlbOrderDetail {
    s.OrderType = &v
    return s
}
func (s *TaobaoWlbOrderdetailDateGetWlbOrderDetail) SetOrderSubType(v string) *TaobaoWlbOrderdetailDateGetWlbOrderDetail {
    s.OrderSubType = &v
    return s
}
func (s *TaobaoWlbOrderdetailDateGetWlbOrderDetail) SetUserId(v int64) *TaobaoWlbOrderdetailDateGetWlbOrderDetail {
    s.UserId = &v
    return s
}
func (s *TaobaoWlbOrderdetailDateGetWlbOrderDetail) SetUserNick(v string) *TaobaoWlbOrderdetailDateGetWlbOrderDetail {
    s.UserNick = &v
    return s
}
func (s *TaobaoWlbOrderdetailDateGetWlbOrderDetail) SetStoreCode(v string) *TaobaoWlbOrderdetailDateGetWlbOrderDetail {
    s.StoreCode = &v
    return s
}
func (s *TaobaoWlbOrderdetailDateGetWlbOrderDetail) SetOrderStatus(v string) *TaobaoWlbOrderdetailDateGetWlbOrderDetail {
    s.OrderStatus = &v
    return s
}
func (s *TaobaoWlbOrderdetailDateGetWlbOrderDetail) SetRemark(v string) *TaobaoWlbOrderdetailDateGetWlbOrderDetail {
    s.Remark = &v
    return s
}
func (s *TaobaoWlbOrderdetailDateGetWlbOrderDetail) SetBuyerNick(v string) *TaobaoWlbOrderdetailDateGetWlbOrderDetail {
    s.BuyerNick = &v
    return s
}
func (s *TaobaoWlbOrderdetailDateGetWlbOrderDetail) SetModifyTime(v util.LocalTime) *TaobaoWlbOrderdetailDateGetWlbOrderDetail {
    s.ModifyTime = &v
    return s
}
func (s *TaobaoWlbOrderdetailDateGetWlbOrderDetail) SetCreateTime(v util.LocalTime) *TaobaoWlbOrderdetailDateGetWlbOrderDetail {
    s.CreateTime = &v
    return s
}
func (s *TaobaoWlbOrderdetailDateGetWlbOrderDetail) SetOrderItemList(v []TaobaoWlbOrderdetailDateGetWlbOrderItem) *TaobaoWlbOrderdetailDateGetWlbOrderDetail {
    s.OrderItemList = &v
    return s
}
func (s *TaobaoWlbOrderdetailDateGetWlbOrderDetail) SetOpenUid(v string) *TaobaoWlbOrderdetailDateGetWlbOrderDetail {
    s.OpenUid = &v
    return s
}
