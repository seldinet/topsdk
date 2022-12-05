package domain

import (
        "topsdk/util"
    )

type TaobaoWlbWmsConsignBillGetInventoryitem struct {
    /*
        批次号     */
    BatchCode  *string `json:"batch_code,omitempty" `

    /*
        生产地区     */
    ProduceArea  *string `json:"produce_area,omitempty" `

    /*
        生产编码，同一商品可能因商家不同有不同编码     */
    ProduceCode  *string `json:"produce_code,omitempty" `

    /*
        商品保质期信息，生产日期     */
    ProduceDate  *util.LocalTime `json:"produce_date,omitempty" `

    /*
        商品保质期信息，失效日期     */
    DueDate  *util.LocalTime `json:"due_date,omitempty" `

    /*
        数量     */
    ItemQty  *int64 `json:"item_qty,omitempty" `

    /*
        库存类型：1 可销售库存 (正品) 101 类型用来定义残次品 201 冻结类型库存 301 在途库存     */
    InventoryType  *int64 `json:"inventory_type,omitempty" `

}

func (s *TaobaoWlbWmsConsignBillGetInventoryitem) SetBatchCode(v string) *TaobaoWlbWmsConsignBillGetInventoryitem {
    s.BatchCode = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetInventoryitem) SetProduceArea(v string) *TaobaoWlbWmsConsignBillGetInventoryitem {
    s.ProduceArea = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetInventoryitem) SetProduceCode(v string) *TaobaoWlbWmsConsignBillGetInventoryitem {
    s.ProduceCode = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetInventoryitem) SetProduceDate(v util.LocalTime) *TaobaoWlbWmsConsignBillGetInventoryitem {
    s.ProduceDate = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetInventoryitem) SetDueDate(v util.LocalTime) *TaobaoWlbWmsConsignBillGetInventoryitem {
    s.DueDate = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetInventoryitem) SetItemQty(v int64) *TaobaoWlbWmsConsignBillGetInventoryitem {
    s.ItemQty = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetInventoryitem) SetInventoryType(v int64) *TaobaoWlbWmsConsignBillGetInventoryitem {
    s.InventoryType = &v
    return s
}
