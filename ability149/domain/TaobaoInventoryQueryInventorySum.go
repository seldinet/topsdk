package domain


type TaobaoInventoryQueryInventorySum struct {
    /*
        总预扣数量     */
    ReserveQuantity  *int64 `json:"reserve_quantity,omitempty" `

    /*
        库存类型：
1：正常 
2：损坏 
3：冻结 
10：质押 
11-20:商家自定义     */
    InventoryType  *int64 `json:"inventory_type,omitempty" `

    /*
        商家仓库编码     */
    StoreCode  *string `json:"store_code,omitempty" `

    /*
        库存类型名称     */
    InventoryTypeName  *string `json:"inventory_type_name,omitempty" `

    /*
        商品后端ID，如果有传sc_item_code,参数可以为0     */
    ScItemId  *int64 `json:"sc_item_id,omitempty" `

    /*
        商品商家编码     */
    ScItemCode  *string `json:"sc_item_code,omitempty" `

    /*
        总物理库存数量     */
    Quantity  *int64 `json:"quantity,omitempty" `

    /*
        总占用数量     */
    OccupyQuantity  *int64 `json:"occupy_quantity,omitempty" `

}

func (s *TaobaoInventoryQueryInventorySum) SetReserveQuantity(v int64) *TaobaoInventoryQueryInventorySum {
    s.ReserveQuantity = &v
    return s
}
func (s *TaobaoInventoryQueryInventorySum) SetInventoryType(v int64) *TaobaoInventoryQueryInventorySum {
    s.InventoryType = &v
    return s
}
func (s *TaobaoInventoryQueryInventorySum) SetStoreCode(v string) *TaobaoInventoryQueryInventorySum {
    s.StoreCode = &v
    return s
}
func (s *TaobaoInventoryQueryInventorySum) SetInventoryTypeName(v string) *TaobaoInventoryQueryInventorySum {
    s.InventoryTypeName = &v
    return s
}
func (s *TaobaoInventoryQueryInventorySum) SetScItemId(v int64) *TaobaoInventoryQueryInventorySum {
    s.ScItemId = &v
    return s
}
func (s *TaobaoInventoryQueryInventorySum) SetScItemCode(v string) *TaobaoInventoryQueryInventorySum {
    s.ScItemCode = &v
    return s
}
func (s *TaobaoInventoryQueryInventorySum) SetQuantity(v int64) *TaobaoInventoryQueryInventorySum {
    s.Quantity = &v
    return s
}
func (s *TaobaoInventoryQueryInventorySum) SetOccupyQuantity(v int64) *TaobaoInventoryQueryInventorySum {
    s.OccupyQuantity = &v
    return s
}
