package request

import (
        "topsdk/util"
    )

type TmallChannelProductsQueryRequest struct {
    /*
        商家产品编码     */
    ProductNumber  *string `json:"product_number,omitempty" required:"false" `
    /*
        商家SKU编码     */
    SkuNumber  *string `json:"sku_number,omitempty" required:"false" `
    /*
        产品Id     */
    ProductIds  *[]int64 `json:"product_ids,omitempty" required:"false" `
    /*
        分页大小     */
    PageSize  *int64 `json:"page_size" required:"true" `
    /*
        产品线Id     */
    ProductLineId  *int64 `json:"product_line_id,omitempty" required:"false" `
    /*
        页码数，从1开始     */
    PageNum  *int64 `json:"page_num" required:"true" `
}

func (s *TmallChannelProductsQueryRequest) SetProductNumber(v string) *TmallChannelProductsQueryRequest {
    s.ProductNumber = &v
    return s
}
func (s *TmallChannelProductsQueryRequest) SetSkuNumber(v string) *TmallChannelProductsQueryRequest {
    s.SkuNumber = &v
    return s
}
func (s *TmallChannelProductsQueryRequest) SetProductIds(v []int64) *TmallChannelProductsQueryRequest {
    s.ProductIds = &v
    return s
}
func (s *TmallChannelProductsQueryRequest) SetPageSize(v int64) *TmallChannelProductsQueryRequest {
    s.PageSize = &v
    return s
}
func (s *TmallChannelProductsQueryRequest) SetProductLineId(v int64) *TmallChannelProductsQueryRequest {
    s.ProductLineId = &v
    return s
}
func (s *TmallChannelProductsQueryRequest) SetPageNum(v int64) *TmallChannelProductsQueryRequest {
    s.PageNum = &v
    return s
}

func (req *TmallChannelProductsQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ProductNumber != nil) {
        paramMap["product_number"] = *req.ProductNumber
    }
    if(req.SkuNumber != nil) {
        paramMap["sku_number"] = *req.SkuNumber
    }
    if(req.ProductIds != nil) {
        paramMap["product_ids"] = util.ConvertBasicList(*req.ProductIds)
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.ProductLineId != nil) {
        paramMap["product_line_id"] = *req.ProductLineId
    }
    if(req.PageNum != nil) {
        paramMap["page_num"] = *req.PageNum
    }
    return paramMap
}

func (req *TmallChannelProductsQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}