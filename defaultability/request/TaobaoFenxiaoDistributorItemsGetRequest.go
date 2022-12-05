package request

import (
        "topsdk/util"
    )

type TaobaoFenxiaoDistributorItemsGetRequest struct {
    /*
        分销商ID 。     */
    DistributorId  *int64 `json:"distributor_id,omitempty" required:"false" `
    /*
        产品ID     */
    ProductId  *int64 `json:"product_id,omitempty" required:"false" `
    /*
        设置开始时间。空为不设置。     */
    StartModified  *util.LocalTime `json:"start_modified,omitempty" required:"false" `
    /*
        设置结束时间,空为不设置。     */
    EndModified  *util.LocalTime `json:"end_modified,omitempty" required:"false" `
    /*
        每页记录数（默认20，最大50）     */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        页码（大于0的整数，默认1）     */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
}

func (s *TaobaoFenxiaoDistributorItemsGetRequest) SetDistributorId(v int64) *TaobaoFenxiaoDistributorItemsGetRequest {
    s.DistributorId = &v
    return s
}
func (s *TaobaoFenxiaoDistributorItemsGetRequest) SetProductId(v int64) *TaobaoFenxiaoDistributorItemsGetRequest {
    s.ProductId = &v
    return s
}
func (s *TaobaoFenxiaoDistributorItemsGetRequest) SetStartModified(v util.LocalTime) *TaobaoFenxiaoDistributorItemsGetRequest {
    s.StartModified = &v
    return s
}
func (s *TaobaoFenxiaoDistributorItemsGetRequest) SetEndModified(v util.LocalTime) *TaobaoFenxiaoDistributorItemsGetRequest {
    s.EndModified = &v
    return s
}
func (s *TaobaoFenxiaoDistributorItemsGetRequest) SetPageSize(v int64) *TaobaoFenxiaoDistributorItemsGetRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoFenxiaoDistributorItemsGetRequest) SetPageNo(v int64) *TaobaoFenxiaoDistributorItemsGetRequest {
    s.PageNo = &v
    return s
}

func (req *TaobaoFenxiaoDistributorItemsGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.DistributorId != nil) {
        paramMap["distributor_id"] = *req.DistributorId
    }
    if(req.ProductId != nil) {
        paramMap["product_id"] = *req.ProductId
    }
    if(req.StartModified != nil) {
        paramMap["start_modified"] = *req.StartModified
    }
    if(req.EndModified != nil) {
        paramMap["end_modified"] = *req.EndModified
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    return paramMap
}

func (req *TaobaoFenxiaoDistributorItemsGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}