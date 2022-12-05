package request


type TaobaoWlbOutInventoryChangeNotifyRequest struct {
    /*
        物流宝商品id或前台宝贝id（由type类型决定）     */
    ItemId  *int64 `json:"item_id" required:"true" `
    /*
        WLB_ITEM--物流宝商品 IC_ITEM--淘宝商品 IC_SKU--淘宝sku     */
    Type  *string `json:"type" required:"true" `
    /*
        库存变化数量     */
    ChangeCount  *int64 `json:"change_count" required:"true" `
    /*
        本次库存变化后库存余额     */
    ResultCount  *int64 `json:"result_count" required:"true" `
    /*
        OUT--出库 IN--入库     */
    OpType  *string `json:"op_type" required:"true" `
    /*
        （1）OTHER： 其他 （2）TAOBAO_TRADE： 淘宝交易 （3）OTHER_TRADE：其他交易 （4）ALLCOATE： 调拨 （5）CHECK:盘点 （6）PURCHASE:采购     */
    Source  *string `json:"source" required:"true" `
    /*
        订单号，如果source为TAOBAO_TRADE,order_source_code必须为淘宝交易号     */
    OrderSourceCode  *string `json:"order_source_code,omitempty" required:"false" `
    /*
        库存变化唯一标识，用于去重，一个外部唯一编码唯一标识一次库存变化。     */
    OutBizCode  *string `json:"out_biz_code" required:"true" `
    /*
        目前非必须，系统会选择默认值     */
    StoreCode  *string `json:"store_code,omitempty" required:"false" `
}

func (s *TaobaoWlbOutInventoryChangeNotifyRequest) SetItemId(v int64) *TaobaoWlbOutInventoryChangeNotifyRequest {
    s.ItemId = &v
    return s
}
func (s *TaobaoWlbOutInventoryChangeNotifyRequest) SetType(v string) *TaobaoWlbOutInventoryChangeNotifyRequest {
    s.Type = &v
    return s
}
func (s *TaobaoWlbOutInventoryChangeNotifyRequest) SetChangeCount(v int64) *TaobaoWlbOutInventoryChangeNotifyRequest {
    s.ChangeCount = &v
    return s
}
func (s *TaobaoWlbOutInventoryChangeNotifyRequest) SetResultCount(v int64) *TaobaoWlbOutInventoryChangeNotifyRequest {
    s.ResultCount = &v
    return s
}
func (s *TaobaoWlbOutInventoryChangeNotifyRequest) SetOpType(v string) *TaobaoWlbOutInventoryChangeNotifyRequest {
    s.OpType = &v
    return s
}
func (s *TaobaoWlbOutInventoryChangeNotifyRequest) SetSource(v string) *TaobaoWlbOutInventoryChangeNotifyRequest {
    s.Source = &v
    return s
}
func (s *TaobaoWlbOutInventoryChangeNotifyRequest) SetOrderSourceCode(v string) *TaobaoWlbOutInventoryChangeNotifyRequest {
    s.OrderSourceCode = &v
    return s
}
func (s *TaobaoWlbOutInventoryChangeNotifyRequest) SetOutBizCode(v string) *TaobaoWlbOutInventoryChangeNotifyRequest {
    s.OutBizCode = &v
    return s
}
func (s *TaobaoWlbOutInventoryChangeNotifyRequest) SetStoreCode(v string) *TaobaoWlbOutInventoryChangeNotifyRequest {
    s.StoreCode = &v
    return s
}

func (req *TaobaoWlbOutInventoryChangeNotifyRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    if(req.ChangeCount != nil) {
        paramMap["change_count"] = *req.ChangeCount
    }
    if(req.ResultCount != nil) {
        paramMap["result_count"] = *req.ResultCount
    }
    if(req.OpType != nil) {
        paramMap["op_type"] = *req.OpType
    }
    if(req.Source != nil) {
        paramMap["source"] = *req.Source
    }
    if(req.OrderSourceCode != nil) {
        paramMap["order_source_code"] = *req.OrderSourceCode
    }
    if(req.OutBizCode != nil) {
        paramMap["out_biz_code"] = *req.OutBizCode
    }
    if(req.StoreCode != nil) {
        paramMap["store_code"] = *req.StoreCode
    }
    return paramMap
}

func (req *TaobaoWlbOutInventoryChangeNotifyRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}