package domain


type TaobaoWlbInventoryDetailGetWlbInventory struct {
    /*
        货主ID     */
    UserId  *int64 `json:"user_id,omitempty" `

    /*
        商品ID     */
    ItemId  *int64 `json:"item_id,omitempty" `

    /*
        仓库编码，关联到仓库类型服务的编码非托管库存(卖家自己管理的库存，物流宝不可见又称自有库存)的所在仓库编码: STORE_SYS_PRIVATE     */
    StoreCode  *string `json:"store_code,omitempty" `

    /*
        可销库存数量(库存总数-拍下预扣数-占用数)     */
    Quantity  *int64 `json:"quantity,omitempty" `

    /*
        冻结(锁定)数量，用来跟踪库存的中间状态，比如前台销售了1件商品，这时lock加1，当商品出库的时候lock再减回去     */
    LockQuantity  *int64 `json:"lock_quantity,omitempty" `

    /*
        系统自动生成     */
    ReserveQuantity  *int64 `json:"reserve_quantity,omitempty" `

    /*
        系统自动生成     */
    OccupyQuantity  *int64 `json:"occupy_quantity,omitempty" `

    /*
        VENDIBLE--可销售库存
FREEZE--冻结库存
ONWAY--在途库存
DEFECT--残次品库存     */
    Type  *string `json:"type,omitempty" `

}

func (s *TaobaoWlbInventoryDetailGetWlbInventory) SetUserId(v int64) *TaobaoWlbInventoryDetailGetWlbInventory {
    s.UserId = &v
    return s
}
func (s *TaobaoWlbInventoryDetailGetWlbInventory) SetItemId(v int64) *TaobaoWlbInventoryDetailGetWlbInventory {
    s.ItemId = &v
    return s
}
func (s *TaobaoWlbInventoryDetailGetWlbInventory) SetStoreCode(v string) *TaobaoWlbInventoryDetailGetWlbInventory {
    s.StoreCode = &v
    return s
}
func (s *TaobaoWlbInventoryDetailGetWlbInventory) SetQuantity(v int64) *TaobaoWlbInventoryDetailGetWlbInventory {
    s.Quantity = &v
    return s
}
func (s *TaobaoWlbInventoryDetailGetWlbInventory) SetLockQuantity(v int64) *TaobaoWlbInventoryDetailGetWlbInventory {
    s.LockQuantity = &v
    return s
}
func (s *TaobaoWlbInventoryDetailGetWlbInventory) SetReserveQuantity(v int64) *TaobaoWlbInventoryDetailGetWlbInventory {
    s.ReserveQuantity = &v
    return s
}
func (s *TaobaoWlbInventoryDetailGetWlbInventory) SetOccupyQuantity(v int64) *TaobaoWlbInventoryDetailGetWlbInventory {
    s.OccupyQuantity = &v
    return s
}
func (s *TaobaoWlbInventoryDetailGetWlbInventory) SetType(v string) *TaobaoWlbInventoryDetailGetWlbInventory {
    s.Type = &v
    return s
}
