package request

import (
        "topsdk/defaultability/domain"
        "topsdk/util"
    )

type TmallItemPriceUpdateRequest struct {
    /*
        商品ID     */
    ItemId  *int64 `json:"item_id" required:"true" `
    /*
        被更新商品价格     */
    ItemPrice  *string `json:"item_price,omitempty" required:"false" `
    /*
        更新SKU价格时候的SKU价格对象；如果没有SKU或者不更新SKU价格，可以不填;查找SKU目前支持ID，属性串和商家编码三种模式，建议选用一种最合适的，切勿滥用，一次调用中如果混合使用，更新结果不可预期！     */
    SkuPrices  *[]domain.TmallItemPriceUpdateUpdateSkuPrice `json:"sku_prices,omitempty" required:"false" `
    /*
        商品价格更新时候的可选参数     */
    Options  *domain.TmallItemPriceUpdateUpdateItemPriceOption `json:"options,omitempty" required:"false" `
}

func (s *TmallItemPriceUpdateRequest) SetItemId(v int64) *TmallItemPriceUpdateRequest {
    s.ItemId = &v
    return s
}
func (s *TmallItemPriceUpdateRequest) SetItemPrice(v string) *TmallItemPriceUpdateRequest {
    s.ItemPrice = &v
    return s
}
func (s *TmallItemPriceUpdateRequest) SetSkuPrices(v []domain.TmallItemPriceUpdateUpdateSkuPrice) *TmallItemPriceUpdateRequest {
    s.SkuPrices = &v
    return s
}
func (s *TmallItemPriceUpdateRequest) SetOptions(v domain.TmallItemPriceUpdateUpdateItemPriceOption) *TmallItemPriceUpdateRequest {
    s.Options = &v
    return s
}

func (req *TmallItemPriceUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.ItemPrice != nil) {
        paramMap["item_price"] = *req.ItemPrice
    }
    if(req.SkuPrices != nil) {
        paramMap["sku_prices"] = util.ConvertStructList(*req.SkuPrices)
    }
    if(req.Options != nil) {
        paramMap["options"] = util.ConvertStruct(*req.Options)
    }
    return paramMap
}

func (req *TmallItemPriceUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}