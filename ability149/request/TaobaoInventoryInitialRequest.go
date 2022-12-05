package request


type TaobaoInventoryInitialRequest struct {
    /*
        商家仓库编码     */
    StoreCode  *string `json:"store_code" required:"true" `
    /*
        商品初始库存信息： [{"scItemId":"商品后端ID，如果有传scItemCode,参数可以为0","scItemCode":"商品商家编码","inventoryType":"库存类型  1：正常,2：损坏,3：冻结,10：质押,11-20:用户自定义","quantity":"数量"}]     */
    Items  *string `json:"items" required:"true" `
}

func (s *TaobaoInventoryInitialRequest) SetStoreCode(v string) *TaobaoInventoryInitialRequest {
    s.StoreCode = &v
    return s
}
func (s *TaobaoInventoryInitialRequest) SetItems(v string) *TaobaoInventoryInitialRequest {
    s.Items = &v
    return s
}

func (req *TaobaoInventoryInitialRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.StoreCode != nil) {
        paramMap["store_code"] = *req.StoreCode
    }
    if(req.Items != nil) {
        paramMap["items"] = *req.Items
    }
    return paramMap
}

func (req *TaobaoInventoryInitialRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}