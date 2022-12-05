package request


type TaobaoInventoryInitialItemRequest struct {
    /*
        后端商品id     */
    ScItemId  *int64 `json:"sc_item_id" required:"true" `
    /*
        商品初始库存信息： [{"storeCode":"必选,商家仓库编号","inventoryType":"可选，库存类型 1：正常,2：损坏,3：冻结,10：质押,11-20:用户自定义,默认为1","quantity":"必选,数量"}]     */
    StoreInventorys  *string `json:"store_inventorys" required:"true" `
}

func (s *TaobaoInventoryInitialItemRequest) SetScItemId(v int64) *TaobaoInventoryInitialItemRequest {
    s.ScItemId = &v
    return s
}
func (s *TaobaoInventoryInitialItemRequest) SetStoreInventorys(v string) *TaobaoInventoryInitialItemRequest {
    s.StoreInventorys = &v
    return s
}

func (req *TaobaoInventoryInitialItemRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ScItemId != nil) {
        paramMap["sc_item_id"] = *req.ScItemId
    }
    if(req.StoreInventorys != nil) {
        paramMap["store_inventorys"] = *req.StoreInventorys
    }
    return paramMap
}

func (req *TaobaoInventoryInitialItemRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}