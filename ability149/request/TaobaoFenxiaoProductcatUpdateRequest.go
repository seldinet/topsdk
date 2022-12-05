package request


type TaobaoFenxiaoProductcatUpdateRequest struct {
    /*
        产品线ID     */
    ProductLineId  *int64 `json:"product_line_id" required:"true" `
    /*
        产品线名称     */
    Name  *string `json:"name,omitempty" required:"false" `
    /*
        最低零售价比例，注意：100.00%，则输入为10000     */
    RetailLowPercent  *int64 `json:"retail_low_percent,omitempty" required:"false" `
    /*
        最高零售价比例，注意：100.00%，则输入为10000     */
    RetailHighPercent  *int64 `json:"retail_high_percent,omitempty" required:"false" `
    /*
        代销默认采购价比例，注意：100.00%，则输入为10000     */
    AgentCostPercent  *int64 `json:"agent_cost_percent,omitempty" required:"false" `
    /*
        经销默认采购价比例，注意：100.00%，则输入为10000     */
    DealerCostPercent  *int64 `json:"dealer_cost_percent,omitempty" required:"false" `
}

func (s *TaobaoFenxiaoProductcatUpdateRequest) SetProductLineId(v int64) *TaobaoFenxiaoProductcatUpdateRequest {
    s.ProductLineId = &v
    return s
}
func (s *TaobaoFenxiaoProductcatUpdateRequest) SetName(v string) *TaobaoFenxiaoProductcatUpdateRequest {
    s.Name = &v
    return s
}
func (s *TaobaoFenxiaoProductcatUpdateRequest) SetRetailLowPercent(v int64) *TaobaoFenxiaoProductcatUpdateRequest {
    s.RetailLowPercent = &v
    return s
}
func (s *TaobaoFenxiaoProductcatUpdateRequest) SetRetailHighPercent(v int64) *TaobaoFenxiaoProductcatUpdateRequest {
    s.RetailHighPercent = &v
    return s
}
func (s *TaobaoFenxiaoProductcatUpdateRequest) SetAgentCostPercent(v int64) *TaobaoFenxiaoProductcatUpdateRequest {
    s.AgentCostPercent = &v
    return s
}
func (s *TaobaoFenxiaoProductcatUpdateRequest) SetDealerCostPercent(v int64) *TaobaoFenxiaoProductcatUpdateRequest {
    s.DealerCostPercent = &v
    return s
}

func (req *TaobaoFenxiaoProductcatUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ProductLineId != nil) {
        paramMap["product_line_id"] = *req.ProductLineId
    }
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

func (req *TaobaoFenxiaoProductcatUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}