package request


type TaobaoItemUpdateListingTmallRequest struct {
    /*
        商品数字ID，该参数必须     */
    NumIid  *int64 `json:"num_iid" required:"true" `
    /*
        需要上架的商品的数量。取值范围:大于零的整数。如果商品有sku，则上架数量默认为所有sku数量总和，不可修改。否则商品数量根据设置数量调整为num     */
    Num  *int64 `json:"num" required:"true" `
}

func (s *TaobaoItemUpdateListingTmallRequest) SetNumIid(v int64) *TaobaoItemUpdateListingTmallRequest {
    s.NumIid = &v
    return s
}
func (s *TaobaoItemUpdateListingTmallRequest) SetNum(v int64) *TaobaoItemUpdateListingTmallRequest {
    s.Num = &v
    return s
}

func (req *TaobaoItemUpdateListingTmallRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.NumIid != nil) {
        paramMap["num_iid"] = *req.NumIid
    }
    if(req.Num != nil) {
        paramMap["num"] = *req.Num
    }
    return paramMap
}

func (req *TaobaoItemUpdateListingTmallRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}