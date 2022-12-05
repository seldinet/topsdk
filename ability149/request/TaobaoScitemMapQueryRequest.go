package request


type TaobaoScitemMapQueryRequest struct {
    /*
        map_type为1：前台ic商品id
map_type为2：分销productid     */
    ItemId  *int64 `json:"item_id" required:"true" `
    /*
        map_type为1：前台ic商品skuid 
map_type为2：分销商品的skuid     */
    SkuId  *int64 `json:"sku_id,omitempty" required:"false" `
}

func (s *TaobaoScitemMapQueryRequest) SetItemId(v int64) *TaobaoScitemMapQueryRequest {
    s.ItemId = &v
    return s
}
func (s *TaobaoScitemMapQueryRequest) SetSkuId(v int64) *TaobaoScitemMapQueryRequest {
    s.SkuId = &v
    return s
}

func (req *TaobaoScitemMapQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.SkuId != nil) {
        paramMap["sku_id"] = *req.SkuId
    }
    return paramMap
}

func (req *TaobaoScitemMapQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}