package domain


type TaobaoWlbItemBatchQueryWlbItemBatchInventory struct {
    /*
        商品ID     */
    ItemId  *int64 `json:"item_id,omitempty" `

    /*
        商品在所有仓库的可销库存总数     */
    TotalQuantity  *int64 `json:"total_quantity,omitempty" `

    /*
        批次库存查询结果     */
    ItemBatch  *[]TaobaoWlbItemBatchQueryWlbItemBatch `json:"item_batch,omitempty" `

    /*
        商品库存查询结果     */
    ItemInventorys  *[]TaobaoWlbItemBatchQueryWlbItemInventory `json:"item_inventorys,omitempty" `

}

func (s *TaobaoWlbItemBatchQueryWlbItemBatchInventory) SetItemId(v int64) *TaobaoWlbItemBatchQueryWlbItemBatchInventory {
    s.ItemId = &v
    return s
}
func (s *TaobaoWlbItemBatchQueryWlbItemBatchInventory) SetTotalQuantity(v int64) *TaobaoWlbItemBatchQueryWlbItemBatchInventory {
    s.TotalQuantity = &v
    return s
}
func (s *TaobaoWlbItemBatchQueryWlbItemBatchInventory) SetItemBatch(v []TaobaoWlbItemBatchQueryWlbItemBatch) *TaobaoWlbItemBatchQueryWlbItemBatchInventory {
    s.ItemBatch = &v
    return s
}
func (s *TaobaoWlbItemBatchQueryWlbItemBatchInventory) SetItemInventorys(v []TaobaoWlbItemBatchQueryWlbItemInventory) *TaobaoWlbItemBatchQueryWlbItemBatchInventory {
    s.ItemInventorys = &v
    return s
}
