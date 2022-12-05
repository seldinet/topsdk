package request


type TaobaoWlbItemCombinationGetRequest struct {
    /*
        要查询的组合商品id     */
    ItemId  *int64 `json:"item_id" required:"true" `
}

func (s *TaobaoWlbItemCombinationGetRequest) SetItemId(v int64) *TaobaoWlbItemCombinationGetRequest {
    s.ItemId = &v
    return s
}

func (req *TaobaoWlbItemCombinationGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    return paramMap
}

func (req *TaobaoWlbItemCombinationGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}