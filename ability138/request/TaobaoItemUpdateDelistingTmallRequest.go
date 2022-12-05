package request


type TaobaoItemUpdateDelistingTmallRequest struct {
    /*
        商品数字ID，该参数必须     */
    NumIid  *int64 `json:"num_iid" required:"true" `
}

func (s *TaobaoItemUpdateDelistingTmallRequest) SetNumIid(v int64) *TaobaoItemUpdateDelistingTmallRequest {
    s.NumIid = &v
    return s
}

func (req *TaobaoItemUpdateDelistingTmallRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.NumIid != nil) {
        paramMap["num_iid"] = *req.NumIid
    }
    return paramMap
}

func (req *TaobaoItemUpdateDelistingTmallRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}