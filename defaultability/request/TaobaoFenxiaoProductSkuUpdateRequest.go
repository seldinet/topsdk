package request


type TaobaoFenxiaoProductSkuUpdateRequest struct {
    /*
        产品ID     */
    ProductId  *int64 `json:"product_id" required:"true" `
    /*
        采购基准价     */
    StandardPrice  *string `json:"standard_price,omitempty" required:"false" `
    /*
        代销采购价     */
    AgentCostPrice  *string `json:"agent_cost_price,omitempty" required:"false" `
    /*
        经销采购价     */
    DealerCostPrice  *string `json:"dealer_cost_price,omitempty" required:"false" `
    /*
        产品SKU库存     */
    Quantity  *int64 `json:"quantity,omitempty" required:"false" `
    /*
        商家编码     */
    SkuNumber  *string `json:"sku_number,omitempty" required:"false" `
    /*
        sku属性     */
    Properties  *string `json:"properties" required:"true" `
}

func (s *TaobaoFenxiaoProductSkuUpdateRequest) SetProductId(v int64) *TaobaoFenxiaoProductSkuUpdateRequest {
    s.ProductId = &v
    return s
}
func (s *TaobaoFenxiaoProductSkuUpdateRequest) SetStandardPrice(v string) *TaobaoFenxiaoProductSkuUpdateRequest {
    s.StandardPrice = &v
    return s
}
func (s *TaobaoFenxiaoProductSkuUpdateRequest) SetAgentCostPrice(v string) *TaobaoFenxiaoProductSkuUpdateRequest {
    s.AgentCostPrice = &v
    return s
}
func (s *TaobaoFenxiaoProductSkuUpdateRequest) SetDealerCostPrice(v string) *TaobaoFenxiaoProductSkuUpdateRequest {
    s.DealerCostPrice = &v
    return s
}
func (s *TaobaoFenxiaoProductSkuUpdateRequest) SetQuantity(v int64) *TaobaoFenxiaoProductSkuUpdateRequest {
    s.Quantity = &v
    return s
}
func (s *TaobaoFenxiaoProductSkuUpdateRequest) SetSkuNumber(v string) *TaobaoFenxiaoProductSkuUpdateRequest {
    s.SkuNumber = &v
    return s
}
func (s *TaobaoFenxiaoProductSkuUpdateRequest) SetProperties(v string) *TaobaoFenxiaoProductSkuUpdateRequest {
    s.Properties = &v
    return s
}

func (req *TaobaoFenxiaoProductSkuUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ProductId != nil) {
        paramMap["product_id"] = *req.ProductId
    }
    if(req.StandardPrice != nil) {
        paramMap["standard_price"] = *req.StandardPrice
    }
    if(req.AgentCostPrice != nil) {
        paramMap["agent_cost_price"] = *req.AgentCostPrice
    }
    if(req.DealerCostPrice != nil) {
        paramMap["dealer_cost_price"] = *req.DealerCostPrice
    }
    if(req.Quantity != nil) {
        paramMap["quantity"] = *req.Quantity
    }
    if(req.SkuNumber != nil) {
        paramMap["sku_number"] = *req.SkuNumber
    }
    if(req.Properties != nil) {
        paramMap["properties"] = *req.Properties
    }
    return paramMap
}

func (req *TaobaoFenxiaoProductSkuUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}