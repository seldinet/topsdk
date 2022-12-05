package request


type TaobaoFenxiaoProductSkuAddRequest struct {
    /*
        产品ID     */
    ProductId  *int64 `json:"product_id" required:"true" `
    /*
        采购基准价，最大值1000000000     */
    StandardPrice  *string `json:"standard_price" required:"true" `
    /*
        代销采购价     */
    AgentCostPrice  *string `json:"agent_cost_price,omitempty" required:"false" `
    /*
        经销采购价     */
    DealerCostPrice  *string `json:"dealer_cost_price,omitempty" required:"false" `
    /*
        sku产品库存，在0到1000000之间，如果不传，则库存为0     */
    Quantity  *int64 `json:"quantity,omitempty" required:"false" `
    /*
        商家编码     */
    SkuNumber  *string `json:"sku_number,omitempty" required:"false" `
    /*
        sku属性     */
    Properties  *string `json:"properties" required:"true" `
}

func (s *TaobaoFenxiaoProductSkuAddRequest) SetProductId(v int64) *TaobaoFenxiaoProductSkuAddRequest {
    s.ProductId = &v
    return s
}
func (s *TaobaoFenxiaoProductSkuAddRequest) SetStandardPrice(v string) *TaobaoFenxiaoProductSkuAddRequest {
    s.StandardPrice = &v
    return s
}
func (s *TaobaoFenxiaoProductSkuAddRequest) SetAgentCostPrice(v string) *TaobaoFenxiaoProductSkuAddRequest {
    s.AgentCostPrice = &v
    return s
}
func (s *TaobaoFenxiaoProductSkuAddRequest) SetDealerCostPrice(v string) *TaobaoFenxiaoProductSkuAddRequest {
    s.DealerCostPrice = &v
    return s
}
func (s *TaobaoFenxiaoProductSkuAddRequest) SetQuantity(v int64) *TaobaoFenxiaoProductSkuAddRequest {
    s.Quantity = &v
    return s
}
func (s *TaobaoFenxiaoProductSkuAddRequest) SetSkuNumber(v string) *TaobaoFenxiaoProductSkuAddRequest {
    s.SkuNumber = &v
    return s
}
func (s *TaobaoFenxiaoProductSkuAddRequest) SetProperties(v string) *TaobaoFenxiaoProductSkuAddRequest {
    s.Properties = &v
    return s
}

func (req *TaobaoFenxiaoProductSkuAddRequest) ToMap() map[string]interface{} {
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

func (req *TaobaoFenxiaoProductSkuAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}