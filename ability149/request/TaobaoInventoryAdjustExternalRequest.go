package request


type TaobaoInventoryAdjustExternalRequest struct {
    /*
        商家外部定单号     */
    BizUniqueCode  *string `json:"biz_unique_code" required:"true" `
    /*
        库存占用返回的操作码.operate_type 为OUTBOUND时，如果是确认事先进行过的库存占用，需要传入当时返回的操作码，并且明细必须与申请时保持一致     */
    OccupyOperateCode  *string `json:"occupy_operate_code,omitempty" required:"false" `
    /*
        test     */
    OperateType  *string `json:"operate_type,omitempty" required:"false" `
    /*
        外部订单类型, BALANCE:盘点、NON_TAOBAO_TRADE:非淘宝交易、ALLOCATE:调拨、OTHERS:其他     */
    BizType  *string `json:"biz_type,omitempty" required:"false" `
    /*
        业务操作时间     */
    OperateTime  *string `json:"operate_time,omitempty" required:"false" `
    /*
        商家仓库编码     */
    StoreCode  *string `json:"store_code" required:"true" `
    /*
        商品初始库存信息： [{"scItemId":"商品后端ID，如果有传scItemCode,参数可以为0","scItemCode":"商品商家编码","inventoryType":"库存类型  1：正常,”direction”: 1: 盘盈 -1: 盘亏,参数可选,"quantity":"数量(正数)"}]     */
    Items  *string `json:"items" required:"true" `
    /*
        test     */
    ReduceType  *string `json:"reduce_type,omitempty" required:"false" `
}

func (s *TaobaoInventoryAdjustExternalRequest) SetBizUniqueCode(v string) *TaobaoInventoryAdjustExternalRequest {
    s.BizUniqueCode = &v
    return s
}
func (s *TaobaoInventoryAdjustExternalRequest) SetOccupyOperateCode(v string) *TaobaoInventoryAdjustExternalRequest {
    s.OccupyOperateCode = &v
    return s
}
func (s *TaobaoInventoryAdjustExternalRequest) SetOperateType(v string) *TaobaoInventoryAdjustExternalRequest {
    s.OperateType = &v
    return s
}
func (s *TaobaoInventoryAdjustExternalRequest) SetBizType(v string) *TaobaoInventoryAdjustExternalRequest {
    s.BizType = &v
    return s
}
func (s *TaobaoInventoryAdjustExternalRequest) SetOperateTime(v string) *TaobaoInventoryAdjustExternalRequest {
    s.OperateTime = &v
    return s
}
func (s *TaobaoInventoryAdjustExternalRequest) SetStoreCode(v string) *TaobaoInventoryAdjustExternalRequest {
    s.StoreCode = &v
    return s
}
func (s *TaobaoInventoryAdjustExternalRequest) SetItems(v string) *TaobaoInventoryAdjustExternalRequest {
    s.Items = &v
    return s
}
func (s *TaobaoInventoryAdjustExternalRequest) SetReduceType(v string) *TaobaoInventoryAdjustExternalRequest {
    s.ReduceType = &v
    return s
}

func (req *TaobaoInventoryAdjustExternalRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.BizUniqueCode != nil) {
        paramMap["biz_unique_code"] = *req.BizUniqueCode
    }
    if(req.OccupyOperateCode != nil) {
        paramMap["occupy_operate_code"] = *req.OccupyOperateCode
    }
    if(req.OperateType != nil) {
        paramMap["operate_type"] = *req.OperateType
    }
    if(req.BizType != nil) {
        paramMap["biz_type"] = *req.BizType
    }
    if(req.OperateTime != nil) {
        paramMap["operate_time"] = *req.OperateTime
    }
    if(req.StoreCode != nil) {
        paramMap["store_code"] = *req.StoreCode
    }
    if(req.Items != nil) {
        paramMap["items"] = *req.Items
    }
    if(req.ReduceType != nil) {
        paramMap["reduce_type"] = *req.ReduceType
    }
    return paramMap
}

func (req *TaobaoInventoryAdjustExternalRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}