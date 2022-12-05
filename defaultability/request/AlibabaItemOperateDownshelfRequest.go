package request


type AlibabaItemOperateDownshelfRequest struct {
    /*
        商品ID     */
    ItemId  *int64 `json:"item_id" required:"true" `
}

func (s *AlibabaItemOperateDownshelfRequest) SetItemId(v int64) *AlibabaItemOperateDownshelfRequest {
    s.ItemId = &v
    return s
}

func (req *AlibabaItemOperateDownshelfRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    return paramMap
}

func (req *AlibabaItemOperateDownshelfRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}