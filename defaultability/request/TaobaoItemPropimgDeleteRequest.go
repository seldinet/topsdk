package request


type TaobaoItemPropimgDeleteRequest struct {
    /*
        商品数字ID，必选     */
    NumIid  *int64 `json:"num_iid" required:"true" `
    /*
        商品属性图片ID     */
    Id  *int64 `json:"id" required:"true" `
}

func (s *TaobaoItemPropimgDeleteRequest) SetNumIid(v int64) *TaobaoItemPropimgDeleteRequest {
    s.NumIid = &v
    return s
}
func (s *TaobaoItemPropimgDeleteRequest) SetId(v int64) *TaobaoItemPropimgDeleteRequest {
    s.Id = &v
    return s
}

func (req *TaobaoItemPropimgDeleteRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.NumIid != nil) {
        paramMap["num_iid"] = *req.NumIid
    }
    if(req.Id != nil) {
        paramMap["id"] = *req.Id
    }
    return paramMap
}

func (req *TaobaoItemPropimgDeleteRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}