package request

import (
        "topsdk/util"
    )

type TaobaoWlbWmsCainiaoBillQueryRequest struct {
    /*
        每页条数。（每页条数不超过50条。默认为50） defalutValue��50    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        页码。（大于0的整数。默认为1） defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        订单类型 201 销售出库 501 退货入库 502 换货出库 503 补发出库904 普通入库 903 普通出库单 306 B2B入库单 305 B2B出库单 601 采购入库 901 退供出库单 701 盘点出库 702 盘点入库 711 库存异动单     */
    OrderType  *string `json:"order_type,omitempty" required:"false" `
    /*
        起始时间，此字段检索订单最后修改时间， 格式 yyyy-MM-dd HH:mm:ss。     */
    EndModifiedTime  *util.LocalTime `json:"end_modified_time" required:"true" `
    /*
        结束时间，此字段检索订单最后修改时间， 格式 yyyy-MM-dd HH:mm:ss。     */
    StartModifiedTime  *util.LocalTime `json:"start_modified_time" required:"true" `
}

func (s *TaobaoWlbWmsCainiaoBillQueryRequest) SetPageSize(v int64) *TaobaoWlbWmsCainiaoBillQueryRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoWlbWmsCainiaoBillQueryRequest) SetPageNo(v int64) *TaobaoWlbWmsCainiaoBillQueryRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoWlbWmsCainiaoBillQueryRequest) SetOrderType(v string) *TaobaoWlbWmsCainiaoBillQueryRequest {
    s.OrderType = &v
    return s
}
func (s *TaobaoWlbWmsCainiaoBillQueryRequest) SetEndModifiedTime(v util.LocalTime) *TaobaoWlbWmsCainiaoBillQueryRequest {
    s.EndModifiedTime = &v
    return s
}
func (s *TaobaoWlbWmsCainiaoBillQueryRequest) SetStartModifiedTime(v util.LocalTime) *TaobaoWlbWmsCainiaoBillQueryRequest {
    s.StartModifiedTime = &v
    return s
}

func (req *TaobaoWlbWmsCainiaoBillQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.OrderType != nil) {
        paramMap["order_type"] = *req.OrderType
    }
    if(req.EndModifiedTime != nil) {
        paramMap["end_modified_time"] = *req.EndModifiedTime
    }
    if(req.StartModifiedTime != nil) {
        paramMap["start_modified_time"] = *req.StartModifiedTime
    }
    return paramMap
}

func (req *TaobaoWlbWmsCainiaoBillQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}