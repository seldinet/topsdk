package domain


type TaobaoWlbItemBatchQueryWlbItemInventory struct {
    /*
        商品id     */
    ItemId  *int64 `json:"item_id,omitempty" `

    /*
        库存数量     */
    Quantity  *int64 `json:"quantity,omitempty" `

    /*
        锁定库存数量     */
    LockQuantity  *int64 `json:"lock_quantity,omitempty" `

    /*
        仓库编码     */
    StoreCode  *string `json:"store_code,omitempty" `

    /*
        SELLALBE 可销售库存
DEFECTIVE 残次
JISHUN 机损
XIANGSHUN 箱损
FREEZE 冻结库存
ONROAD 在途库存     */
    Type  *string `json:"type,omitempty" `

}

func (s *TaobaoWlbItemBatchQueryWlbItemInventory) SetItemId(v int64) *TaobaoWlbItemBatchQueryWlbItemInventory {
    s.ItemId = &v
    return s
}
func (s *TaobaoWlbItemBatchQueryWlbItemInventory) SetQuantity(v int64) *TaobaoWlbItemBatchQueryWlbItemInventory {
    s.Quantity = &v
    return s
}
func (s *TaobaoWlbItemBatchQueryWlbItemInventory) SetLockQuantity(v int64) *TaobaoWlbItemBatchQueryWlbItemInventory {
    s.LockQuantity = &v
    return s
}
func (s *TaobaoWlbItemBatchQueryWlbItemInventory) SetStoreCode(v string) *TaobaoWlbItemBatchQueryWlbItemInventory {
    s.StoreCode = &v
    return s
}
func (s *TaobaoWlbItemBatchQueryWlbItemInventory) SetType(v string) *TaobaoWlbItemBatchQueryWlbItemInventory {
    s.Type = &v
    return s
}
