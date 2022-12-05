package request


type TaobaoWlbItemQueryRequest struct {
    /*
        当前页 defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录 defalutValue��20    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        商品名称     */
    Name  *string `json:"name,omitempty" required:"false" `
    /*
        商品前台销售名字     */
    Title  *string `json:"title,omitempty" required:"false" `
    /*
        商家编码     */
    ItemCode  *string `json:"item_code,omitempty" required:"false" `
    /*
        是否是最小库存单元，只有最小库存单元的商品才可以有库存,值只能给"true","false"来表示;  若值不在范围内，则按true处理     */
    IsSku  *string `json:"is_sku,omitempty" required:"false" `
    /*
        父ID,只有is_sku=1时才能有父ID，商品也可以没有付商品     */
    ParentId  *int64 `json:"parent_id,omitempty" required:"false" `
    /*
        只能输入以下值或空：  ITEM_STATUS_VALID -- 1 表示 有效；  ITEM_STATUS_LOCK -- 2 表示锁住。  若值不在范围内，按ITEM_STATUS_VALID处理     */
    Status  *string `json:"status,omitempty" required:"false" `
    /*
        ITEM类型(只允许输入以下英文或空)  NORMAL 0:普通商品;  COMBINE 1:是否是组合商品  DISTRIBUTION 2:是否是分销商品(货主是别人)  若值不在范围内，则按NORMAL处理     */
    ItemType  *string `json:"item_type,omitempty" required:"false" `
}

func (s *TaobaoWlbItemQueryRequest) SetPageNo(v int64) *TaobaoWlbItemQueryRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoWlbItemQueryRequest) SetPageSize(v int64) *TaobaoWlbItemQueryRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoWlbItemQueryRequest) SetName(v string) *TaobaoWlbItemQueryRequest {
    s.Name = &v
    return s
}
func (s *TaobaoWlbItemQueryRequest) SetTitle(v string) *TaobaoWlbItemQueryRequest {
    s.Title = &v
    return s
}
func (s *TaobaoWlbItemQueryRequest) SetItemCode(v string) *TaobaoWlbItemQueryRequest {
    s.ItemCode = &v
    return s
}
func (s *TaobaoWlbItemQueryRequest) SetIsSku(v string) *TaobaoWlbItemQueryRequest {
    s.IsSku = &v
    return s
}
func (s *TaobaoWlbItemQueryRequest) SetParentId(v int64) *TaobaoWlbItemQueryRequest {
    s.ParentId = &v
    return s
}
func (s *TaobaoWlbItemQueryRequest) SetStatus(v string) *TaobaoWlbItemQueryRequest {
    s.Status = &v
    return s
}
func (s *TaobaoWlbItemQueryRequest) SetItemType(v string) *TaobaoWlbItemQueryRequest {
    s.ItemType = &v
    return s
}

func (req *TaobaoWlbItemQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.Name != nil) {
        paramMap["name"] = *req.Name
    }
    if(req.Title != nil) {
        paramMap["title"] = *req.Title
    }
    if(req.ItemCode != nil) {
        paramMap["item_code"] = *req.ItemCode
    }
    if(req.IsSku != nil) {
        paramMap["is_sku"] = *req.IsSku
    }
    if(req.ParentId != nil) {
        paramMap["parent_id"] = *req.ParentId
    }
    if(req.Status != nil) {
        paramMap["status"] = *req.Status
    }
    if(req.ItemType != nil) {
        paramMap["item_type"] = *req.ItemType
    }
    return paramMap
}

func (req *TaobaoWlbItemQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}