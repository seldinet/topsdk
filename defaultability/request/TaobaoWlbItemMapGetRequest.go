package request


type TaobaoWlbItemMapGetRequest struct {
    /*
        要查询映射关系的物流宝商品id     */
    ItemId  *int64 `json:"item_id" required:"true" `
}

func (s *TaobaoWlbItemMapGetRequest) SetItemId(v int64) *TaobaoWlbItemMapGetRequest {
    s.ItemId = &v
    return s
}

func (req *TaobaoWlbItemMapGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    return paramMap
}

func (req *TaobaoWlbItemMapGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}