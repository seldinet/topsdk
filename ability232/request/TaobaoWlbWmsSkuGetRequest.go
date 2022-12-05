package request


type TaobaoWlbWmsSkuGetRequest struct {
    /*
        菜鸟商品ID,与itemcode必须有一个值不为空     */
    ItemCode  *string `json:"item_code,omitempty" required:"false" `
    /*
        商家商品编码,与itemid必须有一个值不为空     */
    ItemId  *string `json:"item_id,omitempty" required:"false" `
    /*
        货主ID     */
    OwnerUserId  *string `json:"owner_user_id,omitempty" required:"false" `
}

func (s *TaobaoWlbWmsSkuGetRequest) SetItemCode(v string) *TaobaoWlbWmsSkuGetRequest {
    s.ItemCode = &v
    return s
}
func (s *TaobaoWlbWmsSkuGetRequest) SetItemId(v string) *TaobaoWlbWmsSkuGetRequest {
    s.ItemId = &v
    return s
}
func (s *TaobaoWlbWmsSkuGetRequest) SetOwnerUserId(v string) *TaobaoWlbWmsSkuGetRequest {
    s.OwnerUserId = &v
    return s
}

func (req *TaobaoWlbWmsSkuGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemCode != nil) {
        paramMap["item_code"] = *req.ItemCode
    }
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.OwnerUserId != nil) {
        paramMap["owner_user_id"] = *req.OwnerUserId
    }
    return paramMap
}

func (req *TaobaoWlbWmsSkuGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}