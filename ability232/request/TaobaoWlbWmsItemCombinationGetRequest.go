package request


type TaobaoWlbWmsItemCombinationGetRequest struct {
    /*
        货品Id     */
    Itemid  *int64 `json:"itemid" required:"true" `
}

func (s *TaobaoWlbWmsItemCombinationGetRequest) SetItemid(v int64) *TaobaoWlbWmsItemCombinationGetRequest {
    s.Itemid = &v
    return s
}

func (req *TaobaoWlbWmsItemCombinationGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Itemid != nil) {
        paramMap["itemid"] = *req.Itemid
    }
    return paramMap
}

func (req *TaobaoWlbWmsItemCombinationGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}