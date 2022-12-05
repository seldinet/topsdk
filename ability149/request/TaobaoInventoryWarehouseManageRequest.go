package request

import (
        "topsdk/ability149/domain"
        "topsdk/util"
    )

type TaobaoInventoryWarehouseManageRequest struct {
    /*
        仓库信息     */
    WareHouseDto  *domain.TaobaoInventoryWarehouseManageWareHouseDto `json:"ware_house_dto" required:"true" `
}

func (s *TaobaoInventoryWarehouseManageRequest) SetWareHouseDto(v domain.TaobaoInventoryWarehouseManageWareHouseDto) *TaobaoInventoryWarehouseManageRequest {
    s.WareHouseDto = &v
    return s
}

func (req *TaobaoInventoryWarehouseManageRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.WareHouseDto != nil) {
        paramMap["ware_house_dto"] = util.ConvertStruct(*req.WareHouseDto)
    }
    return paramMap
}

func (req *TaobaoInventoryWarehouseManageRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}