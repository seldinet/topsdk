package request

import (
        "topsdk/ability149/domain"
        "topsdk/util"
    )

type TaobaoInventoryMerchantAdjustRequest struct {
    /*
        调整库存对象     */
    InventoryCheck  *domain.TaobaoInventoryMerchantAdjustInventoryCheckDto `json:"inventory_check" required:"true" `
}

func (s *TaobaoInventoryMerchantAdjustRequest) SetInventoryCheck(v domain.TaobaoInventoryMerchantAdjustInventoryCheckDto) *TaobaoInventoryMerchantAdjustRequest {
    s.InventoryCheck = &v
    return s
}

func (req *TaobaoInventoryMerchantAdjustRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.InventoryCheck != nil) {
        paramMap["inventory_check"] = util.ConvertStruct(*req.InventoryCheck)
    }
    return paramMap
}

func (req *TaobaoInventoryMerchantAdjustRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}