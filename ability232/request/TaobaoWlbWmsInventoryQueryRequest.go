package request

import (
        "topsdk/util"
    )

type TaobaoWlbWmsInventoryQueryRequest struct {
    /*
        每页多少条，最大50条     */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        分页的第几页     */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        渠道编码，type=3时字段传值有效。（TB 淘系， OTHERS 其他）     */
    ChannelCode  *string `json:"channel_code,omitempty" required:"false" `
    /*
        失效日期，type=2时字段传值有效。     */
    DueDate  *util.LocalTime `json:"due_date,omitempty" required:"false" `
    /*
        生产日期，type=2时字段传值有效。     */
    ProduceDate  *util.LocalTime `json:"produce_date,omitempty" required:"false" `
    /*
        库存批次号，type=2时字段传值有效。 商品设置为批次管理时，商品产生批次库存。当商品为批次管理时，此字段不传值，返回所有批次库存信息。     */
    BatchCode  *string `json:"batch_code,omitempty" required:"false" `
    /*
        库存查询类型 1-	汇总库存，不区分批次和渠道 2-	批次库存，库存按商品批次维度划分 3-	渠道库存，库存按渠道维度划分 （当前业务不支持批次库存和渠道库存共存，批次库存无渠道属性，渠道库存无批次属性）     */
    Type  *int64 `json:"type,omitempty" required:"false" `
    /*
        库存类型。 (1 正品 101 残次 102 机损 103 箱损 201 冻结库存 301 在途库存 )     */
    InventoryType  *int64 `json:"inventory_type,omitempty" required:"false" `
    /*
        仓库编码     */
    StoreCode  *string `json:"store_code,omitempty" required:"false" `
    /*
        菜鸟商品ID     */
    ItemId  *string `json:"item_id,omitempty" required:"false" `
}

func (s *TaobaoWlbWmsInventoryQueryRequest) SetPageSize(v int64) *TaobaoWlbWmsInventoryQueryRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoWlbWmsInventoryQueryRequest) SetPageNo(v int64) *TaobaoWlbWmsInventoryQueryRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoWlbWmsInventoryQueryRequest) SetChannelCode(v string) *TaobaoWlbWmsInventoryQueryRequest {
    s.ChannelCode = &v
    return s
}
func (s *TaobaoWlbWmsInventoryQueryRequest) SetDueDate(v util.LocalTime) *TaobaoWlbWmsInventoryQueryRequest {
    s.DueDate = &v
    return s
}
func (s *TaobaoWlbWmsInventoryQueryRequest) SetProduceDate(v util.LocalTime) *TaobaoWlbWmsInventoryQueryRequest {
    s.ProduceDate = &v
    return s
}
func (s *TaobaoWlbWmsInventoryQueryRequest) SetBatchCode(v string) *TaobaoWlbWmsInventoryQueryRequest {
    s.BatchCode = &v
    return s
}
func (s *TaobaoWlbWmsInventoryQueryRequest) SetType(v int64) *TaobaoWlbWmsInventoryQueryRequest {
    s.Type = &v
    return s
}
func (s *TaobaoWlbWmsInventoryQueryRequest) SetInventoryType(v int64) *TaobaoWlbWmsInventoryQueryRequest {
    s.InventoryType = &v
    return s
}
func (s *TaobaoWlbWmsInventoryQueryRequest) SetStoreCode(v string) *TaobaoWlbWmsInventoryQueryRequest {
    s.StoreCode = &v
    return s
}
func (s *TaobaoWlbWmsInventoryQueryRequest) SetItemId(v string) *TaobaoWlbWmsInventoryQueryRequest {
    s.ItemId = &v
    return s
}

func (req *TaobaoWlbWmsInventoryQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.ChannelCode != nil) {
        paramMap["channel_code"] = *req.ChannelCode
    }
    if(req.DueDate != nil) {
        paramMap["due_date"] = *req.DueDate
    }
    if(req.ProduceDate != nil) {
        paramMap["produce_date"] = *req.ProduceDate
    }
    if(req.BatchCode != nil) {
        paramMap["batch_code"] = *req.BatchCode
    }
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    if(req.InventoryType != nil) {
        paramMap["inventory_type"] = *req.InventoryType
    }
    if(req.StoreCode != nil) {
        paramMap["store_code"] = *req.StoreCode
    }
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    return paramMap
}

func (req *TaobaoWlbWmsInventoryQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}