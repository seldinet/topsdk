package request

import (
        "topsdk/ability138/domain"
        "topsdk/util"
    )

type TmallItemQuantityUpdateRequest struct {
    /*
        商品id     */
    ItemId  *int64 `json:"item_id" required:"true" `
    /*
        商品库存数；增量编辑方式支持正数、负数。（无SKU商品使用这个字段）     */
    ItemQuantity  *int64 `json:"item_quantity,omitempty" required:"false" `
    /*
        更新SKU库存时候的SKU库存对象；如果没有SKU或者不更新SKU库存，可以不填;查找SKU目前支持ID，属性串和商家编码三种模式，建议选用一种最合适的，切勿滥用，一次调用中如果混合使用，更新结果不可预期！     */
    SkuQuantities  *[]domain.TmallItemQuantityUpdateUpdateSkuQuantity `json:"sku_quantities,omitempty" required:"false" `
    /*
        商品库存更新时候的可选参数     */
    Options  *domain.TmallItemQuantityUpdateUpdateItemQuantityOption `json:"options,omitempty" required:"false" `
}

func (s *TmallItemQuantityUpdateRequest) SetItemId(v int64) *TmallItemQuantityUpdateRequest {
    s.ItemId = &v
    return s
}
func (s *TmallItemQuantityUpdateRequest) SetItemQuantity(v int64) *TmallItemQuantityUpdateRequest {
    s.ItemQuantity = &v
    return s
}
func (s *TmallItemQuantityUpdateRequest) SetSkuQuantities(v []domain.TmallItemQuantityUpdateUpdateSkuQuantity) *TmallItemQuantityUpdateRequest {
    s.SkuQuantities = &v
    return s
}
func (s *TmallItemQuantityUpdateRequest) SetOptions(v domain.TmallItemQuantityUpdateUpdateItemQuantityOption) *TmallItemQuantityUpdateRequest {
    s.Options = &v
    return s
}

func (req *TmallItemQuantityUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.ItemQuantity != nil) {
        paramMap["item_quantity"] = *req.ItemQuantity
    }
    if(req.SkuQuantities != nil) {
        paramMap["sku_quantities"] = util.ConvertStructList(*req.SkuQuantities)
    }
    if(req.Options != nil) {
        paramMap["options"] = util.ConvertStruct(*req.Options)
    }
    return paramMap
}

func (req *TmallItemQuantityUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}