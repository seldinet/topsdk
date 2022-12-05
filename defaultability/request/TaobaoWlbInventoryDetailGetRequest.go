package request

import (
        "topsdk/util"
    )

type TaobaoWlbInventoryDetailGetRequest struct {
    /*
        商品ID     */
    ItemId  *int64 `json:"item_id" required:"true" `
    /*
        库存类型列表，值包括：
VENDIBLE--可销售库存
FREEZE--冻结库存
ONWAY--在途库存
DEFECT--残次品库存
ENGINE_DAMAGE--机损
BOX_DAMAGE--箱损
EXPIRATION--过保     */
    InventoryTypeList  *[]string `json:"inventory_type_list,omitempty" required:"false" `
    /*
        仓库编码     */
    StoreCode  *string `json:"store_code,omitempty" required:"false" `
}

func (s *TaobaoWlbInventoryDetailGetRequest) SetItemId(v int64) *TaobaoWlbInventoryDetailGetRequest {
    s.ItemId = &v
    return s
}
func (s *TaobaoWlbInventoryDetailGetRequest) SetInventoryTypeList(v []string) *TaobaoWlbInventoryDetailGetRequest {
    s.InventoryTypeList = &v
    return s
}
func (s *TaobaoWlbInventoryDetailGetRequest) SetStoreCode(v string) *TaobaoWlbInventoryDetailGetRequest {
    s.StoreCode = &v
    return s
}

func (req *TaobaoWlbInventoryDetailGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.InventoryTypeList != nil) {
        paramMap["inventory_type_list"] = util.ConvertBasicList(*req.InventoryTypeList)
    }
    if(req.StoreCode != nil) {
        paramMap["store_code"] = *req.StoreCode
    }
    return paramMap
}

func (req *TaobaoWlbInventoryDetailGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}