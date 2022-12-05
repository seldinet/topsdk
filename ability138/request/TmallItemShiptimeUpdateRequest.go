package request

import (
        "topsdk/ability138/domain"
        "topsdk/util"
    )

type TmallItemShiptimeUpdateRequest struct {
    /*
        商品ID     */
    ItemId  *int64 `json:"item_id" required:"true" `
    /*
        被更新发货时间（商品级）；格式和具体设置的发货时间格式相关。绝对发货时间填写yyyy-MM-dd;相对发货时间填写数字。发货时间必须在当前时间后三天。如果设置的绝对时间小于当前时间的三天后，会清除该商品的发货时间设置。如果是相对时间小于3，则会提示出错。如果shiptimeType为0，要清除商品上的发货时间，该字段可以填任意字符,也可以不填。     */
    ShipTime  *string `json:"ship_time,omitempty" required:"false" `
    /*
        被更新SKU的发货时间，后台会根据三个子属性去查找匹配的sku，如果找到就默认对sku进行更新，当无匹配sku且更新类型针对sku，会报错。     */
    SkuShipTimes  *[]domain.TmallItemShiptimeUpdateUpdateSkuShipTime `json:"sku_ship_times,omitempty" required:"false" `
    /*
        批量更新商品/SKU发货时间的备选项     */
    Option  *domain.TmallItemShiptimeUpdateUpdateItemShipTimeOption `json:"option" required:"true" `
}

func (s *TmallItemShiptimeUpdateRequest) SetItemId(v int64) *TmallItemShiptimeUpdateRequest {
    s.ItemId = &v
    return s
}
func (s *TmallItemShiptimeUpdateRequest) SetShipTime(v string) *TmallItemShiptimeUpdateRequest {
    s.ShipTime = &v
    return s
}
func (s *TmallItemShiptimeUpdateRequest) SetSkuShipTimes(v []domain.TmallItemShiptimeUpdateUpdateSkuShipTime) *TmallItemShiptimeUpdateRequest {
    s.SkuShipTimes = &v
    return s
}
func (s *TmallItemShiptimeUpdateRequest) SetOption(v domain.TmallItemShiptimeUpdateUpdateItemShipTimeOption) *TmallItemShiptimeUpdateRequest {
    s.Option = &v
    return s
}

func (req *TmallItemShiptimeUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.ShipTime != nil) {
        paramMap["ship_time"] = *req.ShipTime
    }
    if(req.SkuShipTimes != nil) {
        paramMap["sku_ship_times"] = util.ConvertStructList(*req.SkuShipTimes)
    }
    if(req.Option != nil) {
        paramMap["option"] = util.ConvertStruct(*req.Option)
    }
    return paramMap
}

func (req *TmallItemShiptimeUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}