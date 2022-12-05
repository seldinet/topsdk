package request


type TaobaoWlbItemGetRequest struct {
    /*
        商品ID     */
    ItemId  *int64 `json:"item_id" required:"true" `
}

func (s *TaobaoWlbItemGetRequest) SetItemId(v int64) *TaobaoWlbItemGetRequest {
    s.ItemId = &v
    return s
}

func (req *TaobaoWlbItemGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    return paramMap
}

func (req *TaobaoWlbItemGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}