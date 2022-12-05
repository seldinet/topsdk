package request


type TaobaoFenxiaoProductcatAddRequest struct {
    /*
        产品线名称     */
    Name  *string `json:"name" required:"true" `
    /*
        最低零售价比例，注意：100.00%，则输入为10000     */
    RetailLowPercent  *int64 `json:"retail_low_percent" required:"true" `
    /*
        最高零售价比例，注意：100.00%，则输入为10000     */
    RetailHighPercent  *int64 `json:"retail_high_percent" required:"true" `
    /*
        代销默认采购价比例，注意：100.00%，则输入为10000     */
    AgentCostPercent  *int64 `json:"agent_cost_percent" required:"true" `
    /*
        经销默认采购价比例，注意：100.00%，则输入为10000     */
    DealerCostPercent  *int64 `json:"dealer_cost_percent" required:"true" `
}

func (s *TaobaoFenxiaoProductcatAddRequest) SetName(v string) *TaobaoFenxiaoProductcatAddRequest {
    s.Name = &v
    return s
}
func (s *TaobaoFenxiaoProductcatAddRequest) SetRetailLowPercent(v int64) *TaobaoFenxiaoProductcatAddRequest {
    s.RetailLowPercent = &v
    return s
}
func (s *TaobaoFenxiaoProductcatAddRequest) SetRetailHighPercent(v int64) *TaobaoFenxiaoProductcatAddRequest {
    s.RetailHighPercent = &v
    return s
}
func (s *TaobaoFenxiaoProductcatAddRequest) SetAgentCostPercent(v int64) *TaobaoFenxiaoProductcatAddRequest {
    s.AgentCostPercent = &v
    return s
}
func (s *TaobaoFenxiaoProductcatAddRequest) SetDealerCostPercent(v int64) *TaobaoFenxiaoProductcatAddRequest {
    s.DealerCostPercent = &v
    return s
}

func (req *TaobaoFenxiaoProductcatAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Name != nil) {
        paramMap["name"] = *req.Name
    }
    if(req.RetailLowPercent != nil) {
        paramMap["retail_low_percent"] = *req.RetailLowPercent
    }
    if(req.RetailHighPercent != nil) {
        paramMap["retail_high_percent"] = *req.RetailHighPercent
    }
    if(req.AgentCostPercent != nil) {
        paramMap["agent_cost_percent"] = *req.AgentCostPercent
    }
    if(req.DealerCostPercent != nil) {
        paramMap["dealer_cost_percent"] = *req.DealerCostPercent
    }
    return paramMap
}

func (req *TaobaoFenxiaoProductcatAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}