package request


type TaobaoRegionSaleQueryRequest struct {
    /*
        商品id     */
    ItemId  *int64 `json:"item_id" required:"true" `
    /*
        无sku传0     */
    SkuId  *int64 `json:"sku_id" required:"true" `
    /*
        1:国家;2:省;3: 市;4:区县     */
    SaleRegionLevel  *int64 `json:"sale_region_level" required:"true" `
}

func (s *TaobaoRegionSaleQueryRequest) SetItemId(v int64) *TaobaoRegionSaleQueryRequest {
    s.ItemId = &v
    return s
}
func (s *TaobaoRegionSaleQueryRequest) SetSkuId(v int64) *TaobaoRegionSaleQueryRequest {
    s.SkuId = &v
    return s
}
func (s *TaobaoRegionSaleQueryRequest) SetSaleRegionLevel(v int64) *TaobaoRegionSaleQueryRequest {
    s.SaleRegionLevel = &v
    return s
}

func (req *TaobaoRegionSaleQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.SkuId != nil) {
        paramMap["sku_id"] = *req.SkuId
    }
    if(req.SaleRegionLevel != nil) {
        paramMap["sale_region_level"] = *req.SaleRegionLevel
    }
    return paramMap
}

func (req *TaobaoRegionSaleQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}