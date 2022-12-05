package request


type TaobaoScitemGetRequest struct {
    /*
        商品id     */
    ItemId  *int64 `json:"item_id" required:"true" `
}

func (s *TaobaoScitemGetRequest) SetItemId(v int64) *TaobaoScitemGetRequest {
    s.ItemId = &v
    return s
}

func (req *TaobaoScitemGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    return paramMap
}

func (req *TaobaoScitemGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}