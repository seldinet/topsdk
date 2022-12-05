package request

import (
        "topsdk/util"
    )

type TaobaoFenxiaoProductsGetRequest struct {
    /*
        商家编码     */
    OuterId  *string `json:"outer_id,omitempty" required:"false" `
    /*
        产品线ID     */
    ProductcatId  *int64 `json:"productcat_id,omitempty" required:"false" `
    /*
        产品ID列表（最大限制30），用逗号分割，例如：“1001,1002,1003,1004,1005”     */
    Pids  *[]int64 `json:"pids,omitempty" required:"false" `
    /*
        指定查询额外的信息，可选值：skus（sku数据）、images（多图），多个可选值用逗号分割。     */
    Fields  *[]string `json:"fields,omitempty" required:"false" `
    /*
        开始修改时间     */
    StartModified  *util.LocalTime `json:"start_modified,omitempty" required:"false" `
    /*
        结束修改时间     */
    EndModified  *util.LocalTime `json:"end_modified,omitempty" required:"false" `
    /*
        页码（大于0的整数，默认1）     */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        每页记录数（默认20，最大50）     */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        sku商家编码     */
    SkuNumber  *string `json:"sku_number,omitempty" required:"false" `
    /*
        查询产品列表时，查询入参“是否需要授权”
yes:需要授权 
no:不需要授权     */
    IsAuthz  *string `json:"is_authz,omitempty" required:"false" `
    /*
        可根据导入的商品ID列表查询，优先级次于产品ID、sku_numbers，高于其他分页查询条件。最大限制20，用逗号分割，例如：“1001,1002,1003,1004,1005”     */
    ItemIds  *[]int64 `json:"item_ids,omitempty" required:"false" `
}

func (s *TaobaoFenxiaoProductsGetRequest) SetOuterId(v string) *TaobaoFenxiaoProductsGetRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetRequest) SetProductcatId(v int64) *TaobaoFenxiaoProductsGetRequest {
    s.ProductcatId = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetRequest) SetPids(v []int64) *TaobaoFenxiaoProductsGetRequest {
    s.Pids = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetRequest) SetFields(v []string) *TaobaoFenxiaoProductsGetRequest {
    s.Fields = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetRequest) SetStartModified(v util.LocalTime) *TaobaoFenxiaoProductsGetRequest {
    s.StartModified = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetRequest) SetEndModified(v util.LocalTime) *TaobaoFenxiaoProductsGetRequest {
    s.EndModified = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetRequest) SetPageNo(v int64) *TaobaoFenxiaoProductsGetRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetRequest) SetPageSize(v int64) *TaobaoFenxiaoProductsGetRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetRequest) SetSkuNumber(v string) *TaobaoFenxiaoProductsGetRequest {
    s.SkuNumber = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetRequest) SetIsAuthz(v string) *TaobaoFenxiaoProductsGetRequest {
    s.IsAuthz = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetRequest) SetItemIds(v []int64) *TaobaoFenxiaoProductsGetRequest {
    s.ItemIds = &v
    return s
}

func (req *TaobaoFenxiaoProductsGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OuterId != nil) {
        paramMap["outer_id"] = *req.OuterId
    }
    if(req.ProductcatId != nil) {
        paramMap["productcat_id"] = *req.ProductcatId
    }
    if(req.Pids != nil) {
        paramMap["pids"] = util.ConvertBasicList(*req.Pids)
    }
    if(req.Fields != nil) {
        paramMap["fields"] = util.ConvertBasicList(*req.Fields)
    }
    if(req.StartModified != nil) {
        paramMap["start_modified"] = *req.StartModified
    }
    if(req.EndModified != nil) {
        paramMap["end_modified"] = *req.EndModified
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.SkuNumber != nil) {
        paramMap["sku_number"] = *req.SkuNumber
    }
    if(req.IsAuthz != nil) {
        paramMap["is_authz"] = *req.IsAuthz
    }
    if(req.ItemIds != nil) {
        paramMap["item_ids"] = util.ConvertBasicList(*req.ItemIds)
    }
    return paramMap
}

func (req *TaobaoFenxiaoProductsGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}