package request


type TaobaoWlbWmsOrderCancelNotifyRequest struct {
    /*
        订单类型     */
    OrderCode  *string `json:"order_code" required:"true" `
    /*
        单据类型 601普通入库单、501销退入库单、302 调拨入库单、904其他入库单、301 调拨出库单、901普通出库单、903 其他出库单、201 一般交易出库单 202 B2B交易出库单 502 换货出库单 503 补发出库单、1102 仓内加工作业单、 701 盘亏单、702盘盈单、711 库存状态调整单     */
    OrderType  *string `json:"order_type" required:"true" `
    /*
        仓库编码     */
    StoreCode  *string `json:"store_code" required:"true" `
}

func (s *TaobaoWlbWmsOrderCancelNotifyRequest) SetOrderCode(v string) *TaobaoWlbWmsOrderCancelNotifyRequest {
    s.OrderCode = &v
    return s
}
func (s *TaobaoWlbWmsOrderCancelNotifyRequest) SetOrderType(v string) *TaobaoWlbWmsOrderCancelNotifyRequest {
    s.OrderType = &v
    return s
}
func (s *TaobaoWlbWmsOrderCancelNotifyRequest) SetStoreCode(v string) *TaobaoWlbWmsOrderCancelNotifyRequest {
    s.StoreCode = &v
    return s
}

func (req *TaobaoWlbWmsOrderCancelNotifyRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OrderCode != nil) {
        paramMap["order_code"] = *req.OrderCode
    }
    if(req.OrderType != nil) {
        paramMap["order_type"] = *req.OrderType
    }
    if(req.StoreCode != nil) {
        paramMap["store_code"] = *req.StoreCode
    }
    return paramMap
}

func (req *TaobaoWlbWmsOrderCancelNotifyRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}