package domain

import (
        "topsdk/util"
    )

type TaobaoWlbWmsInventoryQueryWmsInventoryQueryItem struct {
    /*
        锁库存数量     */
    LockQuantity  *int64 `json:"lock_quantity,omitempty" `

    /*
        库存数量     */
    Quantity  *int64 `json:"quantity,omitempty" `

    /*
        失效日期，type=2时字段有返回值。     */
    DueDate  *util.LocalTime `json:"due_date,omitempty" `

    /*
        生产日期，type=2时字段有返回值     */
    ProduceDate  *util.LocalTime `json:"produce_date,omitempty" `

    /*
        库存批次号，type=2时字段有返回值。     */
    BatchCode  *string `json:"batch_code,omitempty" `

    /*
        渠道编码，type=3时字段有返回值。（TB 淘系，OTHERS 其他）     */
    ChannelCode  *string `json:"channel_code,omitempty" `

    /*
        库存类型(1 正品 101 残次 102 机损 103 箱损 201 冻结库存 301 在途库存 )     */
    InventoryType  *int64 `json:"inventory_type,omitempty" `

    /*
        商品ID     */
    ItemId  *string `json:"item_id,omitempty" `

    /*
        仓库编码     */
    StoreCode  *string `json:"store_code,omitempty" `

}

func (s *TaobaoWlbWmsInventoryQueryWmsInventoryQueryItem) SetLockQuantity(v int64) *TaobaoWlbWmsInventoryQueryWmsInventoryQueryItem {
    s.LockQuantity = &v
    return s
}
func (s *TaobaoWlbWmsInventoryQueryWmsInventoryQueryItem) SetQuantity(v int64) *TaobaoWlbWmsInventoryQueryWmsInventoryQueryItem {
    s.Quantity = &v
    return s
}
func (s *TaobaoWlbWmsInventoryQueryWmsInventoryQueryItem) SetDueDate(v util.LocalTime) *TaobaoWlbWmsInventoryQueryWmsInventoryQueryItem {
    s.DueDate = &v
    return s
}
func (s *TaobaoWlbWmsInventoryQueryWmsInventoryQueryItem) SetProduceDate(v util.LocalTime) *TaobaoWlbWmsInventoryQueryWmsInventoryQueryItem {
    s.ProduceDate = &v
    return s
}
func (s *TaobaoWlbWmsInventoryQueryWmsInventoryQueryItem) SetBatchCode(v string) *TaobaoWlbWmsInventoryQueryWmsInventoryQueryItem {
    s.BatchCode = &v
    return s
}
func (s *TaobaoWlbWmsInventoryQueryWmsInventoryQueryItem) SetChannelCode(v string) *TaobaoWlbWmsInventoryQueryWmsInventoryQueryItem {
    s.ChannelCode = &v
    return s
}
func (s *TaobaoWlbWmsInventoryQueryWmsInventoryQueryItem) SetInventoryType(v int64) *TaobaoWlbWmsInventoryQueryWmsInventoryQueryItem {
    s.InventoryType = &v
    return s
}
func (s *TaobaoWlbWmsInventoryQueryWmsInventoryQueryItem) SetItemId(v string) *TaobaoWlbWmsInventoryQueryWmsInventoryQueryItem {
    s.ItemId = &v
    return s
}
func (s *TaobaoWlbWmsInventoryQueryWmsInventoryQueryItem) SetStoreCode(v string) *TaobaoWlbWmsInventoryQueryWmsInventoryQueryItem {
    s.StoreCode = &v
    return s
}
