package request

import (
        "topsdk/defaultability/domain"
        "topsdk/util"
    )

type TmallItemOuteridUpdateRequest struct {
    /*
        商品ID     */
    ItemId  *int64 `json:"item_id" required:"true" `
    /*
        商品维度商家编码，如果不修改可以不传；清空请设置空串     */
    OuterId  *string `json:"outer_id,omitempty" required:"false" `
    /*
        商品SKU更新OuterId时候用的数据     */
    SkuOuters  *[]domain.TmallItemOuteridUpdateUpdateSkuOuterId `json:"sku_outers,omitempty" required:"false" `
}

func (s *TmallItemOuteridUpdateRequest) SetItemId(v int64) *TmallItemOuteridUpdateRequest {
    s.ItemId = &v
    return s
}
func (s *TmallItemOuteridUpdateRequest) SetOuterId(v string) *TmallItemOuteridUpdateRequest {
    s.OuterId = &v
    return s
}
func (s *TmallItemOuteridUpdateRequest) SetSkuOuters(v []domain.TmallItemOuteridUpdateUpdateSkuOuterId) *TmallItemOuteridUpdateRequest {
    s.SkuOuters = &v
    return s
}

func (req *TmallItemOuteridUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.OuterId != nil) {
        paramMap["outer_id"] = *req.OuterId
    }
    if(req.SkuOuters != nil) {
        paramMap["sku_outers"] = util.ConvertStructList(*req.SkuOuters)
    }
    return paramMap
}

func (req *TmallItemOuteridUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}