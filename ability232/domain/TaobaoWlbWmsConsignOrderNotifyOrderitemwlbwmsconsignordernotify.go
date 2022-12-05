package domain


type TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify struct {
    /*
        商品优惠金额     */
    DiscountAmount  *int64 `json:"discount_amount,omitempty" `

    /*
        销售价格     */
    ItemPrice  *int64 `json:"item_price,omitempty" `

    /*
        商品成交价格=销售价格-优惠金额     */
    ActualPrice  *int64 `json:"actual_price,omitempty" `

    /*
        订单商品拓展属性数据     */
    ExtendFields  *string `json:"extend_fields,omitempty" `

    /*
        商品数量     */
    ItemQuantity  *int64 `json:"item_quantity,omitempty" `

    /*
        库存类型     */
    InventoryType  *int64 `json:"inventory_type,omitempty" `

    /*
        交易平台商品编码     */
    ItemExtCode  *string `json:"item_ext_code,omitempty" `

    /*
        ERP商品ID     */
    ItemId  *string `json:"item_id,omitempty" `

    /*
        货主名称     */
    OwnerUserName  *string `json:"owner_user_name,omitempty" `

    /*
        货主ID     */
    OwnerUserId  *string `json:"owner_user_id,omitempty" `

    /*
        店铺名称     */
    UserName  *string `json:"user_name,omitempty" `

    /*
        店铺ID     */
    UserId  *string `json:"user_id,omitempty" `

    /*
        平台子交易编码     */
    SubSourceCode  *string `json:"sub_source_code,omitempty" `

    /*
        平台交易订单编码，如果不传入淘系交易订单，子订单系统自动标示此商品为赠品；     */
    OrderSourceCode  *string `json:"order_source_code,omitempty" `

    /*
        ERP订单明细行号ID，数字类型     */
    OrderItemId  *string `json:"order_item_id,omitempty" `

    /*
        商品名称     */
    ItemName  *string `json:"item_name,omitempty" `

    /*
        ERP店铺编码     */
    ShopCode  *string `json:"shop_code,omitempty" `

}

func (s *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify) SetDiscountAmount(v int64) *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify {
    s.DiscountAmount = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify) SetItemPrice(v int64) *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify {
    s.ItemPrice = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify) SetActualPrice(v int64) *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify {
    s.ActualPrice = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify) SetExtendFields(v string) *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify {
    s.ExtendFields = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify) SetItemQuantity(v int64) *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify {
    s.ItemQuantity = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify) SetInventoryType(v int64) *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify {
    s.InventoryType = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify) SetItemExtCode(v string) *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify {
    s.ItemExtCode = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify) SetItemId(v string) *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify {
    s.ItemId = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify) SetOwnerUserName(v string) *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify {
    s.OwnerUserName = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify) SetOwnerUserId(v string) *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify {
    s.OwnerUserId = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify) SetUserName(v string) *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify {
    s.UserName = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify) SetUserId(v string) *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify {
    s.UserId = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify) SetSubSourceCode(v string) *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify {
    s.SubSourceCode = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify) SetOrderSourceCode(v string) *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify {
    s.OrderSourceCode = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify) SetOrderItemId(v string) *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify {
    s.OrderItemId = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify) SetItemName(v string) *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify {
    s.ItemName = &v
    return s
}
func (s *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify) SetShopCode(v string) *TaobaoWlbWmsConsignOrderNotifyOrderitemwlbwmsconsignordernotify {
    s.ShopCode = &v
    return s
}
