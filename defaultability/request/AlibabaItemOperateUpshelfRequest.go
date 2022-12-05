package request


type AlibabaItemOperateUpshelfRequest struct {
    /*
        商品ID     */
    ItemId  *int64 `json:"item_id" required:"true" `
    /*
        商品库存     */
    Quantity  *int64 `json:"quantity,omitempty" required:"false" `
}

func (s *AlibabaItemOperateUpshelfRequest) SetItemId(v int64) *AlibabaItemOperateUpshelfRequest {
    s.ItemId = &v
    return s
}
func (s *AlibabaItemOperateUpshelfRequest) SetQuantity(v int64) *AlibabaItemOperateUpshelfRequest {
    s.Quantity = &v
    return s
}

func (req *AlibabaItemOperateUpshelfRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.Quantity != nil) {
        paramMap["quantity"] = *req.Quantity
    }
    return paramMap
}

func (req *AlibabaItemOperateUpshelfRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}