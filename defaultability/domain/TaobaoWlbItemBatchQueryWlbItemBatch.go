package domain

import (
        "topsdk/util"
    )

type TaobaoWlbItemBatchQueryWlbItemBatch struct {
    /*
        商品批次记录id     */
    Id  *int64 `json:"id,omitempty" `

    /*
        用户id     */
    UserId  *int64 `json:"user_id,omitempty" `

    /*
        商品id     */
    ItemId  *int64 `json:"item_id,omitempty" `

    /*
        商品数量     */
    Quantity  *int64 `json:"quantity,omitempty" `

    /*
        残次数量     */
    DefectQuantity  *int64 `json:"defect_quantity,omitempty" `

    /*
        存储类型     */
    StoreCode  *string `json:"store_code,omitempty" `

    /*
        批次编号     */
    BatchCode  *string `json:"batch_code,omitempty" `

    /*
        生产编号     */
    ProduceCode  *string `json:"produce_code,omitempty" `

    /*
        到期时间     */
    DueDate  *util.LocalTime `json:"due_date,omitempty" `

    /*
        生产日期     */
    ProduceDate  *util.LocalTime `json:"produce_date,omitempty" `

    /*
        接受日期     */
    ReceiveDate  *util.LocalTime `json:"receive_date,omitempty" `

    /*
        保质期     */
    GuaranteePeriod  *string `json:"guarantee_period,omitempty" `

    /*
        天（单位）     */
    GuaranteeUnit  *int64 `json:"guarantee_unit,omitempty" `

    /*
        产地     */
    ProduceArea  *string `json:"produce_area,omitempty" `

    /*
        描述     */
    Remarks  *string `json:"remarks,omitempty" `

    /*
        状态。item_batch_status_open:开放 item_batch_status_lock:冻结 item_batch_status_invalid:无效     */
    Status  *string `json:"status,omitempty" `

}

func (s *TaobaoWlbItemBatchQueryWlbItemBatch) SetId(v int64) *TaobaoWlbItemBatchQueryWlbItemBatch {
    s.Id = &v
    return s
}
func (s *TaobaoWlbItemBatchQueryWlbItemBatch) SetUserId(v int64) *TaobaoWlbItemBatchQueryWlbItemBatch {
    s.UserId = &v
    return s
}
func (s *TaobaoWlbItemBatchQueryWlbItemBatch) SetItemId(v int64) *TaobaoWlbItemBatchQueryWlbItemBatch {
    s.ItemId = &v
    return s
}
func (s *TaobaoWlbItemBatchQueryWlbItemBatch) SetQuantity(v int64) *TaobaoWlbItemBatchQueryWlbItemBatch {
    s.Quantity = &v
    return s
}
func (s *TaobaoWlbItemBatchQueryWlbItemBatch) SetDefectQuantity(v int64) *TaobaoWlbItemBatchQueryWlbItemBatch {
    s.DefectQuantity = &v
    return s
}
func (s *TaobaoWlbItemBatchQueryWlbItemBatch) SetStoreCode(v string) *TaobaoWlbItemBatchQueryWlbItemBatch {
    s.StoreCode = &v
    return s
}
func (s *TaobaoWlbItemBatchQueryWlbItemBatch) SetBatchCode(v string) *TaobaoWlbItemBatchQueryWlbItemBatch {
    s.BatchCode = &v
    return s
}
func (s *TaobaoWlbItemBatchQueryWlbItemBatch) SetProduceCode(v string) *TaobaoWlbItemBatchQueryWlbItemBatch {
    s.ProduceCode = &v
    return s
}
func (s *TaobaoWlbItemBatchQueryWlbItemBatch) SetDueDate(v util.LocalTime) *TaobaoWlbItemBatchQueryWlbItemBatch {
    s.DueDate = &v
    return s
}
func (s *TaobaoWlbItemBatchQueryWlbItemBatch) SetProduceDate(v util.LocalTime) *TaobaoWlbItemBatchQueryWlbItemBatch {
    s.ProduceDate = &v
    return s
}
func (s *TaobaoWlbItemBatchQueryWlbItemBatch) SetReceiveDate(v util.LocalTime) *TaobaoWlbItemBatchQueryWlbItemBatch {
    s.ReceiveDate = &v
    return s
}
func (s *TaobaoWlbItemBatchQueryWlbItemBatch) SetGuaranteePeriod(v string) *TaobaoWlbItemBatchQueryWlbItemBatch {
    s.GuaranteePeriod = &v
    return s
}
func (s *TaobaoWlbItemBatchQueryWlbItemBatch) SetGuaranteeUnit(v int64) *TaobaoWlbItemBatchQueryWlbItemBatch {
    s.GuaranteeUnit = &v
    return s
}
func (s *TaobaoWlbItemBatchQueryWlbItemBatch) SetProduceArea(v string) *TaobaoWlbItemBatchQueryWlbItemBatch {
    s.ProduceArea = &v
    return s
}
func (s *TaobaoWlbItemBatchQueryWlbItemBatch) SetRemarks(v string) *TaobaoWlbItemBatchQueryWlbItemBatch {
    s.Remarks = &v
    return s
}
func (s *TaobaoWlbItemBatchQueryWlbItemBatch) SetStatus(v string) *TaobaoWlbItemBatchQueryWlbItemBatch {
    s.Status = &v
    return s
}
