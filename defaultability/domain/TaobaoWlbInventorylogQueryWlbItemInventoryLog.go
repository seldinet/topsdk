package domain

import (
        "topsdk/util"
    )

type TaobaoWlbInventorylogQueryWlbItemInventoryLog struct {
    /*
        库存变更ID     */
    Id  *int64 `json:"id,omitempty" `

    /*
        用户ID     */
    UserId  *int64 `json:"user_id,omitempty" `

    /*
        库存操作作类型CHU_KU 1-出库RU_KU 2-入库FREEZE 3-冻结THAW 4-解冻CHECK_FREEZE 5-冻结确认CHANGE_KU 6-库存类型变更     */
    OpType  *string `json:"op_type,omitempty" `

    /*
        库存操作者ID     */
    OpUserId  *int64 `json:"op_user_id,omitempty" `

    /*
        商品ID     */
    ItemId  *int64 `json:"item_id,omitempty" `

    /*
        批次号     */
    BatchCode  *string `json:"batch_code,omitempty" `

    /*
        仓库编码     */
    StoreCode  *string `json:"store_code,omitempty" `

    /*
        备注     */
    Remark  *string `json:"remark,omitempty" `

    /*
        订单号     */
    OrderCode  *string `json:"order_code,omitempty" `

    /*
        订单商品ID     */
    OrderItemId  *int64 `json:"order_item_id,omitempty" `

    /*
        处理数量变化值     */
    Quantity  *int64 `json:"quantity,omitempty" `

    /*
        创建日期     */
    GmtCreate  *util.LocalTime `json:"gmt_create,omitempty" `

    /*
        结果值     */
    ResultQuantity  *int64 `json:"result_quantity,omitempty" `

    /*
        VENDIBLE  1-可销售;FREEZE  201-冻结库存;ONWAY  301-在途库存;DEFECT  101-残存品;ENGINE_DAMAGE 102-机损;BOX_DAMAGE 103-箱损     */
    InventType  *string `json:"invent_type,omitempty" `

}

func (s *TaobaoWlbInventorylogQueryWlbItemInventoryLog) SetId(v int64) *TaobaoWlbInventorylogQueryWlbItemInventoryLog {
    s.Id = &v
    return s
}
func (s *TaobaoWlbInventorylogQueryWlbItemInventoryLog) SetUserId(v int64) *TaobaoWlbInventorylogQueryWlbItemInventoryLog {
    s.UserId = &v
    return s
}
func (s *TaobaoWlbInventorylogQueryWlbItemInventoryLog) SetOpType(v string) *TaobaoWlbInventorylogQueryWlbItemInventoryLog {
    s.OpType = &v
    return s
}
func (s *TaobaoWlbInventorylogQueryWlbItemInventoryLog) SetOpUserId(v int64) *TaobaoWlbInventorylogQueryWlbItemInventoryLog {
    s.OpUserId = &v
    return s
}
func (s *TaobaoWlbInventorylogQueryWlbItemInventoryLog) SetItemId(v int64) *TaobaoWlbInventorylogQueryWlbItemInventoryLog {
    s.ItemId = &v
    return s
}
func (s *TaobaoWlbInventorylogQueryWlbItemInventoryLog) SetBatchCode(v string) *TaobaoWlbInventorylogQueryWlbItemInventoryLog {
    s.BatchCode = &v
    return s
}
func (s *TaobaoWlbInventorylogQueryWlbItemInventoryLog) SetStoreCode(v string) *TaobaoWlbInventorylogQueryWlbItemInventoryLog {
    s.StoreCode = &v
    return s
}
func (s *TaobaoWlbInventorylogQueryWlbItemInventoryLog) SetRemark(v string) *TaobaoWlbInventorylogQueryWlbItemInventoryLog {
    s.Remark = &v
    return s
}
func (s *TaobaoWlbInventorylogQueryWlbItemInventoryLog) SetOrderCode(v string) *TaobaoWlbInventorylogQueryWlbItemInventoryLog {
    s.OrderCode = &v
    return s
}
func (s *TaobaoWlbInventorylogQueryWlbItemInventoryLog) SetOrderItemId(v int64) *TaobaoWlbInventorylogQueryWlbItemInventoryLog {
    s.OrderItemId = &v
    return s
}
func (s *TaobaoWlbInventorylogQueryWlbItemInventoryLog) SetQuantity(v int64) *TaobaoWlbInventorylogQueryWlbItemInventoryLog {
    s.Quantity = &v
    return s
}
func (s *TaobaoWlbInventorylogQueryWlbItemInventoryLog) SetGmtCreate(v util.LocalTime) *TaobaoWlbInventorylogQueryWlbItemInventoryLog {
    s.GmtCreate = &v
    return s
}
func (s *TaobaoWlbInventorylogQueryWlbItemInventoryLog) SetResultQuantity(v int64) *TaobaoWlbInventorylogQueryWlbItemInventoryLog {
    s.ResultQuantity = &v
    return s
}
func (s *TaobaoWlbInventorylogQueryWlbItemInventoryLog) SetInventType(v string) *TaobaoWlbInventorylogQueryWlbItemInventoryLog {
    s.InventType = &v
    return s
}
