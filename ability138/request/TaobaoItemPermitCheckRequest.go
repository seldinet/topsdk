package request


type TaobaoItemPermitCheckRequest struct {
    /*
        商品id     */
    ItemId  *int64 `json:"item_id,omitempty" required:"false" `
    /*
        类目id     */
    Cid  *int64 `json:"cid" required:"true" `
    /*
        发布类型。可选值:fixed(一口价),auction(拍卖)     */
    Type  *string `json:"type" required:"true" `
}

func (s *TaobaoItemPermitCheckRequest) SetItemId(v int64) *TaobaoItemPermitCheckRequest {
    s.ItemId = &v
    return s
}
func (s *TaobaoItemPermitCheckRequest) SetCid(v int64) *TaobaoItemPermitCheckRequest {
    s.Cid = &v
    return s
}
func (s *TaobaoItemPermitCheckRequest) SetType(v string) *TaobaoItemPermitCheckRequest {
    s.Type = &v
    return s
}

func (req *TaobaoItemPermitCheckRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.Cid != nil) {
        paramMap["cid"] = *req.Cid
    }
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    return paramMap
}

func (req *TaobaoItemPermitCheckRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}