package domain


type TaobaoWlbOrderitemPageGetWlbOrderItem struct {
    /*
        订单商品备注     */
    Remark  *string `json:"remark,omitempty" `

    /*
        订单商品id     */
    Id  *int64 `json:"id,omitempty" `

    /*
        订单商品用户id     */
    UserId  *int64 `json:"user_id,omitempty" `

    /*
        订单商品用户昵称     */
    UserNick  *string `json:"user_nick,omitempty" `

    /*
        INVENTORY_TYPE_SELL 可销库存
INVENTORY_TYPE_IMPERFECTIONS 残次库存
INVENTORY_TYPE_FREEZE 冻结库存
INVENTORY_TYPE_ON_PASSAGE 在途库存
INVENTORY_TYPE_ENGINE_DAMAGE 机损
INVENTORY_TYPE_BOX_DAMAGE 箱损     */
    InventoryType  *string `json:"inventory_type,omitempty" `

    /*
        物流宝订单id     */
    OrderId  *int64 `json:"order_id,omitempty" `

    /*
        物流宝订单编码     */
    OrderCode  *string `json:"order_code,omitempty" `

    /*
        (1)其它: OTHER (2)淘宝交易: TAOBAO (3)调拨: ALLOCATION (4)盘点:CHECK (5)销售采购:PRUCHASE(6)其它交易：OTHER_TRADE     */
    OrderSubType  *string `json:"order_sub_type,omitempty" `

    /*
        订单号     */
    OrderSubCode  *string `json:"order_sub_code,omitempty" `

    /*
        子交易号     */
    OrderSub2code  *string `json:"order_sub_2code,omitempty" `

    /*
        物流宝订单商品的物流宝商品id     */
    ItemId  *int64 `json:"item_id,omitempty" `

    /*
        订单商品名称     */
    ItemName  *string `json:"item_name,omitempty" `

    /*
        订单商品编码     */
    ItemCode  *string `json:"item_code,omitempty" `

    /*
        货主id     */
    ProviderTpId  *int64 `json:"provider_tp_id,omitempty" `

    /*
        货主nick     */
    ProviderTpNick  *string `json:"provider_tp_nick,omitempty" `

    /*
        计划数量     */
    PlanQuantity  *int64 `json:"plan_quantity,omitempty" `

    /*
        实际数量     */
    RealQuantity  *int64 `json:"real_quantity,omitempty" `

    /*
        商品价格     */
    ItemPrice  *int64 `json:"item_price,omitempty" `

    /*
        物流宝订单确认状态：
NO_NEED_CONFIRM--不需要确认
WAIT_CONFIRM--待确认
CONFIRMED--已确认     */
    ConfirmStatus  *string `json:"confirm_status,omitempty" `

    /*
        商品发布版本号     */
    PublishVersion  *int64 `json:"publish_version,omitempty" `

    /*
        批次号备注     */
    BatchRemark  *string `json:"batch_remark,omitempty" `

    /*
        订单商品图片url     */
    PictureUrl  *string `json:"picture_url,omitempty" `

    /*
        外部实体id     */
    ExtEntityId  *string `json:"ext_entity_id,omitempty" `

}

func (s *TaobaoWlbOrderitemPageGetWlbOrderItem) SetRemark(v string) *TaobaoWlbOrderitemPageGetWlbOrderItem {
    s.Remark = &v
    return s
}
func (s *TaobaoWlbOrderitemPageGetWlbOrderItem) SetId(v int64) *TaobaoWlbOrderitemPageGetWlbOrderItem {
    s.Id = &v
    return s
}
func (s *TaobaoWlbOrderitemPageGetWlbOrderItem) SetUserId(v int64) *TaobaoWlbOrderitemPageGetWlbOrderItem {
    s.UserId = &v
    return s
}
func (s *TaobaoWlbOrderitemPageGetWlbOrderItem) SetUserNick(v string) *TaobaoWlbOrderitemPageGetWlbOrderItem {
    s.UserNick = &v
    return s
}
func (s *TaobaoWlbOrderitemPageGetWlbOrderItem) SetInventoryType(v string) *TaobaoWlbOrderitemPageGetWlbOrderItem {
    s.InventoryType = &v
    return s
}
func (s *TaobaoWlbOrderitemPageGetWlbOrderItem) SetOrderId(v int64) *TaobaoWlbOrderitemPageGetWlbOrderItem {
    s.OrderId = &v
    return s
}
func (s *TaobaoWlbOrderitemPageGetWlbOrderItem) SetOrderCode(v string) *TaobaoWlbOrderitemPageGetWlbOrderItem {
    s.OrderCode = &v
    return s
}
func (s *TaobaoWlbOrderitemPageGetWlbOrderItem) SetOrderSubType(v string) *TaobaoWlbOrderitemPageGetWlbOrderItem {
    s.OrderSubType = &v
    return s
}
func (s *TaobaoWlbOrderitemPageGetWlbOrderItem) SetOrderSubCode(v string) *TaobaoWlbOrderitemPageGetWlbOrderItem {
    s.OrderSubCode = &v
    return s
}
func (s *TaobaoWlbOrderitemPageGetWlbOrderItem) SetOrderSub2code(v string) *TaobaoWlbOrderitemPageGetWlbOrderItem {
    s.OrderSub2code = &v
    return s
}
func (s *TaobaoWlbOrderitemPageGetWlbOrderItem) SetItemId(v int64) *TaobaoWlbOrderitemPageGetWlbOrderItem {
    s.ItemId = &v
    return s
}
func (s *TaobaoWlbOrderitemPageGetWlbOrderItem) SetItemName(v string) *TaobaoWlbOrderitemPageGetWlbOrderItem {
    s.ItemName = &v
    return s
}
func (s *TaobaoWlbOrderitemPageGetWlbOrderItem) SetItemCode(v string) *TaobaoWlbOrderitemPageGetWlbOrderItem {
    s.ItemCode = &v
    return s
}
func (s *TaobaoWlbOrderitemPageGetWlbOrderItem) SetProviderTpId(v int64) *TaobaoWlbOrderitemPageGetWlbOrderItem {
    s.ProviderTpId = &v
    return s
}
func (s *TaobaoWlbOrderitemPageGetWlbOrderItem) SetProviderTpNick(v string) *TaobaoWlbOrderitemPageGetWlbOrderItem {
    s.ProviderTpNick = &v
    return s
}
func (s *TaobaoWlbOrderitemPageGetWlbOrderItem) SetPlanQuantity(v int64) *TaobaoWlbOrderitemPageGetWlbOrderItem {
    s.PlanQuantity = &v
    return s
}
func (s *TaobaoWlbOrderitemPageGetWlbOrderItem) SetRealQuantity(v int64) *TaobaoWlbOrderitemPageGetWlbOrderItem {
    s.RealQuantity = &v
    return s
}
func (s *TaobaoWlbOrderitemPageGetWlbOrderItem) SetItemPrice(v int64) *TaobaoWlbOrderitemPageGetWlbOrderItem {
    s.ItemPrice = &v
    return s
}
func (s *TaobaoWlbOrderitemPageGetWlbOrderItem) SetConfirmStatus(v string) *TaobaoWlbOrderitemPageGetWlbOrderItem {
    s.ConfirmStatus = &v
    return s
}
func (s *TaobaoWlbOrderitemPageGetWlbOrderItem) SetPublishVersion(v int64) *TaobaoWlbOrderitemPageGetWlbOrderItem {
    s.PublishVersion = &v
    return s
}
func (s *TaobaoWlbOrderitemPageGetWlbOrderItem) SetBatchRemark(v string) *TaobaoWlbOrderitemPageGetWlbOrderItem {
    s.BatchRemark = &v
    return s
}
func (s *TaobaoWlbOrderitemPageGetWlbOrderItem) SetPictureUrl(v string) *TaobaoWlbOrderitemPageGetWlbOrderItem {
    s.PictureUrl = &v
    return s
}
func (s *TaobaoWlbOrderitemPageGetWlbOrderItem) SetExtEntityId(v string) *TaobaoWlbOrderitemPageGetWlbOrderItem {
    s.ExtEntityId = &v
    return s
}
