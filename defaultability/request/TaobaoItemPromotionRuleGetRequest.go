package request


type TaobaoItemPromotionRuleGetRequest struct {
    /*
        商品ID     */
    ItemId  *int64 `json:"item_id" required:"true" `
}

func (s *TaobaoItemPromotionRuleGetRequest) SetItemId(v int64) *TaobaoItemPromotionRuleGetRequest {
    s.ItemId = &v
    return s
}

func (req *TaobaoItemPromotionRuleGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    return paramMap
}

func (req *TaobaoItemPromotionRuleGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}