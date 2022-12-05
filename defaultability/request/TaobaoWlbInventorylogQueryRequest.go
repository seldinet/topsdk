package request

import (
        "topsdk/util"
    )

type TaobaoWlbInventorylogQueryRequest struct {
    /*
        商品ID     */
    ItemId  *int64 `json:"item_id,omitempty" required:"false" `
    /*
        仓库编码     */
    StoreCode  *string `json:"store_code,omitempty" required:"false" `
    /*
        单号     */
    OrderCode  *string `json:"order_code,omitempty" required:"false" `
    /*
        起始修改时间,大于等于该时间     */
    GmtStart  *util.LocalTime `json:"gmt_start,omitempty" required:"false" `
    /*
        结束修改时间,小于等于该时间     */
    GmtEnd  *util.LocalTime `json:"gmt_end,omitempty" required:"false" `
    /*
        当前页 defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        分页记录个数 defalutValue��20    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        可指定授权的用户来查询     */
    OpUserId  *int64 `json:"op_user_id,omitempty" required:"false" `
    /*
        库存操作作类型(可以为空) CHU_KU 1-出库 RU_KU 2-入库 FREEZE 3-冻结 THAW 4-解冻 CHECK_FREEZE 5-冻结确认 CHANGE_KU 6-库存类型变更 若值不在范围内，则按CHU_KU处理     */
    OpType  *string `json:"op_type,omitempty" required:"false" `
}

func (s *TaobaoWlbInventorylogQueryRequest) SetItemId(v int64) *TaobaoWlbInventorylogQueryRequest {
    s.ItemId = &v
    return s
}
func (s *TaobaoWlbInventorylogQueryRequest) SetStoreCode(v string) *TaobaoWlbInventorylogQueryRequest {
    s.StoreCode = &v
    return s
}
func (s *TaobaoWlbInventorylogQueryRequest) SetOrderCode(v string) *TaobaoWlbInventorylogQueryRequest {
    s.OrderCode = &v
    return s
}
func (s *TaobaoWlbInventorylogQueryRequest) SetGmtStart(v util.LocalTime) *TaobaoWlbInventorylogQueryRequest {
    s.GmtStart = &v
    return s
}
func (s *TaobaoWlbInventorylogQueryRequest) SetGmtEnd(v util.LocalTime) *TaobaoWlbInventorylogQueryRequest {
    s.GmtEnd = &v
    return s
}
func (s *TaobaoWlbInventorylogQueryRequest) SetPageNo(v int64) *TaobaoWlbInventorylogQueryRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoWlbInventorylogQueryRequest) SetPageSize(v int64) *TaobaoWlbInventorylogQueryRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoWlbInventorylogQueryRequest) SetOpUserId(v int64) *TaobaoWlbInventorylogQueryRequest {
    s.OpUserId = &v
    return s
}
func (s *TaobaoWlbInventorylogQueryRequest) SetOpType(v string) *TaobaoWlbInventorylogQueryRequest {
    s.OpType = &v
    return s
}

func (req *TaobaoWlbInventorylogQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.StoreCode != nil) {
        paramMap["store_code"] = *req.StoreCode
    }
    if(req.OrderCode != nil) {
        paramMap["order_code"] = *req.OrderCode
    }
    if(req.GmtStart != nil) {
        paramMap["gmt_start"] = *req.GmtStart
    }
    if(req.GmtEnd != nil) {
        paramMap["gmt_end"] = *req.GmtEnd
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.OpUserId != nil) {
        paramMap["op_user_id"] = *req.OpUserId
    }
    if(req.OpType != nil) {
        paramMap["op_type"] = *req.OpType
    }
    return paramMap
}

func (req *TaobaoWlbInventorylogQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}