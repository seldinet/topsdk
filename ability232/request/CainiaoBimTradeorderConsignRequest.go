package request


type CainiaoBimTradeorderConsignRequest struct {
    /*
        交易单号     */
    TradeId  *string `json:"trade_id" required:"true" `
    /*
        仓储编码，ERP指定仓库发货时需要传值，编码由菜鸟提供     */
    StoreCode  *string `json:"store_code,omitempty" required:"false" `
    /*
        选择的线路ID非必填字段     */
    ResId  *string `json:"res_id,omitempty" required:"false" `
}

func (s *CainiaoBimTradeorderConsignRequest) SetTradeId(v string) *CainiaoBimTradeorderConsignRequest {
    s.TradeId = &v
    return s
}
func (s *CainiaoBimTradeorderConsignRequest) SetStoreCode(v string) *CainiaoBimTradeorderConsignRequest {
    s.StoreCode = &v
    return s
}
func (s *CainiaoBimTradeorderConsignRequest) SetResId(v string) *CainiaoBimTradeorderConsignRequest {
    s.ResId = &v
    return s
}

func (req *CainiaoBimTradeorderConsignRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.TradeId != nil) {
        paramMap["trade_id"] = *req.TradeId
    }
    if(req.StoreCode != nil) {
        paramMap["store_code"] = *req.StoreCode
    }
    if(req.ResId != nil) {
        paramMap["res_id"] = *req.ResId
    }
    return paramMap
}

func (req *CainiaoBimTradeorderConsignRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}