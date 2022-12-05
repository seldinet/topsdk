package request


type TaobaoInventoryQueryRequest struct {
    /*
        后端商品ID 列表，控制到50个     */
    ScItemIds  *string `json:"sc_item_ids" required:"true" `
    /*
        后端商品的商家编码列表，控制到50个     */
    ScItemCodes  *string `json:"sc_item_codes,omitempty" required:"false" `
    /*
        卖家昵称     */
    SellerNick  *string `json:"seller_nick,omitempty" required:"false" `
    /*
        仓库列表:GLY001^GLY002     */
    StoreCodes  *string `json:"store_codes,omitempty" required:"false" `
}

func (s *TaobaoInventoryQueryRequest) SetScItemIds(v string) *TaobaoInventoryQueryRequest {
    s.ScItemIds = &v
    return s
}
func (s *TaobaoInventoryQueryRequest) SetScItemCodes(v string) *TaobaoInventoryQueryRequest {
    s.ScItemCodes = &v
    return s
}
func (s *TaobaoInventoryQueryRequest) SetSellerNick(v string) *TaobaoInventoryQueryRequest {
    s.SellerNick = &v
    return s
}
func (s *TaobaoInventoryQueryRequest) SetStoreCodes(v string) *TaobaoInventoryQueryRequest {
    s.StoreCodes = &v
    return s
}

func (req *TaobaoInventoryQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ScItemIds != nil) {
        paramMap["sc_item_ids"] = *req.ScItemIds
    }
    if(req.ScItemCodes != nil) {
        paramMap["sc_item_codes"] = *req.ScItemCodes
    }
    if(req.SellerNick != nil) {
        paramMap["seller_nick"] = *req.SellerNick
    }
    if(req.StoreCodes != nil) {
        paramMap["store_codes"] = *req.StoreCodes
    }
    return paramMap
}

func (req *TaobaoInventoryQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}