package request


type TaobaoScitemQueryRequest struct {
    /*
        商品名称     */
    ItemName  *string `json:"item_name,omitempty" required:"false" `
    /*
        商家给商品的一个编码     */
    OuterCode  *string `json:"outer_code,omitempty" required:"false" `
    /*
        仓库编码     */
    WmsCode  *string `json:"wms_code,omitempty" required:"false" `
    /*
        条形码     */
    BarCode  *string `json:"bar_code,omitempty" required:"false" `
    /*
        ITEM类型(只允许输入以下英文或空) NORMAL 0:普通商品; COMBINE 1:是否是组合商品 DISTRIBUTION     */
    ItemType  *int64 `json:"item_type,omitempty" required:"false" `
    /*
        当前页码数     */
    PageIndex  *int64 `json:"page_index,omitempty" required:"false" `
    /*
        分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录     */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
}

func (s *TaobaoScitemQueryRequest) SetItemName(v string) *TaobaoScitemQueryRequest {
    s.ItemName = &v
    return s
}
func (s *TaobaoScitemQueryRequest) SetOuterCode(v string) *TaobaoScitemQueryRequest {
    s.OuterCode = &v
    return s
}
func (s *TaobaoScitemQueryRequest) SetWmsCode(v string) *TaobaoScitemQueryRequest {
    s.WmsCode = &v
    return s
}
func (s *TaobaoScitemQueryRequest) SetBarCode(v string) *TaobaoScitemQueryRequest {
    s.BarCode = &v
    return s
}
func (s *TaobaoScitemQueryRequest) SetItemType(v int64) *TaobaoScitemQueryRequest {
    s.ItemType = &v
    return s
}
func (s *TaobaoScitemQueryRequest) SetPageIndex(v int64) *TaobaoScitemQueryRequest {
    s.PageIndex = &v
    return s
}
func (s *TaobaoScitemQueryRequest) SetPageSize(v int64) *TaobaoScitemQueryRequest {
    s.PageSize = &v
    return s
}

func (req *TaobaoScitemQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemName != nil) {
        paramMap["item_name"] = *req.ItemName
    }
    if(req.OuterCode != nil) {
        paramMap["outer_code"] = *req.OuterCode
    }
    if(req.WmsCode != nil) {
        paramMap["wms_code"] = *req.WmsCode
    }
    if(req.BarCode != nil) {
        paramMap["bar_code"] = *req.BarCode
    }
    if(req.ItemType != nil) {
        paramMap["item_type"] = *req.ItemType
    }
    if(req.PageIndex != nil) {
        paramMap["page_index"] = *req.PageIndex
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    return paramMap
}

func (req *TaobaoScitemQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}