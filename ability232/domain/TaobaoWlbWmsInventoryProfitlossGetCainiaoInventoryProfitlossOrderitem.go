package domain

import (
        "topsdk/util"
    )

type TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossOrderitem struct {
    /*
        商家对商品的编码     */
    ItemCode  *string `json:"item_code,omitempty" `

    /*
        库存类型 1 可销售库存（正品）  101 残次     */
    InventoryType  *string `json:"inventory_type,omitempty" `

    /*
        商品数量     */
    ItemQty  *string `json:"item_qty,omitempty" `

    /*
        商品保质期信息，失效日期     */
    DueDate  *util.LocalTime `json:"due_date,omitempty" `

    /*
        商品保质期信息，生产日期     */
    ProduceDate  *util.LocalTime `json:"produce_date,omitempty" `

    /*
        生产编码，同一商品可能因商家不同有不同编码     */
    ProduceCode  *string `json:"produce_code,omitempty" `

    /*
        生产地区     */
    ProduceArea  *string `json:"produce_area,omitempty" `

    /*
        批次号     */
    BatchCode  *string `json:"batch_code,omitempty" `

    /*
        商品ID     */
    ItemId  *string `json:"item_id,omitempty" `

}

func (s *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossOrderitem) SetItemCode(v string) *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossOrderitem {
    s.ItemCode = &v
    return s
}
func (s *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossOrderitem) SetInventoryType(v string) *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossOrderitem {
    s.InventoryType = &v
    return s
}
func (s *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossOrderitem) SetItemQty(v string) *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossOrderitem {
    s.ItemQty = &v
    return s
}
func (s *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossOrderitem) SetDueDate(v util.LocalTime) *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossOrderitem {
    s.DueDate = &v
    return s
}
func (s *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossOrderitem) SetProduceDate(v util.LocalTime) *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossOrderitem {
    s.ProduceDate = &v
    return s
}
func (s *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossOrderitem) SetProduceCode(v string) *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossOrderitem {
    s.ProduceCode = &v
    return s
}
func (s *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossOrderitem) SetProduceArea(v string) *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossOrderitem {
    s.ProduceArea = &v
    return s
}
func (s *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossOrderitem) SetBatchCode(v string) *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossOrderitem {
    s.BatchCode = &v
    return s
}
func (s *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossOrderitem) SetItemId(v string) *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossOrderitem {
    s.ItemId = &v
    return s
}
