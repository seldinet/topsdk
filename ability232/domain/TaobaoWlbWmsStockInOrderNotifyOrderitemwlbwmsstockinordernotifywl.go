package domain

import (
        "topsdk/util"
    )

type TaobaoWlbWmsStockInOrderNotifyOrderitemwlbwmsstockinordernotifywl struct {
    /*
        订单商品拓展属性     */
    ExtendFields  *string `json:"extend_fields,omitempty" `

    /*
        生产批号     */
    ProduceCode  *string `json:"produce_code,omitempty" `

    /*
        失效日期     */
    DueDate  *util.LocalTime `json:"due_date,omitempty" `

    /*
        生产日期     */
    ProduceDate  *util.LocalTime `json:"produce_date,omitempty" `

    /*
        批次编号     */
    BatchCode  *string `json:"batch_code,omitempty" `

    /*
        采购价格     */
    PurchasePrice  *int64 `json:"purchase_price,omitempty" `

    /*
        商品数量     */
    ItemQuantity  *int64 `json:"item_quantity,omitempty" `

    /*
        库存类型 1 正品 101 残次 102 机损 103 箱损 201 冻结库存 301 在途库存     */
    InventoryType  *int64 `json:"inventory_type,omitempty" `

    /*
        后端商品ID     */
    ItemId  *string `json:"item_id,omitempty" `

    /*
        ERP单据明细ID     */
    OrderItemId  *string `json:"order_item_id,omitempty" `

}

func (s *TaobaoWlbWmsStockInOrderNotifyOrderitemwlbwmsstockinordernotifywl) SetExtendFields(v string) *TaobaoWlbWmsStockInOrderNotifyOrderitemwlbwmsstockinordernotifywl {
    s.ExtendFields = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyOrderitemwlbwmsstockinordernotifywl) SetProduceCode(v string) *TaobaoWlbWmsStockInOrderNotifyOrderitemwlbwmsstockinordernotifywl {
    s.ProduceCode = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyOrderitemwlbwmsstockinordernotifywl) SetDueDate(v util.LocalTime) *TaobaoWlbWmsStockInOrderNotifyOrderitemwlbwmsstockinordernotifywl {
    s.DueDate = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyOrderitemwlbwmsstockinordernotifywl) SetProduceDate(v util.LocalTime) *TaobaoWlbWmsStockInOrderNotifyOrderitemwlbwmsstockinordernotifywl {
    s.ProduceDate = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyOrderitemwlbwmsstockinordernotifywl) SetBatchCode(v string) *TaobaoWlbWmsStockInOrderNotifyOrderitemwlbwmsstockinordernotifywl {
    s.BatchCode = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyOrderitemwlbwmsstockinordernotifywl) SetPurchasePrice(v int64) *TaobaoWlbWmsStockInOrderNotifyOrderitemwlbwmsstockinordernotifywl {
    s.PurchasePrice = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyOrderitemwlbwmsstockinordernotifywl) SetItemQuantity(v int64) *TaobaoWlbWmsStockInOrderNotifyOrderitemwlbwmsstockinordernotifywl {
    s.ItemQuantity = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyOrderitemwlbwmsstockinordernotifywl) SetInventoryType(v int64) *TaobaoWlbWmsStockInOrderNotifyOrderitemwlbwmsstockinordernotifywl {
    s.InventoryType = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyOrderitemwlbwmsstockinordernotifywl) SetItemId(v string) *TaobaoWlbWmsStockInOrderNotifyOrderitemwlbwmsstockinordernotifywl {
    s.ItemId = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyOrderitemwlbwmsstockinordernotifywl) SetOrderItemId(v string) *TaobaoWlbWmsStockInOrderNotifyOrderitemwlbwmsstockinordernotifywl {
    s.OrderItemId = &v
    return s
}
